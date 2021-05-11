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
        $("#stepback").on("click", function () { board.stepBack(); });
        $("#exportimg").on("click", function () { board.exportPicture(); });
    }
    function msgParser(e) {
        var msg = JSON.parse(e.data);
        console.log(msg);
        switch (msg.type) {
            case MsgTypes.Action:
                board.newAction(msg.data);
                break;
            case MsgTypes.SetId:
                switch (msg.data.property) {
                    case "action":
                        board.newActionID(msg.data.id);
                        break;
                    case "drawing":
                        board.newDrawingID(msg.data.id);
                        break;
                }
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
                toPersonalBoard();
                break;
            default:
                toGeneralBoard();
                break;
        }
    }
    function toGeneralBoard() {
        board.clear();
        chat.clear();
        board.isDrawable = false;
        ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 0
            }
        }));
    }
    function toPersonalBoard() {
        board.clear();
        chat.clear();
        board.isDrawable = true;
        ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 1
            }
        }));
    }
    var ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname);
    var board = new ObserverBoard(ws);
    var chat = new BasicChat(document.getElementById("chat"), ws);
    initPage();
    ws.onmessage = msgParser;
})(observer_board || (observer_board = {}));
