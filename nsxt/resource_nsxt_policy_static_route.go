/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_0s"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_1s"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"log"
	"strings"
)

func resourceNsxtPolicyStaticRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtPolicyStaticRouteCreate,
		Read:   resourceNsxtPolicyStaticRouteRead,
		Update: resourceNsxtPolicyStaticRouteUpdate,
		Delete: resourceNsxtPolicyStaticRouteDelete,
		Importer: &schema.ResourceImporter{
			State: resourceNsxtPolicyStaticRouteImport,
		},

		Schema: map[string]*schema.Schema{
			"nsx_id":       getNsxIDSchema(),
			"path":         getPathSchema(),
			"display_name": getDisplayNameSchema(),
			"description":  getDescriptionSchema(),
			"revision":     getRevisionSchema(),
			"tag":          getTagsSchema(),
			"gateway_path": getPolicyGatewayPathSchema(),
			"network": {
				Type:         schema.TypeString,
				Description:  "Network address in CIDR format",
				Required:     true,
				ValidateFunc: validateCidr(),
			},
			"next_hop": {
				Type:        schema.TypeList,
				Description: "Next hop routes for network",
				Required:    true,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_distance": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "Cost associated with next hop route",
							ValidateFunc: validation.IntBetween(1, 255),
							Default:      1,
						},
						"ip_address": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "Next hop gateway IP address",
							ValidateFunc: validateSingleIP(),
						},
						"interface": {
							// NOTE: this is called 'scope' in the golang vapi struct
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Interface path associated with current route",
							ValidateFunc: validatePolicyPath(),
						},
					},
				},
			},
		},
	}
}

func patchNsxtPolicyStaticRoute(connector *client.RestConnector, gwID string, route model.StaticRoutes, isT0 bool) error {
	if isT0 {
		routeClient := tier_0s.NewDefaultStaticRoutesClient(connector)
		return routeClient.Patch(gwID, *route.Id, route)
	}
	routeClient := tier_1s.NewDefaultStaticRoutesClient(connector)
	return routeClient.Patch(gwID, *route.Id, route)
}

func deleteNsxtPolicyStaticRoute(connector *client.RestConnector, gwID string, isT0 bool, routeID string) error {
	if isT0 {
		routeClient := tier_0s.NewDefaultStaticRoutesClient(connector)
		return routeClient.Delete(gwID, routeID)
	}
	routeClient := tier_1s.NewDefaultStaticRoutesClient(connector)
	return routeClient.Delete(gwID, routeID)
}

func getNsxtPolicyStaticRouteByID(connector *client.RestConnector, gwID string, isT0 bool, routeID string) (model.StaticRoutes, error) {
	if isT0 {
		routeClient := tier_0s.NewDefaultStaticRoutesClient(connector)
		return routeClient.Get(gwID, routeID)
	}
	routeClient := tier_1s.NewDefaultStaticRoutesClient(connector)
	return routeClient.Get(gwID, routeID)
}

func resourceNsxtPolicyStaticRouteCreate(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	gwPolicyPath := d.Get("gateway_path").(string)
	isT0, gwID := parseGatewayPolicyPath(gwPolicyPath)

	id := d.Get("nsx_id").(string)
	if id == "" {
		id = newUUID()
	} else {
		_, err := getNsxtPolicyStaticRouteByID(connector, gwID, isT0, id)
		if err == nil {
			return fmt.Errorf("Static Route with nsx_id '%s' already exists", id)
		} else if !isNotFoundError(err) {
			return err
		}
	}

	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	tags := getPolicyTagsFromSchema(d)
	network := d.Get("network").(string)

	var nextHopsStructs []model.RouterNexthop
	nextHops := d.Get("next_hop").([]interface{})
	for _, nextHop := range nextHops {
		nextHopMap := nextHop.(map[string]interface{})
		distance := int64(nextHopMap["admin_distance"].(int))
		ip := nextHopMap["ip_address"].(string)
		scope := nextHopMap["interface"].(string)
		var scopeList []string
		if scope != "" {
			scopeList = append(scopeList, scope)
		}
		hopStuct := model.RouterNexthop{
			AdminDistance: &distance,
			IpAddress:     &ip,
			Scope:         scopeList,
		}
		nextHopsStructs = append(nextHopsStructs, hopStuct)
	}

	routeStruct := model.StaticRoutes{
		Id:          &id,
		DisplayName: &displayName,
		Description: &description,
		Tags:        tags,
		Network:     &network,
	}

	if len(nextHopsStructs) > 0 {
		routeStruct.NextHops = nextHopsStructs
	}

	log.Printf("[INFO] Creating Static Route with ID %s", id)
	err := patchNsxtPolicyStaticRoute(connector, gwID, routeStruct, isT0)
	if err != nil {
		return handleCreateError("Static Route", id, err)
	}

	d.SetId(id)
	d.Set("nsx_id", id)

	return resourceNsxtPolicyStaticRouteRead(d, m)
}

