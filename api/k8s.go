package api

import (
	"errors"
	"strings"
	k8sApi "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	k8sMeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	k8sRest "k8s.io/client-go/rest"
	k8sCmd "k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	clientSet *k8s.Clientset
}

func NewKubeClient(kubeConfig string) (*KubeClient, error) {
	var (
		config *k8sRest.Config
		err    error
	)

	if kubeConfig != "" {
		config, err = k8sCmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, err
		}
	} else {
		config, err = k8sRest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	clentSet, err := k8s.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	c := &KubeClient{
		clientSet: clentSet,
	}

	return c, nil
}

func (c *KubeClient) SecretCreate(clientId, clientSecret, k8sSecretName string) (*k8sApi.Secret, error) {
	sn := strings.Split(k8sSecretName, "/")
	if len(sn) != 2 {
		return nil, errors.New("invalid format of k8s Secret name (namespace/name).")
	}
	secretNamespace := sn[0]
	secretName := sn[1]

	_, err := c.clientSet.CoreV1().Secrets(secretNamespace).Get(secretName, k8sMeta.GetOptions{})
	if err != nil {
		if !k8sErrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		return nil, errors.New("k8s Secret '" + k8sSecretName + "' already exists.")
	}

	secret := &k8sApi.Secret{
		ObjectMeta: k8sMeta.ObjectMeta{
			Namespace: secretNamespace,
			Name:      secretName,
		},
		StringData: map[string]string{
			"id":     clientId,
			"secret": clientSecret,
		},
	}

	return c.clientSet.CoreV1().Secrets(secretNamespace).Create(secret)
}

func (c *KubeClient) SecretExist(k8sSecretName string) (bool, error) {
	sn := strings.Split(k8sSecretName, "/")
	if len(sn) != 2 {
		return false, errors.New("invalid format of k8s Secret name (namespace/name).")
	}
	secretNamespace := sn[0]
	secretName := sn[1]

	_, err := c.clientSet.CoreV1().Secrets(secretNamespace).Get(secretName, k8sMeta.GetOptions{})
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
