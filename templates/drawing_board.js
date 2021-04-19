class DrawingBoard {
    constructor() {
        this.ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname)

        this.ws.onmessage = (e) => { this.msgParser(this, e) }

        // this.msgParser = msgParser
        this.drawing = false
        this.x = 0
        this.y = 0
        this.canvas = document.getElementById('canvas')
        this.ctx = canvas.getContext('2d')
        this.canvas.width = 500
        this.canvas.height = 500
        this.sendBuf = []
        this.isDrawable = true

        this.canvas.addEventListener('mousedown', e => { this.mousedown(e) })
        this.canvas.addEventListener('mousemove', e => { this.mousemove(e) })
        window.addEventListener('mouseup', e => { this.mouseup(e) })
    }
    clear() {
        this.drawing = false
        this.sendBuf = []
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
    }
    msgParser() {
        console.log("from default parser")
    }
    mousedown(e) {
        this.x = e.offsetX
        this.y = e.offsetY
        this.drawing = true
        this.sendBuf.push({
            x: this.x,
            y: this.y
        })
        this.checkBuf()
    }
    mousemove(e) {
        if (this.drawing === true) {
            this.drawLine(this.x, this.y, e.offsetX, e.offsetY)
            this.x = e.offsetX
            this.y = e.offsetY
            this.sendBuf.push({
                x: this.x,
                y: this.y
            })
            this.checkBuf()
        }
    }
    mouseup(e) {
        if (this.drawing === true) {
            this.drawLine(this.x, this.y, e.offsetX, e.offsetY)
            this.sendBuf.push({
                x: this.x,
                y: this.y
            })
            this.sendBuf.push({
                x: e.offsetX,
                y: e.offsetY
            })
            this.x = 0
            this.y = 0
            this.sendPoints()
            this.sendBuf = []
            this.drawing = false
        }
    }

    drawPackage(points) {
        for (let i = 0; i < points.length - 1; i++) {
            this.drawLine(points[i].x, points[i].y, points[i + 1].x, points[i + 1].y)
        }
    }
    sendPoints() {
        console.log("s")
        this.ws.send(JSON.stringify({
            "type": "points",
            "points": this.sendBuf,
        }))
        let pv = this.sendBuf[this.sendBuf.length - 1]
        this.sendBuf = []
        this.sendBuf.push(pv)
    }
    checkBuf() {
        if (this.sendBuf.length >= 5) {
            this.sendPoints()
        }
    }
    drawLine(x1, y1, x2, y2) {
        this.ctx.beginPath()
        this.ctx.strokeStyle = 'black'
        this.ctx.lineWidth = 1
        this.ctx.moveTo(x1, y1)
        this.ctx.lineTo(x2, y2)
        this.ctx.stroke()
        this.ctx.closePath()
    }

}