class ObserverBoard extends BasicBoard {
    constructor(ws: WebSocket) {
        super(ws)
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
    stepBack() {
        if (this.isDrawable) {
            super.stepBack()
        }
    }
}