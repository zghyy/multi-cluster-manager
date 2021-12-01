package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MultiClusterResourceAggregateRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultiClusterResourceAggregateRuleSpec   `json:"spec"`
	Status MultiClusterResourceAggregateRuleStatus `json:"status,omitempty"`
}

type MultiClusterResourceAggregateRuleSpec struct {
	ResourceRef MultiClusterResourceAggregateRuleResourceRef `json:"resourceRef"`
	Rule        MultiClusterResourceAggregateRuleRule        `json:"rule"`
}

type MultiClusterResourceAggregateRuleResourceRef struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type MultiClusterResourceAggregateRuleRule struct {
	Cue string `json:"cue"`
}

type MultiClusterResourceAggregateRuleStatus struct {
	// TODO should define status
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MultiClusterResourceAggregateRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultiClusterResourceAggregateRule `json:"items"`
}
