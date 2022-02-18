package core

import (
	"sync"

	appsv1 "k8s.io/api/apps/v1"
)

type DeploymentHandler struct {
	sync.Map
	manage *WsManage
}

var deploymentHandlerInstance *DeploymentHandler
var deploymentHandlerOnce sync.Once

func GetDeploymentHandler() *DeploymentHandler {
	deploymentHandlerOnce.Do(func() {
		deploymentHandlerInstance = &DeploymentHandler{
			manage: GetWsManage(),
		}
	})
	return deploymentHandlerInstance
}

func (dh *DeploymentHandler) Apply(i interface{}) {
	i.(*K8sInformers).Apps().V1().Deployments().Informer().AddEventHandler(dh)
}

func (dh *DeploymentHandler) OnAdd(obj interface{}) {
	deployment := obj.(*appsv1.Deployment)
	if list, ok := dh.Load(deployment.Namespace); ok {
		list = append(list.([]*appsv1.Deployment), deployment)
		dh.Store(deployment.Namespace, list)
	} else {
		dh.Store(deployment.Namespace, []*appsv1.Deployment{deployment})
	}
	dh.manage.Notice(deployment.Namespace, string(WSTYPE_RESOURCETYPE_DEPLOYMENT), WSACTION_ADD, deployment)
}

func (dh *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	deployment := oldObj.(*appsv1.Deployment)
	if list, ok := dh.Load(deployment.Namespace); ok {
		for i, item := range list.([]*appsv1.Deployment) {
			if item.UID == deployment.UID {
				list.([]*appsv1.Deployment)[i] = newObj.(*appsv1.Deployment)
			}
		}
	}
	dh.manage.Notice(newObj.(*appsv1.Deployment).Namespace, string(WSTYPE_RESOURCETYPE_DEPLOYMENT), WSACTION_UPDATE, deployment)
}

func (dh *DeploymentHandler) OnDelete(obj interface{}) {
	deployment := obj.(*appsv1.Deployment)
	if list, ok := dh.Load(deployment.Namespace); ok {
		for i, item := range list.([]*appsv1.Deployment) {
			if item.UID == deployment.UID {
				newList := append(list.([]*appsv1.Deployment)[:i], list.([]*appsv1.Deployment)[i+1:]...)
				dh.Store(deployment.Namespace, newList)
				break
			}
		}
	}
	dh.manage.Notice(deployment.Namespace, string(WSTYPE_RESOURCETYPE_DEPLOYMENT), WSACTION_DELETE, deployment)
}
