

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

<p>
<h3>Accessing pod metadata and other resources from applications</h3>

<h4>Downward API used in environment variables: downward-api-env.yaml</h4>

Passing metadata through files in a downwardAPI volume

Pod with a downwardAPI volume: downward-api-volume.yaml

Talking to the API server from within a pod

how to talk to it from within a pod, where you (usually) don’t have kubectl. 
Therefore, to talk to the API server from inside a pod, you need to take care of three things:

 Find the location of the API server.
 Make sure you’re talking to the API server and not something impersonating it. 
 Authenticate with the server; otherwise it won’t let you see or do anything.

First, load the token into an environment variable:

```
root@curl:/# TOKEN=$(cat /var/run/secrets/kubernetes.io/ serviceaccount/token)
```

Getting a proper response from the API server

```
root@curl:/# curl -H "Authorization: Bearer $TOKEN" https://kubernetes
```

Listing pods in the pod’s own namespace
```
root@curl:/# NS=$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace)
root@curl:/# curl -H "Authorization: Bearer $TOKEN" https://kubernetes/api/v1/namespaces/$NS/pods
```

Simplifying API server communication with ambassador containers

Instead of talking to the API server directly, the app in the main container can connect to the ambassador through HTTP (instead of HTTPS) 
and let the ambassador proxy handle the HTTPS connection to the API server, taking care of security transparently.
It does this by using the files from the default token’s secret volume.

RUNNING THE CURL POD WITH AN ADDITIONAL AMBASSADOR CONTAINER

A pod with an ambassador container: curl-with-ambassador.yaml

</p>

<p>
<h3>Updating applications running in pods</h3>

Updating applications running in pods
- Deleting old pods and replacing them with new ones
- Spinning up new pods and then deleting the old ones
    change the Service’s label selector and have the Service switch over to the new pods. This is called a blue-green deployment.
    After switching over, and once you’re sure the new version functions correctly, you’re free to delete the old pods by deleting the old ReplicationController.
    You can change a Service’s pod selector with the <i>kubectl set selector</i> command.
- PERFORMING A ROLLING UPDATE
 
 ```  
    #this is deprecated now. Use rollout instead of rolling-update.
    kubectl rolling-update kubia-v1 kubia-v2 --image=luksa/kubia:v2 
 ```

Understanding why kubectl rolling-update is now obsolete

 - it’s perfectly fine for the scheduler to assign a node to my pods after I create them, but Kubernetes modifying the labels of 
    my pods and the label selectors of my ReplicationControllers is something that I don’t expect.
 - what if you lost network connectivity while kubectl was performing the update? The update pro- cess would be interrupted mid-way.
    Pods and ReplicationControllers would end up in an intermediate state.      


CREATING A DEPLOYMENT MANIFEST

A Deployment definition: kubia-deployment-v1.yaml

```
kubectl create -f kubia-deployment-v1.yaml --record
```

You’ll deploy the new version by changing the image in the Deployment specification again:

```
kubectl set image deployment kubia nodejs=luksa/kubia:v3

kubectl rollout status deployment kubia
```

Luckily, Deployments make it easy to roll back to the previously deployed version by telling Kubernetes to undo the last rollout of a Deployment:

```
kubectl rollout undo deployment kubia
```

DISPLAYING A DEPLOYMENT’S ROLLOUT HISTORY

```
kubectl rollout history deployment kubia
```

Remember the --record command-line option you used when creating the Deploy- ment? Without it, the CHANGE-CAUSE column in the revision history would be empty,
 making it much harder to figure out what’s behind each revision.


ROLLING BACK TO A SPECIFIC DEPLOYMENT REVISION

```
kubectl rollout undo deployment kubia --to-revision=1
```

<h4>Controlling the rate of the rollout</h4>

INTRODUCING THE MAXSURGE AND MAXUNAVAILABLE PROPERTIES OF THE ROLLING UPDATE STRATEGY

```
    spec:
      minReadySeconds: 10
      strategy:
        rollingUpdate:
          maxSurge: 1
          maxUnavailable: 0
        type: RollingUpdate

```    

maxSurge : If the desired replica count is set to four, there will never be more than five pod instances running at the same time during an update.

maxUnavailable : Determines how many pod instances can be unavailable relative to the desired replica count during the update.

