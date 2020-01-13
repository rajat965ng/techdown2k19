#Post Installation Steps

### Execute "kubectl get ingress" in Kubernetes Master VM and note "Address"
### Create 'http-abc.conf' file like:
    location / {
        proxy_pass http://$Address/;
    }
### Move 'http-gpssa.conf' file to /etc/nginx/default.d/
### Reload nginx
    sudo systemctl reload nginx