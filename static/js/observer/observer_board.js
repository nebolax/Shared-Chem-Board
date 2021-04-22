"use strict";
function msgParser(board, e) {
    console.log("r");
    var msg = JSON.parse(e.data);
    if (msg.type == "points") {
        board.drawPackage(msg.points);
    }
}
var board = new ObserverBoard(msgParser);
function switchBoard(id) {
    console.log(id);
    switch (id) {
        case 2:
            board.toPersonalBoard();
            break;
        default:
            board.toGeneralBoard();
            break;
    }
}
