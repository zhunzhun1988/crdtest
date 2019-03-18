package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	patrickclient "github.com/zhunzhun1988/crdtest/pkg/client/clientset/versioned"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	exampleClient, err := patrickclient.NewForConfig(cfg)
	if err != nil {
		glog.Errorf("Error building example clientset: %v", err)
		os.Exit(-1)
	}

	list, err := exampleClient.Patrick().Clusters("").List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("Error listing all Clusters: %v", err)
		os.Exit(-1)
	}

	for _, c := range list.Items {
		fmt.Printf("cluster %s with user %q\n", c.Name, c.Spec.Apiserver)
	}
}
