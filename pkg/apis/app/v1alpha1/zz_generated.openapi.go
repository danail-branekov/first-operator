// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"first-operator/first-operator/pkg/apis/app/v1alpha1.PodSet":       schema_pkg_apis_app_v1alpha1_PodSet(ref),
		"first-operator/first-operator/pkg/apis/app/v1alpha1.PodSetSpec":   schema_pkg_apis_app_v1alpha1_PodSetSpec(ref),
		"first-operator/first-operator/pkg/apis/app/v1alpha1.PodSetStatus": schema_pkg_apis_app_v1alpha1_PodSetStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_PodSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodSet is the Schema for the podsets API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("first-operator/first-operator/pkg/apis/app/v1alpha1.PodSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("first-operator/first-operator/pkg/apis/app/v1alpha1.PodSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"first-operator/first-operator/pkg/apis/app/v1alpha1.PodSetSpec", "first-operator/first-operator/pkg/apis/app/v1alpha1.PodSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_PodSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodSetSpec defines the desired state of PodSet",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_app_v1alpha1_PodSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodSetStatus defines the observed state of PodSet",
				Type:        []string{"object"},
			},
		},
	}
}
