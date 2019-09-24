module "instance" {
  source = "../../instance"
  environment = "dev"
  machine_type = "n1-standard-1"
  instance_tags = ["mongo-dev"]
  image = "packer-1569120732"
}