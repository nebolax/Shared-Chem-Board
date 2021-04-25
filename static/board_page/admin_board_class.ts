class AdminBoard extends DrawingBoard {
    curBoardID: number = 0

    constructor(msgParser: (b: AdminBoard, e: MessageEvent) => void) {
        super()
        this.msgParser = this.msgParser = function(b, e) {
            msgParser(b as AdminBoard, e)
        }
    }

    toPersonal(userID: number) {
        super.clear()
        this.ws.send(JSON.stringify({
            "type": "chview",
            "nview": userID
        }))
    }

    toGeneral() {
        super.clear()
        this.ws.send(JSON.stringify({
            "type": "chview",
            "nview": 0
        }))
    }
}