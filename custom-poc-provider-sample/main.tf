terraform {
  required_providers {
    example = {
      source = "local/example"
      version = "0.1.0"
    }
  }
}

provider "example" {}

resource "example_server" "my-server" {}