"use strict";
function toPersonal(viewID) {
    _super.clear.call(this);
    console.log(viewID);
    this.ws.send(JSON.stringify({
        type: MsgTypes.Chview,
        data: {
            nview: viewID
        }
    }));
}
function toGeneral() {
    _super.clear.call(this);
    this.ws.send(JSON.stringify({
        type: MsgTypes.Chview,
        data: {
            nview: 0
        }
    }));
}
var ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname);
