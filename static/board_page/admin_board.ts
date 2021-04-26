module admin_board {

function initPage() {
    $("#views-nav").find("#general-page").on("click", switchView)
}

function msgParser(e: MessageEvent) {
    let msg = JSON.parse(e.data)
    switch(msg.type) {
    case MsgTypes.Points:
        board.drawPackage(msg.data.points)
        break;
    case MsgTypes.ObsStat:
        msg = msg.data
        $("#observers-nav").empty()
        msg.allObsInfo.forEach((el: any) => {
            let templ = <HTMLTemplateElement>document.getElementById("template-obsname")
            let clone = document.importNode(templ.content, true)
            let btn = clone.querySelector("#chviewBtn")!!
            btn.addEventListener("click", switchView)
            btn.innerHTML = el.username
            btn.id = "view" + el.userid
            document.getElementById("observers-nav")?.appendChild(clone)
        });
        break
    case MsgTypes.InpChatMsg:
        chat.newMessage(msg.data)
        break
}
}

function switchView(e: Event) {
    let sourceId = (<HTMLElement>e.target).id
    if (sourceId == "general-page") {
        board.toGeneral()
    } else {
        let nview: number = +sourceId.slice(4)
        board.toPersonal(nview)
    }
}

let ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname)
let board = new AdminBoard(ws)
let chat = new BasicChat(<HTMLDivElement>document.getElementById("chat"), ws)

initPage()
ws.onmessage = msgParser
}