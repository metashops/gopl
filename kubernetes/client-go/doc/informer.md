### Client-go 简介

Client-go 是负责与 k8s APIServer 服务进行交互的客户端库，利用 Client-go 与 k8s APIServer 进行的交互访问，以此来对 k8s 中的各类资源对象进行管理操作，包括内置的资源对象及未来自定义的CRD资源。

### Client-Go 客户端对象

Client-Go 共提供了 4 种与 k8s APIServer 交互的客户端对象。分别是 **RESTClient、DiscoveryClient、ClientSet及DynamicClient**。

* RESTClient：最基础的客户端，主要是对HTTP请求进行封装，支持JSON和Protobuf 格式的数据。
* DiscoveryClient：发现客户端，负责发现 APIServer 支持的资源组、资源版本和资源信息等。如kubectl api-resources
* ClientSet：负责操作 k8s 内置的资源对象，如：Pod、Service等
* DynamicClient：动态客户端，可以对任意的 k8s 资源对象进行通用操作，包括CRD等（需要配置）。

![image-20221030180354498](https://tva1.sinaimg.cn/large/008vxvgGgy1h7nh0s64awj30nq0egglx.jpg)

#### （1）RESTClient

RESTClient 是所有客户端的父类，是最基础的客户端，它提供了 RESTful 对应的方法的封装，如：GET、PUT、POST、DELETE等方法。通过这些封装好的方法与 k8s APIServer 进行交互。

由于是所有客户端的父类，则其他 k8s 内置的资源及CRD 都可以操作。

**栗子：**

```go
package main

import (
	`context`
	`fmt`

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	`k8s.io/client-go/kubernetes/scheme`
	`k8s.io/client-go/rest`
	`k8s.io/client-go/tools/clientcmd`
)

func main() {
	/**
	1. k8s 配置文件
	2. 保证能通过这个配置文件连接到集群
	*/

	// 1、加载配置文件，生成 config 对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、配置 API 路径
	config.APIPath = "api" // pods, /api/v1/pods

	// 3、配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion // Group = "",Version: "v1"

	// 4、配置数据的编解码工具
	config.NegotiatedSerializer = &scheme.Codecs

	// 5、实例化 RESTClient 对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	// 6、定义接收返回值的变量(接收什么类型的数据)
	result := &corev1.PodList{}

  /**
	Get 定义方法，返回值为 *Request 结构体对象，该结构体构建访问APIServer请求用的
	依次执行了Namespace、Resource、VersionedParams构建与APIServer 交互的参数
	Do 方法通过 request 发起请求，然后通过 transformResponse解析请求返回，并绑定到对应的资源对象的结构体，这里表示corev1.PodList{}的对象
	request 先是检查了有没有可用的 client，在这里开始调用net/http包的功能
	
	
	
	 */
  
	// 7、跟APIServer交互
	err = restClient.Get(). // Get 请求
		Namespace("default"). // 指定名称空间
		Resource("pods"). // 指定需要查询的资源，资源名称
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec). // 参数及参数系列化工具
		Do(context.TODO()). // 触发请求
		Into(result) // 写入返回结果
	if err != nil {
		panic(err.Error())
	}
	for _, item := range result.Items {
		fmt.Printf("namespace:%v name:%v\n", item.Namespace, item.Name)
	}
}

```

#### （2）ClientSet

前面介绍了 RESTClient，它虽然可以操作 k8s 的所有资源对象，但是使用起来确实比较复杂，需要配置参数过于频繁。因此，为了更优雅的与 K8S APIServer 进行交互，则需要进一步封装。

前面有过介绍，ClientSet 是基于 RESTClient 的封装，同时 ClientSet 是使用预生成的 API 的对象与 APIServer 进行交互的，这样做更方便我们进行二次开发。

ClientSet 是一组资源对象客户端的集合，例如负责操作 Pods、Services 等资源的 CoreV1Client， 负责操作 Deployments、DaemonSets 等资源的 AppV1Client 等。通过这些资源对象客户端提供操作方法，即可对 k8s 内置的资源对象进行 Create、Update、Get、List、Delete 等操作。

**栗子：**

```go
package main

import (
	`context`
	`fmt`

	metav1 `k8s.io/apimachinery/pkg/apis/meta/v1`
	`k8s.io/client-go/kubernetes`
	`k8s.io/client-go/tools/clientcmd`
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

	// 3、
  // coreV1 返回 CoreV1Client 实例化
  // Pods 调用了 NewPods函数，该函数返回的 PodInterface对象，PodInterface 对象实现了 Pods 资源香港的全部方法。同时在 NewPods 里面还将 RESTClient，实例对象资源赋值给了对应的 Client 属性
	pods, err := clientSet.CoreV1().
		Pods("default").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range pods.Items {
		fmt.Printf("namespace:%v name:%v\n", item.Namespace, item.Name)
	}
}
```

#### （3）DynamicClient

DynamicClient 是一种动态客户端，通过动态指定资源组、资源版本和资源等信息，来操作任意的 k8s 资源对象的一种客户端。即不仅仅是操作 k8s 内置的资源对象，包括自定义的 CRD 资源都可以进行操作。

使用的时候，程序会将所用的版本与类型紧密耦合。而 DynamicClient 使用嵌套的 `Map[string]interface{}` 结构存储 k8s APIServer 的返回值，使用反射机制，在运行的时候，进行数据绑定，这种方式更加灵活，但是既无法获取数据类型的检查和验证。

此外，我们还在介绍了 DynamicClient 之前，还需要了解另外两个重要的知识点， Object.runtime 接口和 Unstructured 的结构体。

* Object.runtime ：k8s 中所有的资源对象，都实现这个接口，其中包含 DeepCopyObject 和 GetObjectKind 的方法，分别用于对象深拷贝和获取对象的具体资源类型。
* Unstructured：包括 `map[string]interface{} ` 类型字段，在处理无法预知结构体的数据时，将数据值存放入 interface{} 中，待运行利用反射判断，该结构体提供了大量的工具方法，便于处理非结构体化的数据。

**栗子：**

```go
package main

import (
	`context`
	`fmt`

	corev1 `k8s.io/api/core/v1`
	metav1 `k8s.io/apimachinery/pkg/apis/meta/v1`
	`k8s.io/apimachinery/pkg/runtime`
	`k8s.io/apimachinery/pkg/runtime/schema`
	`k8s.io/client-go/dynamic`
	`k8s.io/client-go/tools/clientcmd`
)

func main() {
	// 1、加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、实例化动态客户端对象
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3、配置需要调用的GVR
	gvr := schema.GroupVersionResource{
		Group:    "", // 无名资源组，不需要写
		Version:  "v1",
		Resource: "pods",
	}

	// 4、发送请求，且得到返回结果(动态客户端获取到的是非结构化结构体，需要用非结构化变量去存储)
	unStructData, err := dynamicClient.Resource(gvr).Namespace("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 5、unStructData 转换为结构化数据
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unStructData.UnstructuredContent(), podList)
	if err != nil {
		panic(err.Error())
	}

	// 6、打印
	for _, item := range podList.Items {
		fmt.Printf("namespace:%v, name:%v\n", item.Namespace, item.Name)
	}
}

```

解释：动态客户端到 k8s APIServer 请求过程

1. 首先调用 Resource
2. 调用 Namespace
3. 调用 List，首先通过RESTClient 调用 k8s APIServer 的接口返回 Pod 的数据，返回数据格式是二进制的 JSON格式。

#### （4）DiscoveryClient

前面介绍了三种客户端对象，都是针对与资源管理的。而 DiscoveryClient 则是针对于 GVR的。用于查询当前 k8s 集群支持哪些资源组、资源版本、资源信息。

**栗子：**

```go
package main

import (
	`fmt`

	`k8s.io/apimachinery/pkg/runtime/schema`
	`k8s.io/client-go/discovery`
	`k8s.io/client-go/tools/clientcmd`
)

func main() {
	// 1、加载配置文件,生成 config文件
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、实例化DiscoveryClient
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3、发送请求获取GVR数据
	_, apiResource, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}
	for _, list := range apiResource {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err.Error())
		}
		for _, resource := range list.APIResources {
			fmt.Printf("name:%v,group:%v,version:%v\n", resource.Name, gv.Group, gv.Version)
		}
	}
}

```

### GVR 缓存到本地

GVR 数据是一个很少变动的，因此可以将 GVR 的数据缓存到本地，来减少 Client 与 APIServer 交互，毕竟他们的交互还是存在网络的消耗的。

在 discovery/cached 中，有另外两个客户是来实现 GVR 数据缓存到本地文件中的内存中的，分配是 CachedDiscoveryClient 和 MemCacheClient。

其实，我们平时管理 k8s 的集群 kubectl 命令也是使用这种方式来使用我们的 GVR 与 APIServer 交互的，它的缓存文件默认是在 `～/.kube/cache` 中的。

**栗子：**

```go
package main

import (
	`fmt`
	`time`

	`k8s.io/apimachinery/pkg/runtime/schema`
	`k8s.io/client-go/discovery/cached/disk`
	`k8s.io/client-go/tools/clientcmd`
)

func main() {
	// 1、加载配置文件,生成 config文件
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、实例化,将GVR缓存到本地
	cacheDiscoveryClient, err := disk.NewCachedDiscoveryClientForConfig(config, "./cache/discovery", "./cache/http", time.Minute*60)
	if err != nil {
		panic(err.Error())
	}

	// 3、
	_, apiResource, err := cacheDiscoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}
	for _, list := range apiResource {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err.Error())
		}
		for _, resource := range list.APIResources {
			fmt.Printf("name:%v,group:%v,version:%v\n", resource.Name, gv.Group, gv.Version)
		}
	}
}

```

