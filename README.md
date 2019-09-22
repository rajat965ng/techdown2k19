<h2>Terraforming Google Cloud Platform with Mongo using Packer</h2> 

- Use Case
  - Importing a Guest operating system (Ubuntu 16.04) and install Mongo DB on top of it.
  - Provisioning Mongo on VM using infrastructure as code. 
    
- Benefits 
  - Maintenance and Recovery were easy in case of failure conditions.
  - The total cost of ownership was also less due to the reduced need for infrastructure.
  
- Concept
  - Baking Vs Frying
    - <i>Baked images</i> are previously prepared with software and configuration. They are usually bigger as it is bundled with installations and it's dependencies.
      Baked images empowers "Immutable Architecture" because most of the time they don't need extra intervention after instantiation. In case of failure, it's better
      to recreate it, rather than repair. In case of 'Autoscaling' baked images are preferred.
    - <i>Frying</i> is known as provisioning over raw images. In order to be ready-to-use, these lightweight images must be provisioned with software and configurations
      after instantiation. Concern about fried provisioning, is "How to avoid breaking it ?" when executed repeatedly. The package manager like <i>apt</i> usually install
      latest copy of packages unless the version is not specified. Unexpected behaviour can happen with untested latest version of packages.   

- Platform (GCP)
  - Compute Engine
    - To configure your Google Cloud Platform infrastructure
  - Cloud storage
    - The list of GCS paths, e.g. 'gs://mybucket/path/to/file.tar.gz', where the image will be exported. 
  
- Tools
  - Machine Image
    - It is a static unit that has pre-configured operating system and installed software that can quickly create new running machines. Different platforms has their own machine
      formats like, AMI for EC2, VMDK/VMX files for VMware or Compute Engine Images for GCP etc.
    - In this example we are using base image of Ubuntu 16.06. On base image we'll install MongoDB server and bake the bundle to form a new image.   
  - Packer
    - Introduction
      - HashiCorp Packer made it easy to automate and use any type of machine image. It promotes configuration management using automated scripts to install and configure
        software in packer made images.
    - Why use packer ?
      - It is an OpenSource tool that provides a single source configuration which can be use to configure machine images from different providers.
      - Can create multiple images for different platforms in parallel.
    - Advantages
      - Lightning fast infrastructure deployment.
      - Inter provider portability
      - Testable: Post building the image, smoke test can be executed to check if things are working fine. 
    - Code
      - ![Image Json](/image.json)
      - Scripts
        - ![Mongo Installation](/mongo.sh)
        - ![TZ Data](/tzdata.sh)    
  - Terraform
    - Introduction
      - OpenSource tool for building, changing and versioning infrastructure. 
    - Why to use Terraform ?
      - Build by same company as Packer, HashiCorp, Terraform is based on the same principle as of Packer.
      - Terraform lets you to manage infrastructure on GCP using single configuration file in TF format.
    - Advantages
      - Empower Infrastructure as Code
      - Provide Execution Plan
      - Generate Resource Graph
    - Code
      - ![Provider](/iac/provider.tf)
      - ![Instance](/iac/instance.tf)
