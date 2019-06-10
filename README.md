

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
<p>
<h3>Services: enabling clients to discover and talk to pods</h3>

A definition of a service: kubia-svc.yaml

REMOTELY EXECUTING COMMANDS IN RUNNING CONTAINERS

```$xslt
  kubectl exec kubia-7nog1 -- curl -s http://10.111.249.153
```

In kubia-svc.yaml:
 Port 80 is mapped to the container’s port called http.
 Port 443 is mapped to the container’s port, whose name is https.
     

Why the double dash?

The double dash (--) in the command signals the end of command options for kubectl.
Everything after the double dash is the command that should be executed inside the pod.


A example of a service with ClientIP session affinity configured: kubia-svc-session-affinity.yaml


a frontend pod can connect to the backend- database service by opening a connection to the following FQDN:
backend-database.default.svc.cluster.local
backend-database corresponds to the service name, default stands for the name- space the service is defined in, and svc.cluster.local is a configurable cluster domain suffix used in all cluster local service names.


```$xslt
kubectl exec -it kubia-3inly bash

curl http://kubia.default.svc.cluster.local

curl http://kubia.default

curl http://kubia
```

Exposing services to external clients

make a service accessible externally:

Setting a service of type Node port - Every node on a cluster opens up a port on itself. 
The service is not only accessible on cluster IP and port but also through dedicated port on each node.

A NodePort service definition: kubia-svc-nodeport.yaml. The service accessible within the pod using cluster IP.

```$xslt
kubectl exec -it kubia-64v77 -- curl -s http://10.109.196.71
```

Tell kubectl to print out only the node IP instead of the whole service definition:

```$xslt
kubectl get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="ExternalIP")].address}'
```

Setting a service of type LoadBalancer - The service is accessible through load balancer of provisioned cloud provider of kubernetes. 
The load balancer accept all traffic and redirect it to node ports of the cluster. The client can access the service through load balancer's IP.

A LoadBalancer-type service: kubia-svc-loadbalancer.yaml




Creating an Ingress resource - exposing multiple services through a static IP. It operates on HTTP (network layer 7).

An Ingress resource definition: kubia-ingress.yaml

CREATING A TLS CERTIFICATE FOR THE INGRESS

<p>
if the pod runs a web server, it can accept only HTTP traffic and let the Ingress controller take care of everything related to TLS. 
To enable the controller to do that, you need to attach a certificate and a private key to the Ingress. 
The two need to be stored in a Kubernetes resource called a Secret, which is then referenced in the Ingress manifest.
</p>

First, you need to create the private key and certificate:

```
$ openssl genrsa -out tls.key 2048
$ openssl req -new -x509 -key tls.key -out tls.cert -days 360 -subj  /CN=kubia.example.com
```

Then you create the Secret from the two files like this:

```$xslt
$ kubectl create secret tls tls-secret --cert=tls.cert --key=tls.key
```

Ingress handling TLS traffic: kubia-ingress-tls.yaml

Make an entry of <i>kubia.example.com</i> in /etc/hosts file.

Use HTTPS to access TLS enabled ingress.

```$xslt
curl -k -v https://kubia.example.com
```

Introducing readiness probes.

The readiness probe is invoked periodically and determines whether the specific pod should receive client requests or not. 
When a container’s readiness probe returns success, it’s signaling that the container is ready to accept requests.

Following is the syntax use to validate the readiness probe.

```$xslt
spec: 
  containers: 
    - image: luksa/kubia
      name: kubia
      readinessProbe: 
        exec: 
          command: 
            - ls
            - /var/ready
```

The readiness probe will periodically perform the command ls /var/ready inside the container.
The ls command returns exit code zero if the file exists, or a non-zero exit code otherwise.
If the file exists, the readiness probe will succeed; otherwise, it will fail.

<h4>Using a headless service for discovering individual pods</h4>

A headless service: kubia-svc-headless.yaml

</p>  
<p>
<h3>Volumes: attaching disk storage to containers</h3>

a list of several of the available volume types:

emptyDir—A simple empty directory used for storing transient data.

hostPath—Used for mounting directories from the worker node’s filesystem into the pod.

gitRepo—A volume initialized by checking out the contents of a Git repository.

nfs—An NFS share mounted into the pod.

gcePersistentDisk (Google Compute Engine Persistent Disk), awsElastic-BlockStore (Amazon Web Services Elastic Block Store Volume), azureDisk (Microsoft Azure Disk Volume)—Used for mounting cloud provider-specific storage.