minReadySeconds: Use minReadySeconds and readiness probes to have the rollout of a faulty version blocked automatically.
    
</p>
<p>
<h3>StatefulSets: deploying replicated stateful applications</h3>

Replicating stateful pods
    If the pod template includes a volume, which refers to a specific PersistentVolumeClaim, all replicas of the ReplicaSet will 
    use the exact same PersistentVolumeClaim and therefore the same PersistentVolume bound by the claim
    Because the reference to the claim is in the pod template, which is used to stamp out multiple pod replicas, 
    you can’t make each replica use its own separate Persistent- VolumeClaim.

All pods from same ReplicaSet always use the same PersistentVolumeClaim and PersistentVolume.


Running multiple replicas with separate storage for each
    - you maybe use a single ReplicaSet and have each pod instance keep its own persistent state, even though they’re all using the same storage volume.
    - A trick you can use is to have all pods use the same PersistentVolume, but then have a separate file directory inside that volume for each pod. 

Providing a stable identity for each pod
    provide a stable network address for cluster members by creating a dedicated Kubernetes Service for each individual member. 
    Because service IPs are stable, you can then point to each member through its service IP (rather than the pod IP) in the configuration.
    The solution is not only ugly, but it still doesn’t solve everything. 
    The individual pods can’t know which Service they are exposed through (and thus can’t know their stable IP), 
    so they can’t self-register in other pods using that IP. 
    The proper clean and simple way of running these special types of applications in Kubernetes is through a StatefulSet.



Understanding StatefulSets
    A StatefulSet makes sure pods are rescheduled in such a way that they retain their identity and state.
    It also allows you to easily scale the number of pets up and down.
    pods created by the StatefulSet aren’t exact replicas of each other. Each can have its own set of volume.  
    you can reach the pod through its fully qualified domain name, which is 
    
 ```
  podname.svcname.namespacename.svc.cluster.local.
 ```  
  When a pod instance managed by a StatefulSet disappears. the replacement pod gets the same name and hostname as the pod that has disappeared. 

SCALING A STATEFULSET
    StatefulSets also never permit scale-down operations if any of the instances are unhealthy. 
    If an instance is unhealthy, and you scale down by one at the same time.
    
    
Providing stable dedicated storage to each stateful instance
    Because PersistentVolumeClaims map to PersistentVolumes one-to-one, each pod of a StatefulSet needs to reference a different PersistentVolumeClaim to have 
    its own separate PersistentVolume. 
    Because all pod instances are stamped from the same pod template, how can they each refer to a different PersistentVolumeClaim?    
    The StatefulSet has to create the PersistentVolumeClaims as well, the same way it’s creating the pods. The PersistentVolumes for the claims can
    either be provisioned up-front by an administrator or just in time through dynamic provisioning of PersistentVolumes
    
    
    Scaling up a StatefulSet by one creates two or more API objects (the pod and one or more PersistentVolumeClaims referenced by the pod). Scaling down, however, deletes only the pod, leaving the claims alone. The reason for this is obvious, if you consider what happens when a claim is deleted. After a claim is deleted, the PersistentVolume it was bound to gets recycled or deleted and its contents are lost.
    Because stateful pods are meant to run stateful applications, which implies that the data they store in the volume is important, deleting the claim on scale-down of a Stateful- Set could be catastrophic—especially since triggering a scale-down is as simple as decreasing the replicas field of the StatefulSet. For this reason, you’re required to delete PersistentVolumeClaims manually to release the underlying PersistentVolume.

Deploying the app through a StatefulSet
    To deploy your app, you’ll need to create two (or three) different types of objects:
     PersistentVolumes for storing your data files (you’ll need to create these only if the cluster doesn’t support dynamic provisioning of PersistentVolumes).
     A governing Service required by the StatefulSet. 
     The StatefulSet itself.    

CREATING THE PERSISTENT VOLUMES
    If you’re using Google Kubernetes Engine, you’ll first need to create the actual GCE Persistent Disks like this:
            $ gcloud compute disks create --size=1GiB --zone us-central1-a pv-a
            $ gcloud compute disks create --size=1GiB --zone us-central1-a pv-b
            $ gcloud compute disks create --size=1GiB --zone us-central1-a pv-c

