package main

import (
	`fmt`
	`time`

	`k8s.io/apimachinery/pkg/labels`
	`k8s.io/client-go/informers`
	`k8s.io/client-go/kubernetes`
	`k8s.io/client-go/tools/clientcmd`
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/apple/.kube/config")
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 初始化 SharedInformerFactor
	sharedInformerFactor := informers.NewSharedInformerFactory(clientSet, 0)

	// 查询 Pod
	podInformer := sharedInformerFactor.Core().V1().Pods()

	// 生成 Indexer ，便宜查询数据
	indexer := podInformer.Lister()

	// start informer
	sharedInformerFactor.Start(nil)

	// 等待数据同步完成
	sharedInformerFactor.WaitForCacheSync(nil)

	for _ = range time.Tick(time.Second * 2) {
		// 利用Indexer查询数据
		podList, err := indexer.List(labels.Everything())
		if err != nil {
			panic(err)
		}

		// 循环读数据
		for _, items := range podList {
			fmt.Printf("name:%v", items.Name)
		}
	}

}
