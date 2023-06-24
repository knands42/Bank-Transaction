resource "random_password" "random_password" {
  length           = 16
  special          = true
  override_special = "_%@"
}

resource "aws_db_instance" "postgres_instance" {
 identifier             = "${var.prefix}-postgres"
  allocated_storage      = 10
  engine                 = "postgres"
  engine_version         = "12.11"
  instance_class         = "db.t2.micro"
  db_name                = var.database_name
  username               = var.database_username
  password               = random_password.random_password.result
  parameter_group_name   = "default.postgres12"
}

