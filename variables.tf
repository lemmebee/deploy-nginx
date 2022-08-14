variable "region" {
  type        = string
  default     = "eu-west-3"
  description = "aws region"

}

variable "namespace" {
  type        = string
  default     = "nginx-ns"
  description = "kubernetes namespace"
}

variable "tag" {
  type        = string
  default     = "terraform-eks-cluster"
  description = "tag for aws resources"
}

variable "kubernetes_version" {
  type        = string
  default     = "1.22"
  description = "kubernetes version for eks cluster"
}

variable "nginx_name" {
  type        = string
  default     = "nginx"
  description = "nginx helm release name"
}

variable "nginx_repository" {
  type        = string
  default     = "https://charts.bitnami.com/bitnami"
  description = "nginx helm repository"
}

variable "nginx_chart" {
  type        = string
  default     = "nginx"
  description = "nginx helm chart"
}

variable "nginx_version" {
  type        = string
  default     = "4.1.0"
  description = "nginx chart version"
}

variable "replica" {
  type        = string
  default     = "3"
  description = "number of nginx replica"
}
