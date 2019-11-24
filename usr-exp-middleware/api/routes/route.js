
module.exports = function (app) {
    var userList = require('../controller/controller')


    app.route('/users')
        .get(userList.list_all_users)
        .post(userList.create_user)
}