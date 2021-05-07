"use strict";
var MsgTypes;
(function (MsgTypes) {
    MsgTypes[MsgTypes["Points"] = 0] = "Points";
    MsgTypes[MsgTypes["ObsStat"] = 1] = "ObsStat";
    MsgTypes[MsgTypes["Chview"] = 2] = "Chview";
    MsgTypes[MsgTypes["OutChatMsg"] = 3] = "OutChatMsg";
    MsgTypes[MsgTypes["InpChatMsg"] = 4] = "InpChatMsg";
})(MsgTypes || (MsgTypes = {}));
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
        this.sendBuf = [];
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
        this.sendBuf = [];
        this.snap.clear();
    };
    BasicBoard.prototype.generalDraw = function (e) {
        this.x = e.offsetX;
        this.y = e.offsetY;
        this.sendBuf.push({
            x: this.x,
            y: this.y
        });
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
            this.sendPoints();
            this.drawing = false;
            this.allGroups.push(this.curGroup);
            this.curGroup = this.snap.group();
            console.log(this.allGroups);
        }
    };
    BasicBoard.prototype.canvasBack = function () {
    };
    BasicBoard.prototype.drawPackage = function (points) {
        for (var i = 0; i < points.length - 1; i++) {
            this.snap.line(points[i].x, points[i].y, points[i + 1].x, points[i + 1].y);
        }
    };
    BasicBoard.prototype.sendPoints = function () {
        this.ws.send(JSON.stringify({
            type: MsgTypes.Points,
            data: {
                points: this.sendBuf
            }
        }));
        this.sendBuf = [];
    };
    return BasicBoard;
}());
