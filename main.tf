provider "camilocot" {
  api_user     = "user"
  api_password = "password"
  api_url      = "https://terraform.free.beeceptor.com"
}

resource "camilocot_dns" "dns-server1" {
  address  = "2.2.3.1"
  hostname = "test.com"
}

