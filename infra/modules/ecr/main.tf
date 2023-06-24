resource "aws_ecr_repository" "my_ecr_repo" {
  name = "${var.prefix}-${var.ecr_repository}"

  image_scanning_configuration {
    scan_on_push = false
  }
}