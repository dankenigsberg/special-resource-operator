// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/sro/v1alpha1.SpecialResource":                schema_pkg_apis_sro_v1alpha1_SpecialResource(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceArtifacts":       schema_pkg_apis_sro_v1alpha1_SpecialResourceArtifacts(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceBuilArgs":        schema_pkg_apis_sro_v1alpha1_SpecialResourceBuilArgs(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceClaims":          schema_pkg_apis_sro_v1alpha1_SpecialResourceClaims(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceDependsOn":       schema_pkg_apis_sro_v1alpha1_SpecialResourceDependsOn(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceDriverContainer": schema_pkg_apis_sro_v1alpha1_SpecialResourceDriverContainer(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceGit":             schema_pkg_apis_sro_v1alpha1_SpecialResourceGit(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceImages":          schema_pkg_apis_sro_v1alpha1_SpecialResourceImages(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceNode":            schema_pkg_apis_sro_v1alpha1_SpecialResourceNode(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourcePaths":           schema_pkg_apis_sro_v1alpha1_SpecialResourcePaths(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceRunArgs":         schema_pkg_apis_sro_v1alpha1_SpecialResourceRunArgs(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceSource":          schema_pkg_apis_sro_v1alpha1_SpecialResourceSource(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceSpec":            schema_pkg_apis_sro_v1alpha1_SpecialResourceSpec(ref),
		"./pkg/apis/sro/v1alpha1.SpecialResourceStatus":          schema_pkg_apis_sro_v1alpha1_SpecialResourceStatus(ref),
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResource is the Schema for the specialresources API",
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
							Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/sro/v1alpha1.SpecialResourceSpec", "./pkg/apis/sro/v1alpha1.SpecialResourceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceArtifacts(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceArtifacts defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"hostPaths": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourcePaths"),
									},
								},
							},
						},
					},
					"images": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceImages"),
									},
								},
							},
						},
					},
					"claims": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceClaims"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/sro/v1alpha1.SpecialResourceClaims", "./pkg/apis/sro/v1alpha1.SpecialResourceImages", "./pkg/apis/sro/v1alpha1.SpecialResourcePaths"},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceBuilArgs(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceBuilArgs defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "value"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceClaims(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceClaims defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mountPath": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "mountPath"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceDependsOn(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceDependsOn defines the desired state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceDriverContainer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceDriverContainer defines the desired state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"source": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("./pkg/apis/sro/v1alpha1.SpecialResourceSource"),
						},
					},
					"buildArgs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceBuilArgs"),
									},
								},
							},
						},
					},
					"runArgs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceRunArgs"),
									},
								},
							},
						},
					},
					"artifacts": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceArtifacts"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/sro/v1alpha1.SpecialResourceArtifacts", "./pkg/apis/sro/v1alpha1.SpecialResourceBuilArgs", "./pkg/apis/sro/v1alpha1.SpecialResourceRunArgs", "./pkg/apis/sro/v1alpha1.SpecialResourceSource"},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceGit(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceGit defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"ref": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"uri": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"ref", "uri"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceImages(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceImages defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"pullsecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourcePaths"),
									},
								},
							},
						},
					},
				},
				Required: []string{"name", "kind", "namespace", "pullsecret", "path"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/sro/v1alpha1.SpecialResourcePaths"},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceNode(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceNode defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"selector": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"selector"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourcePaths(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourcePaths defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"sourcePath": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"destinationDir": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"sourcePath", "destinationDir"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceRunArgs(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceRunArgs defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "value"},
			},
		},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceSource defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"git": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceGit"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/sro/v1alpha1.SpecialResourceGit"},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceSpec defines the desired state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"driverContainer": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("./pkg/apis/sro/v1alpha1.SpecialResourceDriverContainer"),
						},
					},
					"node": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceNode"),
						},
					},
					"dependsOn": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/sro/v1alpha1.SpecialResourceDependsOn"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/sro/v1alpha1.SpecialResourceDependsOn", "./pkg/apis/sro/v1alpha1.SpecialResourceDriverContainer", "./pkg/apis/sro/v1alpha1.SpecialResourceNode"},
	}
}

func schema_pkg_apis_sro_v1alpha1_SpecialResourceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpecialResourceStatus defines the observed state of SpecialResource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"state"},
			},
		},
	}
}