create the PersistentVolumes from the persistent-volumes-gcepd.yaml file,
CREATING THE GOVERNING SERVICE, Headless service to be used in the StatefulSet: kubia-service-headless.yaml
CREATING THE STATEFULSET MANIFEST, StatefulSet manifest: kubia-statefulset.yaml


<h4>Discovering peers in a StatefulSet</h4>
  
  SRV records are used to point to hostnames and ports of servers providing a specific service. Kubernetes creates SRV records to point to the hostnames of the pods
   backing a headless service.
  
  You’re going to list the SRV records for your stateful pods by running the dig DNS lookup tool inside a new temporary pod.
  
  ```
  kubectl run -it srvlookup --image=tutum/dnsutils --rm --restart=Never -- dig SRV kubia.default.svc.cluster.local
  ```
  
  The pod runs a single container from the tutum/dnsutils image and runs the following command:
  
  ```
  dig SRV kubia.default.svc.cluster.local
  ```  
  
</p>
<p>
<h3>Understanding Kubernetes internals</h3>

Checking the status of the Control Plane components

```
kubectl get componentstatuses
```

Kubernetes components running as pods

```
kubectl get po -o custom-columns=POD:metadata.name,NODE:spec.nodeName  --sort-by spec.nodeName -n kube-system
```

How Kubernetes uses etcd

```
Kubernetes uses etcd, which is a fast, distributed, and consistent key-value store. Because it’s distributed, you can run more than
one etcd instance to provide both high availability and better performance.
Component that talks to etcd directly is the Kubernetes API server. Other components read and write data to etcd indirectly through the API server.

Few benefits:
Robust optimistic locking system as well as validation.
Abstracting away the actual storage mechanism from all the other components, it’s much simpler to replace it in the future.
```
HOW RESOURCES ARE STORED IN ETCD

```
Kubernetes stores all its data in etcd under /registry.
```

ENSURING CONSISTENCY WHEN ETCD IS CLUSTERED

```
For ensuring high availability, you’ll usually run more than a single instance of etcd. Multiple etcd instances will need to remain consistent. 
Such a distributed system needs to reach a consensus on what the actual state is. etcd uses the RAFT consensus algorithm to achieve this, 
which ensures that at any given moment, each node’s state is either what the majority of the nodes agrees is the current state or is one of 
the previously agreed upon states.

Clients connecting to different nodes of an etcd cluster will either see the actual current state or one of the states from the past.
```

WHY THE NUMBER OF ETCD INSTANCES SHOULD BE AN ODD NUMBER

```
Having two instances requires both instances to be present to have a majority. If either of them fails, the etcd cluster can’t transition to 
a new state because no majority exists. Having two instances is worse than having only a single instance. By having two, the chance of the whole 
cluster fail- ing has increased by 100%, compared to that of a single-node cluster failing.
```

What the API server does

```
When creating a resource from a JSON file, for example, kubectl posts the file’s contents to the API server through an HTTP POST request.

First, the API server needs to authenticate the client sending the request.
Depending on the authentication method, the user can be extracted from the client’s certificate or an HTTP header, such as Authorization.
The plugin extracts the client’s username, user ID, and groups the user belongs to. This data is then used in the next stage, which is authorization.

VALIDATING AND/OR MODIFYING THE RESOURCE IN THE REQUEST WITH ADMISSION CONTROL PLUGINS.

VALIDATING THE RESOURCE AND STORING IT PERSISTENTLY.

```

Understanding the Scheduler

```
All the Scheduler does is update the pod definition through the API server. The API server then notifies the Kubelet that the pod has been scheduled.
 As soon as the Kubelet on the target node sees the pod has been scheduled to its node, it creates and runs the pod’s containers.
 
 he actual task of selecting the best node for the pod.
 
 UNDERSTANDING THE DEFAULT SCHEDULING ALGORITHM
 
 Filtering the list of all nodes to obtain a list of acceptable nodes the pod can be scheduled to.
 
    1. FINDING ACCEPTABLE NODES
        Can the node fulfill the pod’s requests for hardware resources?
        Is the node running out of resources (is it reporting a memory or a disk pres- sure condition)?
        If the pod requests to be scheduled to a specific node (by name), is this the node?
        Does the node have a label that matches the node selector in the pod specification (if one is defined)?
        If the pod requests to be bound to a specific host port is that port already taken on this node or not?
        If the pod requests a certain type of volume, can this volume be mounted for this pod on this node, or is another pod on the node already using
          the same volume?
        Does the pod tolerate the taints of the node? 
        Does the pod specify node and/or pod affinity or anti-affinity rules? If yes, would scheduling the pod to this node break those rules?   
    
    2. SELECTING THE BEST NODE FOR THE POD
        a.) Suppose you have a two-node cluster. Both nodes are eli- gible, but one is already running 10 pods, while the other, for whatever reason,
         isn’t running any pods right now. It’s obvious the Scheduler should favor the second node in this case.
        
        b.) If these two nodes are provided by the cloud infrastructure, it may be bet- ter to schedule the pod to the first node and relinquish the 
          second node back to the cloud provider to save money.         
```

