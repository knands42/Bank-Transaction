data "aws_availability_zones" "available" {
  exclude_names = ["us-east-1e"]
}