cinder, cephfs, iscsi, flocker, glusterfs, quobyte, rbd, flexVolume, vsphereVolume, photonPersistentDisk, scaleIO—Used for mounting other types of network storage.

configMap, secret, downwardAPI—Special types of volumes used to expose certain Kubernetes resources and cluster information to the pod.

persistentVolumeClaim—A way to use a pre- or dynamically provisioned persistent storage. 


A pod with two containers sharing the same volume: fortune-pod.yaml

```$xslt
kubectl port-forward fortune 8080:80
```

The emptyDir you used as the volume was created on the actual disk of the worker node hosting your pod, so its performance depends on the type of the node’s disks.

But you can tell Kubernetes to create the emptyDir on a tmpfs filesystem (in memory instead of on disk). To do this, set the emptyDir’s medium to Memory like this:

```$xslt
volumes:
  - name: html
    emptyDir:
      medium: Memory
```

A pod using a gitRepo volume: gitrepo-volume-pod.yaml


<h4>Using persistent storage</h4>

When an application running in a pod needs to persist data to disk and have that same data available even when the pod is rescheduled to another node, 
you can’t use any of the volume types we’ve mentioned so far. 
Because this data needs to be accessible from any cluster node, it must be stored on some type of network-attached storage (NAS).

A pod using a gcePersistentDisk volume: mongodb-pod-gcepd.yaml

If you’re using Minikube, you can’t use a GCE Persistent Disk, but you can deploy mongodb-pod-hostpath.yaml, which uses a hostPath volume instead of a GCE PD.

This is against the basic idea of Kubernetes, which aims to hide the actual infrastructure from both the application and its developer,
leaving them free from worrying about the specifics of the infrastructure and making apps portable across a wide array of cloud providers and on-premises 
data centers.

To enable apps to request storage in a Kubernetes cluster without having to deal with infrastructure specifics, two new resources were introduced. 
They are PersistentVolumes and PersistentVolumeClaims.

PersistentVolumes are provisioned by cluster admins and consumed by pods through PersistentVolumeClaims.

A gcePersistentDisk PersistentVolume: mongodb-pv-gcepd.yaml

CREATING A PERSISTENT VOLUME CLAIM

A PersistentVolumeClaim: mongodb-pvc.yaml

<h4>Dynamic provisioning of PersistentVolumes</h4>

The cluster admin, instead of creating PersistentVolumes, can deploy a Persistent- Volume provisioner and define one or more StorageClass objects to 
let users choose what type of PersistentVolume they want. The users can refer to the StorageClass in their PersistentVolumeClaims and the provisioner 
will take that into account when provisioning the persistent storage.

A StorageClass definition: storageclass-fast-gcepd.yaml

</p>
<p>
<h3>ConfigMaps and Secrets: configuring applications</h3>

The Kubernetes resource for storing configuration data is called a ConfigMap.

Any sensitive information includes credentials, private encryption keys, and similar data that needs to be kept secure. 
This type of information needs to be handled with special care, which is why Kubernetes offers another type of first-class object called a Secret.

Overriding the command and arguments in Kubernetes.
To do that, you set the properties command and args in the container specification, as shown in the following listing.

```$xslt
kind: Pod
spec:
  containers:
  - image: some/image
    command: ["/bin/command"]
    args: ["arg1", "arg2", "arg3"]
```

Setting environment variables for a container

```$xslt
kind: Pod
spec:
 containers:
 - image: luksa/fortune:env
   env:
   - name: INTERVAL
value: "30"
   name: html-generator
```

Referring to other environment variables in a variable’s value. Referring to an environment variable inside another one.

```$xslt
env:
- name: FIRST_VAR
  value: "foo"
- name: SECOND_VAR
  value: "$(FIRST_VAR)bar"
```


Values effectively hardcoded in the pod definition means you need to have separate pod definitions for your production and your development pods. 
To reuse the same pod definition in multiple environments, it makes sense to decouple the configuration from the pod descriptor.
You can do that using a ConfigMap resource and using it as a source for environment variable values using the valueFrom instead of the value field.

<h4>Introducing ConfigMaps</h4>

The contents of the map are instead passed to containers as either environment variables or as files in a volume.

Creating a ConfigMap

