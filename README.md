<h1>Designing Distributed Systems<h1>
<h3>Patterns and Paradigms for Scalable, Reliable Services</h3>
<h4>By Brendan Burns</h4><br>
<h2>Single Node Pattern</h2>

<h3>Side Car Pattern</h3>


 ![genric sidecar](singleNodePatterns/sideCar/generic-sidecar.png)
 
 
 An Example Sidecar: Adding HTTPS to a Legacy Service
 
 ![https sidecar](singleNodePatterns/sideCar/https-sidecar.png)
 
 Building a Simple PaaS with Sidecars
 
 - Imagine building a simple platform as a service (PaaS) built around the git workflow.  
 - Once you deploy this PaaS, simply pushing new code up to a Git repository results in that code being deployed to the running servers.
 
 ![paas sidecar](singleNodePatterns/sideCar/paas-sidecar.png)
 
 
 
 <h3>Ambassador Pattern</h3>
 
 ![generic ambassador](singleNodePatterns/ambassador/generic-ambassador.png)

Using an Ambassador to Shard a Service
 
 - Sometimes the data that you want to store in a storage layer becomes too big for a single machine to handle. In such situations, you need to shard your storage layer. Sharding splits up the layer into multiple disjoint pieces, each hosted by a separate machine.
 - When adapting an existing application to a sharded backend, you can introduce an ambassador container that contains all of the logic needed to route requests to the appropriate storage shard

