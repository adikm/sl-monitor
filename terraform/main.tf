data "google_client_openid_userinfo" "me" {}

terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.64.0"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "3.1.0"
    }
  }
  required_version = ">= 0.15.0"
}
