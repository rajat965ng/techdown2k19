const express = require('express');

const PORT = process.env.PORT || 3000;
const app = express();

app.listen(PORT);

console.log("GraphQL server API up and running at localhost ",PORT);

const graphqlHttp = require('express-graphql');


const schema = require('./src/graphql/schema.js');

app.use('/library', graphqlHttp({
    schema: schema,
    graphiql: true
}));