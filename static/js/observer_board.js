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
    function msgParser(e) {
        var msg = JSON.parse(e.data);
        switch (msg.type) {
            case MsgTypes.Points:
                board.drawPackage(msg.data.points);
                break;
            case MsgTypes.InpChatMsg:
                chat.newMessage(msg.data);
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
    var ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname);
    var board = new ObserverBoard(ws);
    var chat = new BasicChat(document.getElementById("chat"), ws);
    initPage();
    ws.onmessage = msgParser;
})(observer_board || (observer_board = {}));
