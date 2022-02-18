package core

type WsMessage struct {
	Code    string `json:"code"`
	Type    string `json:"type"`
	Action  string `json:"action"`
	Content map[string]interface{}
}

type WsAction string

const (
	WSACTION_SUBSCRIBE WsAction = "subscribe"
	WSACTION_UPDATE    WsAction = "update"
	WSACTION_ADD       WsAction = "add"
	WSACTION_DELETE    WsAction = "delete"
)

type ResourceType string

const (
	WSTYPE_RESOURCETYPE_DEPLOYMENT ResourceType = "deployment"
	WSTYPE_RESOURCETYPE_EVENT      ResourceType = "event"
	WSTYPE_RESOURCETYPE_POD        ResourceType = "pod"
	WSTYPE_RESOURCETYPE_NAMESPACE  ResourceType = "namespace"
	WSTYPE_RESOURCETYPE_INGRESS    ResourceType = "ingress"
	WSTYPE_RESOURCETYPE_SERVICE    ResourceType = "service"
)

type NamespaceAlias string

const (
	NAMESPACE_DEV NamespaceAlias = "dev"
)
