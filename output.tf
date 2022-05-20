output "cluster_endpoint" {
  description = "eks cluster endpoint"
  value       = data.aws_eks_cluster.cluster.endpoint
}

output "cluster_name" {
  description = "eks cluster name"
  value       = aws_eks_cluster.this.id
}


output "region" {
  description = "aws region"
  value       = var.region
}


output "nginx_release_namespace" {
  description = "nginx release namespace"
  value       = helm_release.nginx.namespace
}
