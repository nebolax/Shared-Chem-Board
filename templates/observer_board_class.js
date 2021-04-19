class ObserverBoard extends DrawingBoard {
    constructor(msgParser) {
        super(msgParser)
        this.msgParser = msgParser
        this.isDrawable = false
    }
    mousedown(e) {
        if (this.isDrawable) {
            super.mousedown(e)
        }
    }
    mousemove(e) {
        if (this.isDrawable) {
            super.mousemove(e)
        }
    }
    mouseup(e) {
        if (this.isDrawable) {
            super.mouseup(e)
        }
    }

    toGeneralBoard() {
        this.isDrawable = false
        this.ws.send(JSON.stringify({
            "type": "toGeneral"
        }))
    }
    toPersonalBoard() {
        this.isDrawable = true
        this.ws.send(JSON.stringify({
            "type": "toPersonal"
        }))
    }
}