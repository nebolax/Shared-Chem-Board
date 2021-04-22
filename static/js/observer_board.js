"use strict";
var observer_board;
(function (observer_board) {
    function initPage() {
        $("#general-board").on("click", function () {
            switchBoard(1);
        });
        $("#personal-board").on("click", function () {
            switchBoard(2);
        });
    }
    function msgParser(board, e) {
        console.log("r");
        var msg = JSON.parse(e.data);
        if (msg.type == "points") {
            board.drawPackage(msg.points);
        }
    }
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
    initPage();
    var board = new ObserverBoard(msgParser);
})(observer_board || (observer_board = {}));