Implementing a Sharded Redis

 - Redis is a fast key-value store that can be used as a cache or for more persistent storage.
 - Begin by deploying a sharded Redis service to a Kubernetes cluster.
 - <b>twemproxy</b> is a lightweight, highly performant proxy for memcached and Redis, which was originally developed by Twitter and is open source and available on GitHub.
 - configure twemproxy to point to the replicas created.
 
 
 Using an Ambassador for Service Brokering
 
 - Building a portable application requires that the application know how to introspect its environment and find the appropriate MySQL service to connect to. 
 - This process is called service discovery, and the system that performs this discovery and linking is commonly called a service broker.
 
 ![brokering ambassador](singleNodePatterns/ambassador/service-broking-ambassador.png)
 
 
 Using an Ambassador to Do Request Splitting
 
 - In many production systems, it is advantageous to be able to perform request splitting, where some fraction of all requests are not serviced by the main production service but rather are redirected to a different implementation of the service. 
 - Most often, this is used to perform experiments with new, beta versions of the service to determine if the new version of the software is reliable or comparable in performance to the currently deployed version.
 - To implement our request-splitting experiment, we’re going to use the nginx web server.
 
 ```yaml
        worker_processes  5;
        error_log  error.log;
        pid        nginx.pid;
        worker_rlimit_nofile 8192;
        events {
          worker_connections  1024;
        }
        http {
            upstream backend {
                ip_hash;
                server web weight=9;
                server experiment;
            }
            server {
                listen localhost:80;
                location / {
                    proxy_pass http://backend;
                }
              } 
            }
 ```
 - Using IP hashing in this configuration. This is important because it ensures that the user doesn’t flip-flop back and forth between the experiment and the main site. 
 - This assures that every user has a consistent experience with the application.
 - The weight parameter is used to send 90% of the traffic to the main existing application, while 10% of the traffic is redirected to the experiment.
 
 - As with other examples, we’ll deploy this configuration as a ConfigMap object in Kubernetes:
 
 ```
 kubectl create configmaps --from-file=nginx.conf
 ``` 
 
 
 <h3>Adapter Pattern</h3>
 
 - In the adapter pattern, the adapter container is used to modify the interface of the application container so that it conforms to some predefined interface that is expected of all applications. For exam‐ ple, an adapter might ensure that an application implements a consistent monitoring interface. Or it might ensure that log files are always written to stdout or any number of other conventions.
 
 
 ![generic adapter](singleNodePatterns/adapter/generic-adapter.png)
 
 
 Monitoring
 
 - Applying the adapter pattern to monitoring, we see that the application container is simply the application that we want to monitor. The adapter container contains the tools for transforming the monitoring interface exposed by the application container into the interface expected by the general purpose monitoring system.
 
 
 Using Prometheus for Monitoring
 
 - Prometheus is a monitoring aggregator, which collects metrics and aggregates them into a single time-series database. 
 - On top of this database, Prometheus provides visualization and query language for introspecting the collected metrics. 
 - To collect metrics from a variety of different systems, Prometheus expects every container to expose a specific metrics API. 
 - This enables Prometheus to monitor a wide variety of different programs through a single interface.
 
 - Redis key-value store, do not export metrics in a format that is compatible with Prometheus.
 - The adapter pattern is quite useful for taking an existing service like Redis and adapting it to the Prometheus metrics-collection interface.
 - Provide an adapter that implements the Prometheus interface. The following image is an adapter for Redis metrics to Prometheus conversion.
        
  ```
  - image: oliver006/redis_exporter
  ```       
  
  
  Logging
  
  - In the world of containerized applications where there is a general expectation that your containers will log to stdout, because that is what is available via commands like docker logs or kubectl logs.
  - Different application containers can log information in different formats, but the adapter container can transform that data into a single structured representation that can be consumed by your log aggregator.
  
  Normalizing Different Logging Formats with Fluentd
  
  - fluentd is one of the more popular open source logging agents available. One of its major features is a rich set of community-supported plugins that enable a great deal of flexibility in monitoring a variety of applications.
  - Redis is a popular key-value store; one of the commands it offers is the SLOWLOG command. This command lists recent queries that exceeded a particular time interval. Such information is quite useful in debugging your application’s performance. 
  - Unfortunately, SLOWLOG is only available as a command on the Redis server, which means that it is difficult to use retrospectively if a problem happens when someone isn’t available to debug the server.
  - To fix this limitation, we can use fluentd and the adapter pattern to add slow-query logging to Redis.
  - Use the adapter pattern with a redis container as the main application container, and the fluentd container as our adapter container. 
  - In this case, we will also use the fluent-plugin-redis-slowlog fluentd plugin to listen to the slow queries. 
  - We can configure this plugin using the following snippet:

   ```
   <source>
         type redis_slowlog
         host localhost
         port 6379
         tag redis.slowlog
   </source>
   ```
   
   - A similar exercise can be done to monitor logs from the Apache Storm system.
   - we deploy a fluentd adapter with the fluent-plugin-storm plugin enabled.
   
   ```
   <source>
         type storm
         tag storm
         url http://localhost:8080
         window 600
         sys 0
   </source>
   ```
   
   <hr>
   <h2>Serving Patterns</h2>
   
   <h3>Replicated Load-Balanced Services</h3>

   Stateless Services
   
   ![generic replication](servingPatterns/replicatedLoadBalancedServices/generic-replication.png)
   
   - Stateless services include things like static content servers and complex middleware systems that receive and aggregate responses from numerous different backend systems. 
   - No matter how small your service is, you need at least two replicas to provide a service with a “highly available” service level agreement (SLA).
   - Consider trying to deliver a three-nines (99.9% availability). In a three-nines service, you get 1.4 minutes of downtime per day (24 × 60 × 0.001).
   - Assuming that you have a service that never crashes, that still means you need to be able to do a software upgrade in less than 1.4 minutes in order to hit your SLA with a single instance.
   - If your team is really embracing contin‐ uous delivery and you’re pushing a new version of software every hour, you need to be able to do a software rollout in 3.6 seconds to achieve your 99.9% uptime SLA with a single instance.
   - Any longer than that and you will have more than 0.01% downtime from those 3.6 seconds.
   - That way, while you are doing a rollout, or in the—unlikely, I’m sure—event that your software crashes, your users will be served by the other replica of the service and never know anything was going on. 
   - Horizontally scalable systems handle more and more users by adding more replicas;
   
   ![horizontal replication](servingPatterns/replicatedLoadBalancedServices/horizontal-replication.png)
   
   Readiness Probes for Load Balancing
   
   - A readiness probe determines when an application is ready to serve user requests.
   - When building an application for a replicated service pattern, be sure to include a special URL that implements this readiness check.
   
   ```yaml
        spec:
          containers:
          - name: server
            image: brendanburns/dictionary-server
            ports:
            - containerPort: 8080
            readinessProbe:
              httpGet:
                path: /ready
                port: 8080
              initialDelaySeconds: 5
              periodSeconds: 5
   ```
   
   Session Tracked Services
   
   - Often there are reasons for wanting to ensure that a particular user’s requests always end up on the same machine. Sometimes this is because you are caching that user’s data in memory, so landing on the same machine ensures a higher cache hit rate.
   - This session tracking is performed by hashing the source and destination IP addresses and using that key to identify the server that should service the requests. So long as the source and destination IP addresses remain constant, all requests are sent to the same replica.
   - IP-based session tracking works within a cluster (internal IPs) but generally doesn’t work well with external IP addresses because of network address translation (NAT). For external session tracking, application-level tracking (e.g., via cookies) is preferred.
   - Session tracking is accomplished via a consistent hashing function. When the number of replicas changes, the mapping of a particular user to a replica may change.
   - Consistent hashing functions minimize the number of users that actually change which replica they are mapped to, reducing the impact of scaling on your application.
   
   Introducing a Caching Layer
   
   - The simplest form of caching for web applications is a caching web proxy. The caching proxy is simply an HTTP server that maintains user requests in memory state. If two users request the same web page, only one request will go to your backend; the other will be serviced out of memory in the cache.
   
   ![cache service](servingPatterns/replicatedLoadBalancedServices/cache-service.png)
   
   ```yaml
          spec:
             containers:
             - name: cache
               image: brendanburns/varnish
               command:
               - varnishd
               - -F
               - -f
               - /etc/varnish-config/default.vcl
               - -a
               - 0.0.0.0:8080
               - -s
               # This memory allocation should match the memory request above 
               - malloc,2G
               resources:
                 requests:
                   # We'll use two gigabytes for each varnish cache
                   memory: 2Gi
   ```
   
   
   Expanding the Caching Layer
   
   Rate Limiting and Denial-of-Service Defense
   
   - Accidentally running a load test against a production installation. Thus, it makes sense to add general denial-of-service defense via rate limiting.
   - A best practice to have a relatively small rate limit for anonymous access and then force users to log in to obtain a higher rate limit. 
   - Requiring a login provides auditing to determine who is responsible for the unexpected load.
   - Also offers a barrier to would-be attackers who need to obtain multiple identities to launch a successful attack.
   - When a user hits the rate limit, the server will return the 429 error code indicating that too many requests have been issued.
   
   SSL Termination

   - Each individual internal service should use its own certificate to ensure that each layer can be rolled out independently. 
   - Thus we want to add a third layer to our stateless application pattern, which will be a replicated layer of nginx servers that will handle SSL termination for HTTPS traffic and forward traffic on to our cache.
   
   ![complete replicated stateless service](servingPatterns/replicatedLoadBalancedServices/complete-load-balanced-service.png)
   
   - nginx con‐ figuration to serve SSL:
   
   ```
       events {
             worker_connections  1024;
       }
       http {
         server {
           listen 443 ssl;
           server_name my-domain.com www.my-domain.com;
           ssl on;
           ssl_certificate         /etc/certs/tls.crt;
           ssl_certificate_key     /etc/certs/tls.key;
           location / {
               proxy_pass http://varnish-service:80;
               proxy_set_header Host $host;
               proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
               proxy_set_header X-Forwarded-Proto $scheme;
               proxy_set_header X-Real-IP $remote_addr;        
           } 
         }
       }
   ```
   