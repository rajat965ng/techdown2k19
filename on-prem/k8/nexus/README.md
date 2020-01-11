## Admin Credentials
    Username: admin  
    Password: admin1234
    
## Anonymous User Credentials    
    Username: abc
    Password: abc123
    
    
## Commands
    Go to every node of k8 cluster and make entry of every node IP in '/etc/docker/daemon.json'
    Execute 'service docker restart'
    
    Login
        docker login -u admin -p admin1234 10.150.16.171:32000  [where 10.150.16.171 is IP of the node(any)]
    
    
    
    Grant access to newly created user to push docker image on repo
        Setting -> Security -> Roles -> Create Role -> Nexus Role
            Role ID: nx-docker-xcs
            Role Name: nx-docker-xcs
            Role Description: Accessing docker repository
            Previliges:
                nx-repository-view-docker-*-add
                nx-repository-view-docker-*-browse
                nx-repository-view-docker-*-edit
                nx-repository-view-docker-*-read
                
        Users -> abc
            Roles -> Granted
                nx-anonymous
                nx-docker-xcs        
    
    Enable docker Realm to let docker client communicate with nexus
        Setting -> Security -> Realms
                Enable 'Docker Bearer Token Realm'            