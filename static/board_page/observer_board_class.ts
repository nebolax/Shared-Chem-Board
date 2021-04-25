class ObserverBoard extends DrawingBoard {
    constructor(msgParser: (b: ObserverBoard, e: MessageEvent) => void) {
        super()
        this.msgParser = function(b, e) {
            msgParser(b as ObserverBoard, e)
        }
        this.isDrawable = false
    }
    mousedown(e: MouseEvent) {
        if (this.isDrawable) {
            super.mousedown(e)
        }
    }
    mousemove(e: MouseEvent) {
        if (this.isDrawable) {
            super.mousemove(e)
        }
    }
    mouseup(e: MouseEvent) {
        if (this.isDrawable) {
            super.mouseup(e)
        }
    }

    toGeneralBoard() {
        super.clear()
        this.isDrawable = false
        this.ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 0
            }
        }))
    }
    toPersonalBoard() {
        super.clear()
        this.isDrawable = true
        this.ws.send(JSON.stringify({
            type: MsgTypes.Chview,
            data: {
                nview: 1
            }
        }))
    }
}