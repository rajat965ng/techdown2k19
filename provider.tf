provider "google" {
  credentials = "${file("account/gaganapp-ff9be8369f89.json")}"
  project     = "gaganapp"
  region      = "us-central1"
}