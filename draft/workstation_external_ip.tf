#
# Retrieve the external IP of this workstation
#
# This configuration is not required and is only provided as an example to
# easily fetch the external IP of your local workstation. This can be used as
# the client_ip value for a Namecheap provider.
#

provider "http" {}

data "http" "workstation_external_ip" {
  url = "http://ipv4.icanhazip.com"
}

locals {
  # Override with variable or a hardcoded value if necessary.
  workstation_external_ip = "${data.http.workstation_external_ip.body}"
}
