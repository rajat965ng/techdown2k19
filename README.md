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