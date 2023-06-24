variable "secrets_value" {
  default = {
    PROFILE                = "local"
    DB_DRIVER              = "postgres"
    SERVER_ADDRESS         = "0.0.0.0:8080"
    PASSWORD_HASH_SALT     = "12345678901234567890123456789012"
    PASSWORD_HASH_TIME     = "1"
    TOKEN_TYPE             = "passeto"
    TOKEN_SYMMETRIC_KEY    = "0b7e151628aed2a6abf7158809cf4f3c"
    ACCESS_TOKEN_DURATION  = "15m"
    REFRESH_TOKEN_DURATION = "24h"
  }

  type = map(string)
}

variable "database_name" {
  type    = string
  default = "simple_bank"
}

variable "database_username" {
  type    = string
  default = "masteruser"
}

variable "database_password" {
  type = string
}

variable "database_host" {
  type = string
}

variable "database_port" {
  type    = number
  default = 5432
}
