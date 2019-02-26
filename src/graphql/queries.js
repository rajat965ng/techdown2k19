
const _ = require('lodash');

const {
    GraphQLList,
    GraphQLObjectType,
    GraphQLNonNull
} = require('graphql');

const {
    BookType,
    findByBookName
} = require('./types.js');

const Books = require('../../data/books.js');

const BookQueryType  = new GraphQLObjectType({
    name: 'BookQueryType',
    description: 'Query Schema for BookType',
    fields: {
        books: {
            type: new GraphQLList(BookType),
            resolve: () => Books
        },
        findByBookName:{
            type: BookType,
            args: {
                input: {type: new GraphQLNonNull(findByBookName)}
            },
            resolve: (source,{ input }) => {
                console.log("input is",input);
                return _.find(Books,b => b.name == input.name);
            }
        }
    }
});


module.exports = BookQueryType;