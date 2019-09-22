provider "google" {
  credentials = "${file("service-account.json")}"
  project     = "project-name"
  region      = "us-central1"
}

