package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	klient "github.com/vitu1234/iot-operator/pkg/client/clientset/versioned"
	kinfFac "github.com/vitu1234/iot-operator/pkg/client/informers/externalversions"
	controller "github.com/vitu1234/iot-operator/pkg/controller"
	"github.com/vitu1234/iot-operator/pkg/ocfclient"

	"github.com/jessevdk/go-flags"

	Options "github.com/vitu1234/iot-operator/pkg/apis/iot.dev/v1alpha1"
	OCFClient "github.com/vitu1234/iot-operator/pkg/ocfclient"
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

	clientset, err := kubernetes.NewForConfig(config) // clientset because it is used to interact with clients from different API versions
	if err != nil {
		fmt.Printf("Error getting standard clientset: %s\n", err.Error())
	}

	// fmt.Println(clientset)
	// ocfdevices, err := klientset.IotV1alpha1().OCFDevices("").List(context.Background(), metav1.ListOptions{})

	// if err != nil {
	// 	log.Printf("sError getting ocfdevices: %s \n", err)
	// }

	// fmt.Println(ocfdevices)
	var opts Options.Options
	parser := flags.NewParser(&opts, flags.Default)
	_, err = parser.Parse()
	if err != nil {
		fmt.Println("Parsing command options has failed : " + err.Error())
	}

	ocfclient.ReadCommandOptions(opts)

	fmt.Println("we here ")
	discoveryTimeout := opts.DiscoveryTimeout
	if discoveryTimeout <= 0 {
		discoveryTimeout = time.Second * 5
	}

	// Create OCF Client
	client := OCFClient.OCFClient{}
	err = client.Initialize()
	if err != nil {
		fmt.Println("OCF Client has failed to initialize : " + err.Error())
	}

	res, err := client.Discover(discoveryTimeout)
	if err != nil {
		println("\nDiscovering devices has failed : " + err.Error())
	}
	println("\nDiscovered devices : \n" + res)

	infoFactory := kinfFac.NewSharedInformerFactory(klientset, 20*time.Minute)

	ch := make(chan struct{})
	c := controller.NewController(clientset, klientset, infoFactory.Iot().V1alpha1().OCFDevices())

	infoFactory.Start(ch)
	if err := c.Run(ch); err != nil {
		log.Printf("Error running controller: %s\n", err.Error())
	}

}
