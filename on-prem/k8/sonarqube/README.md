# SonarQube Setup

## Verification endpoint
    kubectl get svc
    
    NAME             TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)              AGE
    sonar-pg         ClusterIP   10.43.91.221   <none>        8080/TCP,5432/TCP    29h
    sonarqube        NodePort    10.43.23.215   <none>        80:32690/TCP         29h

    curl http://10.43.23.215/sonar/api/server/version

## Commands
### To access postgress
    export PGPASSWORD='gpssapassword'; psql -h 'localhost' -U 'gpssa' -d 'sonarqube' ;
### To list users
    \du
### To list databases    
    \l    
### To list tables
    \dt    
### To Describe tables
    \d table_name or \d+ table_name to find the information on columns of a table.        