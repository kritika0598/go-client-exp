package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/kritika/.kube/config", "location to your kube config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error %s building from flags\n\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error %s  Getting InClusterConfig\n\n", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error %s, creating client set \n\n", err.Error())
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error %s, While listenting all the pods from default namespace\n\n", err.Error())
	}
	fmt.Println("PODS")
	for _, pod := range pods.Items {
		fmt.Println("%s", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error %s, While listenting all the deployments from default namespace\n\n", err.Error())
	}
	fmt.Println("DEPLOYMENTS")
	for _, deployment := range deployments.Items {
		fmt.Println("%s", deployment.Name)
	}
}
