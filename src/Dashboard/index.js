const express = require("express");
let path = require("path");
let cors = require('cors');
const request = require('request');

const port = 3000;
const app = express();

app.use(cors());

// Sub Node Dashboard
app.get("/", function (req, res) {
    res.header("Access-Control-Allow-Origin", "*");
    req.header("Access-Control-Allow-Origin", "*");
    res.sendFile(path.join(__dirname + "/static" + "/index.html"))
});

// Gets information from Master Node and passes it to Dashboard
app.get("/getUpdate", function (req, res) {
    request('http://35.243.155.9:3141/api/getInfo', {
        json: true
    }, (err, _, body) => {
        if (err) {
            return console.log(err);
        }
        console.log(body);
        res.json(body);
    });
});

app.listen(port, "127.0.0.1", function () {
    console.log("Listening on port %s...", port);
});