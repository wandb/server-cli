locals {
  wandb_license = "{{.License}}"
  wandb_image   = "{{.Docker.Image}}"
  wandb_version = "{{.Docker.Version}}"
  wandb_apikey  = "{{.APIKey}}"
  deployment_id = "{{.DeploymentID}}"

  database_engine_version = "{{.Database.Version}}"
  database_instance_class = "{{.Database.Size}}"

  deletion_protection = {{.DeletionProtection}}

  namespace   = "{{.Namespace}}"
  domain_name = "{{.DomainName}}"

  zone_id = "{{.AWS.ZoneID}}"

  region              = "{{.Region}}"
  acm_certificate_arn = "{{.AWS.ACMCertificateARN}}"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.60"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.6"
    }
  }
  
  # Use deploy as your backend for storing terraform state.
  backend "http" {
    address  = "https://deploy.wandb.ai/api/terraform/state/${local.deployment_id}"
    username = "apikey"
    password = local.wandb_apikey
  }
}

provider "aws" {
  region = local.region

  default_tags {
    tags = {
      GithubRepo = "terraform-aws-wandb"
      GithubOrg  = "wandb"
      App        = "wandb"
    }
  }
}

# Deploy AWS Infrasturcture
module "wandb_infra" {
  source  = "wandb/wandb/aws"
  version = "{{.Modules.AWS.Version}}"

  namespace          = local.namespace
  public_access      = local.public_access
  external_dns       = local.external_dns
  use_internal_queue = local.use_internal_queue

  acm_certificate_arn = local.acm_certificate_arn

  allowed_inbound_cidr      = ["0.0.0.0/0"]
  allowed_inbound_ipv6_cidr = ["::/0"]

  kubernetes_public_access       = true
  kubernetes_public_access_cidrs = ["0.0.0.0/0"]

  database_engine_version = local.database_engine_version
  database_instance_class = local.database_instance_class

  deletion_protection = local.deletion_protection

  domain_name = local.domain_name
  zone_id     = local.zone_id

  create_elasticache = true
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.app_cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.app_cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.app_cluster.token
}

# Deploy Wandb local to k8 cluster
module "wandb_app" {
  source  = "wandb/wandb/kubernetes"
  version = "{{.Modules.Kubernetes.Version}}"

  wandb_image   = local.wandb_image
  wandb_version = local.wandb_version

  license                    = local.wandb_license
  host                       = module.wandb_infra.url
  bucket                     = "s3://${module.wandb_infra.bucket_name}"
  bucket_queue               = local.use_internal_queue ? "internal://" : "sqs://${module.wandb_infra.bucket_queue_name}"
  bucket_aws_region          = module.wandb_infra.bucket_region
  bucket_kms_key_arn         = module.wandb_infra.kms_key_arn
  database_connection_string = "mysql://${module.wandb_infra.database_connection_string}"
  redis_connection_string    = "redis://${module.wandb_infra.elasticache_connection_string}?tls=true&ttlInSeconds=604800"

  service_port = module.wandb_infra.internal_app_port

  # If we dont wait, tf will start trying to deploy while the node group is
  # still spinning up
  depends_on = [module.wandb_infra]
}

output "url" {
  value = module.wandb_infra.url
}