Introducing the controllers running in the Controller Manager

```
    you need other active components to make sure the actual state of the system converges toward the desired state, as specified in the resources deployed through the API server.
     This work is done by controllers running inside the Controller Manager.
```

What the Kubelet does

```
Its initial job is to register the node it’s running on by creating a Node resource in the API server.
Then it needs to continuously monitor the API server for Pods that have been scheduled to the node, and start the pod’s containers.
It does this by telling the configured container runtime (which is Docker, CoreOS’ rkt, or some- thing else) to run a container from a specific container image. 
The Kubelet then con- stantly monitors running containers and reports their status, events, and resource consumption to the API server.

The Kubelet is also the component that runs the container liveness probes, restart- ing containers when the probes fail.
Lastly, it terminates containers when their Pod is deleted from the API server and notifies the server that the pod has terminated.
```

The role of the Kubernetes Service Proxy

```
The kube-proxy makes sure connections to the service IP and port end up at one of the pods backing that service. 
When a service is backed by more than one pod, the proxy performs load balancing across those pods.

The initial implementation of the kube-proxy was the userspace proxy. It used an actual server process to accept connections and proxy them to the pods. 
To intercept connections destined to the service IPs, the proxy configured iptables rules (iptables is the tool for managing the Linux kernel’s packet 
filtering features) to redirect the connections to the proxy server. This mode is called the userspace proxy mode. 
Balanced connections across pods in a true round-robin fashion.

The kube-proxy implementation only uses iptables rules to redirect packets to a randomly selected backend pod without passing them through an actual 
proxy server. This mode is called the iptables proxy mode. It selects pods randomly.
```

The chain of events

```
Imagine you prepared the YAML file containing the Deployment manifest and you’re about to submit it to Kubernetes through kubectl.
kubectl sends the manifest to the Kubernetes API server in an HTTP POST request. 
The API server validates the Deployment specification, stores it in etcd, and returns a response to kubectl.

All API server clients watching the list of Deployments through the API server’s watch mechanism are notified of the newly created 
Deployment resource immediately after it’s created.

As a new Deployment object is detected by the Deployment controller, it creates a ReplicaSet for the current specification of the Deployment.

The newly created ReplicaSet is then picked up by the ReplicaSet controller. The controller takes into consideration the replica count and pod selector 
defined in the ReplicaSet and verifies whether enough existing Pods match the selector. The controller then creates the Pod resources based on the pod 
template in the ReplicaSet.

These newly created Pods are now stored in etcd, but they each still lack one important thing—they don’t have an associated node yet.
The Scheduler watches for Pods like this, and when it encounters one, chooses the best node for the Pod and assigns the Pod to the node.

The Kubelet, watching for changes to Pods on the API server, sees a new Pod scheduled to its node, so it inspects the Pod definition and instructs Docker, 
or whatever container runtime it’s using, to start the pod’s containers. The container runtime then runs the containers.
```

Running highly available clusters

```
 - RUNNING MULTIPLE INSTANCES TO REDUCE THE LIKELIHOOD OF DOWNTIME
 - USING LEADER-ELECTION FOR NON-HORIZONTALLY SCALABLE APPS
    it’s a way for multiple app instances running in a distributed environment to come to an agreement on which is the leader. 
     leader is either the only one performing tasks, while all others are waiting for the leader to fail and then becoming leaders themselves.
     the leader being the only instance performing writes, while all the others are providing read-only access to their data.
     This ensures two instances are never doing the same job.     
```
</p>
<p>
<h3>Understanding authentication</h3>

