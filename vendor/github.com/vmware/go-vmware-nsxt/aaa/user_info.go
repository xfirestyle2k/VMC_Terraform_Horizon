/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package aaa

// Authenticated User Info
type UserInfo struct {

	// Permissions
	Roles []NsxRole `json:"roles,omitempty"`

	// User Name
	UserName string `json:"user_name,omitempty"`
}
