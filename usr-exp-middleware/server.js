var express = require('express'),
    cors = require('cors'),
    app = express(),
    port = process.env.PORT || 9000,
    mongoose = require('mongoose'),
    User = require("./api/model/model"),
    bodyParser = require('body-parser');

mongoose.Promise = global.Promise;
mongoose.connect('mongodb://localhost/UsrDb');

app.use(cors())
app.use(bodyParser.urlencoded({extended: true}));
app.use(bodyParser.json());

var routes = require('./api/routes/route');
routes(app);

app.listen(port);

console.log("The server is running on port: " + port);