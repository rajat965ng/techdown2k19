<h2>Terraforming Google Cloud Platform with Mongo using Packer</h2> 

- <h4>Use Case</h4>
  - Importing a Guest operating system (Ubuntu 16.04) and install Mongo DB on top of it.
  - Provisioning Mongo on VM using infrastructure as code. 
    
- <h4>Benefits</h4> 
  - Maintenance and Recovery were easy in case of failure conditions.
  - The total cost of ownership was also less due to the reduced need for infrastructure.
  
- <h4>Concept</h4>
  - <b>Baking Vs Frying</b>
    - <b><i>Baked images</i></b> are previously prepared with software and configuration. They are usually bigger as it is bundled with installations and it's dependencies.
      Baked images empowers "Immutable Architecture" because most of the time they don't need extra intervention after instantiation. In case of failure, it's better
      to recreate it, rather than repair. In case of 'Autoscaling' baked images are preferred.
    - <b><i>Frying</i></b> is known as provisioning over raw images. In order to be ready-to-use, these lightweight images must be provisioned with software and configurations
      after instantiation. Concern about fried provisioning, is "How to avoid breaking it ?" when executed repeatedly. The package manager like <i>apt</i> usually install
      latest copy of packages unless the version is not specified. Unexpected behaviour can happen with untested latest version of packages.   

- <h4>Platform (GCP)</h4>
  - <h4>Compute Engine</h4>
    - To configure your Google Cloud Platform infrastructure
  - <h4>Cloud storage</h4>
    - The list of GCS paths, e.g. 'gs://mybucket/path/to/file.tar.gz', where the image will be exported. 
  
- <h4>Tools</h4>
  - <h4>Machine Image</h4>
    - It is a static unit that has pre-configured operating system and installed software that can quickly create new running machines. Different platforms has their own machine
      formats like, AMI for EC2, VMDK/VMX files for VMware or Compute Engine Images for GCP etc.
    - In this example we are using base image of Ubuntu 16.06. On base image we'll install MongoDB server and bake the bundle to form a new image.   
  - <h4>Packer</h4>
    - <b>Introduction</b>
      - HashiCorp Packer made it easy to automate and use any type of machine image. It promotes configuration management using automated scripts to install and configure
        software in packer made images.
    - <b>Why use packer ?</b>
      - It is an OpenSource tool that provides a single source configuration which can be use to configure machine images from different providers.
      - Can create multiple images for different platforms in parallel.
    - <b>Advantages</b>
      - Lightning fast infrastructure deployment.
      - Inter provider portability
      - Testable: Post building the image, smoke test can be executed to check if things are working fine. 
    - <b>Code</b>
      - [image.json]
        ```json
        {
          "variables": {
            "project_id": "project_id",
            "zone": "us-central1-a"
          },
          "builders": [
            {
              "type": "googlecompute",
              "account_file": "service-account.json",
              "project_id": "{{user `project_id`}}",
              "source_image": "ubuntu-1604-xenial-v20190913",
              "ssh_username": "root",
              "zone": "{{user `zone`}}"
            }
          ],
          "provisioners": [{
            "type": "shell",
            "scripts": ["tzdata.sh", "mongo.sh"]
          }],
          "post-processors": [
            {
              "type": "googlecompute-export",
              "paths": [
                "gs://dummy-bucket/file1.tar.gz"
              ],
              "keep_input_artifact": true
            }
          ]
        }
        ```
      - Scripts
        - [mongo.sh]
        ```bash
        #!/usr/bin/env bash
        sudo apt-get install apt-transport-https -y
        wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add - ;
        echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list ;
        sudo apt-get update -y;
        sudo apt-get install -y mongodb-org;
        echo "mongodb-org hold" | sudo dpkg --set-selections ;
        echo "mongodb-org-server hold" | sudo dpkg --set-selections ;
        echo "mongodb-org-shell hold" | sudo dpkg --set-selections ;
        echo "mongodb-org-mongos hold" | sudo dpkg --set-selections ;
        echo "mongodb-org-tools hold" | sudo dpkg --set-selections ;
        ```
        - [tzdata.sh]
        ```bash
        #!/usr/bin/env bash
        sudo adduser --disabled-password --gecos "packer" packer;
        su packer;
        export DEBIAN_FRONTEND=noninteractive;
        sudo apt-get update -y;
        ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime;
        sudo apt-get install -y tzdata;
        sudo dpkg-reconfigure --frontend noninteractive tzdata;
        ```            
  - <h4>Terraform</h4>
    - <b>Introduction</b>
      - OpenSource tool for building, changing and versioning infrastructure. 
    - <b>Why to use Terraform ?</b>
      - Build by same company as Packer, HashiCorp, Terraform is based on the same principle as of Packer.
      - Terraform lets you to manage infrastructure on GCP using single configuration file in TF format.
    - <b>Advantages</b>
      - Empower Infrastructure as Code
      - Provide Execution Plan
      - Generate Resource Graph
    - <b>Code</b>
      - [provider.tf]
      ```hcl-terraform
        provider "google" {
          credentials = "${file("service-account.json")}"
          project     = "project-name"
          region      = "us-central1"
        }
      ```
      - [instance.tf]
      ```hcl-terraform
        resource "google_compute_instance" "mongo-server" {
          name         = "mongo-server"
          machine_type = "n1-standard-1"
          zone         = "us-central1-a"
        
          tags = ["mongo"]
        
          boot_disk {
            initialize_params {
              image = "packer-1569120732"
            }
          }
        
        
          network_interface {
            network = "default"
        
            access_config {
              // Ephemeral IP
            }
          }
        
        
          metadata_startup_script = "sudo service mongod start"
        
          service_account {
            scopes = ["userinfo-email", "compute-ro", "storage-ro"]
          }
        }
      ```