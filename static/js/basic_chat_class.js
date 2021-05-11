"use strict";
var Observer = /** @class */ (function () {
    function Observer(userid, username) {
        this.userid = userid;
        this.username = username;
    }
    return Observer;
}());
var User = /** @class */ (function () {
    function User(msg) {
        this.nickname = msg.nickname;
    }
    return User;
}());
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
    TimeStamp.prototype.time = function () {
        return this.hour + ":" + this.minute;
    };
    TimeStamp.prototype.date = function () {
        return this.day + "." + this.month + "." + this.year;
    };
    return TimeStamp;
}());
var ChatMessage = /** @class */ (function () {
    function ChatMessage(msg) {
        this.id = msg.id;
        this.sender = new User(msg.senderinfo);
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
        this.msgInput.value = "";
        if (msgText == null || msgText == undefined || msgText == "") {
            alert("Вы должны ввести хотя бы какой-то текст");
        }
        else {
            var outMsg = {
                type: MsgTypes.OutChatMsg,
                data: {
                    text: msgText
                }
            };
            this.ws.send(JSON.stringify(outMsg));
        }
    };
    BasicChat.prototype.loadHistory = function (msgHist) {
        var _this = this;
        this.clear();
        this.history = msgHist.history;
        this.history.forEach(function (el) {
            _this.newMessage(el);
        });
    };
    BasicChat.prototype.newMessage = function (inpMsg) {
        console.log(inpMsg);
        var msg = new ChatMessage(inpMsg);
        this.history.push(msg);
        var clone = document.importNode(this.msgTemplate.content, true);
        clone.querySelector(".chatmsg-text").innerHTML = msg.content.text;
        var timestamp = msg.timestamp.time();
        clone.querySelector(".chatmsg-info").innerHTML = msg.sender.nickname + "  -  " + timestamp;
        this.chatContainer.appendChild(clone);
    };
    return BasicChat;
}());
