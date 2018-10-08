provider "camilocot" {
  api_user     = "user"
  api_password = "password"
  api_url      = "https://test.com"
}

resource "camilocot_dns" "dns-server1" {
  address  = "2.2.3.4"
  hostname = "test.com"
}

resource "camilocot_dns" "dns-server2" {
  address  = "3.2.3.4"
  hostname = "test.com"
}

resource "camilocot_dns" "dns-server4" {
  address  = "1.2.3.4"
  hostname = "test.com"
}