func resourceNsxtPolicyStaticRouteRead(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining Static Route ID")
	}

	gwPolicyPath := d.Get("gateway_path").(string)
	isT0, gwID := parseGatewayPolicyPath(gwPolicyPath)

	obj, err := getNsxtPolicyStaticRouteByID(connector, gwID, isT0, id)
	if err != nil {
		return handleReadError(d, "Static Route", id, err)
	}

	d.Set("display_name", obj.DisplayName)
	d.Set("description", obj.Description)
	setPolicyTagsInSchema(d, obj.Tags)
	d.Set("nsx_id", id)
	d.Set("path", obj.Path)
	d.Set("revision", obj.Revision)
	d.Set("network", obj.Network)

	var nextHopMaps []map[string]interface{}
	for _, nextHop := range obj.NextHops {
		nextHopMap := make(map[string]interface{})

		iface := ""
		scope := nextHop.Scope
		if len(scope) > 0 {
			iface = scope[0]
		}
		nextHopMap["interface"] = iface
		nextHopMap["ip_address"] = *nextHop.IpAddress
		nextHopMap["admin_distance"] = *nextHop.AdminDistance

		nextHopMaps = append(nextHopMaps, nextHopMap)
	}

	d.Set("next_hop", nextHopMaps)

	d.SetId(id)

	return nil
}

func resourceNsxtPolicyStaticRouteUpdate(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining Static Route ID")
	}

	gwPolicyPath := d.Get("gateway_path").(string)
	isT0, gwID := parseGatewayPolicyPath(gwPolicyPath)

	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	tags := getPolicyTagsFromSchema(d)
	network := d.Get("network").(string)

	var nextHopsStructs []model.RouterNexthop
	nextHops := d.Get("next_hop").([]interface{})
	for _, nextHop := range nextHops {
		nextHopMap := nextHop.(map[string]interface{})
		distance := int64(nextHopMap["admin_distance"].(int))
		ip := nextHopMap["ip_address"].(string)
		scope := nextHopMap["interface"].(string)
		var scopeList []string
		if scope != "" {
			scopeList = append(scopeList, scope)
		}
		hopStuct := model.RouterNexthop{
			AdminDistance: &distance,
			IpAddress:     &ip,
			Scope:         scopeList,
		}
		nextHopsStructs = append(nextHopsStructs, hopStuct)
	}

	routeStruct := model.StaticRoutes{
		Id:          &id,
		DisplayName: &displayName,
		Description: &description,
		Tags:        tags,
		Network:     &network,
	}

	if len(nextHopsStructs) > 0 {
		routeStruct.NextHops = nextHopsStructs
	}

	log.Printf("[INFO] Updating Static Route with ID %s", id)
	err := patchNsxtPolicyStaticRoute(connector, gwID, routeStruct, isT0)
	if err != nil {
		return handleUpdateError("Static Route", id, err)
	}

	d.SetId(id)
	d.Set("nsx_id", id)

	return resourceNsxtPolicyStaticRouteRead(d, m)
}

func resourceNsxtPolicyStaticRouteDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining Static Route ID")
	}

	gwPolicyPath := d.Get("gateway_path").(string)
	isT0, gwID := parseGatewayPolicyPath(gwPolicyPath)

	err := deleteNsxtPolicyStaticRoute(getPolicyConnector(m), gwID, isT0, id)
	if err != nil {
		return handleDeleteError("Static Route", id, err)
	}

	return nil
}

func resourceNsxtPolicyStaticRouteImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	importID := d.Id()
	s := strings.Split(importID, "/")
	if len(s) != 2 {
		return nil, fmt.Errorf("Please provide <gateway-id>/<static-route-id> as an input")
	}

	gwID := s[0]
	connector := getPolicyConnector(m)
	t0Client := infra.NewDefaultTier0sClient(connector)
	t0gw, err := t0Client.Get(gwID)
	if err != nil {
		if !isNotFoundError(err) {
			return nil, err
		}
		t1Client := infra.NewDefaultTier1sClient(connector)
		t1gw, err := t1Client.Get(gwID)
		if err != nil {
			return nil, err
		}
		d.Set("gateway_path", t1gw.Path)
	} else {
		d.Set("gateway_path", t0gw.Path)
	}

	d.SetId(s[1])

	return []*schema.ResourceData{d}, nil
}
