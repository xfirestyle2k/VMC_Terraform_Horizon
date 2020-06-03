/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type ControlConnStatus struct {

	// IP address of the control Node.
	ControlNodeIp string `json:"control_node_ip,omitempty"`

	// Failure status of the control Node for e.g CONNECTION_REFUSED,INCOMPLETE_HOST_CERT.
	FailureStatus string `json:"failure_status,omitempty"`

	// Status of the control Node for e.g  UP, DOWN.
	Status string `json:"status,omitempty"`
}