```$xslt
kubectl create configmap fortune-config --from-literal=sleep-interval=25

kubectl create configmap myconfigmap --from-literal=foo=bar --from-literal=bar=baz --from-literal=one=two

kubectl create configmap my-config
 --from-file=foo.json
 --from-file=bar=foobar.conf
 --from-file=config-opts/
 --from-literal=some=thing
```
Pod with env var from a config map: fortune-pod-env-configmap.yaml

Passing all entries of a ConfigMap as environment variables at once

```$xslt
spec:
  containers:
  - image: some-image
    envFrom:
    - prefix: CONFIG_
      configMapRef:
        name: my-config-map
```

Using a configMap volume to expose ConfigMap entries as files

An Nginx config with enabled gzip compression: my-nginx-config.conf

A pod with ConfigMap entries mounted as files: fortune-pod-configmap-volume.yaml

Seeing if nginx responses have compression enabled

```$xslt
kubectl port-forward fortune-configmap-volume 8080:80 &

curl -H "Accept-Encoding: gzip" -I localhost:8080
```

EXAMINING THE MOUNTED CONFIGMAP VOLUME’S CONTENTS

```$xslt
kubectl exec fortune-configmap-volume -c web-server ls /etc/nginx/conf.d
```

EXPOSING CERTAIN CONFIGMAP ENTRIES IN THE VOLUME

```$xslt
volumes:
- name: config
  configMap:
    name: fortune-config
    items:
    - key: my-nginx-config.conf
      path: gzip.conf
```

The directory then only contains the files from the mounted filesystem, whereas the original files in that directory are inaccessible 
for as long as the filesystem is mounted. This would most likely break the whole container, because all of the original files 
that should be in the /etc directory would no longer be there.

MOUNTING INDIVIDUAL CONFIGMAP ENTRIES AS FILES WITHOUT HIDING OTHER FILES IN THE DIRECTORY.

use the 'subPath' property to mount it there without affecting any other files in that directory.

```$xslt
spec:
  containers:
  - image: some/image
    volumeMounts:
    - name: myvolume
      mountPath: /etc/someconfig.conf  #You’re mounting into a file, not a directory.
      subPath: myconfig.conf           #Instead of mounting the whole volume, you’re only mounting the myconfig.conf entry.
```

SETTING THE FILE PERMISSIONS FOR FILES IN A CONFIGMAP VOLUME

```$xslt
volumes:
- name: config
  configMap:
    name: fortune-config
    defaultMode: "6600"  #By default, the permissions on all files in a configMap volume are set to 644.
```

EDITING A CONFIGMAP

```$xslt
$ kubectl edit configmap fortune-config
```

SIGNALING NGINX TO RELOAD THE CONFIG

```$xslt
kubectl exec fortune-configmap-volume -c web-server -- nginx -s reload
```

<h4>Using Secrets to pass sensitive data to containers</h4>

Creating a Secret

```$xslt
openssl genrsa -out https.key 2048
openssl req -new -x509 -key https.key -out https.cert -days 3650 -subj /CN=www.kubia-example.com

echo bar > foo

```
create an additional dummy file called foo and make it contain the string bar.

Now you can use kubectl create secret to create a Secret from the three files:

```$xslt
kubectl create secret generic fortune-https --from-file=https.key --from-file=https.cert --from-file=foo

kubectl get secret fortune-https -o yaml
```
contents of a Secret’s entries are shown as Base64-encoded strings


Adding plain text entries to a Secret using the stringData field

```$xslt
kind: Secret
apiVersion: v1
stringData:
  foo: plain text
data:
  https.cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCekNDQ...
  https.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcE...
  
```

YAML definition of the fortune-https pod: fortune-pod-https.yaml

MODIFYING THE FORTUNE-CONFIG CONFIGMAP TO ENABLE HTTPS

```$xslt
kubectl edit configmap fortune-config


data:
  my-nginx-config.conf: |-
    server {
      listen        80;
      listen        443 ssl;
      server_name   www.kubia-example.com;



      ssl_certificate     certs/https.cert;
      ssl_certificate_key certs/https.key;
      ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
      ssl_ciphers         HIGH:!aNULL:!MD5;


      location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
      }
    }

```

CREATING A SECRET FOR AUTHENTICATING WITH A DOCKER REGISTRY

kubectl create secret docker-registry mydockerhubsecret \
  --docker-username=myusername --docker-password=mypassword \
  --docker-email=my.email@provider.com

</p>