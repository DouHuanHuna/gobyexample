package main

import (
	"context"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := "/home/douzengrui/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("无法加载 kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("无法创建 Kubernetes 客户端: %v", err)
	}

	podName := "example-pod"
	namespace := "default"
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})

	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s 不存在\n", podName)
		} else {
			fmt.Printf("获取 Pod 时发生错误: %v\n", err)
		}
	} else {
		fmt.Printf("Pod %s 已找到，状态为: %v\n", podName, pod.Status.Phase)
	}
}
