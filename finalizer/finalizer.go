package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
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

	pod, err := clientset.CoreV1().Pods("redis-operators").Get(context.TODO(), "redis-standalone-0", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("无法获取 Pod: %v", err)
	}

	fmt.Println(pod.GetObjectMeta().GetFinalizers())

	if pod.GetDeletionTimestamp() != nil {
		fmt.Println("Pod 已经被删除", pod.GetDeletionTimestamp())
		pod.ObjectMeta.Finalizers = []string{}
		_, err := clientset.CoreV1().Pods("redis-operators").Update(context.TODO(), pod, metav1.UpdateOptions{})
		if err != nil {
			fmt.Println("finalizer 删除失败", err)
		}
	} else {
		fmt.Printf("Pod 未被删除")
	}
}
