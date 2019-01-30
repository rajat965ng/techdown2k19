
variable "project_name" {
  default = "myproject"
}


variable "vpc_cidr" {
  default = "10.0.0.0/16"
}

variable "subnet_cidr" {
  default = "10.0.1.0/16"
}

variable "RDS_CIDR" {
  default = "0.0.0.0/0"
}