Creating ServiceAccounts

```
kubectl create serviceaccount foo
```

pods can authenticate by sending the contents of the file /var/run/secrets/kubernetes.io/serviceaccount/token, 
which is mounted into each container’s filesystem through a secret volume.

Every pod is associated with a ServiceAccount, which represents the identity of the app running in the pod. 
The token file holds the ServiceAccount’s authentication token. When an app uses this token to connect to the API server, 
the authentication plugin authenticates the ServiceAccount and passes the ServiceAccount’s username back to the API server core. 
ServiceAccount usernames are formatted like this:

```
system:serviceaccount:<namespace>:<service account name>
```

The API server passes this username to the configured authorization plugins, which determine whether the action the app is trying 
to perform is allowed to be performed by the ServiceAccount.


ServiceAccount with an image pull Secret: sa-image-pull-secrets.yaml

```yaml
        apiVersion: v1
        kind: ServiceAccount
        metadata:
          name: my-service-account
        imagePullSecrets:
        - name: my-dockerhub-secret
```

Assigning a ServiceAccount to a pod

by setting the name of the ServiceAccount in the spec.serviceAccountName field in the pod definition.

Pod using a non-default ServiceAccount: curl-custom-sa.yaml


<h4>Securing the cluster with role-based access control</h4>

 RBAC prevents unau- thorized users from viewing or modifying the cluster state. 
 The default Service- Account isn’t allowed to view cluster state, let alone modify it in any way, unless you grant it additional privileges.

The RBAC authorization rules are configured through four resources, which can be grouped into two groups:
 Roles and ClusterRoles, which specify which verbs can be performed on which resources.
 RoleBindings and ClusterRoleBindings, which bind the above roles to specific users, groups, or ServiceAccounts.

Role and RoleBinding are namespaced resources
ClusterRole and ClusterRoleBinding are cluster-level resources 

```
If you’re using Minikube, you also may need to enable RBAC by starting Minikube with --extra-config=apiserver.Authorization.Mode=RBAC
```

CREATING THE NAMESPACES AND RUNNING THE PODS

```
$ kubectl create ns foo

$ kubectl run test --image=luksa/kubectl-proxy -n foo

$ kubectl create ns bar

$ kubectl run test --image=luksa/kubectl-proxy -n bar

$ kubectl get po -n bar

$ kubectl exec -it test-5cd6f69956-fsg7h -n bar sh

# curl localhost:8001/api/v1/namespaces/bar/services

response: 
 "status": "Failure",
  "message": "services is forbidden: User \"system:serviceaccount:bar:default\" cannot list resource \"services\" in API group \"\" in the namespace \"bar\"",

```
The default permissions for a ServiceAccount don’t allow it to list or modify any resources.

First, you’ll need to create a Role resource.

A definition of a Role: service-reader.yaml

To grant the rights, run the following command:

```
kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=your.email@address.com
```

BINDING A ROLE TO A SERVICEACCOUNT

```
kubectl create rolebinding test --role=service-reader --serviceaccount=bar:default -n bar
``` 

To bind a Role to a user instead of a ServiceAccount, use the --user argument to specify the username. To bind it to a group, use --group.

Now execute the following and see the response.

```
$ kubectl exec -it test-5cd6f69956-fsg7h -n bar sh

# curl localhost:8001/api/v1/namespaces/bar/services
```


Using ClusterRoles and ClusterRoleBindings.

certain resources aren’t namespaced at all (this includes Nodes, PersistentVolumes, Namespaces, and so on). 
We’ve also mentioned the API server exposes some URL paths that don’t represent resources (/healthz for example). 
Regular Roles can’t grant access to those resources or non- resource URLs, but ClusterRoles can.

Let’s look at how to allow your pod to list PersistentVolumes in your cluster. First, you’ll create a ClusterRole called pv-reader:

```
kubectl create clusterrole pv-reader --verb=get,list --resource=persistentvolumes

kubectl create clusterrolebinding pv-test --clusterrole=pv-reader  --serviceaccount=bar:default

curl localhost:8001/api/v1/persistentvolumes

```

ALLOWING ACCESS TO NON-RESOURCE URLS

The default system:discovery ClusterRole

