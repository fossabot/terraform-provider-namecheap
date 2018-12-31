provider "namecheap" {
  token     = "${var.namecheap_token}"
  user      = "${var.namecheap_user}"
  client_ip = "${local.workstation_external_ip}" # Optional, but useful
}

data "namecheap_domain" "example_com" {
  name = "example.com"
}

resource "namecheap_nameservers" "example_nameservers" {
  domain      = "${namecheap_domain.example_com.self_link}"
  nameservers = ["ns1.myprovider.com", "ns2.myprovider.com"]
}
