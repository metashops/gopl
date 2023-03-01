package main

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	virtv1 "kubevirt.io/client-go/api/v1"
	"kubevirt.io/client-go/kubecli"
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
	pods, err := clientSet.
		CoreV1().
		Pods("default").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range pods.Items {
		fmt.Printf("namespace:%v name:%v\n", item.Namespace, item.Name)
	}

	// 获取PVC
	pvcName := "my-pvc"
	pvc, err := clientSet.CoreV1().PersistentVolumeClaims("default").Get(context.TODO(), pvcName, metav1.GetOptions{})
	// Annotations的作用是提供有关对象的附加信息和元数据
	pvc.Annotations["mycompany.com/cdi"] = "none"
	pvc.Annotations["mycompany.com/cdi"] = "kubevirt"
	anntationMap := map[string]interface{}{
		"metadata": map[string]map[string]string{
			"annotations": pvc.Annotations,
		},
	}

	annotationsBytes, err := json.Marshal(anntationMap)

	ctx := context.Background()

	// 更新PVC对象Annotations
	_, err := clientSet.CoreV1().PersistentVolumeClaims("default").Patch(ctx, pvcName, types.MergePatchType, annotationsBytes)

	// AddVolume
	virtClient := kubecli.KubevirtClient
	err = virtClient.VirtualMachine(namespace).AddVolume(vmId, &virtv1.AddVolumeOptions{
		Name: volume,
		Disk: &virtv1.Disk{
			Name: volume,
			DiskDevice: virtv1.DiskDevice{
				Disk: &virtv1.DiskTarget{
					Bus: diskBus,
				},
			},
		},
		VolumeSource: vsPersist,
	})

}
