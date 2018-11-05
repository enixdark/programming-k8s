package controller

import (
	"fmt"
	"log"
	"time"

	api_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	clientset "github.com/interma/programming-k8s/pkg/client/clientset/versioned"
)

const maxRetries = 5
const namespace = "default"

// Controller object
type PodsStatsController struct {
	KubeClient kubernetes.Interface
	CrClient   clientset.Interface

	queue    workqueue.RateLimitingInterface
	informer cache.SharedIndexInformer
}

func CreatePodsStatsController(kc kubernetes.Interface, cc clientset.Interface) *PodsStatsController {
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return kc.CoreV1().Pods(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return kc.CoreV1().Pods(namespace).Watch(options)
			},
		},
		&api_v1.Pod{},
		0, //Skip resync
		cache.Indexers{},
	)

	c := newPodsStatsController(kc, cc, informer)

	return c
}

func newPodsStatsController(kc kubernetes.Interface, cc clientset.Interface,
	informer cache.SharedIndexInformer) *PodsStatsController {

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Printf("processing pod add: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			log.Printf("processing pod update: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			log.Printf("processing pod delete: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	return &PodsStatsController{kc, cc, queue, informer}
}

func (c *PodsStatsController) HasSynced() bool {
	return c.informer.HasSynced()
}

// Run starts the kubewatch controller
func (c *PodsStatsController) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	log.Printf("controller start")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("sync timeout"))
		return
	}

	log.Printf("controller synced and ready")

	//only one worker
	wait.Until(c.runWorker, time.Second, stopCh)
}

func (c *PodsStatsController) runWorker() {
	for c.processNextItem() {
		// continue looping
	}
}

func (c *PodsStatsController) processNextItem() bool {
	key, quit := c.queue.Get()

	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.processItem(key.(string))
	if err == nil {
		// No error, reset the ratelimit counters
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		log.Printf("Error processing %s (will retry): %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		log.Printf("Error processing %s (giving up): %v", key, err)
		c.queue.Forget(key)
		utilruntime.HandleError(err)
	}

	return true
}

func (c *PodsStatsController) processItem(key string) error {
	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}
	log.Printf("processItem: %v\n", obj)

	if exists {
		pod, _ := obj.(*api_v1.Pod)

		requestCpu := pod.Spec.Containers[0].Resources.Requests.Cpu() //assuming only one container here
		log.Printf("processItem: %s request cpu: %v\n", pod.Name, requestCpu)
	}

	return nil
}
