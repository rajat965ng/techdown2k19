#Post Installation Steps

### Execute "kubectl get ingress" and note "Address"
### Create '.conf' file like:
    location /dev {
        proxy_pass http://$Address/;
    }
### Move '.conf' file to /etc/nginx/default.d/
### Reload nginx
    sudo systemctl reload nginx