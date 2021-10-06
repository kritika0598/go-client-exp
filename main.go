package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/kritika/.kube/config", "location to your kube config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//fmt.Println(config)
	if err != nil {
		//handle error
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle error
	}
	pods, err := clientset.CoreV1().Pods("sms-test").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		//handle error
	}
	fmt.Println("PODS")

	for _, pod := range pods.Items {
		fmt.Println("%s", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("sms-test").List(context.Background(), metav1.ListOptions{})
	fmt.Println("DEPLOYMENTS")
	for _, deployment := range deployments.Items {
		fmt.Println("%s", deployment.Name)
	}

}
