provider "google" {
  project = var.projectName
  region  = var.region
  zone    = var.zone
}

module "project_services" {
  source  = "terraform-google-modules/project-factory/google//modules/project_services"
  version = "14.2.0"

  project_id = var.projectName

  activate_apis = [
    "compute.googleapis.com",
    "oslogin.googleapis.com",
  ]

  disable_services_on_destroy = false
  disable_dependent_services  = false
}

# VPC

resource "google_compute_network" "vpc_network" {
  name = "slmonitor-network"
}

resource "google_compute_firewall" "allow_ssh" {
  name          = "allow-ssh"
  network       = google_compute_network.vpc_network.name
  target_tags   = ["allow-ssh"]
  source_ranges = ["0.0.0.0/0"]

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }
}

# VM

resource "google_compute_address" "static_ip" {
  name = "slmonitor-instance"
}

resource "google_compute_instance" "vm_instance" {
  name         = "slmonitor-instance"
  machine_type = "e2-micro"
  tags         = ["web", "dev", "allow-ssh"]

  metadata = {
    ssh-keys = "${data.google_client_openid_userinfo.me.email}:${tls_private_key.ssh.public_key_openssh}"
  }


  boot_disk {
    initialize_params {
      image = "ubuntu-minimal-2204-lts"
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.name
    access_config {
      nat_ip = google_compute_address.static_ip.address
    }
  }

  allow_stopping_for_update = true
}

output "public_ip" {
  description = "The IP address of the VM instance."
  value       = google_compute_address.static_ip.address
}

output "username" {
  description = "username"
  value       = data.google_client_openid_userinfo.me.email
}


# REDIS

resource "google_redis_instance" "slmonitor_cache" {
  name           = "slmonitor"
  tier           = "BASIC"
  memory_size_gb = 1
  region         = var.region
  redis_version  = "REDIS_6_X"
}
output "cache_host" {
  description = "The IP address of the cache instance."
  value       = google_redis_instance.slmonitor_cache.host
}