terraform {
  backend "gcs" {
    prefix = "terraform/state"
  }
  required_providers {
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "3.73.0"
    }
  }
}

variable "project" {
  type        = string
  description = "the GCP project where the cluster will be created"
}

module "zonal-cluster" {
  source     = "../../../modules/region"
  region     = "%%REGION%%"
  zone       = "%%ZONE%%"
  project    = var.project
}