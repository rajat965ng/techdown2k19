resource "google_compute_instance" "mongo-server" {
  name         = "mongo-${var.environment}-server"
  machine_type = "${var.machine_type}"
  zone         = "us-central1-a"

  tags = "${var.instance_tags}"

  boot_disk {
    initialize_params {
      image = "${var.image}"
    }
  }


  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }


  metadata_startup_script = "sudo service mongod start"

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}