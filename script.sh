echo hi
sudo mkdir -p /var/opt/jfrog/artifactory
sudo useradd artifactory
sudo usermod -aG docker artifactory
sudo chmod -R 777 /var/opt/jfrog/artifactory
docker run -d -v /var/opt/jfrog/artifactory:/var/opt/jfrog/artifactory -p 8081:8081 docker.bintray.io/jfrog/artifactory-pro:latest