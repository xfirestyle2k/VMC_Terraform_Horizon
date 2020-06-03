/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vmware/go-vmware-nsxt/manager"
	"net/http"
)

func dataSourceNsxtMacPool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtMacPoolRead,

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
		},
	}
}

func dataSourceNsxtMacPoolRead(d *schema.ResourceData, m interface{}) error {
	// Read Mac Pool by name or id
	nsxClient := m.(nsxtClients).NsxtClient
	if nsxClient == nil {
		return dataSourceNotSupportedError()
	}

	objID := d.Get("id").(string)
	objName := d.Get("display_name").(string)
	var obj manager.MacPool
	if objID != "" {
		// Get by id
		objGet, resp, err := nsxClient.PoolManagementApi.ReadMacPool(nsxClient.Context, objID)

		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("Mac pool %s was not found", objID)
		}
		if err != nil {
			return fmt.Errorf("Error while reading Mac pool %s: %v", objID, err)
		}
		obj = objGet
	} else if objName != "" {
		// Get by full name
		// TODO use 2nd parameter localVarOptionals for paging
		objList, _, err := nsxClient.PoolManagementApi.ListMacPools(nsxClient.Context, nil)
		if err != nil {
			return fmt.Errorf("Error while reading Mac pool: %v", err)
		}
		// go over the list to find the correct one
		found := false
		for _, objInList := range objList.Results {
			if objInList.DisplayName == objName {
				if found {
					return fmt.Errorf("Found multiple Mac pool with name '%s'", objName)
				}
				obj = objInList
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Mac pool with name '%s' was not found among %d pools", objName, len(objList.Results))
		}
	} else {
		return fmt.Errorf("Error obtaining Mac pool ID or name during read")
	}

	d.SetId(obj.Id)
	d.Set("display_name", obj.DisplayName)
	d.Set("description", obj.Description)

	return nil
}
