package pkg

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getClient() *kubernetes.Clientset {
	var kubeconfig string;
	if conf := os.Getenv("KUBECONFIG"); conf != "" {
		kubeconfig = conf
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err.Error())
		}
		kubeconfig = filepath.Join(home, ".kube", "config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client
}
