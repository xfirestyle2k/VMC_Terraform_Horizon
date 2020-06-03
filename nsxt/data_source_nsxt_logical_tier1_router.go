/* Copyright © 2018 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vmware/go-vmware-nsxt/manager"
	"net/http"
	"strings"
)

func dataSourceNsxtLogicalTier1Router() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtLogicalTier1RouterRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "Unique ID of this resource",
				Optional:    true,
				Computed:    true,
			},
			"display_name": {
				Type:        schema.TypeString,
				Description: "The display name of this resource",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of this resource",
				Optional:    true,
				Computed:    true,
			},
			"edge_cluster_id": {
				Type:        schema.TypeString,
				Description: "The ID of the edge cluster connected to this router",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceNsxtLogicalTier1RouterRead(d *schema.ResourceData, m interface{}) error {
	// Read a logical tier1 router by name or id
	nsxClient := m.(nsxtClients).NsxtClient
	if nsxClient == nil {
		return dataSourceNotSupportedError()
	}

	objID := d.Get("id").(string)
	objName := d.Get("display_name").(string)
	var obj manager.LogicalRouter
	if objID != "" {
		// Get by id
		objGet, resp, err := nsxClient.LogicalRoutingAndServicesApi.ReadLogicalRouter(nsxClient.Context, objID)

		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("Logical tier1 router %s was not found", objID)
		}
		if err != nil {
			return fmt.Errorf("Error while reading logical tier1 router %s: %v", objID, err)
		}
		if objGet.RouterType != "TIER1" {
			return fmt.Errorf("Logical router %s is not a tier1 router", objID)
		}
		obj = objGet
	} else if objName == "" {
		return fmt.Errorf("Error obtaining logical tier1 router ID or name during read")
	} else {
		// Get by full name/prefix
		// TODO use 2nd parameter localVarOptionals for paging
		objList, _, err := nsxClient.LogicalRoutingAndServicesApi.ListLogicalRouters(nsxClient.Context, nil)
		if err != nil {
			return fmt.Errorf("Error while reading logical tier1 routers: %v", err)
		}
		// go over the list to find the correct one (prefer a perfect match. If not - prefix match)
		var perfectMatch []manager.LogicalRouter
		var prefixMatch []manager.LogicalRouter
		for _, objInList := range objList.Results {
			if objInList.RouterType == "TIER1" {
				if strings.HasPrefix(objInList.DisplayName, objName) {
					prefixMatch = append(prefixMatch, objInList)
				}
				if objInList.DisplayName == objName {
					perfectMatch = append(perfectMatch, objInList)
				}
			}
		}
		if len(perfectMatch) > 0 {
			if len(perfectMatch) > 1 {
				return fmt.Errorf("Found multiple logical tier1 routers with name '%s'", objName)
			}
			obj = perfectMatch[0]
		} else if len(prefixMatch) > 0 {
			if len(prefixMatch) > 1 {
				return fmt.Errorf("Found multiple logical tier1 routers with name starting with '%s'", objName)
			}
			obj = prefixMatch[0]
		} else {
			return fmt.Errorf("Logical tier1 router with name '%s' was not found", objName)
		}
	}

	d.SetId(obj.Id)
	d.Set("display_name", obj.DisplayName)
	d.Set("description", obj.Description)
	d.Set("edge_cluster_id", obj.EdgeClusterId)

	return nil
}
