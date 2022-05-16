package test

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

var namespace string = "nginx-ns"
var kubeOptions *k8s.KubectlOptions = k8s.NewKubectlOptions("", "", namespace)

// Setup a TLS configuration to submit with the helper, a blank struct is acceptable
var tlsConfig tls.Config = tls.Config{}

//
func newClientset(cluster *eks.Cluster) (*kubernetes.Clientset, error) {
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return nil, err
	}
	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(cluster.Name),
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		return nil, err
	}
	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(cluster.CertificateAuthority.Data))
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(
		&rest.Config{
			Host:        aws.StringValue(cluster.Endpoint),
			BearerToken: tok.Token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: ca,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

//

// Returns pod name
func awaitPods(t *testing.T, kubeOptions *k8s.KubectlOptions, filter string) string {
	var podName string

	pods := k8s.ListPods(t, kubeOptions, v1.ListOptions{FieldSelector: "status.phase=Running"})

	for _, pod := range pods {

		// Await all deployed pods to be available and ready
		k8s.WaitUntilPodAvailable(t, kubeOptions, pod.Name, 60, 1*time.Second)

		if strings.Contains(pod.Name, filter) {
			podName := pod.Name
			return podName
		}
	}

	return podName
}

// Returns service name
func awaitServices(t *testing.T, kubeOptions *k8s.KubectlOptions, filter string) string {
	var serviceName string

	services := k8s.ListServices(t, kubeOptions, v1.ListOptions{FieldSelector: "metadata.namespace=nginx-ns"})

	for _, service := range services {

		// Await all deployed services to be available and ready
		k8s.WaitUntilServiceAvailable(t, kubeOptions, service.Name, 60, 1*time.Second)

		if strings.Contains(service.Name, filter) {
			serviceName := service.Name
			return serviceName
		}
	}
	return serviceName
}

func verifyNginx(statusCode int, body string) bool {
	if statusCode != 200 {
		return false
	}

	return strings.Contains(body, "Welcome to nginx")
}

func TestingNginxDeployment(t *testing.T) {

	var filter string = "nginx"

	nginxPod := awaitPods(t, kubeOptions, filter)

	awaitServices(t, kubeOptions, filter)

	tunnel := k8s.NewTunnel(kubeOptions, k8s.ResourceTypePod, nginxPod, 0, 80)
	defer tunnel.Close()
	tunnel.ForwardPort(t)

	// Try to access the nginx service on the local port, retrying until we get a good response for up to 5 minutes
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		fmt.Sprintf("http://%s", tunnel.Endpoint()),
		&tlsConfig,
		60,
		5*time.Second,
		verifyNginx,
	)

}

func TestInfrastructure(t *testing.T) {

	terraformOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOpts)

	terraform.InitAndApply(t, terraformOpts)

	//
	clusterName := terraform.Output(t, terraformOpts, "cluster_name")
	region := terraform.Output(t, terraformOpts, "region")

	fmt.Println("CEREBRAL:::TERRATEST_OUTPUT::clusterName", clusterName)
	fmt.Println("CEREBRAL:::TERRATEST_OUTPUT::region", region)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	eksSvc := eks.New(sess)

	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}

	result, err := eksSvc.DescribeCluster(input)
	assert.NoError(t, err)

	fmt.Println("CEREBRAL:::TERRATEST_OUTPUT::result", result)

	clientset, err := newClientset(result.Cluster)
	assert.NoError(t, err)
	fmt.Println("CEREBRAL:::TERRATEST_OUTPUT::clientset", clientset)
	//

	// TestingNginxDeployment(t)

}
