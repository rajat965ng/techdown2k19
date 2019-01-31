<h1>Terraform: Infrastructure As Code</h1>
<br>
<h2>Use Case</h2>
<p>This code creates a Red Hat Instance on AWS platform and deploy an Apache Tomcat server on it. Tomcat GUI can be accessible on 8080(HTTP) using public IP </p>
<br>
<h2>Following are the architectural components for setting up simple instance on AWS.</h2>
<ul>
  <li>VPC (Virtual Private Cloud)</li>
  <li>Subnet</li>
  <li>Security Group</li>
  <li>Internet Gateway</li>
  <li>Route Table</li>
  <li>Route Table Associations</li>
  <li>AWS Instance along with remote-exec provisioner</li>
  <li>Elastic IP</li>
  <li>Provider</li>
</ul>
