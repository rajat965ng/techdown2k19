var mongoose  = require('mongoose'),
    User = mongoose.model('User');

exports.list_all_users = function (req,res) {
    User.find({}, (err,user) => {
        if (err)
            res.send(err)
        res.json(user)
    })
};

exports.create_user = function (req,res) {
    var new_user = new User(req.body);
    new_user.save((err,user) => {
        if (err)
            res.send(err)
        res.json(user)
    })
}