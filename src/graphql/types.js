const {
    GraphQLString,
    GraphQLNonNull,
    GraphQLObjectType,
    GraphQLInputObjectType
} = require('graphql');

const BookType = new GraphQLObjectType({
    name: 'BookType',
    description: 'Book List',
    fields: {
        id: {type: GraphQLString},
        name: {type: GraphQLString},
        author: {type: GraphQLString}
    }
})

const findByBookName = new GraphQLInputObjectType({
    name: 'findByBookName',
    description: 'find a book by its name',
    type: BookType,
    fields: {
        name: {type: GraphQLNonNull(GraphQLString)}
    }
})

const BookCreateType = new GraphQLInputObjectType({
    name: 'BookCreateType',
    description: 'Add a book to the list',
    type: BookType,
    fields: {
        name: {type: new GraphQLNonNull(GraphQLString)},
        author: {type: new GraphQLNonNull(GraphQLString)}
    }
})

const BookUpdateType = new GraphQLInputObjectType({
    name: 'BookUpdateType',
    description: 'Update a book in the list',
    type: BookType,
    fields: {
        id: {type: new GraphQLNonNull(GraphQLString)},
        name: {type: new GraphQLNonNull(GraphQLString)},
        author: {type: new GraphQLNonNull(GraphQLString)}
    }
})

const BookDeleteType = new GraphQLInputObjectType({
    name: 'BookDeleteType',
    description: 'Delete a book by Id in the list',
    type: BookType,
    fields: {
        id: {type: new GraphQLNonNull(GraphQLString)}
    }
})

module.exports = {
    BookType,
    findByBookName,
    BookCreateType,
    BookUpdateType,
    BookDeleteType
}