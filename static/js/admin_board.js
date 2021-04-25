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
                msg = msg.data;
                var clone = $("#view0").clone();
                clone.attr("id", "view" + msg.userID);
                clone.find("#chviewBtn").html(msg.username);
                clone.find("#chviewBtn").on("click", switchView);
                $("#observers-nav").append(clone);
                break;
        }
    }
    function switchView(e) {
        var nview = +e.target.parentElement.id.slice(4);
        board.toPersonal(nview);
    }
    var board = new AdminBoard(msgParser);
    $("#view0").find("#chviewBtn").on("click", switchView);
})(admin_board || (admin_board = {}));
