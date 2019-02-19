<h2>Creating EKS using Terraform</h2>
<br>
<h3>Points to remember</h3>
<ul>
<li>Enable ingress in master security group</li>
<li>Copy Kube config content generated within console, into ~/.kube/config file or create a file inside ~/.kube and point the path of that file to the environment variable KUBECONFIG</li>
<li>Copy authentication config map content generated within console in config-map.yaml and execute <b>kubectl apply -f config-map.yaml</b></li>
<li>Invoke <b>kubectl get nodes -watch</b> to check if nodes are up</li>
<li>Post making deployments and creating services execute <b>kubectl get services -o wide</b> to view external IPs</li>
</ul>
<br>
<h3>For deploying kubectl dashboard execute following file in <b>kubectl apply</b></h3>
<ul>
<li>dashboard-account-rbac.yaml</li>
<li>deploy-dashboard.yaml</li>
<li>deploy-heapster.yaml</li>
<li>deploy-dashboard.yaml</li>
<li>deploy-influxdb.yaml</li>
<li>deploy-heapster-rbac.yaml</li>
<li>admin-service-account.yaml</li>
</ul>

<p>
<b>Execute the following commands to view the svc,IPs etc.</b>
<li>
<ul>kubectl get all --namespace kube-system --selector=k8s-app=kubernetes-dashboard</ul>
<ul>kubectl get all --namespace kube-system --selector=task=monitoring</ul>
</li>
</p

<p>
<h3>Execute following to generate access <i>Tokens</i></h3>
<h4>kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep eks-course-admin | awk '{print $1}')</h4>
</p>
<br>

# Setup sample guestbook app
* based on https://github.com/kubernetes/examples/tree/master/guestbook

## Redis master
deploy the master Redis pod and a _service_ on top of it:
```
kubectl apply -f redis-master.yaml
kubectl get pods
kubectl get services
```

## Redis slaves
deploy the Redis slave pods and a _service_ on top of it:
```
kubectl apply -f redis-slaves.yaml
kubectl get pods
kubectl get services
```

## Frontend app
deploy the PHP Frontend pods and a _service_ of type **LoadBalancer** on top of it, to expose the loadbalanced service to the public via ELB:
```
kubectl apply -f frontend.yaml
```
some checks:
```
kubectl get pods
kubectl get pods -l app=guestbook
kubectl get pods -l app=guestbook -l tier=frontend
```
check AWS mgm console for the ELB which has been created !!!

## Access from outside the cluster
grab the public DNS of the frontend service LoadBalancer (ELB):
```
kubectl describe service frontend
```
copy the name and paste it into your browser !!!


## kubectl cmds for scaling pods
scaling a deployment:
```
kubectl scale --replicas <number-of-replicas> deployment <name-of-deployment>
```
e.g. set no of replicas for _frontend_ service to _5_ :
```
kubectl scale --replicas 5 deployment frontend
```

## commands to check state
* get all pods incl additional info like e.g. k8s worker node the pod is running on
```
kubectl get pods -o wide
```
* state of service(s)
```
kubectl get services
```
details of a particular service:
```
kubectl describe service <servicename>
```
