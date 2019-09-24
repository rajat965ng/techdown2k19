output "instance_ip" {
  value = "${google_compute_instance.mongo-server.network_interface}"
}