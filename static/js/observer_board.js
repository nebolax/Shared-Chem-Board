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
    function msgParser(b, e) {
        var msg = JSON.parse(e.data);
        switch (msg.type) {
            case MsgTypes.Points:
                board.drawPackage(msg.data.points);
                break;
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
