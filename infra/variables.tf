variable "prefix" {
  type    = string
  default = "simplebank"
}

variable "cluster_name" {
  type    = string
  default = "simplebank-cluster"
}

variable "ecr_repository" {
  type    = string
  default = "simplebank-ecr"
}

variable "retention_days" {
  type    = number
  default = 1
}

variable "node_desired_size" {
  type    = number
  default = 1
}

variable "node_max_size" {
  type    = number
  default = 1
}

variable "node_min_size" {
  type    = number
  default = 1
}

variable "vpc_cidr_block" {
  type    = string
  default = "10.0.0.0/16"
}