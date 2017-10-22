package tester

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestRun struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   TestRunSpec
	Status TestRunStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestRunList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []TestRun
}

type TestRunSpec struct {
	Selector *metav1.LabelSelector
	MaxJobs  int
	MaxFail  int
}

type TestRunStatus struct {
	Status  string
	Message string
	Success bool
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestTemplate struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec TestTemplateSpec
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestTemplateList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []TestTemplate
}

type TestTemplateSpec struct {
	Template corev1.PodTemplateSpec
}
