"use strict";
var admin_board;
(function (admin_board) {
    function initPage() {
        $("#views-nav").find("#general-page").on("click", switchView);
    }
    function msgParser(e) {
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
            case MsgTypes.InpChatMsg:
                chat.newMessage(msg.data);
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
    var ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname);
    var board = new AdminBoard(ws);
    var chat = new BasicChat(document.getElementById("chat"), ws);
    initPage();
    ws.onmessage = msgParser;
})(admin_board || (admin_board = {}));
