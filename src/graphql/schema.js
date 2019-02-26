const {
    GraphQLSchema
} = require('graphql');


const BookQueryType = require('./queries.js');
const BookMutationType = require('./mutations.js');

const BookSchema = new GraphQLSchema({
    query: BookQueryType,
    mutation: BookMutationType
})



module.exports = BookSchema