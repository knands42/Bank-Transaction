module "my-vpc" {
  source         = "./modules/vpc"
  prefix         = var.prefix
  vpc_cidr_block = var.vpc_cidr_block
}

module "my-ecr" {
  source         = "./modules/ecr"
  prefix         = var.prefix
  ecr_repository = var.ecr_repository
}

module "my-rds-postgres" {
  source            = "./modules/rds"
  prefix            = var.prefix
  database_name     = var.database_name
  database_username = var.database_username
}

module "my-secret-manager" {
  source            = "./modules/secret_manager"
  database_name     = var.database_name
  database_username = var.database_username
  database_port     = var.database_port
  database_password = module.my-rds-postgres.database_password
  database_host     = module.my-rds-postgres.database_host
}

# module "my-eks" {
#   source            = "./modules/eks"
#   prefix            = var.prefix
#   vpc_id            = module.my-vpc.vpc_id
#   cluster_name      = var.cluster_name
#   retention_days    = var.retention_days
#   node_desired_size = var.node_desired_size
#   node_max_size     = var.node_max_size
#   node_min_size     = var.node_min_size
# }
