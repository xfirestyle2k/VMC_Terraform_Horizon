/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type StandardHostSwitchSpec struct {
	ResourceType string `json:"resource_type"`

	// Transport Node host switches
	HostSwitches []StandardHostSwitch `json:"host_switches"`
}
