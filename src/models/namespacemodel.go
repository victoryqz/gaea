package models

import "time"

type NamespaceModel struct {
	Name              string
	Phase             string
	UID               string
	Labels            map[string]string
	Conditions        []string
	Deployments       []string
	CreationTimestamp time.Time
}
