"use strict";
var ChatMsgContent = /** @class */ (function () {
    function ChatMsgContent(msgContent) {
        this.text = msgContent.text;
    }
    return ChatMsgContent;
}());
var TimeStamp = /** @class */ (function () {
    function TimeStamp(msgStamp) {
        this.year = msgStamp.year;
        this.month = msgStamp.month;
        this.day = msgStamp.day;
        this.hour = msgStamp.hour;
        this.minute = msgStamp.minute;
    }
    return TimeStamp;
}());
var ChatMessage = /** @class */ (function () {
    function ChatMessage(msg) {
        this.id = msg.id;
        this.senderid = msg.senderid;
        this.timestamp = new TimeStamp(msg.timestamp);
        this.content = new ChatMsgContent(msg.content);
    }
    return ChatMessage;
}());
var BasicChat = /** @class */ (function () {
    function BasicChat(chatTag, ws) {
        var _a;
        this.history = [];
        this.chatTag = chatTag;
        this.ws = ws;
        (_a = chatTag.querySelector("#send-new-chat-msg")) === null || _a === void 0 ? void 0 : _a.addEventListener("click", this.sendMessage);
    }
    BasicChat.prototype.sendMessage = function (e) {
        var _a;
        var msgText = (_a = this.chatTag.querySelector("#new-chat-msg-text")) === null || _a === void 0 ? void 0 : _a.textContent;
        if (msgText == null || msgText == undefined) {
            alert("Вы должны ввести хотя бы какой-то текст");
        }
        else {
            this.ws.send(JSON.stringify({
                type: MsgTypes.OutChatMsg,
                data: {
                    text: msgText
                }
            }));
        }
    };
    BasicChat.prototype.loadHistory = function (msgHist) {
        this.history = msgHist.history;
    };
    BasicChat.prototype.newMessage = function (msg) {
        var _a;
        this.history.push(msg);
        var templ = this.chatTag.querySelector("#template-chatmsg");
        var clone = document.importNode(templ.content, true);
        clone.querySelector(".chatmsg-text").innerHTML = msg.content.text;
        (_a = this.chatTag.firstChild) === null || _a === void 0 ? void 0 : _a.appendChild(clone);
    };
    return BasicChat;
}());
