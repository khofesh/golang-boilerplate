## TLS

```shell
mkdir ~/.proglog
make gencert
make test
```

## observability

```shell
cd internal/server
go test -v -debug=true
```

result

```shell
[fahmad@ryzen server]$ go test -v -debug=true
=== RUN   TestServer
=== RUN   TestServer/produce/consume_a_message_to/from_the_log_succeeeds
    server_test.go:60: metrics log file: /tmp/metrics-779215398.log
    server_test.go:60: traces log file: /tmp/traces-443977549.log

```

## raft

turns out the following line is still needed, otherwise the test will fail.

```shell
go mod edit -replace github.com/hashicorp/raft-boltdb=github.com/travisjeffery/raft-boltdb@v1.0.0
```

## resolver.go and resolver_test.go

`Endpoint` is deprecated, so in `resolver.go` I replaced with

```go
r.resolverConn, err = grpc.Dial(target.URL.Host, dialOpts...)
if err != nil {
	return nil, err
}
```

and in `resolver_test.go`

```go
_, err = r.Build(
	resolver.Target{
		URL: url.URL{
			Host: l.Addr().String(),
		},
	},
	conn,
	opts,
)
require.NoError(t, err)
```

I'm using the `Host` param

## minikube

I'm using minikube here, because I'm used to it.

some problem I found (Fedora 35, BTRFS, and LUKS):

```shell
❌  Problems detected in kubelet:
    May 08 14:50:12 minikube kubelet[70792]: E0508 14:50:12.639233   70792 kubelet.go:1423] "Failed to start ContainerManager" err="failed to get rootfs info: failed to get device for dir \"/var/lib/kubelet\": could not find device with major: 0, minor: 34 in cached partitions map"
    May 08 14:50:13 minikube kubelet[70990]: E0508 14:50:13.680934   70990 kubelet.go:1423] "Failed to start ContainerManager" err="failed to get rootfs info: failed to get device for dir \"/var/lib/kubelet\": could not find device with major: 0, minor: 34 in cached partitions map"
    May 08 14:50:14 minikube kubelet[71198]: E0508 14:50:14.887413   71198 kubelet.go:1423] "Failed to start ContainerManager" err="failed to get rootfs info: failed to get device for dir \"/var/lib/kubelet\": could not find device with major: 0, minor: 34 in cached partitions map"
    May 08 14:50:15 minikube kubelet[71406]: E0508 14:50:15.889173   71406 kubelet.go:1423] "Failed to start ContainerManager" err="failed to get rootfs info: failed to get device for dir \"/var/lib/kubelet\": could not find device with major: 0, minor: 34 in cached partitions map"
    May 08 14:50:16 minikube kubelet[71623]: E0508 14:50:16.941373   71623 kubelet.go:1423] "Failed to start ContainerManager" err="failed to get rootfs info: failed to get device for dir \"/var/lib/kubelet\": could not find device with major: 0, minor: 34 in cached partitions map"
```

workaround:

```shell
minikube start --feature-gates="LocalStorageCapacityIsolation=false"
```

https://github.com/kubernetes/minikube/issues/10182

or this way:

- update to the latest version

```shell
[fahmad@ryzen ~]$ curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-latest.x86_64.rpm
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 15.6M  100 15.6M    0     0   929k      0  0:00:17  0:00:17 --:--:--  695k
[fahmad@ryzen ~]$  sudo rpm -Uvh minikube-latest.x86_64.rpm
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:minikube-1.25.2-0                ################################# [ 50%]
Cleaning up / removing...
   2:minikube-1.23.2-0                ################################# [100%]
```

start minikube

```shell
minikube start --extra-config=kubelet.cgroup-driver=systemd -v=5 --kubernetes-version=1.24.0
```

result:

```shell
[fahmad@ryzen ~]$ minikube start --extra-config=kubelet.cgroup-driver=systemd -v=5 --kubernetes-version=1.24.0
😄  minikube v1.25.2 on Fedora 35
❗  Specified Kubernetes version 1.24.0 is newer than the newest supported version: v1.23.4-rc.0
✨  Automatically selected the docker driver. Other choices: kvm2, none, ssh
👍  Starting control plane node minikube in cluster minikube
🚜  Pulling base image ...
    > gcr.io/k8s-minikube/kicbase: 379.06 MiB / 379.06 MiB  100.00% 1.65 MiB p/
🔥  Creating docker container (CPUs=2, Memory=3900MB) ...
🐳  Preparing Kubernetes v1.24.0 on Docker 20.10.12 ...
    ▪ kubelet.cgroup-driver=systemd
    ▪ kubelet.housekeeping-interval=5m
    > kubelet.sha256: 64 B / 64 B [--------------------------] 100.00% ? p/s 0s
    > kubectl.sha256: 64 B / 64 B [--------------------------] 100.00% ? p/s 0s
    > kubeadm.sha256: 64 B / 64 B [--------------------------] 100.00% ? p/s 0s
    > kubeadm: 42.31 MiB / 42.31 MiB [---------------] 100.00% 1.34 MiB p/s 32s
    > kubectl: 43.59 MiB / 43.59 MiB [---------------] 100.00% 1.21 MiB p/s 36s
    > kubelet: 110.95 MiB / 110.95 MiB [------------] 100.00% 1.82 MiB p/s 1m1s
    ▪ Generating certificates and keys ...
    ▪ Booting up control plane ...
    ▪ Configuring RBAC rules ...
🔎  Verifying Kubernetes components...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
🌟  Enabled addons: default-storageclass, storage-provisioner
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default

```

https://github.com/kubernetes/minikube/issues/7923#issuecomment-997442546

https://kubernetes.io/releases/

https://minikube.sigs.k8s.io/docs/commands/start/

https://minikube.sigs.k8s.io/docs/start/

**I prefer the latter**

## build the docker image

```shell
minikube start
eval $(minikube -p minikube docker-env)
make build-docker
```

## helm

```shell
sudo snap install helm --classic
```

https://helm.sh/docs/intro/install/

install nginx

```shell
helm install my-nginx bitnami/nginx --set service.type=NodePort
NAME: my-nginx
LAST DEPLOYED: Tue May 10 09:05:08 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: nginx
CHART VERSION: 10.2.1
APP VERSION: 1.21.6

** Please be patient while the chart is being deployed **

NGINX can be accessed through the following DNS name from within your cluster:

    my-nginx.default.svc.cluster.local (port 80)

To access NGINX from outside the cluster, follow the steps below:

1. Get the NGINX URL by running these commands:

    export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services my-nginx)
    export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
    echo "http://${NODE_IP}:${NODE_PORT}"
```

confirm if the nginx is running:

```shell
export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services my-nginx)

export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")

echo "http://${NODE_IP}:${NODE_PORT}"

curl "http://${NODE_IP}:${NODE_PORT}"
```

https://stackoverflow.com/a/65653598
https://docs.bitnami.com/kubernetes/get-started-kubernetes/

## install proglog helm chart

in vscode disable `insertSpaces`:

```json
"[yaml]": {
  "editor.insertSpaces": false,
  "editor.tabSize": 2,
  "editor.autoIndent": false,
  "gitlens.codeLens.scopes": ["document"],
  "editor.quickSuggestions": {
    "other": true,
    "comments": false,
    "strings": true
  }
}
```

deploy

```shell
[fahmad@ryzen proglog]$ helm install proglog deploy/proglog
NAME: proglog
LAST DEPLOYED: Wed May 11 11:29:13 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```
