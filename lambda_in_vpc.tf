provider "aws" {
  region = "us-east-2"
}

module "lambda_function_in_vpc" {
  source = "terraform-aws-modules/lambda/aws"

  function_name = "hello-world"
  description   = "Hello World function in VPC"
  handler       = "hello-world.handler"
  runtime       = "nodejs12.x"

  source_path = "hello-world"

  vpc_subnet_ids         = module.vpc.intra_subnets
  vpc_security_group_ids = [module.vpc.default_security_group_id]
  attach_network_policy = true
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "hello-world-vpc"
  cidr = var.cidr

  # Specify at least one of: intra_subnets, private_subnets, or public_subnets
  azs           = var.azs
  intra_subnets = var.subnets
}