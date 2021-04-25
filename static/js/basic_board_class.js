"use strict";
var Point = /** @class */ (function () {
    function Point() {
        this.x = 0;
        this.y = 0;
    }
    return Point;
}());
var DrawingBoard = /** @class */ (function () {
    function DrawingBoard() {
        var _this = this;
        this.msgParser = function () {
            console.log("from default parser");
        };
        this.ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname);
        this.ws.onmessage = function (e) { _this.msgParser(_this, e); };
        this.drawing = false;
        this.x = 0;
        this.y = 0;
        this.canvas = document.getElementById('canvas');
        this.ctx = this.canvas.getContext('2d');
        this.canvas.width = 500;
        this.canvas.height = 500;
        this.sendBuf = [];
        this.isDrawable = true;
        this.canvas.addEventListener('mousedown', function (e) { _this.mousedown(e); });
        this.canvas.addEventListener('mousemove', function (e) { _this.mousemove(e); });
        window.addEventListener('mouseup', function (e) { _this.mouseup(e); });
    }
    DrawingBoard.prototype.clear = function () {
        this.drawing = false;
        this.sendBuf = [];
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
    };
    DrawingBoard.prototype.mousedown = function (e) {
        this.x = e.offsetX;
        this.y = e.offsetY;
        this.drawing = true;
        this.sendBuf.push({
            x: this.x,
            y: this.y
        });
        this.checkBuf();
    };
    DrawingBoard.prototype.mousemove = function (e) {
        if (this.drawing === true) {
            this.drawLine(this.x, this.y, e.offsetX, e.offsetY);
            this.x = e.offsetX;
            this.y = e.offsetY;
            this.sendBuf.push({
                x: this.x,
                y: this.y
            });
            this.checkBuf();
        }
    };
    DrawingBoard.prototype.mouseup = function (e) {
        if (this.drawing === true) {
            this.drawLine(this.x, this.y, e.offsetX, e.offsetY);
            this.sendBuf.push({
                x: this.x,
                y: this.y
            });
            this.sendBuf.push({
                x: e.offsetX,
                y: e.offsetY
            });
            this.x = 0;
            this.y = 0;
            this.sendPoints();
            this.sendBuf = [];
            this.drawing = false;
        }
    };
    DrawingBoard.prototype.drawPackage = function (points) {
        for (var i = 0; i < points.length - 1; i++) {
            this.drawLine(points[i].x, points[i].y, points[i + 1].x, points[i + 1].y);
        }
    };
    DrawingBoard.prototype.sendPoints = function () {
        this.ws.send(JSON.stringify({
            type: MsgTypes.Points,
            data: {
                points: this.sendBuf,
            }
        }));
        var pv = this.sendBuf[this.sendBuf.length - 1];
        this.sendBuf = [];
        this.sendBuf.push(pv);
    };
    DrawingBoard.prototype.checkBuf = function () {
        if (this.sendBuf.length >= 5) {
            this.sendPoints();
        }
    };
    DrawingBoard.prototype.drawLine = function (x1, y1, x2, y2) {
        this.ctx.beginPath();
        this.ctx.strokeStyle = 'black';
        this.ctx.lineWidth = 1;
        this.ctx.moveTo(x1, y1);
        this.ctx.lineTo(x2, y2);
        this.ctx.stroke();
        this.ctx.closePath();
    };
    return DrawingBoard;
}());
var MsgTypes;
(function (MsgTypes) {
    MsgTypes[MsgTypes["Points"] = 0] = "Points";
    MsgTypes[MsgTypes["ObsStat"] = 1] = "ObsStat";
    MsgTypes[MsgTypes["Chview"] = 2] = "Chview";
})(MsgTypes || (MsgTypes = {}));
