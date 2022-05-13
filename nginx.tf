resource "helm_release" "nginx" {
  # Kubernetes 1.19+
  # Helm 3.2.0+
  create_namespace = true
  chart            = var.nginx_chart
  name             = var.nginx_name
  namespace        = var.namespace
  repository       = var.nginx_repository

  set {
    name  = "replicaCount"
    value = "3"
  }

  depends_on = [
    aws_eks_cluster.this
  ]
}
