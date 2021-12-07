package multi_cluster_resource_binding

import (
	"github.com/sirupsen/logrus"
	"context"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/record"
	controllerruntime "sigs.k8s.io/controller-runtime"
	multiclusterv1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ControllerName is the controller name that will be used when reporting events.
const ControllerName = "multi-cluster-resource-binding-controller"


type MultiClusterResourceBindingController struct {
	client.Client                     // used to operate ResourceBinding resources.
	DynamicClient   dynamic.Interface // used to fetch arbitrary resources.
	EventRecorder   record.EventRecorder
	RESTMapper      meta.RESTMapper
	//OverrideManager overridemanager.OverrideManager
}

func (c *MultiClusterResourceBindingController) Reconcile(ctx context.Context, req controllerruntime.Request) (controllerruntime.Result, error) {
	logrus.Info("Reconciling MultiClusterResourceBinding %s",req.NamespacedName.String())

	multiclusterresourcebinding := &multiclusterv1alpha1.MultiClusterResourceBinding{}
	err := c.Client.Get(context.TODO(),req.NamespacedName,multiclusterresourcebinding)
	if err != nil{
		//TODO 删除逻辑
		logrus.Error("MultiClusterResourceBinding %s in %s not found",req.NamespacedName.String())
	}
	//TODO 正常更新下发

}
