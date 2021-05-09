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
var ObserverBoard = /** @class */ (function (_super) {
    __extends(ObserverBoard, _super);
    function ObserverBoard(ws) {
        var _this = _super.call(this, ws) || this;
        _this.isDrawable = false;
        return _this;
    }
    ObserverBoard.prototype.mousedown = function (e) {
        if (this.isDrawable) {
            _super.prototype.mousedown.call(this, e);
        }
    };
    ObserverBoard.prototype.mousemove = function (e) {
        if (this.isDrawable) {
            _super.prototype.mousemove.call(this, e);
        }
    };
    ObserverBoard.prototype.mouseup = function (e) {
        if (this.isDrawable) {
            _super.prototype.mouseup.call(this, e);
        }
    };
    ObserverBoard.prototype.stepBack = function () {
        if (this.isDrawable) {
            _super.prototype.stepBack.call(this);
        }
    };
    return ObserverBoard;
}(BasicBoard));
