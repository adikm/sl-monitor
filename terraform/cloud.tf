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
  source_ranges = [var.ip_access]

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

output "vm_ip" {
  description = "The IP address of the VM instance."
  value       = google_compute_address.static_ip.address
}


# PostgreSQL

resource "google_sql_database_instance" "postgresql" {
  name                = "${var.projectName}-db1"
  deletion_protection = false
  project             = var.projectName
  region              = var.region
  database_version    = "POSTGRES_11"

  settings {
    tier              = "db-f1-micro"
    activation_policy = "ALWAYS"
    disk_autoresize   = false
    disk_size         = "10"
    disk_type         = "PD_SSD"

    location_preference {
      zone = var.zone
    }

    maintenance_window {
      day  = "7"  # sunday
      hour = "3" # 3am
    }

    backup_configuration {
      enabled    = true
      start_time = "00:00"
    }

    ip_configuration {
      ipv4_enabled = true
      authorized_networks {
        value = var.ip_access
      }
    }
  }
}

resource "google_sql_database" "postgresql_db" {
  provider = google-beta

  name     = "slmonitor"
  project  = var.projectName
  instance = google_sql_database_instance.postgresql.name
  charset  = "UTF8"
}

resource "google_sql_user" "postgresql_user" {
  name     = "slmonitor-user"
  project  = var.projectName
  instance = google_sql_database_instance.postgresql.name
  password = "slmonitor-pwd"
}

output db_instance_ip {
  description = "The IP address of the master database instance"
  value       = google_sql_database_instance.postgresql.ip_address[0]
}

# REDIS

resource "google_redis_instance" "slmonitor_cache" {
  name               = "slmonitor"
  tier               = "BASIC"
  memory_size_gb     = 1
  region             = var.region
  redis_version      = "REDIS_6_X"
  authorized_network = google_compute_network.vpc_network.name
}
output "cache_ip" {
  description = "The IP address of the cache instance."
  value       = google_redis_instance.slmonitor_cache.host
}

# CONTAINER REGISTRY

resource "google_container_registry" "registry" {
  project  = var.projectName
  location = "EU"
}

# CLOUD RUN

resource "google_project_service" "run_api" {
  service = "run.googleapis.com"

  disable_on_destroy = true
}

resource "google_cloud_run_service" "run_service" {
  name     = var.projectName
  location = var.region

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello:latest"
#        image = "gcr.io/slmonitor/slmonitor-app:latest"
        env {
          name = "DB_HOST"
          value = google_sql_database_instance.postgresql.ip_address[0].ip_address
        }
        env {
          name = "CACHE_HOST"
          value = google_redis_instance.slmonitor_cache.host
        }
        env {
          name = "TRAFFIC_API_AUTH_KEY"
          value = "0e7862ebcacf4d7a90c2a90a443bca3f"
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  depends_on = [google_project_service.run_api]
}

resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.run_service.name
  location = google_cloud_run_service.run_service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

output "cloudrun_url" {
  value = google_cloud_run_service.run_service.status[0].url
}