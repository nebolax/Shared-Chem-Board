enum MsgTypes {
    Points = 0,
    ObsStat,
    Chview,
    OutChatMsg,
    InpChatMsg
}

class Point {
    x: number = 0;
    y: number = 0;
}

class BasicBoard {
    ws: WebSocket;
    snap: Snap.Paper;
    curGroup: Snap.Paper;
    allGroups: Snap.Paper[];

    drawing: boolean;
    x: number;
    y: number;
    sendBuf: Point[];
    isDrawable: boolean;

    constructor(ws: WebSocket) {
        this.ws = ws
        this.drawing = false
        this.x = 0
        this.y = 0
        this.sendBuf = []
        this.isDrawable = true
        this.snap = Snap("#svg")
        this.snap.attr({
                 strokeWidth: 2,
                 stroke: "#000"
         })
         this.curGroup = this.snap.group()
         this.allGroups = []

        this.snap.mousedown(e => { this.mousedown(e) })
        this.snap.mousemove(e => { this.mousemove(e) })
        window.addEventListener('mouseup', e => { this.mouseup(e) })
    }
    clear() {
        this.x = 0
        this.y = 0
        this.drawing = false
        this.sendBuf = []
        this.snap.clear()
    }
    generalDraw(e: MouseEvent) {
        this.x = e.offsetX
        this.y = e.offsetY
        this.sendBuf.push({
            x: this.x,
            y: this.y
        })
    }
    mousedown(e: MouseEvent) {
       this.generalDraw(e)
       this.drawing = true
    } 
    mousemove(e: MouseEvent) {
        if (this.drawing === true) {
            this.curGroup.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY))
           this.generalDraw(e)
        }
    }
    mouseup(e: MouseEvent) {
        if (this.drawing === true) {
            this.curGroup.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY))
            this.generalDraw(e)
            this.sendPoints()
            this.drawing = false
            this.allGroups.push(this.curGroup)
            this.curGroup = this.snap.group()
            console.log(this.allGroups)
        }
    }
    canvasBack() {

    }
    drawPackage(points: Point[]) {
        for (let i = 0; i < points.length - 1; i++) {
            this.snap.line(points[i].x, points[i].y, points[i + 1].x, points[i + 1].y)
        }
    }
    sendPoints() {
        this.ws.send(JSON.stringify({
            type: MsgTypes.Points,
            data: {
                points: this.sendBuf
            }
        }))
        this.sendBuf = []
    }
}