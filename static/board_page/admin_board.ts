module admin_board {

function msgParser(board: AdminBoard, e: MessageEvent) {
    let msg = JSON.parse(e.data)
    switch(msg.type) {
    case MsgTypes.Points:
        board.drawPackage(msg.data.points)
        break;
    case MsgTypes.ObsStat:
        let clone = $("#observer-bar").clone()
        clone.attr("id", "user" + msg.userID)
        clone.html(msg.username)
        $("observers-nav").append(clone)
        break;
}
}

let board = new AdminBoard(msgParser)

}