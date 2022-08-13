variable "region" {
  type        = string
  default     = "eu-west-3"
  description = "aws region"

}

variable "kubernetes_version" {
  type        = string
  default     = "1.22"
  description = "kubernetes version for eks cluster"
}
