data "google_compute_image" "cos_image" {
  family  = "cos-69-lts"
  project = "gce-uefi-images"
}


variable "var_project" {
  default = "jfrog-deployment"
}
variable "var_env" {
  default = "dev"
}
variable "var_company" {
  default = "com"
}
variable "uc1_private_subnet" {
  default = "10.26.1.0/24"
}
variable "uc1_public_subnet" {
  default = "10.26.2.0/24"
}
variable "ue1_private_subnet" {
  default = "10.26.3.0/24"
}
variable "ue1_public_subnet" {
  default = "10.26.4.0/24"
}


#----------Setup VPC from here----------------#
resource "google_compute_network" "jfrog-vpc" {
  name = "${var.var_company}-${var.var_env}-vpc"
  auto_create_subnetworks = false
  routing_mode = "GLOBAL"
}

resource "google_compute_firewall" "allow-internal" {
  name = "${var.var_company}-fw-allow-internal"
  network = "${google_compute_network.jfrog-vpc.name}"
  allow {
    protocol = "icmp"
  }
  allow {
    protocol = "tcp"
    ports = ["0-65535"]
  }
  allow {
    protocol = "udp"
    ports = ["0-65535"]
  }
  source_ranges = [
    "${var.uc1_private_subnet}",
    "${var.uc1_public_subnet}",
    "${var.ue1_public_subnet}",
    "${var.ue1_private_subnet}"
  ]
}

resource "google_compute_firewall" "allow-http" {
  name = "${var.var_company}-fw-allow-http"
  network = "${google_compute_network.jfrog-vpc.name}"
  allow {
    protocol = "tcp"
    ports = ["80","8080","8081","7000","7002","7003","8000"]
  }
  target_tags = ["http"]
}

resource "google_compute_firewall" "allow-bastion" {
  name = "${var.var_company}-fw-allow-bastion"
  network = "${google_compute_network.jfrog-vpc.name}"
  allow {
    protocol = "tcp"
    ports = ["22"]
  }
  target_tags = ["ssh"]
}

#----------Setup Network from here----------------#

resource "google_compute_subnetwork" "us_east_public_subnet" {
  ip_cidr_range = "${var.ue1_public_subnet}"
  name = "${var.var_company}-${var.var_env}-us-east1-a-pub-net"
  network = "${google_compute_network.jfrog-vpc.self_link}"
  region = "us-east1"
}

resource "google_compute_subnetwork" "us_central_public_subnet" {
  ip_cidr_range = "${var.uc1_public_subnet}"
  name = "${var.var_company}-${var.var_env}-us-central1-a-pub-net"
  network = "${google_compute_network.jfrog-vpc.self_link}"
  region = "us-central1"
}

resource "google_compute_subnetwork" "us_east_private_subnet" {
  ip_cidr_range = "${var.ue1_private_subnet}"
  name = "${var.var_company}-${var.var_env}-us-east1-a-priv-net"
  network = "${google_compute_network.jfrog-vpc.self_link}"
  region = "us-east1"
}

resource "google_compute_subnetwork" "us_central_private_subnet" {
  ip_cidr_range = "${var.uc1_private_subnet}"
  name = "${var.var_company}-${var.var_env}-us-central1-a-priv-net"
  network = "${google_compute_network.jfrog-vpc.self_link}"
  region = "us-central1"
}


#----------Setup Instances from here----------------#

resource "google_compute_address" "ip_address" {
  name = "global-appserver-address"
}

resource "google_compute_address" "xray_ip_address" {
  name = "xray-appserver-address"
}

resource "google_compute_instance" "jfrog-default" {
  machine_type = "n1-standard-2"
  zone = "us-central1-a"
  name = "${var.var_company}-${var.var_env}-us-central-jfrog-instance"
  tags = ["ssh","http"]
  "boot_disk" {
    initialize_params {
      image = "${data.google_compute_image.cos_image.self_link}"
    }
  }

  metadata_startup_script = "${file("script.sh")}"

  "network_interface" {
    subnetwork = "${google_compute_subnetwork.us_central_public_subnet.name}"
    access_config {
      nat_ip = "${google_compute_address.ip_address.address}"
    }
  }
}

resource "google_compute_instance" "xray-default" {
  machine_type = "n1-standard-4"
  zone = "us-central1-a"
  name = "${var.var_company}-${var.var_env}-us-central-xray-instance"
  tags = ["ssh","http"]
  "boot_disk" {
    initialize_params {
      image = "${data.google_compute_image.cos_image.self_link}"
    }
  }


  "network_interface" {
    subnetwork = "${google_compute_subnetwork.us_central_public_subnet.name}"
    access_config {
      nat_ip = "${google_compute_address.xray_ip_address.address}"
    }
  }
}