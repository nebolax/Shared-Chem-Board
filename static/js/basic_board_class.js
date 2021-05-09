"use strict";
var MsgTypes;
(function (MsgTypes) {
    MsgTypes[MsgTypes["Action"] = 0] = "Action";
    MsgTypes[MsgTypes["SetId"] = 1] = "SetId";
    MsgTypes[MsgTypes["ObsStat"] = 2] = "ObsStat";
    MsgTypes[MsgTypes["Chview"] = 3] = "Chview";
    MsgTypes[MsgTypes["OutChatMsg"] = 4] = "OutChatMsg";
    MsgTypes[MsgTypes["InpChatMsg"] = 5] = "InpChatMsg";
})(MsgTypes || (MsgTypes = {}));
var DrawingTypes;
(function (DrawingTypes) {
    DrawingTypes[DrawingTypes["FreeMouse"] = 0] = "FreeMouse";
})(DrawingTypes || (DrawingTypes = {}));
var ActionTypes;
(function (ActionTypes) {
    ActionTypes[ActionTypes["NewDrawing"] = 0] = "NewDrawing";
    ActionTypes[ActionTypes["DrawingDeleted"] = 1] = "DrawingDeleted";
})(ActionTypes || (ActionTypes = {}));
var Modes;
(function (Modes) {
    Modes[Modes["Drawing"] = 0] = "Drawing";
    Modes[Modes["Dragging"] = 1] = "Dragging";
})(Modes || (Modes = {}));
var Point = /** @class */ (function () {
    function Point() {
        this.x = 0;
        this.y = 0;
    }
    return Point;
}());
var Drawing = /** @class */ (function () {
    function Drawing(fig) {
        this.id = 0;
        this.type = DrawingTypes.FreeMouse;
        this.fig = fig;
        this.data = [];
    }
    Drawing.prototype.computePlainData = function () {
        switch (this.type) {
            case DrawingTypes.FreeMouse:
                var cords = [];
                for (var i = 0; i < this.fig.children().length; i++) {
                    var cattrs = this.fig.children()[i].toJSON().attr;
                    cords.push({
                        x: cattrs.x1,
                        y: cattrs.y1
                    });
                    if (i == this.fig.children().length - 1) {
                        cords.push({
                            x: cattrs.x2,
                            y: cattrs.y2
                        });
                    }
                }
                this.data = cords;
                break;
        }
    };
    return Drawing;
}());
var Action = /** @class */ (function () {
    function Action(drawing) {
        this.id = 0;
        this.type = ActionTypes.NewDrawing;
        this.drawing = drawing;
    }
    return Action;
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
        this.curDrawing = new Drawing(this.snap.group());
        this.allDrawings = [];
        this.actions = [];
        this.snap.mousedown(function (e) { _this.mousedown(e); });
        this.snap.mousemove(function (e) { _this.mousemove(e); });
        window.addEventListener('mouseup', function (e) { _this.mouseup(e); });
    }
    BasicBoard.prototype.exportPicture = function () {
        // let pic = this.snap.outerS
    };
    BasicBoard.prototype.newDrawingID = function (id) {
        for (var i = 0; i < this.allDrawings.length; i++) {
            var el = this.allDrawings[i];
            if (el.id == 0) {
                el.id = id;
                break;
            }
        }
    };
    BasicBoard.prototype.newActionID = function (id) {
        for (var i = 0; i < this.actions.length; i++) {
            var el = this.actions[i];
            if (el.id == 0) {
                el.id = id;
                break;
            }
        }
    };
    BasicBoard.prototype.clear = function () {
        this.x = 0;
        this.y = 0;
        this.drawing = false;
        this.snap.clear();
        this.curDrawing = new Drawing(this.snap.group());
        this.allDrawings = [];
        this.actions = [];
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
        if (this.drawing) {
            this.curDrawing.fig.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY));
            this.generalDraw(e);
        }
    };
    BasicBoard.prototype.mouseup = function (e) {
        if (this.drawing) {
            this.drawing = false;
            this.curDrawing.fig.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY));
            this.generalDraw(e);
            var action = {
                id: 0,
                type: ActionTypes.NewDrawing,
                drawing: this.curDrawing
            };
            this.allDrawings.push(this.curDrawing);
            this.actions.push(action);
            this.sendAction(action);
            this.curDrawing = new Drawing(this.snap.group());
        }
    };
    BasicBoard.prototype.stepBack = function () {
        if (this.allDrawings.length > 0) {
            if (this.allDrawings[this.allDrawings.length - 1].id > 0) {
                var last = this.allDrawings.pop();
                last === null || last === void 0 ? void 0 : last.fig.remove();
                this.ws.send(JSON.stringify({
                    type: MsgTypes.Action,
                    data: {
                        id: last === null || last === void 0 ? void 0 : last.id,
                        type: ActionTypes.DrawingDeleted
                    }
                }));
            }
        }
    };
    BasicBoard.prototype.newAction = function (msg) {
        switch (msg.type) {
            case ActionTypes.NewDrawing:
                switch (msg.drawing.type) {
                    case DrawingTypes.FreeMouse:
                        var drawing = new Drawing(this.snap.group());
                        drawing.id = msg.drawing.id;
                        for (var i = 0; i < msg.drawing.data.length - 1; i++) {
                            drawing.fig.append(this.snap.line(msg.drawing.data[i].x, msg.drawing.data[i].y, msg.drawing.data[i + 1].x, msg.drawing.data[i + 1].y));
                        }
                        this.allDrawings.push(drawing);
                        msg.drawing = drawing;
                        this.actions.push(msg);
                        break;
                }
                break;
            case ActionTypes.DrawingDeleted:
                var res = [];
                for (var i = 0; i < this.allDrawings.length; i++) {
                    if (this.allDrawings[i].id != msg.id) {
                        res.push(this.allDrawings[i]);
                    }
                    else {
                        this.allDrawings[i].fig.remove();
                    }
                }
                this.allDrawings = res;
                break;
        }
    };
    BasicBoard.prototype.sendAction = function (action) {
        action.drawing.computePlainData();
        var ts = JSON.stringify({
            type: MsgTypes.Action,
            data: action
        }, function (key, val) { return key == "fig" ? undefined : val; });
        this.ws.send(ts);
    };
    return BasicBoard;
}());
