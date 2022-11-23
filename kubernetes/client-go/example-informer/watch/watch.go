package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/apple/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	// instance clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// call monitor function, return a interface type
	watch, err := clientSet.CoreV1().Pods("default").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("start...")
	// need to monitor return of data
	for {
		select {
		case e, _ := <-watch.ResultChan():
			// e.Type changed event type(e.g ADDED...)
			// e.Object after change of data
			fmt.Println(e.Type, e.Object)
		}
	}
}
