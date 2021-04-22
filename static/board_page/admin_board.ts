module admin_board {

function msgParser(board: AdminBoard, e: MessageEvent) {
    console.log("r")
    let msg = JSON.parse(e.data)
    if (msg.type == "points") {
        board.drawPackage(msg.points)
    } else if (msg.type == "newObserver") {
        let el = document.createElement("div")
        el.innerHTML = "<button></button>"
    }
}

let board = new AdminBoard(msgParser)

let e = document.createElement("div")
e.innerHTML = "<p>Hi there</p>"
}