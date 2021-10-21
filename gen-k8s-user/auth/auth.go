package auth

import (
	"context"

	"gen-k8s-user/types"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	appmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	TEST_ENDPOINT = "https://10.214.39.62:5443"
)

func NewK8sClient(c *kubernetes.Clientset) K8sClient {
	return K8sClient{
		client: c,
	}

}

type K8sClient struct {
	client *kubernetes.Clientset
}

// 创建账号
func (k *K8sClient) CreateAccount(name, namespace string) (sa *corev1.ServiceAccount, err error) {
	labels := map[string]string{
		"member-name": name,
		"groups-name": "rcmd",
	}
	metasa := corev1.ServiceAccount{ObjectMeta: appmetav1.ObjectMeta{Name: name, Labels: labels}}
	sa, err = k.client.CoreV1().ServiceAccounts(namespace).Create(context.TODO(), &metasa, appmetav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		sa, err = k.client.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), name, appmetav1.GetOptions{})
	}
	return
}

// 绑定角色
func (k *K8sClient) BindRole(role, namespace string, sa *corev1.ServiceAccount) error {
	sub := rbacv1.Subject{
		Kind:      "ServiceAccount",
		Name:      sa.Name,
		Namespace: sa.Namespace,
	}
	ref := rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     role,
	}
	roleb := rbacv1.RoleBinding{
		ObjectMeta: appmetav1.ObjectMeta{Name: role},
		Subjects:   []rbacv1.Subject{sub},
		RoleRef:    ref,
	}
	_, err := k.client.RbacV1().RoleBindings(namespace).Create(context.TODO(), &roleb, appmetav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		err = nil
	}
	return err
}

// 生成kubeconfig
func (k *K8sClient) GenKubeConfig(sa *corev1.ServiceAccount) (kubeconfig types.KubeConfig, err error) {
	cluster := clientcmdapi.Cluster{
		Server: TEST_ENDPOINT,
	}
	user := clientcmdapi.AuthInfo{}
	for _, item := range sa.Secrets {
		scr, err := k.client.CoreV1().Secrets(sa.Namespace).Get(context.TODO(), item.Name, appmetav1.GetOptions{})
		if err != nil {
			return kubeconfig, err
		}
		for k, v := range scr.Data {
			if k == "ca.crt" {
				// ca := base64.StdEncoding.EncodeToString(v)
				cluster.CertificateAuthorityData = v
			}
			if k == "token" {
				user.Token = string(v)
			}
		}
	}
	kubecluster := types.ConfigCluster{
		Name:    "test",
		Cluster: &cluster,
	}
	kubeauth := types.ConfigAuthInfo{
		Name:     sa.Name,
		AuthInfo: &user,
	}
	kubecontext := types.ConfigContext{
		Name: "test",
		Context: &clientcmdapi.Context{
			Cluster:  "test",
			AuthInfo: sa.Name,
		},
	}
	kubeconfig = types.KubeConfig{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters:   []*types.ConfigCluster{&kubecluster},
		AuthInfos:  []*types.ConfigAuthInfo{&kubeauth},
		Contexts:   []*types.ConfigContext{&kubecontext},
	}

	return kubeconfig, err
}
