provider "nsxt" {
  host                 = var.host
  vmc_token            = var.vmc_token
  allow_unverified_ssl = true
  enforcement_point    = "vmc-enforcementpoint"
}


###################### creating all Services ######################

// creating Services TCP 8443:
resource "nsxt_policy_service" "Blast_TCP8443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast TCP8443"

  l4_port_set_entry {
    display_name      = "TCP8443"
    description       = "TCP port 8443 entry"
    protocol          = "TCP"
    destination_ports = ["8443"]
  }
}

// creating Services UDP 8443:
resource "nsxt_policy_service" "Blast_UDP8443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast UDP8443"

  l4_port_set_entry {
    display_name      = "UDP8443"
    description       = "UDP port 8443 entry"
    protocol          = "UDP"
    destination_ports = ["8443"]
  }
}

// creating Services TCP 9443:
resource "nsxt_policy_service" "Blast_TCP9443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast TCP9443"

  l4_port_set_entry {
    display_name      = "TCP9443"
    description       = "TCP port 9443 entry"
    protocol          = "TCP"
    destination_ports = ["9443"]
  }
}

// creating Services TCP 22443:
resource "nsxt_policy_service" "Blast_TCP22443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast TCP22443"

  l4_port_set_entry {
    display_name      = "TCP22443"
    description       = "TCP port 22443 entry"
    protocol          = "TCP"
    destination_ports = ["22443"]
  }
}

// creating Services UDP 22443:
resource "nsxt_policy_service" "Blast_UDP22443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast UDP22443"

  l4_port_set_entry {
    display_name      = "UDP22443"
    description       = "UDP port 22443 entry"
    protocol          = "UDP"
    destination_ports = ["22443"]
  }
}

// creating Services TCP 4172:
resource "nsxt_policy_service" "PCoIP_TCP4172" {
  description  = "PCoIP service provisioned by Terraform"
  display_name = "PCoIP TCP4172"

  l4_port_set_entry {
    display_name      = "TCP4172"
    description       = "TCP port 4172 entry"
    protocol          = "TCP"
    destination_ports = ["4172"]
  }
}

// creating Services UDP 4172:
resource "nsxt_policy_service" "PCoIP_UDP4172" {
  description  = "PCoIP service provisioned by Terraform"
  display_name = "PCoIP UDP4172"

  l4_port_set_entry {
    display_name      = "UDP4172"
    description       = "UDP port 4172 entry"
    protocol          = "UDP"
    destination_ports = ["4172"]
  }
}

// creating Services UDP 443:
resource "nsxt_policy_service" "Blast_UDP443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast UDP443"

  l4_port_set_entry {
    display_name      = "UDP433"
    description       = "UDP port 433 entry"
    protocol          = "UDP"
    destination_ports = ["433"]
  }
}

// creating Services TCP 443:
resource "nsxt_policy_service" "Blast_TCP443" {
  description  = "Blast service provisioned by Terraform"
  display_name = "Blast TCP443"

  l4_port_set_entry {
    display_name      = "TCP433"
    description       = "TCP port 433 entry"
    protocol          = "TCP"
    destination_ports = ["433"]
  }
}

// creating Services TCP 443:
resource "nsxt_policy_service" "CDR_MMR_TCP9427" {
  description  = "CDR/MMR service provisioned by Terraform"
  display_name = "CDR/MMR TCP9427"

  l4_port_set_entry {
    display_name      = "TCP9427"
    description       = "TCP port 9427 entry"
    protocol          = "TCP"
    destination_ports = ["9427"]
  }
}

// creating Services TCP 3389:
resource "nsxt_policy_service" "RDP_TCP3389" {
  description  = "RDP service provisioned by Terraform"
  display_name = "RDP TCP3389"

  l4_port_set_entry {
    display_name      = "TCP3389"
    description       = "TCP port 8443 entry"
    protocol          = "TCP"
    destination_ports = ["3389"]
  }
}

// creating Services TCP 32111:
resource "nsxt_policy_service" "USB_TCP32111" {
  description  = "USB service provisioned by Terraform"
  display_name = "USB TCP3389"

  l4_port_set_entry {
    display_name      = "TCP32111"
    description       = "TCP port 32111 entry"
    protocol          = "TCP"
    destination_ports = ["32111"]
  }
}

// creating Services TCP 1433:
resource "nsxt_policy_service" "EventDB_TCP1433" {
  description  = "USB service provisioned by Terraform"
  display_name = "EventDB_TCP1433"

  l4_port_set_entry {
    display_name      = "TCP1433"
    description       = "TCP port 1433 entry"
    protocol          = "TCP"
    destination_ports = ["1433"]
  }
}

// creating Services TCP 3091:
resource "nsxt_policy_service" "vROPS_TCP3091" {
  description  = "vROPS service provisioned by Terraform"
  display_name = "vROPS_TCP3091"

  l4_port_set_entry {
    display_name      = "TCP3091"
    description       = "TCP port 3091 entry"
    protocol          = "TCP"
    destination_ports = ["3091"]
  }
}

// creating Services TCP 3101:
resource "nsxt_policy_service" "vROPS_TCP3101" {
  description  = "vROPS service provisioned by Terraform"
  display_name = "vROPS_TCP3101"

  l4_port_set_entry {
    display_name      = "TCP3101"
    description       = "TCP port 3101 entry"
    protocol          = "TCP"
    destination_ports = ["3101"]
  }
}

