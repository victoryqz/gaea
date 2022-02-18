package models

import "time"

type DeploymentModel struct {
	Name                string
	Namespace           string
	GenerateName        string
	Labels              map[string]string
	MatchLabels         map[string]string
	Replicas            int32
	AvailableReplicas   int32
	UnavailableReplicas int32
	Conditions          []string
	Pods                []string
	CreationTimestamp   time.Time
}
