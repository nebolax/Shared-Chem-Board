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
        var _this = this;
        this.history = [];
        this.chatTag = chatTag;
        this.msgTemplate = this.chatTag.querySelector("#template-chatmsg");
        this.msgInput = this.chatTag.querySelector("#new-chat-msg-text");
        this.chatContainer = this.chatTag.querySelector("#chat-container");
        this.ws = ws;
        chatTag.querySelector("#send-new-chat-msg").addEventListener("click", function () { _this.sendMessage(); });
    }
    BasicChat.prototype.clear = function () {
        this.history = [];
        this.chatContainer.innerHTML = "";
    };
    BasicChat.prototype.sendMessage = function () {
        var msgText = this.msgInput.value;
        console.log("text: " + msgText);
        this.msgInput.value = "";
        if (msgText == null || msgText == undefined) {
            alert("Вы должны ввести хотя бы какой-то текст");
        }
        else {
            var outMsg = {
                type: MsgTypes.OutChatMsg,
                data: {
                    text: msgText
                }
            };
            console.log(outMsg);
            this.ws.send(JSON.stringify(outMsg));
        }
    };
    BasicChat.prototype.loadHistory = function (msgHist) {
        this.history = msgHist.history;
    };
    BasicChat.prototype.newMessage = function (msg) {
        this.history.push(msg);
        var clone = document.importNode(this.msgTemplate.content, true);
        clone.querySelector(".chatmsg-text").innerHTML = msg.content.text;
        this.chatContainer.appendChild(clone);
    };
    return BasicChat;
}());
