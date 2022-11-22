/*
Copyright 2022 yf.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	etcdv1alpha1 "github.com/metashops/etcd-operator/api/v1alpha1"
)

// EtcdClusterReconciler reconciles a EtcdCluster object
type EtcdClusterReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=etcd.yf.io,resources=etcdclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=etcd.yf.io,resources=etcdclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=etcd.yf.io,resources=etcdclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the EtcdCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *EtcdClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// 获取 EtcdCluster instance
	var etcdCluster etcdv1alpha1.EtcdCluster
	if err := r.Client.Get(ctx, req.NamespacedName, &etcdCluster); err != nil {
		// EtcdCluster was deleted，Ignore
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// already given EtcdCluster instance
	// 创建/更新对应 statefulSet and Headless SVC object
	// 使用已经已提供好的 CreateOrUpdate
	// 调谐： 观察当前的状态和期望状态进行对比

	// CreateOrUpdate service
	var svc corev1.Service
	svc.Name = etcdCluster.Name // this is CRD name: etcdcluster-sample
	svc.Namespace = etcdCluster.Namespace
	op, err := ctrl.CreateOrUpdate(ctx, r, &svc, func() error {
		// 调谐必须在这个函数中去实现,实际是拼装我们的 Service
		MutateHeadlessSvc(&etcdCluster, &svc)
		return controllerutil.SetControllerReference(&etcdCluster, &svc, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, err
	}
	log.Log.Info("CreateOrUpdate Result", "Service", op)

	// CreateOrUpdate service
	var sts appsv1.StatefulSet
	sts.Name = etcdCluster.Name // this is CRD name: etcdCluster-sample
	sts.Namespace = etcdCluster.Namespace
	or, err := ctrl.CreateOrUpdate(ctx, r, &sts, func() error {
		// 调谐必须在这个函数中去实现,实际是拼装我们的 Service
		MutateStatefulSet(&etcdCluster, &sts)
		return controllerutil.SetControllerReference(&etcdCluster, &sts, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, err
	}
	log.Log.Info("CreateOrUpdate Result", "StatefulSet", or)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EtcdClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&etcdv1alpha1.EtcdCluster{}).
		Complete(r)
}
