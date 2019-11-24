var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var UserSchema = new Schema({
    name: {
        type: String,
        required: 'Kindly enter the name of user'
    },
    department: {
        type: String,
        required: "Kindly enter department"
    },
    email: {
        type: String
    },
    username: {
        type: String,
        required: "Kindly enter username"
    },
    password: {
        type: String,
        required: "Kindly enter password"
    },
})

module.export = mongoose.model("User", UserSchema)