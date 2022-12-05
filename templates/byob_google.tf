locals {
  project_id    = "{{.Google.ProjectID}}"
  region        = "{{.Region}}"
  bucket_prefix = "{{.BucketPrefix}}"

  service_account_email = "deploy@wandb-production.iam.gserviceaccount.com"
  deletion_protection   = false
}

provider "google" {
  project = local.project_id
  region  = local.region
}

module "project_factory_project_services" {
  source                      = "terraform-google-modules/project-factory/google//modules/project_services"
  version                     = "~> 13.0"
  project_id                  = null
  disable_dependent_services  = false
  disable_services_on_destroy = false
  activate_apis = [
    "pubsub.googleapis.com",  // File Storage
    "storage.googleapis.com", // Cloud Storage
  ]
}

module "resources" {
  source  = "wandb/wandb/google//modules/storage/bucket"
  version = "1.12.2"

  namespace = local.bucket_prefix

  bucket_location     = local.bucket_location
  deletion_protection = local.deletion_protection
  create_queue        = false

  service_account = { "email" : local.service_account_email }

  depends_on = [module.project_factory_project_services]
}

resource "google_storage_bucket_iam_member" "admin" {
  bucket = module.resources.bucket_name
  member = "serviceAccount:${local.service_account_email}"
  role   = "roles/storage.admin"
}

output "bucket_name" {
  value = module.resources.bucket_name
}
