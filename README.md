# Terraform automated whitelist firewall ruleset for Horizon with VMC on AWS

Explanation can be found here: https://securefever.com/blog/terraform-blueprint-for-a-horizon7-ruleset-with-vmc-on-aws

In this repository I created a automated Horizon Whitelist Firewall Ruleset, which can be used with
Terraform and VMC on AWS.

To get familiar with VMC and VMC NSX-T I highly recommend to take a look on the Blog posts from Nicolas Vibert:
https://nicovibert.com/2020/02/04/terraform-provider-for-nsx-t-policy-and-vmware-cloud-on-aws/

Everything I did, I learned from him :)! @Nico Thank you! You are awesome!

# About the main.tf

This Terraform code only apply Horizon related Services, groups and Distributed Firewall Sections and Rules. Ports and Rules for DNS, NTP etc. are missing.
This code should support you to build a whitelist Firewall Horizon Plattform or get a idea how you can secure your Horizon Environment.

I will keep on working on this script!

# Support

if you have any problems with the script, you always can reach out to me and I will try to support and help you as soon as possible!
