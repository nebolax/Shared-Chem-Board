"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var AdminBoard = /** @class */ (function (_super) {
    __extends(AdminBoard, _super);
    function AdminBoard(msgParser) {
        var _this = _super.call(this) || this;
        _this.curBoardID = 0;
        _this.msgParser = _this.msgParser = function (b, e) {
            msgParser(b, e);
        };
        return _this;
    }
    AdminBoard.prototype.toPersonal = function (viewID) {
        _super.prototype.clear.call(this);
        console.log(viewID);
        this.ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: viewID
            }
        }));
    };
    AdminBoard.prototype.toGeneral = function () {
        _super.prototype.clear.call(this);
        this.ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 0
            }
        }));
    };
    return AdminBoard;
}(DrawingBoard));
