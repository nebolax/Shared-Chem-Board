class AdminBoard extends DrawingBoard {
    constructor(msgParser) {
        super()
        this.msgParser = msgParser
        this.curBoardID = 0
    }

    toPersonal(userID) {
        super.clear()
        this.ws.send(JSON.stringify({
            "type": "switchBoard",
            "userID": userID
        }))
    }

    toGeneral() {
        super.clear()
        this.ws.send(JSON.stringify({
            "type": "switchBoard",
            "userID": 0
        }))
    }
}