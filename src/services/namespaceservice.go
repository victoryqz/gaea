package services

import (
	"sync"

	"gaea.olympus.io/src/core"
	"gaea.olympus.io/src/models"
	corev1 "k8s.io/api/core/v1"
)

type NamespaceService struct {
	handler *core.NamespaceHandler
}

var namespaceHandlerInstance *NamespaceService
var namespaceHandlerOnce sync.Once

func GetNamespaceService() *NamespaceService {
	namespaceHandlerOnce.Do(func() {
		namespaceHandlerInstance = &NamespaceService{
			handler: core.GetNamespaceHandler(),
		}
	})

	return namespaceHandlerInstance
}

func (ns *NamespaceService) ShowAllNamespace() []*models.NamespaceModel {
	namespaces := ns.handler.GetNamespaces()
	result := []*models.NamespaceModel{}
	for _, item := range namespaces {
		result = append(result, ns.Convert(item))
	}

	return result
}

func (ns *NamespaceService) Convert(namespace *corev1.Namespace) *models.NamespaceModel {
	return &models.NamespaceModel{
		Name:              namespace.Name,
		Labels:            namespace.Labels,
		Phase:             string(namespace.Status.Phase),
		CreationTimestamp: namespace.CreationTimestamp.Time,
		UID:               string(namespace.UID),
		Conditions:        ns.ConvertConditionsToStringArray(namespace.Status.Conditions),
		Deployments:       nil,
	}
}

func (ns *NamespaceService) ConvertConditionsToStringArray(conditions []corev1.NamespaceCondition) []string {
	arr := []string{}
	for _, item := range conditions {
		arr = append(arr, item.Message)
	}
	return arr
}