// creating Services TCP 3100:
resource "nsxt_policy_service" "vROPS_TCP3100" {
  description  = "vROPS service provisioned by Terraform"
  display_name = "vROPS_TCP3100"

  l4_port_set_entry {
    display_name      = "TCP3100"
    description       = "TCP port 3100 entry"
    protocol          = "TCP"
    destination_ports = ["3100"]
  }
}

###################### creating all Groups ######################

// creating Group for UAG_external:
resource "nsxt_policy_group" "UAG_external" {
  display_name = "UAG_external"
  description  = "Created from Terraform UAG_external"
  domain       = "cgw"
}

// creating Group for UAG_internal:
resource "nsxt_policy_group" "UAG_internal" {
  display_name = "UAG_internal"
  description  = "Created from Terraform UAG_internal"
  domain       = "cgw"
}

// creating Group for ConnectionServer:
resource "nsxt_policy_group" "ConnectionServer" {
  display_name = "ConnectionServer"
  description  = "Created from Terraform ConnectionServer"
  domain       = "cgw"
}

// creating Group for VDI-Environment:
resource "nsxt_policy_group" "VDI_Clients" {
  display_name = "VDI_Clients"
  description  = "Created from Terraform VDI_Clients"
  domain       = "cgw"
}

// creating Group for AppVolumes Manager:
resource "nsxt_policy_group" "AppVol_MGMT" {
  display_name = "AppVol_MGMT"
  description  = "Created from Terraform AppVol_MGMT"
  domain       = "cgw"
}

// creating Group for Event_Database:
resource "nsxt_policy_group" "Event_Database" {
  display_name = "Event_Database"
  description  = "Created from Terraform AppVol_MGMT"
  domain       = "cgw"
}

// creating Group for Admin_VMs:
resource "nsxt_policy_group" "Admin_VMs" {
  display_name = "Admin_VMs"
  description  = "Created from Terraform Admin_VMs"
  domain       = "cgw"
}

// creating Group for vROPS:
resource "nsxt_policy_group" "vROPS" {
  display_name = "vROPS"
  description  = "Created from Terraform vROPS"
  domain       = "cgw"
}

// creating Group for RFC_1918:
resource "nsxt_policy_group" "RFC_1918" {
  display_name = "RFC_1918"
  description  = "Created from Terraform RFC_1918"
  domain       = "cgw"

    criteria {
    ipaddress_expression {
      ip_addresses = ["192.168.0.0/16", "172.16.0.0/16", "10.0.0.0/8"]
    }
  }
}

###################### creating CGW Security Rules ######################

###################### creating Rules for Unified Access Gateway external ######################

resource "nsxt_policy_security_policy" "UAG_external" {
  domain       = "cgw"
  display_name = "UAG_external"
  description  = "Terraform UAG_external Rule"
  category     = "Environment"

  rule {
    display_name       = "UAG_external_Inbound"
    source_groups      = []
    destination_groups = ["${nsxt_policy_group.UAG_external.path}"]
    action             = "ALLOW"
    services           = ["${nsxt_policy_service.Blast_TCP443.path}", "${nsxt_policy_service.Blast_TCP8443.path}", "${nsxt_policy_service.Blast_UDP443.path}", "${nsxt_policy_service.PCoIP_TCP4172.path}", "${nsxt_policy_service.PCoIP_UDP4172.path}", "${nsxt_policy_service.Blast_TCP9443.path}"]
    logged             = true
    }
   rule {
      display_name       = "UAG_external_Outbound"
      source_groups      = ["${nsxt_policy_group.UAG_external.path}"]
      destination_groups = ["${nsxt_policy_group.VDI_Clients.path}"]
      action             = "ALLOW"
      services           = ["${nsxt_policy_service.Blast_TCP22443.path}", "${nsxt_policy_service.Blast_TCP8443.path}"]
      logged             = true
    }
}

###################### creating Rules for Unified Access Gateway internal ######################

resource "nsxt_policy_security_policy" "UAG_internal" {
  domain       = "cgw"
  display_name = "UAG_internal"
  description  = "Terraform UAG_internal Rule"
  category     = "Environment"

  rule {
    display_name       = "UAG_internal_Inbound"
    source_groups      = ["${nsxt_policy_group.RFC_1918.path}"]
    destination_groups = ["${nsxt_policy_group.UAG_external.path}"]
    action             = "ALLOW"
    services           = ["${nsxt_policy_service.Blast_TCP443.path}", "${nsxt_policy_service.Blast_TCP8443.path}", "${nsxt_policy_service.Blast_UDP443.path}", "${nsxt_policy_service.PCoIP_TCP4172.path}", "${nsxt_policy_service.PCoIP_UDP4172.path}", "${nsxt_policy_service.Blast_TCP9443.path}"]
    logged             = true
  }
  rule {
    display_name       = "UAG_internal_Outbound"
    source_groups      = ["${nsxt_policy_group.UAG_internal.path}"]
    destination_groups = ["${nsxt_policy_group.VDI_Clients.path}"]
    action             = "ALLOW"
    services           = ["${nsxt_policy_service.Blast_TCP22443.path}", "${nsxt_policy_service.Blast_TCP8443.path}"]
    logged             = true
  }
}
