"use strict";
function msgParser(board, e) {
    console.log("r");
    var msg = JSON.parse(e.data);
    if (msg.type == "points") {
        board.drawPackage(msg.points);
    }
    else if (msg.type == "newObserver") {
        var el = document.createElement("div");
        el.innerHTML = "<button></button>";
    }
}
var board = new AdminBoard(msgParser);
var e = document.createElement("div");
e.innerHTML = "<p>Hi there</p>";
