module admin_board {

function msgParser(board: AdminBoard, e: MessageEvent) {
    let msg = JSON.parse(e.data)
    switch(msg.type) {
    case MsgTypes.Points:
        board.drawPackage(msg.data.points)
        break;
    case MsgTypes.ObsStat:
        msg = msg.data
        let clone = $("#view0").clone()
        clone.attr("id", "view" + msg.userID)
        clone.find("#chviewBtn").html(msg.username)
        clone.find("#chviewBtn").on("click", switchView)
        $("#observers-nav").append(clone)
        break
}
}

function switchView(e: Event) {
    let nview: number = +(<HTMLElement>e.target).parentElement!!.id.slice(4)
    board.toPersonal(nview)
}

let board = new AdminBoard(msgParser)
$("#view0").find("#chviewBtn").on("click", switchView)

}