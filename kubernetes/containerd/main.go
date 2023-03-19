package main

import (
	"context"
	`fmt`

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
)

func main() {
	// 连接到本地的 Containerd 实例
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 创建一个新的命名空间，以便我们可以管理容器和映像
	ctx := namespaces.WithNamespace(context.Background(), "my-namespace")

	// 创建一个新的容器
	container, err := client.NewContainer(
		ctx,
		"my-container",
		containerd.WithImage("my-image"),
		containerd.WithNewSpec(oci.WithImageConfig("my-image")),
	)
	if err != nil {
		panic(err)
	}
	defer container.Delete(ctx, containerd.WithSnapshotCleanup)

	// 启动容器
	task, err := container.NewTask(ctx, oci.WithProcessArgs("my-app"))
	if err != nil {
		panic(err)
	}
	defer task.Delete(ctx)

	// 等待容器退出
	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		panic(err)
	}

	// 输出容器的退出状态
	status := <-exitStatusC
	fmt.Printf("Container exited with status %d\n", status.ExitCode())
}
