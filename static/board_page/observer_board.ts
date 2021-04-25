module observer_board {

function initPage() {
    $("#general-board").on("click", () => {
        switchBoard(1)
    })
    $("#personal-board").on("click", () => {
        switchBoard(2)
    })
}

function msgParser(b: ObserverBoard, e: MessageEvent) {
    let msg = JSON.parse(e.data)
    switch(msg.type) {
    case MsgTypes.Points:
        board.drawPackage(msg.data.points)
        break;
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