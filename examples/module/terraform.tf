variable "region" {
  type    = "string"
}

locals {
  key_name  = "${uuid()}"
}

provider "aws" {
  version = "2.32.0"
  region  = "${var.region}"
}

provider "tls" {
  version = "2.1.1"
}

module "key_pair" {
  source  = "../../"
  key_name  = "${local.key_name}"
}

output "key_name" {
  value = "${local.key_name}"
}

output "fingerprint" {
  value = "${module.key_pair.fingerprint}"
}


