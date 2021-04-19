class canvasDataH {
    constructor() {
        this.ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname)
        this.ws.onmessage = parseMessage

        this.isDrawing = false
        this.x = 0
        this.y = 0
        this.canvas = document.getElementById('canvas')
        this.ctx = canvas.getContext('2d')
        this.canvas.width = 500
        this.canvas.height = 500
        this.sendBuf = []

        this.canvas.addEventListener('mousedown', e => {
            this.x = e.offsetX
            this.y = e.offsetY
            this.isDrawing = true
            this.sendBuf.push({
                x: this.x,
                y: this.y
            })
            this.checkBuf()
        })
        this.canvas.addEventListener('mousemove', e => {
            if (this.isDrawing === true) {
                this.drawLine(this.x, this.y, e.offsetX, e.offsetY)
                this.x = e.offsetX
                this.y = e.offsetY
                this.sendBuf.push({
                    x: this.x,
                    y: this.y
                })
                this.checkBuf()
            }
        })
        window.addEventListener('mouseup', e => {
            if (this.isDrawing === true) {
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
                this.isDrawing = false
            }
        })
    }
    drawPackage(points) {
        for (let i = 0; i < points.length - 1; i++) {
            this.drawLine(points[i].x, points[i].y, points[i + 1].x, points[i + 1].y)
        }
    }
    sendPoints() {
        this.ws.send(JSON.stringify({
            "type": "points",
            "points": this.sendBuf,
        }))
        let pv = this.sendBuf[this.sendBuf.length - 1]
        this.sendBuf = []
        this.sendBuf.push(pv)
        console.log("sent")
    }
    checkBuf() {
        if (this.sendBuf.length >= 5) {
            this.sendPoints()
        }
    }
    drawLine(x1, y1, x2, y2) {
        console.log("x")
        this.ctx.beginPath()
        this.ctx.strokeStyle = 'black'
        this.ctx.lineWidth = 1
        this.ctx.moveTo(x1, y1)
        this.ctx.lineTo(x2, y2)
        this.ctx.stroke()
        this.ctx.closePath()
    }
}

let dh = new canvasDataH()

function parseMessage(e) {
    console.log("got package")
    let msg = JSON.parse(e.data)
    if (msg.type == "points") {
        dh.drawPackage(msg.points)
    }
}