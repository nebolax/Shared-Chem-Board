class AdminBoard extends DrawingBoard {
    curBoardID: number = 0

    constructor(msgParser: (b: AdminBoard, e: MessageEvent) => void) {
        super()
        this.msgParser = this.msgParser = function(b, e) {
            msgParser(b as AdminBoard, e)
        }
    }

    toPersonal(viewID: number) {
        super.clear()
        console.log(viewID)
        this.ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: viewID
            }
        }))
    }

    toGeneral() {
        super.clear()
        this.ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 0
            }
        }))
    }
}