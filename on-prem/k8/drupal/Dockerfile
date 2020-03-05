FROM docker
USER root
RUN docker pull drupal:8
ENTRYPOINT ["docker run -p 8080:80 drupal:8"]
EXPOSE 8080