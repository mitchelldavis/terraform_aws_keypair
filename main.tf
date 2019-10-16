variable "key_name" {
  type  = "string"
}

variable "rsa_bits" {
  type    = "string"
  default = "4096"
}

# AWS Key Pairs only support RSA Key Pairs.
resource "tls_private_key" "_" {
  algorithm   = "RSA"
  rsa_bits    = "${var.rsa_bits}"
}

resource "aws_key_pair" "_" {
  key_name   = "${var.key_name}"
  public_key = "${tls_private_key._.public_key_openssh}"
}

# See https://www.terraform.io/docs/providers/tls/r/private_key.html
output "private_key_pem" {
  value = "${tls_private_key._.private_key_pem}"
}

output "public_key_pem" {
  value = "${tls_private_key._.public_key_pem}"
}

output "public_key_openssh" {
  value = "${tls_private_key._.public_key_openssh}"
}

output "public_key_fingerprint_md5" {
  value = "${tls_private_key._.public_key_fingerprint_md5}"
}

output "fingerprint" {
  value = "${aws_key_pair._.fingerprint}"
}
