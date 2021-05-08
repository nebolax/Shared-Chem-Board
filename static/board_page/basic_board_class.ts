enum MsgTypes {
    Drawing = 0,
    ObsStat,
    Chview,
    OutChatMsg,
    InpChatMsg
}

enum DrawingTypes {
    FreeMouse = 0
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
    isDrawable: boolean;

    constructor(ws: WebSocket) {
        this.ws = ws
        this.drawing = false
        this.x = 0
        this.y = 0
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
        this.snap.clear()
    }
    generalDraw(e: MouseEvent) {
        this.x = e.offsetX
        this.y = e.offsetY
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
            this.drawing = false
            this.allGroups.push(this.curGroup)
            this.sendDrawing(DrawingTypes.FreeMouse, this.curGroup)
            this.curGroup = this.snap.group()
        }
    }
    canvasBack() {

    }
    drawPackage(msg: any) {
        switch (msg.type) {
        case DrawingTypes.FreeMouse:
            for (let i = 0; i < msg.data.length - 1; i++) {
                this.snap.line(msg.data[i].x, msg.data[i].y, msg.data[i + 1].x, msg.data[i + 1].y)
            }
            break;
    }
    }
    sendDrawing(type: DrawingTypes, fig: Snap.Paper) {
        switch (type) {
            case DrawingTypes.FreeMouse:
                let cords: Point[] = []
                for(let i = 0; i < fig.children().length; i++) {
                    let cattrs = fig.children()[i].toJSON().attr
                    cords.push({
                        x: cattrs.x1,
                        y: cattrs.y1
                    }) 
                    if (i == fig.children().length - 1) {
                        cords.push({
                            x: cattrs.x2,
                            y: cattrs.y2
                        })
                    }
                }
                this.ws.send(JSON.stringify({
                    type: MsgTypes.Drawing,
                    data: {
                        type: DrawingTypes.FreeMouse,
                        data: cords
                    }
                }))
            break;
        }
    }
}