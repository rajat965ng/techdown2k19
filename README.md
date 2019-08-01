<h2>Configure Vault with Nginx</h2>

- Implement RBAC on Kubernetes 
  - Apply rolebinding.yaml

- Deploy Vault 
  - Apply vault.yaml
  - It also exposes a dns with the name 'vault' for intra-cluster communication.
  - Following are the configurations that needs to be done at vaults container.
    - Create vault policy for the secret 'myapp'. This should be executed from container lifecycle post start placeholder.
    
    ```
    cat <<EOF | vault policy write myapp-kv-ro -
    path "secret/myapp/*" { capabilities = ["read", "list"] }
    path "secret/data/myapp/*" { capabilities = ["read", "list"] }
    EOF
    ```
    - Enable userpass auth mode for a test-user.
    
    ```
    vault auth enable userpass
    
    vault write auth/userpass/users/test-user \
            password=training \
            policies=myapp-kv-ro
            
    
    vault login -method=userpass \
            username=test-user \
            password=training
            
    vault kv get secret/myapp/config
                    
    ```
    - Retrieve kubernetes cert using service account role created in rolebinding.yaml
    
    ```
    kubectl get sa vault-auth -o jsonpath="{.secrets[*]['name']}"
    
    O/P: vault-auth-token-bwhtd , That's why we use RBACs in kubernetes
    
    kubectl get secret vault-auth-token-bwhtd -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo
    ```
    - Enable kubernetes authentication on vault using the cert obtained in previous step. And specify the role ('example') that will issue auth token to every agent.
    
    ```
    vault auth enable kubernetes
    
    vault write auth/kubernetes/config  kubernetes_host=https://192.168.99.100:8443   kubernetes_ca_cert="-----BEGIN CERTIFICATE-----
    MIIC5zCCAc+gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p
    a3ViZUNBMB4XDTE5MDYwMjEwMTMzMloXDTI5MDUzMTEwMTMzMlowFTETMBEGA1UE
    AxMKbWluaWt1YmVDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAM0o
    6jVyamqgbuJYWTyx7thDVY0mCFBjNORGLac+noPuZZKu57aCHTDvPr92LsrJXnSO
    D50QKSD60vIRVXDmmGyEx9EKvRaH/O5Ac6zBFOirR7id78LuPv/nE3NVojYqzMaQ
    KisHBVzwZIYlKlmiTEIAyUjUF0MSflTYAw+nguG5tc15ntB+VTZ7PQd1vGzF8spb
    sUT0qh2M8L1stmDuVhJ2e8jCCpOtuWTBTrAxJivTKozlXSdhG4l98s67QFWgcFmB
    WPWL+EtdZsY6uJg/o5aQGXCpnpWHK2QhntLPwkg4h6aGDJCeW0cnmQcHXNL1Pa6l
    pJUG1vGtaRRdGtv93pUCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQW
    MBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
    DQEBCwUAA4IBAQA/z880o/fiUGHS73OWu0GKCpIEpKPcGBxMhR4v8vZ2gHt5TW2X
    JL/p9lMY7m3u65SvfoZ3MPCCZsCQfYCHu4VUbPm+NmuoMz1gd+WTkOTO95+ZVK3H
    mtdCgFR09qhm8ELLuzTATJB5T+Wy+4Gqk6yiEpvNp/GsbYxumQ4kCZg3qLjvufSY
    TuWLJsuzmWBYsEOpMfnYhaoLCVD4VXDq0cRfhvDATRX5VkLjLZu/siRzav07pqmW
    OhSgborFxqA+tq/kmKWlKq4oocx/eFEVvI895picxXvAZPSlZovkCMWZOE17KK8w
    +e7YH+63/91lHqNli+kE7vKNg2BP8U3t+U5e
    -----END CERTIFICATE-----"
    
    
    vault write auth/kubernetes/role/example \
            bound_service_account_names=vault-auth \
            bound_service_account_namespaces=default \
            policies=myapp-kv-ro \
            ttl=24h
    ```
    - Test the kubernetes auth method to ensure that you can authenticate with Vault.

    ```    
    kubectl run --generator=run-pod/v1 tmp --rm -i --tty --serviceaccount=vault-auth --image alpine:3.7
    
    apk update
    apk add curl jq
    
    VAULT_ADDR=http://vault:8200
    
    curl -s $VAULT_ADDR/v1/sys/health | jq
    
    
    KUBE_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
    echo $KUBE_TOKEN
    
    curl --request POST \
            --data '{"jwt": "'"$KUBE_TOKEN"'", "role": "example"}' \
            $VAULT_ADDR/v1/auth/kubernetes/login | jq
    ```
- Deploy Nginx microservice
  - All vault auth agent config and nginx config are placed in config.yaml
  - Deploy microservice using deployment.yaml  
    
    
