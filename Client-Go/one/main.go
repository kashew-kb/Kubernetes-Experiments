package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig := flag.String("KubeConfig", "/Users/deadman/.kube/arthashastra.yml", "Location to your KubeConnfig File")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		// Handle Error
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// Handle Error
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Handle Error
	}

	for _, pod := range pods.Items {
		fmt.Printf(pod.Name)
	}

}
