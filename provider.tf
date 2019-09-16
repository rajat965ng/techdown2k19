provider "google" {
  credentials = "${file("account/lforlogging-23fd021fe1c6.json")}"
  project     = "lforlogging"
  region      = "us-central1"
}