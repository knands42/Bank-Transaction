output "database_password" {
  value = random_password.random_password.result
}

output "database_host" {
  value = aws_db_instance.postgres_instance.address
}