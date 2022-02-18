package core

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
)

type NamespaceHandler struct {
	sync.Map
	manage *WsManage
}

var namespaceHandlerInstance *NamespaceHandler
var namespaceHandlerOnce sync.Once

func GetNamespaceHandler() *NamespaceHandler {
	namespaceHandlerOnce.Do(func() {
		namespaceHandlerInstance = &NamespaceHandler{
			manage: GetWsManage(),
		}
	})

	return namespaceHandlerInstance
}

func (nh *NamespaceHandler) Apply(i interface{}) {
	i.(*K8sInformers).Core().V1().Namespaces().Informer().AddEventHandler(nh)
}

func (nh *NamespaceHandler) OnAdd(obj interface{}) {
	namespace := obj.(*corev1.Namespace)
	nh.Store(namespace.Name, namespace)
	nh.manage.Notice(namespace.Name, string(WSTYPE_RESOURCETYPE_NAMESPACE), WSACTION_ADD, namespace)
}

func (nh *NamespaceHandler) OnUpdate(oldObj, newObj interface{}) {
	oldNamespace := oldObj.(*corev1.Namespace)
	newNamespace := newObj.(*corev1.Namespace)
	nh.Store(oldNamespace.Name, newNamespace)
	nh.manage.Notice(newNamespace.Name, string(WSTYPE_RESOURCETYPE_NAMESPACE), WSACTION_UPDATE, newNamespace)
}

func (nh *NamespaceHandler) OnDelete(obj interface{}) {
	namespace := obj.(*corev1.Namespace)
	nh.Delete(namespace.Name)
	nh.manage.Notice(namespace.Name, string(WSTYPE_RESOURCETYPE_NAMESPACE), WSACTION_DELETE, obj.(*corev1.Namespace))
}

func (nh *NamespaceHandler) GetNamespaceByName(name string) *corev1.Namespace {
	if namespace, ok := nh.Load(name); ok {
		return namespace.(*corev1.Namespace)
	}
	return nil
}

func (nh *NamespaceHandler) GetNamespaces() []*corev1.Namespace {
	result := []*corev1.Namespace{}

	nh.Range(func(key, value interface{}) bool {
		result = append(result, value.(*corev1.Namespace))
		return true
	})

	return result
}
