package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	"github.com/projectcalico/api/pkg/lib/numorstring"
)

func main() {
	// Create a new config based on kubeconfig file.
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "/Users/thientran/.kube/config", "absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Build a clientset based on the provided kubeconfig file.
	cs, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	tcp := numorstring.ProtocolFromString("TCP")
	// Test create network-policies
	np := v3.NetworkPolicy{
		TypeMeta: v1.TypeMeta{
			Kind:       "NetworkPolicy",
			APIVersion: "projectcalico.org/v3",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-policy-deny-all-egress",
			Namespace: "default",
		},
		Spec: v3.NetworkPolicySpec{
			Order: new(float64),
			Egress: []v3.Rule{
				{
					Action:   v3.Deny,
					Protocol: &tcp,
					Destination: v3.EntityRule{
						Selector: "app == 'nginx'",
						Ports: []numorstring.Port{
							{
								MinPort:  7740,
								MaxPort:  7749,
								PortName: "custom-port",
							},
						},
					},
				},
			},
			Selector:               "app == 'mongodb'",
			Types:                  []v3.PolicyType{v3.PolicyTypeEgress},
			ServiceAccountSelector: "",
		},
	}
	if createNetworkpolicies(cs, context.Background(), &np, "default") {
		fmt.Println("Successfully created a test network policy...")
	}

	// Get list of network policies in default namespace
	fmt.Println("Network policies in default namespace....")
	for idx, np := range getNetworkpolicies(cs, context.Background(), "default") {
		fmt.Printf("-------- Network Policy #%d --------\n", idx+1)
		fmt.Printf("%s\n\n", np)
	}
}

func getNetworkpolicies(cs *clientset.Clientset, ctx context.Context, namespace string) []string {
	networkpolicies := []string{}
	list, err := cs.ProjectcalicoV3().NetworkPolicies(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, np := range list.Items {
		stringJson, _ := json.MarshalIndent(np, "", "\t")
		networkpolicies = append(networkpolicies, string(stringJson))
	}
	return networkpolicies
}

func getGlobalNetworkpolicies(cs *clientset.Clientset, ctx context.Context) []string {
	globalNetworkpolicies := []string{}
	list, err := cs.ProjectcalicoV3().GlobalNetworkPolicies().List(ctx, v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, np := range list.Items {
		stringJson, _ := json.MarshalIndent(np, "", "\t")
		globalNetworkpolicies = append(globalNetworkpolicies, string(stringJson))
	}
	return globalNetworkpolicies
}

func createNetworkpolicies(cs *clientset.Clientset, ctx context.Context, np *v3.NetworkPolicy, namespace string) bool {
	_, err := cs.ProjectcalicoV3().NetworkPolicies(namespace).Create(ctx, np, v1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating network policy: %v\n", err)
		return false
	}
	return true
}
