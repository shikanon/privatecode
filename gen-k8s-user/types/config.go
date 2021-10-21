package types

import (
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// 生成的kubeconfig数据结构
type KubeConfig struct {
	// Legacy field from pkg/api/types.go TypeMeta.
	// TODO(jlowdermilk): remove this after eliminating downstream dependencies.
	// +k8s:conversion-gen=false
	// +optional
	Kind string `json:"kind,omitempty"`
	// Legacy field from pkg/api/types.go TypeMeta.
	// TODO(jlowdermilk): remove this after eliminating downstream dependencies.
	// +k8s:conversion-gen=false
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
	// Preferences holds general information to be use for cli interactions
	Preferences clientcmdapi.Preferences `json:"preferences"`
	// Clusters is a map of referencable names to cluster configs
	Clusters []*ConfigCluster `json:"clusters"`
	// AuthInfos is a map of referencable names to user configs
	AuthInfos []*ConfigAuthInfo `json:"users"`
	// Contexts is a map of referencable names to context configs
	Contexts []*ConfigContext `json:"contexts"`
	// CurrentContext is the name of the context that you would like to use by default
	CurrentContext string `json:"current-context"`
}

type ConfigCluster struct {
	Name    string                `json:"name"`
	Cluster *clientcmdapi.Cluster `json:"cluster"`
}

type ConfigAuthInfo struct {
	Name     string                 `json:"name"`
	AuthInfo *clientcmdapi.AuthInfo `json:"user"`
}

type ConfigContext struct {
	Name    string                `json:"name"`
	Context *clientcmdapi.Context `json:"context"`
}
