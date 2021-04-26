class AdminBoard extends BasicBoard {
    curBoardID: number = 0

    constructor(ws: WebSocket) {
        super(ws)
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