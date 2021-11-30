package common

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen=true

type Condition struct {
	Timestamp metav1.Time `json:"timestamp"`
	Message   string      `json:"message"`
	Reason    string      `json:"reason"`
	Type      string      `json:"type"`
}
