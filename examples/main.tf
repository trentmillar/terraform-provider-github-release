terraform {
  required_providers {
    github-release = {
      versions = ["0.1"]
      source = "github.com/trentmillar/github-release"
    }
  }
}

provider "github-release" {}

module "rel" {
  source = "./release"
}
