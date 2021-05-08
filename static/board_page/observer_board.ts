module observer_board {

function initPage() {
    $("#general-board").on("click", () => {
        switchBoard(1)
    })
    $("#personal-board").on("click", () => {
        switchBoard(2)
    })
}

function msgParser(e: MessageEvent) {
    let msg = JSON.parse(e.data)
    switch(msg.type) {
    case MsgTypes.Action:
        board.drawPackage(msg.data.points)
        break;
    case MsgTypes.InpChatMsg:
        chat.newMessage(msg.data)
        break
    }
}

function switchBoard(id: number) {
    console.log(id)
    switch (id) {
        case 2:
            toPersonalBoard()
            break
        default:
            toGeneralBoard()
            break
    }
}

function toGeneralBoard() {
    board.clear()
    chat.clear()
    board.isDrawable = false
    ws.send(JSON.stringify({
        type: MsgTypes.Chview,
        data: {
            nview: 0
        }
    }))
}
function toPersonalBoard() {
    board.clear()
    chat.clear()
    board.isDrawable = true
    ws.send(JSON.stringify({
        type: MsgTypes.Chview,
        data: {
            nview: 1
        }
    }))
}

let ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname)
let board = new ObserverBoard(ws)
let chat = new BasicChat(<HTMLDivElement>document.getElementById("chat"), ws)

initPage()
ws.onmessage = msgParser
}