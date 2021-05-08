"use strict";
var MsgTypes;
(function (MsgTypes) {
    MsgTypes[MsgTypes["Drawing"] = 0] = "Drawing";
    MsgTypes[MsgTypes["ObsStat"] = 1] = "ObsStat";
    MsgTypes[MsgTypes["Chview"] = 2] = "Chview";
    MsgTypes[MsgTypes["OutChatMsg"] = 3] = "OutChatMsg";
    MsgTypes[MsgTypes["InpChatMsg"] = 4] = "InpChatMsg";
})(MsgTypes || (MsgTypes = {}));
var DrawingTypes;
(function (DrawingTypes) {
    DrawingTypes[DrawingTypes["FreeMouse"] = 0] = "FreeMouse";
})(DrawingTypes || (DrawingTypes = {}));
var Point = /** @class */ (function () {
    function Point() {
        this.x = 0;
        this.y = 0;
    }
    return Point;
}());
var BasicBoard = /** @class */ (function () {
    function BasicBoard(ws) {
        var _this = this;
        this.ws = ws;
        this.drawing = false;
        this.x = 0;
        this.y = 0;
        this.isDrawable = true;
        this.snap = Snap("#svg");
        this.snap.attr({
            strokeWidth: 2,
            stroke: "#000"
        });
        this.curGroup = this.snap.group();
        this.allGroups = [];
        this.snap.mousedown(function (e) { _this.mousedown(e); });
        this.snap.mousemove(function (e) { _this.mousemove(e); });
        window.addEventListener('mouseup', function (e) { _this.mouseup(e); });
    }
    BasicBoard.prototype.clear = function () {
        this.x = 0;
        this.y = 0;
        this.drawing = false;
        this.snap.clear();
    };
    BasicBoard.prototype.generalDraw = function (e) {
        this.x = e.offsetX;
        this.y = e.offsetY;
    };
    BasicBoard.prototype.mousedown = function (e) {
        this.generalDraw(e);
        this.drawing = true;
    };
    BasicBoard.prototype.mousemove = function (e) {
        if (this.drawing === true) {
            this.curGroup.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY));
            this.generalDraw(e);
        }
    };
    BasicBoard.prototype.mouseup = function (e) {
        if (this.drawing === true) {
            this.curGroup.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY));
            this.generalDraw(e);
            this.drawing = false;
            this.allGroups.push(this.curGroup);
            this.sendDrawing(DrawingTypes.FreeMouse, this.curGroup);
            this.curGroup = this.snap.group();
        }
    };
    BasicBoard.prototype.canvasBack = function () {
    };
    BasicBoard.prototype.drawPackage = function (msg) {
        switch (msg.type) {
            case DrawingTypes.FreeMouse:
                for (var i = 0; i < msg.data.length - 1; i++) {
                    this.snap.line(msg.data[i].x, msg.data[i].y, msg.data[i + 1].x, msg.data[i + 1].y);
                }
                break;
        }
    };
    BasicBoard.prototype.sendDrawing = function (type, fig) {
        switch (type) {
            case DrawingTypes.FreeMouse:
                var cords = [];
                for (var i = 0; i < fig.children().length; i++) {
                    var cattrs = fig.children()[i].toJSON().attr;
                    cords.push({
                        x: cattrs.x1,
                        y: cattrs.y1
                    });
                    if (i == fig.children().length - 1) {
                        cords.push({
                            x: cattrs.x2,
                            y: cattrs.y2
                        });
                    }
                }
                this.ws.send(JSON.stringify({
                    type: MsgTypes.Drawing,
                    data: {
                        type: DrawingTypes.FreeMouse,
                        data: cords
                    }
                }));
                break;
        }
    };
    return BasicBoard;
}());
