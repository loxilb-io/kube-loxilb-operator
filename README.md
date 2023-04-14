# kube-loxilb-operator
kube-loxilb-operator automates the deployment and management kube-loxilb. 


## Description
kube-loxilb is loxilb's implementation of kubernetes service load-balancer spec which includes support for load-balancer class, advanced IPAM (shared or exclusive)   
The github page of kube-loxilb: <https://github.com/loxilb-io/kube-loxilb>

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

If you want to know more information about:   
  - How to kube-loxilb: <https://loxilb-io.github.io/loxilbdocs/kube-loxilb/>   
  - LoxiLB documentation: <https://loxilb-io.github.io/loxilbdocs/>

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/kube-loxilb-operator:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/kube-loxilb-operator:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

