"use strict";
var admin_board;
(function (admin_board) {
    function msgParser(board, e) {
        var msg = JSON.parse(e.data);
        switch (msg.type) {
            case MsgTypes.Points:
                board.drawPackage(msg.data.points);
                break;
            case MsgTypes.ObsStat:
                var clone = $("#observer-bar").clone();
                clone.attr("id", "user" + msg.userID);
                clone.html(msg.username);
                $("observers-nav").append(clone);
                break;
        }
    }
    var board = new AdminBoard(msgParser);
})(admin_board || (admin_board = {}));
