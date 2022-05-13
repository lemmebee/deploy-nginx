output "cluster_endpoint" {
  value = data.aws_eks_cluster.cluster.endpoint
}

output "nginx_release_namespace" {
  description = "nginx release namespace"
  value       = helm_release.nginx.namespace
}
