resource "aws_secretsmanager_secret" "simple_bank_secrets" {
  name = "simplebank"
}

resource "aws_secretsmanager_secret_version" "simple_bank_secrets_version" {
  secret_id     = aws_secretsmanager_secret.simple_bank_secrets.id
  secret_string = jsonencode(merge(var.secrets_value, {
    "DB_SOURCE" : "postgresql://${var.database_username}:${var.database_password}@${var.database_host}:${var.database_port}/${var.database_name}?sslmode=${var.database_ssl_mode}"
  }))
}

