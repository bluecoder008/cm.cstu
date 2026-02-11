package v1

import (
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)   

// ClassScheduleSpec defines the desired state of ClassSchedule
type ClassScheduleSpec struct {
    CourseName string        `json:"courseName,omitempty"`
    University string        `json:"university,omitempty"`
    Schedule   []ScheduleItem `json:"schedule,omitempty"`
}

type ScheduleItem struct {
    Day  string `json:"day,omitempty"`
    Time string `json:"time,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ClassSchedule struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   ClassScheduleSpec   `json:"spec,omitempty"`
    Status ClassScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ClassScheduleList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []ClassSchedule `json:"items"`
}

// ClassScheduleStatus defines the observed state (can be empty for now)
type ClassScheduleStatus struct {
    // Define observed fields here if needed
}

func init() {
        SchemeBuilder.Register(&ClassSchedule{}, &ClassScheduleList{})
}       

