<h2>Configure Vault with Nginx</h2>



```
cat <<EOF | vault policy write myapp-kv-ro -
path "secret/myapp/*" { capabilities = ["read", "list"] }
path "secret/data/myapp/*" { capabilities = ["read", "list"] }
EOF

vault auth enable userpass

vault write auth/userpass/users/test-user \
        password=training \
        policies=myapp-kv-ro
        

vault login -method=userpass \
        username=test-user \
        password=training
        
vault kv get secret/myapp/config
                
```
```
kubectl get sa vault-auth -o jsonpath="{.secrets[*]['name']}"

O/P: vault-auth-token-bwhtd , That's why we use RBACs in kubernetes

kubectl get secret vault-auth-token-bwhtd -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo


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

```yaml

kubectl run --generator=run-pod/v1 tmp --rm -i --tty --serviceaccount=vault-auth --image alpine:3.7

apk update
apk add curl jq

VAULT_ADDR=http://vault:8200

curl -s $VAULT_ADDR/v1/sys/health | jq


KUBE_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
echo $KUBE_TOKEN
```

Test the kubernetes auth method to ensure that you can authenticate with Vault.

```
curl --request POST \
        --data '{"jwt": "'"$KUBE_TOKEN"'", "role": "example"}' \
        $VAULT_ADDR/v1/auth/kubernetes/login | jq
```

Leverage Vault Agent Auto-Auth
    - Edit config.yaml file to create config map as per your requirement.
    - Use config map in your deployment/pod. 
    
    
curl --header "X-VAULT-TOKEN: `cat /home/vault/.vault-token`" http://vault:8200/v1/secret/data/myapp/config | jq .data.data.username
    
export username=$(curl --header "X-VAULT-TOKEN: `cat /home/vault/.vault-token`" http://vault:8200/v1/secret/data/myapp/config | jq .data.data.username) && sed -i "s/username/$username/g" /usr/share/nginx/html/index.html    

