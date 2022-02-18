package core

import (
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
)

type K8sInformers struct {
	informers.SharedInformerFactory
}

var k8sInformersInstance *K8sInformers
var k8sInformersOnce sync.Once

func GetK8sInformers() *K8sInformers {
	k8sInformersOnce.Do(func() {
		k8sInformersInstance = &K8sInformers{
			SharedInformerFactory: informers.NewSharedInformerFactory(
				GetK8sClient(),
				time.Second*3,
			),
		}
	})

	return k8sInformersInstance
}

func (ki *K8sInformers) Mount(applys ...IApply) *K8sInformers {
	for _, item := range applys {
		item.Apply(ki)
	}
	return ki
}

func (ki *K8sInformers) Launch() *K8sInformers {
	ki.Start(wait.NeverStop)
	return ki
}