```
kubectl get clusterrole system:discovery -o yaml

kubectl get clusterrolebinding system:discovery -o yaml
```


USING CLUSTERROLES TO GRANT ACCESS TO RESOURCES IN SPECIFIC NAMESPACES

With the first command, you’re trying to list pods across all namespaces. With the sec- ond, you’re trying to list pods in the foo namespace.
 The server doesn’t allow you to do either. Now, let’s see what happens when you create a ClusterRoleBinding and bind it to the pod’s ServiceAccount:

```
kubectl create clusterrolebinding view-test --clusterrole=view  --serviceaccount=foo:default
```

Granting authorization permissions wisely

giving all your ServiceAccounts the cluster-admin ClusterRole is a bad idea. As is always the case with security, it’s best to give everyone only t
he permissions they need to do their job and not a single permission more (principle of least privilege).

CREATING SPECIFIC SERVICEACCOUNTS FOR EACH POD

It’s a good idea to create a specific ServiceAccount for each pod (or a set of pod replicas) and then associate it with a tailor-made Role 
(or a ClusterRole) through a RoleBinding (not a ClusterRoleBinding, because that would give the pod access to resources in other namespaces, 
which is probably not what you want).

EXPECTING YOUR APPS TO BE COMPROMISED

You should expect unwanted persons to eventually get their hands on the ServiceAccount’s authentication token, so you should always constrain
 the ServiceAccount to prevent them from doing any real damage.

</p>
<p>
<h3>Securing cluster nodes and the network</h3>

Using the node’s network namespace in a pod

A pod may need to use the node’s network adapters instead of its own virtual network adapters. 
 This can be achieved by setting the <i>hostNetwork</i> property in the pod spec to <i>true</i>.

A pod using the node’s network namespace: pod-with-host-network.yaml

Binding to a host port without using the host’s network namespace

- when a pod is using a hostPort, a connection to the node’s port is forwarded directly to the pod running on that node, whereas with a NodePort service, 
a connection to the node’s port is forwarded to a randomly selected pod (possibly on another node).

- pods using a hostPort, the node’s port is only bound on nodes that run such pods, whereas NodePort services bind the port on all nodes, 
even on those that don’t run such a pod 

- If a host port is used, only a single pod instance can be scheduled to a node.

Binding a pod to a port in the node’s port space: kubia-hostport.yaml

Using the node’s PID and IPC namespaces

When you set them to true, the pod’s containers will use the node’s PID and IPC namespaces, allowing processes running in the containers to see all the other
 processes on the node or communicate with them through IPC.

Using the host’s PID and IPC namespaces: pod-with-host-pid-and-ipc.yaml

<h4>Configuring the container’s security context.</h4>

Running a container as a specific user

Running containers as a specific user: pod-as-user-guest.yaml

Preventing a container from running as root

```yaml
spec:
  containers:
  - name: main
    image: alpine
    command: ["/bin/sleep", "999999"]
    securityContext:
      runAsNonRoot: true
```

To get full access to the node’s kernel, the pod’s container runs in privileged mode. This is achieved by setting the privileged property in the 
container’s securityContext property to true. 


Adding individual kernel capabilities to a container

Eg. a container usually isn’t allowed to change the system time. If you want to allow the container to change the system time, you can add a 
capability called CAP_SYS_TIME to the container’s capabilities list.

```yaml
securityContext:
  capabilities:
    add:
    - SYS_TIME
```
Run the command in this new pod’s container, the system time is changed successfully:
```
kubectl exec -it pod-add-settime-capability -- date +%T -s "12:00:00"
```

Dropping capabilities from a container

```
ecurityContext:
  capabilities:
    drop:
    - CHOWN
```

Preventing processes from writing to the container’s filesystem

You may want to prevent the processes running in the container from writing to the container’s filesystem, and only allow them to write to mounted volumes.
This is done by set- ting the container’s securityContext.readOnlyRootFilesystem property to true.

```
spec:
  containers:
  - name: main
    image: alpine
    command: ["/bin/sleep", "999999"]
    securityContext:
      readOnlyRootFilesystem: true
    volumeMounts:
    - name: my-volume
      mountPath: /volume
      readOnly: false
  volumes:
  - name: my-volume
emptyDir:
```

Sharing volumes when containers run as different users.

