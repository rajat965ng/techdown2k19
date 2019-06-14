

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