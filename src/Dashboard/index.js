const express = require("express");
let path = require("path");

const port = 3000;
const app = express();

// app.use((req, res, next) => {
//     res.header("Access-Control-Allow-Origin", "*");
//     res.header("Access-Control-Allow-Headers", "*");

//     if (req.method == "OPTIONS"){
//         res.header("Access-Control-Allow-Methods", "*");

//         return res.status(200).json({});
//     }
//     next()
// });

app.get("/", function (req, res) {
    res.header("Access-Control-Allow-Origin", "*");
    req.header("Access-Control-Allow-Origin", "*");
    res.sendFile(path.join(__dirname + "/static" + "/index.html"))
});

app.listen(port, "127.0.0.1", function () {
    console.log("Listening on port %s...", port);
});