// console.log('ws://' + window.location.host + "/ws" + window.location.pathname)
let ws = new WebSocket('ws://' + window.location.host + "/ws" + window.location.pathname)
ws.onmessage = function(e) {
    console.log("got message")
    var msg = JSON.parse(e.data).points
    console.log(msg)
    for (let i = 0; i < msg.length - 1; i++) {
        drawLine(msg[i].x, msg[i].y, msg[i + 1].x, msg[i + 1].y)
    }
}

// When true, moving the mouse draws on the canvas
let isDrawing = false
let x = 0
let y = 0

const canvas = document.getElementById('canvas')
const ctx = canvas.getContext('2d')

canvas.width = 500
canvas.height = 500

let sendBuf = []

// event.offsetX, event.offsetY gives the (x,y) offset from the edge of the canvas.

// Add the event listeners for mousedown, mousemove, and mouseup
canvas.addEventListener('mousedown', e => {
    x = e.offsetX
    y = e.offsetY
    isDrawing = true
    sendBuf.push({
        x: x,
        y: y
    })
    checkBuf()
})
canvas.addEventListener('mousemove', e => {
    if (isDrawing === true) {
        drawLine(x, y, e.offsetX, e.offsetY)
        x = e.offsetX
        y = e.offsetY
        sendBuf.push({
            x: x,
            y: y
        })
        checkBuf()
    }
})
window.addEventListener('mouseup', e => {
    if (isDrawing === true) {
        drawLine(x, y, e.offsetX, e.offsetY)
        sendBuf.push({
            x: x,
            y: y
        })
        sendBuf.push({
            x: e.offsetX,
            y: e.offsetY
        })
        x = 0
        y = 0
        sendPoints()
        sendBuf = []
        isDrawing = false
    }
})

function sendPoints() {
    ws.send(JSON.stringify({
        "points": sendBuf
    }))
    pv = sendBuf[sendBuf.length - 1]
    sendBuf = []
    sendBuf.push(pv)
    console.log("sent")
}

function checkBuf() {
    if (sendBuf.length >= 5) {
        sendPoints()
    }
}

function drawLine(x1, y1, x2, y2) {
    console.log("x")
    ctx.beginPath()
    ctx.strokeStyle = 'black'
    ctx.lineWidth = 1
    ctx.moveTo(x1, y1)
    ctx.lineTo(x2, y2)
    ctx.stroke()
    ctx.closePath()
}