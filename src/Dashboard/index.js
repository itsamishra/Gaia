const express = require("express");
let path = require("path");

const port = 3000;
const app = express();

console.log(path.join("static", "index.html"));

app.get("/", function(req, res){
    res.sendFile(path.join(__dirname+"/static"+"/index.html"))
});

app.listen(port, "127.0.0.1", function(){
    console.log("Listening...");
});