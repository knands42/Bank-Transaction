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


module "new-eks" {
  source            = "./modules/eks"
  prefix            = var.prefix
  vpc_id            = module.my-vpc.vpc_id
  cluster_name      = var.cluster_name
  retention_days    = var.retention_days
  node_desired_size = var.node_desired_size
  node_max_size     = var.node_max_size
  node_min_size     = var.node_min_size
}
