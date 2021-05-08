"use strict";
var admin_board;
(function (admin_board) {
    function initPage() {
        $("#views-nav").find("#general-page").on("click", switchView);
        $("#stepback").on("click", function () { board.stepBack(); });
        $("#dragbtn").on("click", function () { board.switchDragMode(); });
    }
    function msgParser(e) {
        var msg = JSON.parse(e.data);
        switch (msg.type) {
            case MsgTypes.Action:
                console.log(msg);
                board.drawPackage(msg.data);
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
            toGeneral();
        }
        else {
            var nview = +sourceId.slice(4);
            toPersonal(nview);
        }
    }
    function toPersonal(viewID) {
        board.clear();
        chat.clear();
        ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: viewID
            }
        }));
    }
    function toGeneral() {
        board.clear();
        chat.clear();
        ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 0
            }
        }));
    }
    var ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname);
    var board = new AdminBoard(ws);
    var chat = new BasicChat(document.getElementById("chat"), ws);
    initPage();
    ws.onmessage = msgParser;
})(admin_board || (admin_board = {}));
