# kubernetes_learning
- Install calico apiserver and generate some resources: `/test_calico`
- Interact with calico API V3 using generated clientset: `https://github.com/projectcalico/api`

```
$ go run main.go
Successfully created a test network policy...
Network policies in default namespace....
-------- Network Policy #1 --------
{
        "kind": "NetworkPolicy",
        "apiVersion": "projectcalico.org/v3",
        "metadata": {
                "name": "deny-all-ingress",
                "namespace": "default",
                "uid": "f14af472-6e3b-43b4-a515-52917e7f573f",
                "resourceVersion": "2642",
                "creationTimestamp": "2023-02-09T02:52:38Z",
                "annotations": {
                        "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"projectcalico.org/v3\",\"kind\":\"NetworkPolicy\",\"metadata\":{\"annotations\":{},\"name\":\"deny-all-ingress\",\"namespace\":\"default\"},\"spec\":{\"ingress\":[{\"action\":\"Deny\",\"destination\":{\"ports\":[8080,80]},\"protocol\":\"TCP\",\"source\":{\"selector\":\"app == 'mongodb'\"}}],\"selector\":\"app == 'nginx'\",\"types\":[\"Ingress\"]}}\n"
                },
                "managedFields": [
                        {
                                "manager": "kubectl-client-side-apply",
                                "operation": "Update",
                                "apiVersion": "projectcalico.org/v3",
                                "time": "2023-02-09T02:52:38Z",
                                "fieldsType": "FieldsV1",
                                "fieldsV1": {
                                        "f:metadata": {
                                                "f:annotations": {
                                                        ".": {},
                                                        "f:kubectl.kubernetes.io/last-applied-configuration": {}
                                                }
                                        },
                                        "f:spec": {
                                                "f:ingress": {},
                                                "f:selector": {},
                                                "f:types": {}
                                        }
                                }
                        }
                ]
        },
        "spec": {
                "ingress": [
                        {
                                "action": "Deny",
                                "protocol": "TCP",
                                "source": {
                                        "selector": "app == 'mongodb'"
                                },
                                "destination": {
                                        "ports": [
                                                8080,
                                                80
                                        ]
                                }
                        }
                ],
                "selector": "app == 'nginx'",
                "types": [
                        "Ingress"
                ]
        }
}

-------- Network Policy #2 --------
{
        "kind": "NetworkPolicy",
        "apiVersion": "projectcalico.org/v3",
        "metadata": {
                "name": "test-policy-deny-all-egress",
                "namespace": "default",
                "uid": "703db325-11da-46f2-92ab-c95d0da64c93",
                "resourceVersion": "13864",
                "creationTimestamp": "2023-03-25T21:54:04Z",
                "managedFields": [
                        {
                                "manager": "main",
                                "operation": "Update",
                                "apiVersion": "projectcalico.org/v3",
                                "time": "2023-03-25T21:54:04Z",
                                "fieldsType": "FieldsV1",
                                "fieldsV1": {
                                        "f:spec": {
                                                "f:egress": {},
                                                "f:order": {},
                                                "f:selector": {},
                                                "f:types": {}
                                        }
                                }
                        }
                ]
        },
        "spec": {
                "order": 0,
                "egress": [
                        {
                                "action": "Deny",
                                "protocol": "TCP",
                                "source": {},
                                "destination": {
                                        "selector": "app == 'nginx'",
                                        "ports": [
                                                "custom-port"
                                        ]
                                }
                        }
                ],
                "selector": "app == 'mongodb'",
                "types": [
                        "Egress"
                ]
        }
}
```