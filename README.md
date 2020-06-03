# Terraform automated whitelist firewall ruleset for Horizon and VMC on AWS

This is the repository for the automated Horizon Whitelist Firewall Ruleset, which can be used with
Terraform.

In this script I using VMC on AWS and the Terraform Provider for NSX-T.

# About Rules

After you applyed the Ruleset. Some Rules depending on your Environment. Following you can find some explainations.

If you using internal Unified Access Gateways, you are able to delete Section: "Internal_Client_Connection"
It is only necessary if you haven't deployed internal Unified access Gateways. 

# Terraform NSX-T Provider
The newest version from Terraform NSX-T Provider can always be found here: https://github.com/terraform-providers/terraform-provider-nsxt/

Currently the NSX-T Provider for VMC on AWS is not official supported.

For general information about Terraform, visit the [official
website][tf-website] and the [GitHub project page][tf-github].

[tf-website]: https://terraform.io/
[tf-github]: https://github.com/hashicorp/terraform

To get familiar with VMC and VMC NSX-T I highly recommend to take a look on the Blog posts of Nicolas Vibert
https://nicovibert.com/2020/02/04/terraform-provider-for-nsx-t-policy-and-vmware-cloud-on-aws/

Everything I leared, I learned from him :)! @Nico Thank you! You are awesome!

# Support

if you have any problems with the script, you always can reach out to me and I will try support and help you as soon as possible!
