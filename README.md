# Terraform automated whitelist firewall ruleset for Horizon with VMC on AWS

This is the repository for the automated Horizon Whitelist Firewall Ruleset, which can be used with
Terraform.

In this script I using VMC on AWS and the Terraform Provider for NSX-T.

# About Rules

This Terraform script only apply VMware related connections! Ports and Rules for DNS, NTP, Syslog etc. are missing.

This script should only support you to build a whitelist Firewall Horizon Plattform.

# Terraform NSX-T Provider
The newest version from Terraform NSX-T Provider can always be found here: https://github.com/terraform-providers/terraform-provider-nsxt/

Currently the NSX-T Provider for VMC on AWS is not official supported.

For general information about Terraform, visit the [official
website][tf-website] and the [GitHub project page][tf-github].

[tf-website]: https://terraform.io/
[tf-github]: https://github.com/hashicorp/terraform

To get familiar with VMC and VMC NSX-T I highly recommend to take a look on the Blog posts of Nicolas Vibert:
https://nicovibert.com/2020/02/04/terraform-provider-for-nsx-t-policy-and-vmware-cloud-on-aws/

Everything I did, I learned from him :)! @Nico Thank you! You are awesome!

# Support

if you have any problems with the script, you always can reach out to me and I will try support and help you as soon as possible!
