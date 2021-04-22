module observer_board {

function initPage() {
    $("#general-board").on("click", () => {
        switchBoard(1)
    })
    $("#personal-board").on("click", () => {
        switchBoard(2)
    })
}

function msgParser(board: ObserverBoard, e: MessageEvent) {
    console.log("r")
    let msg = JSON.parse(e.data)
    if (msg.type == "points") {
        board.drawPackage(msg.points)
    }
}

function switchBoard(id: number) {
    console.log(id)
    switch (id) {
        case 2:
            board.toPersonalBoard()
            break
        default:
            board.toGeneralBoard()
            break
    }
}

initPage()
let board = new ObserverBoard(msgParser)

}