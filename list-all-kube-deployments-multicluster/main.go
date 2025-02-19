package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type DeploymentTree map[string]map[string][]string

func main() {
	configAccess := clientcmd.NewDefaultPathOptions()

	config, err := configAccess.GetStartingConfig()
	if err != nil {
		fmt.Printf("Error %s, getting starting config", err.Error())
		os.Exit(1)
	}

	deploymentTree := make(DeploymentTree)

	for name := range config.Contexts {
        deploymentTree[name] = make(map[string][]string)

		config.CurrentContext = name
		err = clientcmd.ModifyConfig(configAccess, *config, true)
		if err != nil {
			fmt.Printf("Error %s, modifying config", err.Error())
			os.Exit(1)
		}

		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("error getting user home dir: %v\n", err)
			os.Exit(1)
		}
		kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

		kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err != nil {
			fmt.Printf("Error getting kubernetes config: %v\n", err)
			os.Exit(1)
		}

		clientset, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			fmt.Printf("error getting kubernetes config: %v\n", err)
			continue
		}

		namespaces, err := ListNamespaces(clientset)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, namespace := range namespaces.Items {
            deployments, err := ListDeployments(namespace.Name, clientset)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }

            for _, deployment := range deployments.Items {
			    deploymentTree[name][namespace.Name] = append(deploymentTree[name][namespace.Name], deployment.Name)
            }
		}
	}

	jsonDeployemntTree, err := json.Marshal(deploymentTree)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	fmt.Println(string(jsonDeployemntTree))
}

func ListNamespaces(client kubernetes.Interface) (*v1.NamespaceList, error) {
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, err
	}

	return namespaces, nil
}

func ListDeployments(namespace string, client kubernetes.Interface) (*appsv1.DeploymentList, error) {
	deployments, err := client.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, err
	}
	return deployments, nil
}
