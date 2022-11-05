variable "project_id" {
  type = string
}

variable "zone" {
  type = string
}

variable "builder_sa" {
  type = string
}

source "googlecompute" "test-image" {
  project_id                  = var.project_id
  source_image_family         = "ubuntu-2204-lts"
  zone                        = var.zone
  image_description           = "Created with Packer from Cloudbuild"
  ssh_username                = "root"
  tags                        = ["packer"]
  impersonate_service_account = var.builder_sa
  temporary_key_pair_type     = "ed25519"
}

build {
  sources = ["sources.googlecompute.test-image"]
}
