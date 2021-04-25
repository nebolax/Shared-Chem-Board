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
                $("#observers-nav").empty();
                msg.allObsInfo.forEach(function (el) {
                    var _a;
                    var templ = document.getElementById("template-obsname");
                    var clone = document.importNode(templ.content, true);
                    var btn = clone.querySelector("#chviewBtn");
                    btn.addEventListener("click", switchView);
                    btn.innerHTML = el.username;
                    btn.id = "view" + el.userid;
                    (_a = document.getElementById("observers-nav")) === null || _a === void 0 ? void 0 : _a.appendChild(clone);
                });
                break;
        }
    }
    function switchView(e) {
        var sourceId = e.target.id;
        if (sourceId == "general-page") {
            board.toGeneral();
        }
        else {
            var nview = +sourceId.slice(4);
            board.toPersonal(nview);
        }
    }
    var board = new AdminBoard(msgParser);
    $("#views-nav").find("#general-page").on("click", switchView);
})(admin_board || (admin_board = {}));
