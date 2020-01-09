mkdir -p /etc/pki/trust/anchors/
mv cacert.pem /etc/pki/trust/anchors/
update-ca-trust
service docker reload