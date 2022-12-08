package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件，生成 config 对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}
	// 2、实例化 ClientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3、初始化shared informer factory以及pod informer
	// 注：一个应用中往往会在多个地方对同一种资源对象都有informer的需求，所以就有了共享informer，
	// 即SharedInformerFactory
	factory := informers.NewSharedInformerFactory(clientSet, time.Second*30)
	podInformer := factory.Core().V1().Pods()
	informer := podInformer.Informer()

	// 注册informer的自定义ResourceEventHandler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) {},
		UpdateFunc: func(oldObj, newObj interface{}) {},
		DeleteFunc: func(obj interface{}) {},
	})

	// 启动shared informer factory，开始informer的list & watch操作
	stopper := make(chan struct{})
	go factory.Start(stopper)

	// 等待informer从kube-apiserver同步资源完成，即informer的list操作获取的对象都存入到informer中的indexer本地缓存中
	// 或者调用factory.WaitForCacheSync(stopper)
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	// 创建lister
	podLister := podInformer.Lister()
	// 从informer中的indexer本地缓存中获取对象
	podList, err := podLister.List(labels.Everything())
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range podList {
		fmt.Printf("namespace:%v name:%v\n", item.Namespace, item.Name)
	}
}
