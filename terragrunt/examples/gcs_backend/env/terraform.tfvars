terragrunt = {
  remote_state {
    backend = "gcs"

    config {
      bucket = "test-terragrunt-cloudbuild" # Replace this with your own bucket
      prefix = "${path_relative_to_include()}"
    }
  }

  terraform {
    extra_arguments "-var-file" {
      commands = ["${get_terraform_commands_that_need_vars()}"]

      optional_var_files = [
        "${get_tfvars_dir()}/${find_in_parent_folders("account.tfvars", "ignore")}",
      ]
    }
  }
}
