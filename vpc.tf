provider "aws" {
  region = var.region
}

data "aws_region" "current" {
}

data "aws_availability_zones" "available" {
}

resource "aws_vpc" "this" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    "Name" = "terraform-eks-cluster"
  }
}

resource "aws_subnet" "this" {
  count = 2

  availability_zone       = data.aws_availability_zones.available.names[count.index]
  cidr_block              = "10.0.${count.index}.0/24"
  vpc_id                  = aws_vpc.this.id
  map_public_ip_on_launch = true

  tags = {
    "Name" = "terraform-eks-cluster"
  }
}

resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id

  tags = {
    Name = "terraform-eks-cluster"
  }
}

resource "aws_route_table" "this" {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }
}

resource "aws_route_table_association" "this" {
  count = 2

  subnet_id      = aws_subnet.this[count.index].id
  route_table_id = aws_route_table.this.id
}
