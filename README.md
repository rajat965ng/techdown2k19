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