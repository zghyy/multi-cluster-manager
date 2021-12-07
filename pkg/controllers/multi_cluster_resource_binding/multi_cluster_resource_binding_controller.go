package multi_cluster_resource_binding

import (
	"fmt"
	multiclusterv1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	clientset "harmonycloud.cn/multi-cluster-manager/pkg/client/clientset/versioned"
	mcmscheme "harmonycloud.cn/multi-cluster-manager/pkg/client/clientset/versioned/scheme"
	informers "harmonycloud.cn/multi-cluster-manager/pkg/client/informers/externalversions/multicluster/v1alpha1"
	listers "harmonycloud.cn/multi-cluster-manager/pkg/client/listers/multicluster/v1alpha1"
	_ "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"time"
)

const controllerAgentName = "multi-cluster-resource-binding-controller"

type MultiClusterResourceBindingController struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// sampleclientset is a clientset for our own API group
	sampleclientset clientset.Interface

	multiClusterResourceBindingLister listers.MultiClusterResourceBindingLister
	multiClusterResourceBindingSynced cache.InformerSynced

	multiClusterResourceLister listers.MultiClusterResourceLister
	multiClusterResourceSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

func NewMultiClusterResourceBindingController(
	kubeclientset kubernetes.Interface,
	sampleclientset clientset.Interface,
	multiClusterResourceBindingInformer informers.MultiClusterResourceBindingInformer,
	multiClusterResourceInformers informers.MultiClusterResourceInformer) *MultiClusterResourceBindingController {

	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	utilruntime.Must(mcmscheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &MultiClusterResourceBindingController{
		kubeclientset:                     kubeclientset,
		sampleclientset:                   sampleclientset,
		multiClusterResourceBindingLister: multiClusterResourceBindingInformer.Lister(),
		multiClusterResourceBindingSynced: multiClusterResourceBindingInformer.Informer().HasSynced,
		multiClusterResourceLister:        multiClusterResourceInformers.Lister(),
		multiClusterResourceSynced:        multiClusterResourceInformers.Informer().HasSynced,
		workqueue:                         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Foos"),
		recorder:                          recorder,
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	multiClusterResourceBindingInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleMultiClusterResourceBinding,
		UpdateFunc: func(old, new interface{}) {
			controller.handleMultiClusterResourceBinding(new)
		},
		DeleteFunc: controller.handleMultiClusterResourceBinding,
	})

	//TODO multiClusterResourceHandler

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *MultiClusterResourceBindingController) Run(workers int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting Foo controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.multiClusterResourceBindingSynced, c.multiClusterResourceSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	// Launch two workers to process Foo resources
	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *MultiClusterResourceBindingController) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *MultiClusterResourceBindingController) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *MultiClusterResourceBindingController) enqueue(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
	}
	c.workqueue.Add(key)
}

func (c *MultiClusterResourceBindingController) handleMultiClusterResourceBinding(obj interface{}) {
	multiClusterResourceBinding, ok := obj.(multiclusterv1alpha1.MultiClusterResourceBinding)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("error decoding object into MultiClusterResourceBinding"))
		return
	}
	c.enqueue(multiClusterResourceBinding)
}

func (c *MultiClusterResourceBindingController) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	multiClusterResourceBinding, err := c.multiClusterResourceBindingLister.MultiClusterResourceBindings(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("MultiClusterResourceBinding '%s' in work queue no longer exists", key))
			return nil
		}
		return err
	}

	for _, resource := range multiClusterResourceBinding.Spec.Resources {
		var multiClusterResource multiclusterv1alpha1.MultiClusterResource
		multiClusterResource, err = c.multiClusterResourceLister.MultiClusterResources(multiClusterResourceBinding.Namespace).Get(resource.Name)
		if err != nil {
			return err
		}

	}

	return nil
}
