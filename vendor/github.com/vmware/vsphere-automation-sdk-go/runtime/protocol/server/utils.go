/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package server

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
	"net/http"
	"strings"
)

func CopyHeadersToContexts(ctx *core.ExecutionContext, r *http.Request) {

	appCtx := ctx.ApplicationContext()
	secCtx := ctx.SecurityContext()

	vapiAppCtxConstants := map[string]string{
		"opid":                "opId",
		"actid":               "actId",
		"$showunreleasedapis": "$showUnreleasedAPIs",
		"$useragent":          "$userAgent",
		"$donotroute":         "$doNotRoute",
		"vmwaresessionid":     "vmwareSessionId",
		"activationid":        "ActivationId",
	}

	for key, value := range r.Header {
		// req.Header returns a list of values for each key (name)
		val := value[0]

		keyLowerCase := strings.ToLower(key)
		s := strings.Split(keyLowerCase, lib.VAPI_HEADER_PREFIX)
		if len(s) > 1 {
			// Override values in appCtx with headers with "vapi-ctx-" prefix
			// The values from HTTP headers override the body values.
			// if there are multiple values for the same header, the first entry will be chosen.
			if vapiAppCtxKey, ok := vapiAppCtxConstants[s[1]]; ok {
				appCtx.SetProperty(vapiAppCtxKey, &val)
			} else {
				appCtx.SetProperty(s[1], &val)
			}
		} else {
			switch keyLowerCase {
			case lib.HTTP_ACCEPT_LANGUAGE:
				appCtx.SetProperty(lib.HTTP_ACCEPT_LANGUAGE, &val)
			case lib.VAPI_SESSION_HEADER:
				secCtx.SetProperty(security.SESSION_ID, val)
				secCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
			}
		}
	}
	// When the request has $useragent header, it will override the custom one if present.
	if userAgentVal, ok := r.Header["User-Agent"]; ok {
		appCtx.SetProperty("$userAgent", &userAgentVal[0])
	}
}
