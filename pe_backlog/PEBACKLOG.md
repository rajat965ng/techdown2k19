# CLOUDIFICATION
## Open backlog for future cloud platform engineering teams.

![Dynasty Warriors - Landscape, Army, Horse, Spear](image.jpg)

> While having almost 9 years of work experience in writing code, from the past 4 years I've been working on cloud migration projects. 
  Despite the organizational designation, I performed in the equal capacity of Senior Dev, Tech lead and Architect at the same time in almost every other assignment. 
  I got a chance to interact with people having designations like Engineering Managers, Solution Architect, Principal Consultants, and Evangelists. 
  Some of them came from the background of core infrastructure and platform provisioning, while others have prior experience of application development and later on they worked in platform engineering as well. I'm penning down, my thoughts on creating an open backlog for cloudification of any existing architecture 
  based on my experience and knowledge exchange with people who came across.

---
*The intent of this write up is to have a high-level product backlog to create sprint stories and kick-start development.*

---
#####Identify cloud platform: eg. AWS, GCP, Azure, etc.
  ##### 8 criteria to choose a cloud provider 
     1. Certifications & Standards
     2. Technologies & Service Roadmap
     3. Data Security, Data Governance and Business policies
     4. Service Dependencies & Partnerships
     5. Contracts, Commercials & SLAs
     6. Reliability & Performance
     7. Migration Support, Vendor Lock in & Exit Planning
     8. Business health & Company profile


#####Identify tools
  ##### Container orchestration 
     Eg. AKS, GKE, Openshift, etc.
  ##### Solution for securing the network
     Eg. OpenVPN, Bastion host (aws, gcp) etc.  
  ##### Source code repository management tools
     Eg. GitLab, Gogs, Bitbucket, etc.
  ##### CI/CD pipeline
     Eg. Jenkins, Teamcity, GitLab, CircleCI, etc.  
  ##### Infrastructure as a code
     Eg. Terraform, AWS CloudFormation, Ansible, Chef, Puppet Enterprise, Google Cloud Deployment Manager, Azure Automation, SaltStack, etc.
  ##### Image baking
     Eg. Packer (for VM), Docker (for Container)  
  ##### Setup Artifactory/ Repository
     Eg. Nexus, JFrog, etc.
  ##### Artifact scanning
     Eg. Nexus Auditor, Jfrog Xray, Qualys, etc.
  ##### Secrets Manager
     Eg. Vault, Google KMS, AWS Secrets Manager, etc.       
  ##### Logging
     Eg. ELK stack, Splunk, Fluentd, StackDriver, CloudWatch, etc.
  ##### Monitoring
     Eg. Prometheus, Grafana, etc.   
  ##### Cloud security and compliance risks
     Eg. RedLock Cloud, Google Apigee Sense, Amazon VPC PrivateLink, Duo Security, etc.
  
  
##### Strategy formulation across all tools
  ###### High Availability 
  ###### Disaster Recovery 
  ###### Backups 
  ###### On-boarding process 
  ###### RBAC (Role Back Access Control)
     
#####Primary tasks**
  #####Identify the "start the world" solution. 
     How to plant the first seed of your platform in a cloud provider? 
      It can be done through a local machine or by some external build pipeline eg. Google cloud build, Azure DevOps, AWS Code pipeline, etc.
  #####Implement VPN or Jump Servers.
     IP Whitelisting.
     Implement cloud security and compliance risk tools.
  #####Setup your DEV environment
     1.  Automate infrastructure provisioning.
     2.  Create shared VPC.
     3.  Create VPC peering in the case of multiple VPCs.
     4.  Active Directory's LDAP to/from Synchronization with the cloud platform.
     5.  Set up CI/CD tools.
     6.  Provision artifactory.
     7.  Create a backend for the terraform state. Eg. cloud storage, Terraform Enterprise, etc.
     8.  Writing custom Terraform providers (if not available) for any tool.
     9.  Create CI/CD pipelines.
     10. Implement Sentinel policies if eligible.
     11. Form and implement IAC testing strategy.
     12. Bake VM images.
     13. Publish docker gold images. 

---
*Subject to maturity of team and project requirements these tasks can be modified.
 Many times developers face 'Where to start from ?' situation when brainstorming together in their formative discussions.
 Hopefully, this high-level open backlog can help them to kick-start the discussion and transform it into achievable user stories.*
---