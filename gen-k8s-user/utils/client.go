package utils

import (
	"fmt"
	"io/ioutil"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func NewKubeClient() (*KubeClient, error) {
	k := &KubeClient{}
	err := k.findAllContextFromConfig()
	if err != nil {
		return k, err
	}
	return k, nil

}

type KubeClient struct {
	contexts []string
	current  string
}

// 获取指定contextname的clientset
func (k *KubeClient) GetClientSet(env string) (cli *kubernetes.Clientset, err error) {
	for _, c := range k.contexts {
		if c == env {
			cfg, err := config.GetConfigWithContext(env)
			if err != nil {
				return nil, err
			}
			cli, err = kubernetes.NewForConfig(cfg)
			if err != nil {
				return nil, err
			}
			return cli, err
		}
	}
	return nil, fmt.Errorf("can not find the env in the kubeconfig file, only be %v", k.contexts)
}

// 搜索 kubeconfig 的 context
func (k *KubeClient) findAllContextFromConfig() (err error) {
	lr := clientcmd.NewDefaultClientConfigLoadingRules()
	// lr.Precedence 为 kubeconfig 文件
	for _, filename := range lr.Precedence {
		kubeconfigBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		config, err := clientcmd.Load(kubeconfigBytes)
		if err != nil {
			return err
		}
		// 当前context
		k.current = config.CurrentContext
		// 所有的 context
		for key := range config.Contexts {
			k.contexts = append(k.contexts, key)
		}
	}
	return
}
