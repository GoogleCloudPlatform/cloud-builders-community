terragrunt = {
  terraform {
    source = "../../../modules/pubsub"
  }

  include = {
    path = "${find_in_parent_folders()}"
  }
}

name = "prod-example"
