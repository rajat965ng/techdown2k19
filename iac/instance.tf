resource "google_compute_instance" "mongo-server" {
  name         = "mongo-server"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  tags = ["mongo"]

  boot_disk {
    initialize_params {
      image = "packer-1569120732"
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