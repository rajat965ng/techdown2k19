variable "machine_type" {
  type = "string"
  description = "type of the machine"
}

variable "image" {
  type = "string"
  description = "image created by packer"
}

variable "instance_tags" {
  type = "list"
}

variable "environment" {
  type = "string"
}