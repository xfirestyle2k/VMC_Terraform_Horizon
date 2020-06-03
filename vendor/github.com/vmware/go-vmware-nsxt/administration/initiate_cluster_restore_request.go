/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

type InitiateClusterRestoreRequest struct {

	// Unique id of the backed-up configuration from which the appliance will be restored
	NodeId string `json:"node_id,omitempty"`

	// Timestamp of the backed-up configuration from which the appliance will be restored
	Timestamp int64 `json:"timestamp,omitempty"`
}
