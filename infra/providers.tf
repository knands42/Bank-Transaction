terraform {
  required_version = ">=0.13.1"
  required_providers {
    aws   = ">=3.54.0"
    local = ">=2.1.0"
  }
  # backend "s3" {
  #   bucket = "terraform-locks-caiofernandes00"
  #   key    = "global/s3/terraform.tfstate"
  #   region = "us-east-1"

  #   dynamodb_table = "terraform-locks-caiofernandes00"
  #   encrypt        = true
  # }
}

provider "aws" {
  region = "us-east-1"
}