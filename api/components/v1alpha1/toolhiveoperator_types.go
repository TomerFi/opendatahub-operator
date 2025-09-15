package v1alpha1

import (
	"github.com/opendatahub-io/opendatahub-operator/v2/api/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ToolHiveOperatorComponentName = "toolhiveoperator"
	// value should match whats set in the XValidation below
	ToolHiveOperatorInstanceName = "default-" + ToolHiveOperatorComponentName
	ToolHiveOperatorKind         = "ToolHiveOperator"
)

// Check that the component implements common.PlatformObject.
var _ common.PlatformObject = (*ToolHiveOperator)(nil)

// ToolHiveOperatorCommonSpec defines the shared desired state of ToolHiveOperator
type ToolHiveOperatorCommonSpec struct {
	// Spec fields exposed to the DSC API
}

// ToolHiveOperatorSpec defines the desired state of ToolHiveOperator
type ToolHiveOperatorSpec struct {
	ToolHiveOperatorCommonSpec `json:",inline"`
}

// ToolHiveOperatorCommonStatus defines the shared observed state of ToolHiveOperator
type ToolHiveOperatorCommonStatus struct {
	common.ComponentReleaseStatus `json:",inline"`
}

// ToolHiveOperatorStatus defines the observed state of ToolHiveOperator
type ToolHiveOperatorStatus struct {
	common.Status                `json:",inline"`
	ToolHiveOperatorCommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'default-toolhiveoperator'",message="ToolHiveOperator name must be default-toolhiveoperator"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`,description="Ready"
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`,description="Reason"

// ToolHiveOperator is the Schema for the ToolHiveOperators API
type ToolHiveOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToolHiveOperatorSpec   `json:"spec,omitempty"`
	Status ToolHiveOperatorStatus `json:"status,omitempty"`
}

func (c *ToolHiveOperator) GetStatus() *common.Status {
	return &c.Status.Status
}

func (c *ToolHiveOperator) GetConditions() []common.Condition {
	return c.Status.GetConditions()
}

func (c *ToolHiveOperator) SetConditions(conditions []common.Condition) {
	c.Status.SetConditions(conditions)
}

func (c *ToolHiveOperator) GetReleaseStatus() *[]common.ComponentRelease {
	return &c.Status.Releases
}

func (c *ToolHiveOperator) SetReleaseStatus(releases []common.ComponentRelease) {
	c.Status.Releases = releases
}

// +kubebuilder:object:root=true

// ToolHiveOperatorList contains a list of ToolHiveOperator
type ToolHiveOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolHiveOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolHiveOperator{}, &ToolHiveOperatorList{})
}

type DSCToolHiveOperator struct {
	common.ManagementSpec      `json:",inline"`
	ToolHiveOperatorCommonSpec `json:",inline"`
}

// DSCToolHiveOperatorStatus contains the observed state of the ToolHiveOperator exposed in the DSC instance
type DSCToolHiveOperatorStatus struct {
	common.ManagementSpec         `json:",inline"`
	*ToolHiveOperatorCommonStatus `json:",inline"`
}
