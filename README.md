go mod init github.com/vitu1234/iot-operator

Type: OCFDevice
Group: iot.dev
Version: v1alpha1

Create directory Structure
    pkg/apis/iot.dev/v1alpha1

Create types.go and add the tags on the type
    // +genclient
    // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
    // +kubebuilder:subresource:status

Register the type using schemes in register.go
    - use code-generator to register resources
    - clone https://github.com/kubernetes/code-generator into ~/go/src/k8s.io
    - use the code-generator to resolve deepcopy object not available for the type
    - generates clientset, informers, listers and deepcopy objects
    - create doc.go to specify global tags for the global apis
    - code-generator gen command
        1. export GOPATH=~/go
        2. ~/go/src/k8s.io/code-generator | git cloned to this location
        3. cd ~/go/src/github.com/vitu1234/kluster
        4. execDir=~/go/src/k8s.io/code-generator
        5. "${execDir}"/generate-groups.sh all github.com/vitu1234/iot-operator/pkg/client github.com/vitu1234/iot-operator/pkg/apis iot.dev:v1alpha1 --go-header-file "${execDir}"/hack/boilerplate.go.txt

            WHERE; project path: github.com/vitu1234/iot-operator
    - create the CRDs so the server can recognize the CRs
    - Generate CRDs with controller gen
        - controller-gen
            generate CRD for the type
            1. export GOPATH=~/go
            2. go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.6.0
            3. export PATH=$PATH:$GOPATH/bin
            4. source ~/.bashrc
            5. controller-gen --version
            6. controller-gen command run in in project path: 
                controller-gen paths=github.com/vitu1234/iot-operator/pkg/apis/iot.dev/v1alpha1 crd:trivialVersions=true crd:crdVersions=v1 output:crd:artifacts:config=manifests
