resource "aws_vpc" "myvpc" {
  enable_dns_support = true
  enable_dns_hostnames = true
  cidr_block = "${var.vpc_cidr}"
  tags {
    Name = "${var.project_name}-vpc"
  }
}

data "aws_availability_zones" "available" {}

resource "aws_subnet" "mysubnet" {
  vpc_id = "${aws_vpc.myvpc.id}"
  availability_zone = "${data.aws_availability_zones.available.names[0]}"
  cidr_block = "${var.subnet_cidr}"
  map_public_ip_on_launch = true
  tags {
    Name = "${var.project_name}-subnet"
  }
}

resource "aws_security_group" "mydev" {

  vpc_id      = "${aws_vpc.myvpc.id}"
  name = "${var.project_name}-sg-dev"

  tags{
   Name = "${var.project_name}-sg-dev"
 }

 ingress {
   from_port = 8080
   to_port = 8080
   protocol = "tcp"
   cidr_blocks = ["${var.RDS_CIDR}"]
 }


  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["${var.RDS_CIDR}"]
  }


  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

}

resource "aws_internet_gateway" "igw" {
  vpc_id            = "${aws_vpc.myvpc.id}"
  tags {
    Name = "${var.project_name}-internet-gateway"
  }
}

resource "aws_eip" "lb" {
  depends_on = ["aws_internet_gateway.igw"]
  instance = "${aws_instance.myinstance.id}"
  vpc      = true
  tags {
    Name = "${var.project_name}-elastic-ip"
  }
}

resource "aws_route_table" "myroute" {
  vpc_id            = "${aws_vpc.myvpc.id}"
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.igw.id}"
  }
  tags {
    Name = "${var.project_name}-route-table"
  }
}

resource "aws_route_table_association" "to_myroute_mysubnet" {
  subnet_id      = "${aws_subnet.mysubnet.id}"
  route_table_id = "${aws_route_table.myroute.id}"
}

resource "aws_instance" "myinstance" {
  depends_on = ["aws_security_group.mydev"]
  ami = "ami-011b3ccf1bd6db744"
  instance_type = "t2.micro"
  availability_zone = "${data.aws_availability_zones.available.names[0]}"
  key_name = "name_of_private_key"
  subnet_id = "${aws_subnet.mysubnet.id}"
  security_groups = ["${aws_security_group.mydev.id}"]
  vpc_security_group_ids = ["${aws_security_group.mydev.id}"]

  provisioner "remote-exec" {
    inline = ["sudo yum install tomcat tomcat-admin-webapps tomcat-webapps -y", "sudo service tomcat start"]
  }
  connection {
    type = "ssh"
    user = "ec2-user"
    private_key = "${file("newck.pem")}"
  }

  tags {
    Name = "${var.project_name}-ec2-instance"
  }
}

output "pubip" {
  value = "${aws_instance.myinstance.public_ip}"
}