You may need to run the two containers as two different users. If those two containers use a volume to share files, they may not necessarily be able to 
read or write files of one another.

Kubernetes allows you to specify supplemental groups for all the pods running in the container, allowing them to share files, 
regardless of the user IDs they’re running as.

This is done using the following two properties:
 fsGroup
 supplementalGroups

```
spec:
  securityContext:
    fsGroup: 555
    supplementalGroups: [666, 777]
```

<h4>Restricting the use of security-related features in pods</h4>

The cluster admin can restrict the use of the previously described security-related features by creating one or more PodSecurityPolicy resources.
PodSecurityPolicy is a cluster-level (non-namespaced) resource, which defines what security-related features users can or can’t use in their pods.

A PodSecurityPolicy resource defines things like the following:
 Whether a pod can use the host’s IPC, PID, or Network namespaces 
 Which host ports a pod can bind to
 What user IDs a container can run as
 Whether a pod with privileged containers can be created
 Which kernel capabilities are allowed, which are added by default and which are always dropped
 What SELinux labels a container can use
 Whether a container can use a writable root filesystem or not
 Which filesystem groups the container can run as
 Which volume types a pod can use

An example PodSecurityPolicy: pod-security-policy.yaml

USING THE MUSTRUNAS RULE

To only allow containers to run as user ID 2 and constrain the default filesystem group and supplemental group IDs to be anything 
from 2–10 or 20– 30 (all inclusive)

```yaml
runAsUser:
  rule: MustRunAs
  ranges:
  - min: 2
    max: 2
fsGroup:
  rule: MustRunAs
  ranges:
  - min: 2
    max: 10
  - min: 20
    max: 30
supplementalGroups:
  rule: MustRunAs
  ranges:
  - min: 2
    max: 10
  - min: 20
    max: 30
```

If the pod spec tries to set either of those fields to a value outside of these ranges, the pod will not be accepted by the API server.

Dockerfile with a USER directive: kubia-run-as-user-5/Dockerfile

```
FROM node:7
ADD app.js /app.js
USER 5
ENTRYPOINT ["node", "app.js"]
```

Specifying capabilities in a PodSecurityPolicy

```yaml
apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
spec:
  allowedCapabilities:
  - SYS_TIME
  defaultAddCapabilities:
  - CHOWN
  requiredDropCapabilities:
  - SYS_ADMIN
- SYS_MODULE
```
<h4>Isolating the pod network</h4>

Enabling network isolation in a namespace

By default, pods in a given namespace can be accessed by anyone. First, you’ll need to change that. You’ll create a default-deny NetworkPolicy, 
which will prevent all clients from connecting to any pod in your namespace.

A default-deny NetworkPolicy: network-policy-default-deny.yaml

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
   name: default-deny
spec:
  podSelector:      
```
Empty pod selector matches all pods in the same namespace


Allowing only some pods in the namespace to connect to a server pod

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: postgres-netpolicy
spec:
  podSelector:
    matchLabels:
      app: database
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: webserver
    ports:
    - port: 5432
```
</p>
<p>
<h3>Requesting resources for a pod’s containers</h3>

Creating pods with resource requests

A pod with resource requests: requests-pod.yaml


<h4>Understanding pod QoS classes<h4>

```
Imagine having two pods, where pod A is using, let’s say, 90% of the node’s memory and then pod B suddenly requires 
more memory than what it had been using up to that point and the node can’t provide the required amount of memory. 
Which container should be killed? Should it be pod B, because its request for memory can’t be satisfied, or should
pod A be killed to free up memory, so it can be provided to pod B?
```

Kubernetes does this by categorizing pods into three Quality of Service (QoS) classes:
 BestEffort (the lowest priority) : It’s assigned to pods that don’t have any requests or limits set at all. 
They may get almost no CPU time at all and will be the first ones killed when memory needs to be freed for other pods
A BestEffort pod has no memory limits set, its containers may use as much memory as they want, if enough memory is available.


 Burstable: This includes single-container pods where the container’s limits don’t match its requests and all pods where at 
least one container has a resource request specified, but not the limit. Burstable pods get the amount of resources they request,
but are allowed to use addi- tional resources (up to the limit) if needed.


 Guaranteed (the highest): This class is given to pods whose containers’ requests are equal to the limits for all resources. 
