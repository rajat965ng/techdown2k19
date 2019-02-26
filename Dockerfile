FROM yarnpkg/node-yarn:latest
WORKDIR /opt/graphqlapp
COPY  . .
ENTRYPOINT ["yarn","start"]