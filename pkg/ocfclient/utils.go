// package ocfclient

// import (
// 	"bufio"
// 	"context"
// 	"crypto"
// 	"crypto/ecdsa"
// 	"crypto/elliptic"
// 	"crypto/rand"
// 	"crypto/tls"
// 	"crypto/x509"
// 	"encoding/pem"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"time"

// 	Options "github.com/vitu1234/iot-operator/pkg/apis/iot.dev/v1alpha1"

// 	"github.com/jessevdk/go-flags"
// 	local "github.com/plgd-dev/device/v2/client"
// 	"github.com/plgd-dev/device/v2/client/core"
// 	"github.com/plgd-dev/device/v2/pkg/codec/json"
// 	"github.com/plgd-dev/device/v2/pkg/security/generateCertificate"
// 	"github.com/plgd-dev/device/v2/pkg/security/signer"
// 	"github.com/plgd-dev/kit/v2/security"
// )

// func main() {
// 	var opts Options.Options
// 	parser := flags.NewParser(&opts, flags.Default)
// 	_, err := parser.Parse()
// 	if err != nil {
// 		fmt.Println("Parsing command options has failed : " + err.Error())
// 	}

// 	// Read Command Options
// 	ReadCommandOptions(opts)

// 	// Create OCF Client
// 	client := OCFClient{}
// 	err = client.Initialize()
// 	if err != nil {
// 		fmt.Println("OCF Client has failed to initialize : " + err.Error())
// 	}

// 	// Console Input
// 	scanner(client, opts.DiscoveryTimeout)
// }

// func scanner(client OCFClient, discoveryTimeout time.Duration) {
// 	if discoveryTimeout <= 0 {
// 		discoveryTimeout = time.Second * 5
// 	}

// 	scanner := bufio.NewScanner(os.Stdin)
// 	printMenu()
// 	for scanner.Scan() {
// 		selMenu, _ := strconv.ParseInt(scanner.Text(), 10, 32)
// 		switch selMenu {
// 		case 0:
// 			printMenu()
// 		case 1:
// 			res, err := client.Discover(discoveryTimeout)
// 			if err != nil {
// 				println("\nDiscovering devices has failed : " + err.Error())
// 				break
// 			}
// 			println("\nDiscovered devices : \n" + res)
// 		case 2:
// 			// Select Device
// 			print("\nInput device ID : ")
// 			scanner.Scan()
// 			deviceID := scanner.Text()
// 			res, err := client.OwnDevice(deviceID)
// 			if err != nil {
// 				println("\nTransferring Ownership has failed : " + err.Error())
// 				break
// 			}
// 			println("\nTransferring Ownership of " + deviceID + " was successful : \n" + res)
// 		case 3:
// 			// Select Device
// 			print("\nInput device ID : ")
// 			scanner.Scan()
// 			deviceID := scanner.Text()
// 			res, err := client.GetResources(deviceID)
// 			if err != nil {
// 				println("\nGetting Resources has failed : " + err.Error())
// 				break
// 			}
// 			println("\nResources of " + deviceID + " : \n" + res)
// 		case 4:
// 			// Select Device
// 			print("\nInput device ID : ")
// 			scanner.Scan()
// 			deviceID := scanner.Text()
// 			res, err := client.GetResources(deviceID)
// 			if err != nil {
// 				println("\nGetting Resources has failed : " + err.Error())
// 				break
// 			}
// 			println("\nResources of " + deviceID + " : \n" + res)

// 			// Select Resource
// 			print("\nInput resource href : ")
// 			scanner.Scan()
// 			href := scanner.Text()
// 			aRes, err := client.GetResource(deviceID, href)
// 			if err != nil {
// 				println("\nGetting Resource has failed : " + err.Error())
// 				break
// 			}
// 			println("\nResource properties of " + deviceID + href + " : \n" + aRes)
// 		case 5:
// 			// Select Device
// 			print("\nInput device ID : ")
// 			scanner.Scan()
// 			deviceID := scanner.Text()
// 			res, err := client.GetResources(deviceID)
// 			if err != nil {
// 				println("\nGetting Resources has failed : " + err.Error())
// 				break
// 			}
// 			println("\nResources of " + deviceID + " : \n" + res)

// 			// Select Resource
// 			print("\nInput resource href : ")
// 			scanner.Scan()
// 			href := scanner.Text()
// 			aRes, err := client.GetResource(deviceID, href)
// 			if err != nil {
// 				println("\nGetting Resource has failed : " + err.Error())
// 				break
// 			}
// 			println("\nResource properties of " + deviceID + href + " : \n" + aRes)

// 			// Select Property
// 			print("\nInput property name : ")
// 			scanner.Scan()
// 			key := scanner.Text()
// 			// Input Value of the property
// 			print("\nInput property value : ")
// 			scanner.Scan()
// 			value := scanner.Text()

// 			// Update Property of the Resource
// 			jsonString := "{\"" + key + "\": " + value + "}"
// 			var data interface{}
// 			err = json.Decode([]byte(jsonString), &data)
// 			if err != nil {
// 				println("\nDecoding resource property has failed : " + err.Error())
// 				break
// 			}
// 			dataBytes, err := json.Encode(data)
// 			if err != nil {
// 				println("\nEncoding resource property has failed : " + err.Error())
// 				break
// 			}
// 			println("\nProperty data to update : " + string(dataBytes))
// 			err = client.UpdateResource(deviceID, href, data)
// 			if err != nil {
// 				println("\nUpdating resource property has failed : " + err.Error())
// 				break
// 			}
// 			println("\nUpdating resource property of " + deviceID + href + " was successful.")
// 		case 6:
// 			// Select Device
// 			print("\nInput device ID : ")
// 			scanner.Scan()
// 			deviceID := scanner.Text()
// 			err := client.DisownDevice(deviceID)
// 			if err != nil {
// 				println("\nOff-boarding has failed : " + err.Error())
// 				break
// 			}
// 			println("\nOff-boarding " + deviceID + " was successful.")
// 		case 99:
// 			// Close Client
// 			if errC := client.Close(); errC != nil {
// 				println("\nCannot close client: %v", errC)
// 			}
// 			os.Exit(0)
// 		}
// 		printMenu()
// 	}
// }

//	func printMenu() {
//		fmt.Println("\n#################### OCF Client for D2D ####################")
//		fmt.Println("[0] Display this menu")
//		fmt.Println("--------------------------------------------------------------")
//		fmt.Println("[1] Discover devices")
//		fmt.Println("[2] Transfer ownership to the device (On-boarding)")
//		fmt.Println("[3] Retrieve resources of the device")
//		fmt.Println("[4] Retrieve a resource of the device")
//		fmt.Println("[5] Update a resource of the device")
//		fmt.Println("[6] Reset ownership of the device (Off-boarding)")
//		fmt.Println("--------------------------------------------------------------")
//		fmt.Println("[99] Exit")
//		fmt.Println("##############################################################")
//		fmt.Print("\nSelect menu : ")
//	}
package ocfclient