Requests and limits need to be set for both CPU and memory.
They need to be set for each container.
They need to be equal.

Understanding which process gets killed when memory is low

First in line to get killed are pods in the BestEffort class, followed by Burstable pods, and finally Guaranteed pods, 
which only get killed if system processes need memory.


UNDERSTANDING HOW CONTAINERS WITH THE SAME QOS CLASS ARE HANDLED

Each running process has an OutOfMemory (OOM) score. The system selects the process to kill by comparing OOM scores of all the running processes.
When memory needs to be freed, the process with the highest score gets killed.

OOM scores are calculated from two things: the percentage of the available memory the process is consuming and a fixed OOM score adjustment,
 which is based on the pod’s QoS class and the container’s requested memory.


<h4>Setting default requests and limits for pods per namespace</h4>

If you don’t set them, the container is at the mercy of all other containers that do specify resource requests and limits. 
It’s a good idea to set requests and limits on every container.


Introducing the LimitRange resource

It allows you to specify (for each namespace) not only the minimum and maximum limit you can set on a container for each resource, 
but also the default resource requests for containers that don’t specify requests explicitly

A LimitRange resource: limits.yaml

</p>
<p>
<h3>Automatic scaling of pods and cluster nodes</h3>

<h4>Horizontal pod autoscaling</h4>

It’s performed by the Horizontal controller, which is enabled and configured by creating a HorizontalPodAutoscaler (HPA) resource.

The autoscaling process can be split into three steps:
 Obtain metrics of all the pods managed by the scaled resource object.
    - pod and node met- rics are collected by an agent called cAdvisor, which runs in the Kubelet on each node, and then aggregated by the 
    cluster-wide component called Heapster. The horizontal pod autoscaler controller gets the metrics of all the pods by querying Heapster 
    through REST calls. 

 Calculate the number of pods required to bring the metrics to (or close to) the specified target value.
    - calculating the required replica count is simple. All it takes is summing up the metrics values of all the pods, dividing that by the 
    target value set on the HorizontalPodAutoscaler resource, and then rounding it up to the next-larger integer. 


 Update the replicas field of the scaled resource.
    - Autoscaler controller modifies the replicas field of the scaled resource through the Scale sub-resource (Deployment, ReplicaSet, StatefulSet, 
    or ReplicationController). It enables the Autoscaler to do its work without knowing any details of the resource it’s scaling, except for what’s 
    exposed through the Scale sub-resource


CREATING A HORIZONTALPODAUTOSCALER BASED ON CPU USAGE

1. create a Deployment, each instance requesting 100 millicores of CPU.
2. After creating the Deployment, to enable horizontal autoscaling of its pods, you need to create a HorizontalPodAutoscaler (HPA) object and point
    it to the Deployment. 
    You could prepare and post the YAML manifest for the HPA.
    An easier way exists—using the kubectl autoscale command:

```
  $ kubectl autoscale deployment kubia --cpu-percent=30 --min=1 --max=5 
   
  $ kubectl get hpa
```
Watching multiple resources in parallel

```
$ watch -n 1 kubectl get hpa,deployment
```

MODIFYING THE TARGET METRIC VALUE ON AN EXISTING HPA OBJECT

```yaml
spec:
  maxReplicas: 5
  metrics:
  - resource:
      name: cpu
      targetAverageUtilization: 60
    type: Resource

```

<h4>Horizontal scaling of cluster nodes</h4>

Cluster autoscaling is currently available on
 Google Kubernetes Engine (GKE)
 Google Compute Engine (GCE)
 Amazon Web Services (AWS)
 Microsoft Azure

enable the Cluster Autoscaler like this:

```
$ gcloud container clusters update kubia --enable-autoscaling  --min-nodes=3 --max-nodes=5
```

Limiting service disruption during cluster scale-down

Certain services require that a minimum number of pods always keeps running; this is especially true for quorum-based clustered applications. 
For this reason, Kubernetes provides a way of specifying the minimum number of pods that need to keep running while performing these types of operations.
This is done by creating a Pod- DisruptionBudget resource.

If you want to ensure three instances of your kubia pod are always running (they have the label app=kubia), create the PodDisruptionBudget resource like this:

```
$ kubectl create pdb kubia-pdb --selector=app=kubia --min-available=3
```



</p>