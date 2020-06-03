/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EvpnTunnelEndpoints
 * Used by client-side stubs.
 */

package locale_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type EvpnTunnelEndpointsClient interface {

    // Delete evpn tunnel endpoint configuration.
    //
    // @param tier0IdParam tier0 id (required)
    // @param localeServicesIdParam locale services id (required)
    // @param tunnelEndpointIdParam tunnel endpoint id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier0IdParam string, localeServicesIdParam string, tunnelEndpointIdParam string) error

    // Read evpn tunnel endpoint configuration.
    //
    // @param tier0IdParam tier0 id (required)
    // @param localeServicesIdParam locale services id (required)
    // @param tunnelEndpointIdParam tunnel endpoint id (required)
    // @return com.vmware.nsx_policy.model.EvpnTunnelEndpointConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServicesIdParam string, tunnelEndpointIdParam string) (model.EvpnTunnelEndpointConfig, error)

    // List all evpn tunnel endpoint configuration.
    //
    // @param tier0IdParam (required)
    // @param localeServicesIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.EvpnTunnelEndpointConfigListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServicesIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EvpnTunnelEndpointConfigListResult, error)

    // Create a evpn tunnel endpoint config if the tunnel-endpoint-id is not already present, otherwise update the tunnel endpoint configuration.
    //
    // @param tier0IdParam tier0 id (required)
    // @param localeServicesIdParam locale services id (required)
    // @param tunnelEndpointIdParam tunnel endpoint id (required)
    // @param evpnTunnelEndpointConfigParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServicesIdParam string, tunnelEndpointIdParam string, evpnTunnelEndpointConfigParam model.EvpnTunnelEndpointConfig) error

    // Create or update evpn tunnel endpoint configuration.
    //
    // @param tier0IdParam tier0 id (required)
    // @param localeServicesIdParam locale services id (required)
    // @param tunnelEndpointIdParam tunnel endpoint id (required)
    // @param evpnTunnelEndpointConfigParam (required)
    // @return com.vmware.nsx_policy.model.EvpnTunnelEndpointConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, localeServicesIdParam string, tunnelEndpointIdParam string, evpnTunnelEndpointConfigParam model.EvpnTunnelEndpointConfig) (model.EvpnTunnelEndpointConfig, error)
}
