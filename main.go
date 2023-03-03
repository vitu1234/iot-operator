package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	klient "github.com/vitu1234/iot-operator/pkg/client/clientset/versioned"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//if failed to find kubeconfig location because the code is now running in cluster, do the following
		fmt.Println(err)
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error getting incluster config: %s\n", err.Error())
		}

	}

	klientset, err := klient.NewForConfig(config) // clientset because it is used to interact with clients from different API versions
	if err != nil {
		fmt.Printf("Error getting custom klientset: %s\n", err.Error())
	}

	//clientset for k8s native resources

	// clientset, err := kubernetes.NewForConfig(config) // clientset because it is used to interact with clients from different API versions
	// if err != nil {
	// 	fmt.Printf("Error getting standard clientset: %s\n", err.Error())
	// }

	// fmt.Println(clientset)
	fmt.Println(klientset)

}
