

The idea is to setup kubernetes and spin a container on it locally

1. Create a docker image for a sample node application. Following is the app's docker file.
```
FROM node
ADD app.js /app.js
ENTRYPOINT ["node", "app.js"]
```

2. Setting up Kubernetes cluster on MacOS

By Minikube

1. To check if virtualization is supported on macOS, run the following command on your terminal.
   
   ```$xslt
        sysctl -a | grep machdep.cpu.features
   ```
   
   If you see VMX in the output, the VT-x feature is supported on your OS.
   
2.  Install Minikube
    
    ```$xslt
    brew cask install minikube
    ```   
    
3. To start Minikube
    
    ```$xslt
    minikube start
    ```    
 
4. Installing Kubernetes Client (kubectl)
   
   ```$xslt
    curl -LO https://storage.googleapis.com/kubernetes-release/release
    ➥ /$(curl -s https://storage.googleapis.com/kubernetes-release/release ➥ /stable.txt)/bin/darwin/amd64/kubectl
    ➥ && chmod +x kubectl
    ➥ && sudo mv kubectl /usr/local/bin/
    ```     
    
5. To verify your cluster is working
    
    ```$xslt
     kubectl cluster-info
    ```    
  
6. Deploying container with JSON or YAML
    
    ```$xslt
     kubectl run kubia --image=luksa/kubia --port=8080 --generator=run/v1
    ``` 
    
    --image=luksa/kubia [name of the image]
    --port=8080 [your app is listening on port 8080]
    --genrator [creates a Replication Controller instead of Deployment]
    
    Delete replication controller by following command:
    
    ```$xslt
    kubectl delete rc kubia
    ```
    
    By creating a service of type LoadBalancer, an external load balancer gets 
    created and you can connect the pod using load balancer's public IP.
    
   ```$xslt
    kubectl expose rc kubia --type=LoadBalancer --name kubia-http
   ``` 
   
   Minikube doesn’t support LoadBalancer services, so the service will never get an external IP.
   But you can access the service anyway through its external port.
   
   Access service running on Minikube by  following command.
   
   ```$xslt
    minikube service kubia-http
   ```
   
7.  INCREASING THE DESIRED REPLICA COUNT

    ```$xslt
     kubectl scale rc kubia --replicas=3
    ```  
    
    Describing a pod with kubectl describe
    
    ```$xslt
    kubectl describe pod <<pod_name>>
    ```
   
8. Creating pods with YAML descriptors with name kubia-manual.yaml

   ```yaml
      
     apiVersion: v1
     kind: Pod
     metadata:
       name: kubia-manual
     spec:
       containers:
       - image: rajat965ng/kubia:v1
         name: kubia
         ports:
         - containerPort: 8080
           protocol: TCP
  
   ```   
   
   Execute following command to execute descriptor
   
   ```$xslt
   kubectl create -f kubia-manual.yaml 
   ```
   
   RETRIEVING A POD’S LOG WITH KUBECTL LOGS
   
   ```$xslt
   kubectl logs kubia-manual
   ```   
   
   Retrieving logs for specific container in a multi-container pod.
   
   ```$xslt
    kubectl logs kubia-manual -c kubia
   ```
   
   FORWARDING A LOCAL NETWORK PORT TO A PORT IN THE POD
   
   ```$xslt
    kubectl port-forward kubia-manual 80880:8080
   ```
   
   To adding and editing the annotation in existing pod
   ```$xslt
    kubectl annotate pod kubia-manual someannotation="a quick brown fox jumps over the lazy dog"
   ```
   
   To select pos under specific namespace
 
   ```$xslt
    kubectl get pods -n custom-namespace
   ```
   
<p>
<h3>Replication and other controllers: deploying managed pods</h3><br>

Creating an HTTP-based liveness probe

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: kubia-liveness
spec:
  containers:
    - image: luksa/kubia-unhealthy
      name: kubia-liveness
      livenessProbe:
        httpGet:
          path: /
          port: 8080
        initialDelaySeconds: 15    
```

```kubectl get po kubia-liveness```

The RESTARTS column shows that the pod’s container has been restarted<br>


A YAML definition of a ReplicationController: kubia-rc.yaml

SCALING UP A REPLICATIONCONTROLLER<br>

One way:
```$xslt
 kubectl scale rc kubia --replicas=10
```

Second way:
```$xslt
 kubectl edit rc kubia
 
 When the text editor opens, find the spec.replicas field and change its value to 10
```

Deleting a ReplicationController <br>

```$xslt
 kubectl delete rc kubia --cascade=false
```

When deleting a ReplicationController with kubectl delete, you can keep its pods running by passing the --cascade=false option to the command.

A YAML definition of a ReplicaSet: kubia-replicaset.yaml

A matchExpressions selector: kubia-replicaset-matchexpressions.yaml

A YAML for a DaemonSet: kubia-daemonset.yaml

```$xslt
 kubectl get ds
```
you’ll need to list the nodes first, because you’ll need to know the node’s name when labeling it:

```$xslt
 kubectl get node
```

add the disk=ssd label to one of your nodes like this:

```$xslt
 kubectl label node minikube disk=ssd
```

A YAML definition of a Job: kubia-batch-job.yaml

YAML for a CronJob resource: kubia-cron-job.yaml

</p>   