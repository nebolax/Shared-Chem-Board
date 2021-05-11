class Observer {
    userid: number;
    username: string;

    constructor(userid: number, username: string) {
        this.userid = userid
        this.username = username
    }
}

class User {
    nickname: string;

    constructor(msg: any) {
        this.nickname = msg.nickname
    }
}

class ChatMsgContent {
    text: string;

    constructor(msgContent: any) {
        this.text = msgContent.text
    }
}

class TimeStamp {
    year: number;
    month: number;
    day: number;
    hour: number;
    minute: number;

    constructor(msgStamp: any) {
        this.year = msgStamp.year
        this.month = msgStamp.month
        this.day = msgStamp.day
        this.hour = msgStamp.hour
        this.minute = msgStamp.minute
    }
    time() {
        return this.hour + ":" + this.minute
    }
    date() {
        return this.day + "." + this.month + "." + this.year
    }
}

class ChatMessage {
    id: number;
    sender: User;
    timestamp: TimeStamp;
    content: ChatMsgContent;

    constructor(msg: any) {
        this.id = msg.id
        this.sender = new User(msg.senderinfo)
        this.timestamp = new TimeStamp(msg.timestamp)
        this.content = new ChatMsgContent(msg.content)
    }
}

class BasicChat {
    history: ChatMessage[] = [];
    chatTag: HTMLDivElement;
    ws: WebSocket;
    msgTemplate: HTMLTemplateElement;
    msgInput: HTMLInputElement;
    chatContainer: HTMLDivElement;

    constructor(chatTag: HTMLDivElement, ws: WebSocket) {
        this.chatTag = chatTag
        this.msgTemplate = <HTMLTemplateElement>this.chatTag.querySelector("#template-chatmsg")
        this.msgInput = <HTMLInputElement>this.chatTag.querySelector("#new-chat-msg-text")!!
        this.chatContainer = <HTMLDivElement>this.chatTag.querySelector("#chat-container")
        this.ws = ws
        chatTag.querySelector("#send-new-chat-msg")!!.addEventListener("click", () => { this.sendMessage() })
    }

    clear() {
        this.history = []
        this.chatContainer.innerHTML = ""
    }
    sendMessage() {
        let msgText = this.msgInput.value
        this.msgInput.value = ""
        if (msgText == null || msgText == undefined || msgText == "") {
            alert("Вы должны ввести хотя бы какой-то текст")
        } else {
            let outMsg = {
                type: MsgTypes.OutChatMsg,
                data: {
                    text: msgText
                }
            }
            this.ws.send(JSON.stringify(outMsg))
        }
    }
    loadHistory(msgHist: any) {
        this.clear()
        this.history = msgHist.history
        this.history.forEach(el => {
            this.newMessage(el)
        });
    }
    newMessage(inpMsg: any) {
        console.log(inpMsg)
        let msg = new ChatMessage(inpMsg)
        this.history.push(msg)
        let clone = document.importNode(this.msgTemplate.content, true)
        clone.querySelector(".chatmsg-text")!!.innerHTML = msg.content.text
        let timestamp = msg.timestamp.time()
        clone.querySelector(".chatmsg-info")!!.innerHTML = msg.sender.nickname + "  -  " + timestamp
        this.chatContainer.appendChild(clone)
    }
}