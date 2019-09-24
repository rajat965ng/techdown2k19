module "instance" {
  source = "../../instance"
  environment = "prod"
  machine_type = "n1-standard-2"
  instance_tags = ["mongo-prod"]
  image = "packer-1569120732"
}