package models

type PodModel struct {
	Name        string
	Namespace   string
	Phase       string
	PodIP       string
	HostIP      string
	QOSClass    string
	NodeName    string
	Images      []string
	Labels      map[string]string
	Annotations map[string]string
}
