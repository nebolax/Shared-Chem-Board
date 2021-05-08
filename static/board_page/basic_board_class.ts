enum MsgTypes {
    Action = 0,
    SetId,
    ObsStat,
    Chview,
    OutChatMsg,
    InpChatMsg
}

enum DrawingTypes {
    FreeMouse = 0
}

enum ActionTypes {
    NewDrawing = 0,
    DrawingMoved,
    DrawingDeleted
}

enum Modes {
    Drawing = 0,
    Dragging
}

class Point {
    x: number = 0;
    y: number = 0;
}

class Drawing {
    id: number;
    type: DrawingTypes;
    data: any[];
    fig: Snap.Paper;

    constructor(fig: Snap.Paper) {
        this.id = 0
        this.type = DrawingTypes.FreeMouse
        this.fig = fig
        this.data = []
    }

    computePlainData() {
        switch (this.type) {
        case DrawingTypes.FreeMouse:
            let cords: Point[] = []
            for(let i = 0; i < this.fig.children().length; i++) {
                let cattrs = this.fig.children()[i].toJSON().attr
                cords.push({
                    x: cattrs.x1,
                    y: cattrs.y1
                 }) 
                if (i == this.fig.children().length - 1) {
                    cords.push({
                        x: cattrs.x2,
                     y: cattrs.y2
                    })
                }
            }
        this.data = cords
        break;
        }
    }
}

class Action {
    id: number;
    type: ActionTypes;
    drawing: Drawing;

    constructor(drawing: Drawing) {
        this.id = 0
        this.type =  ActionTypes.NewDrawing
        this.drawing = drawing
    }
}

class BasicBoard {
    ws: WebSocket;
    snap: Snap.Paper;
    curDrawing: Drawing;
    allDrawings: Drawing[];
    actions: Action[];

    drawing: boolean;
    x: number;
    y: number;
    isDrawable: boolean;
    isDragMode: boolean;

    constructor(ws: WebSocket) {
        this.ws = ws
        this.drawing = false
        this.x = 0
        this.y = 0
        this.isDrawable = true
        this.isDragMode = false
        this.snap = Snap("#svg")
        this.snap.attr({
                 strokeWidth: 2,
                 stroke: "#000"
         })
         this.curDrawing = new Drawing(this.snap.group())

         this.allDrawings = []
         this.actions = []

        this.snap.mousedown(e => { this.mousedown(e) })
        this.snap.mousemove(e => { this.mousemove(e) })
        window.addEventListener('mouseup', e => { this.mouseup(e) })
    }
    newDrawingID(id: number) {
        for (let i = 0; i < this.allDrawings.length; i++) {
            let el = this.allDrawings[i]
            if (el.id == 0) {
                el.id = id
                break
            }
        }
        console.log("drawings: " + this.allDrawings.map(val => { return val.id }))
    }
    newActionID(id: number) {
        for (let i = 0; i < this.actions.length; i++) {
            let el = this.actions[i]
            if (el.id == 0) {
                el.id = id
                break
            }
        }
        console.log("actions: " + this.actions.map(val => { return val.id }))
    }
    switchDragMode() {
        this.isDragMode = !this.isDragMode
        if (this.isDragMode) {
            this.snap.drag()
        } else {
            this.snap.undrag()
        }
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
        if (!this.isDragMode) {
            this.generalDraw(e)
            this.drawing = true
        }
    } 
    mousemove(e: MouseEvent) {
        if (this.drawing && !this.isDragMode) {
            this.curDrawing.fig.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY))
            this.generalDraw(e)
        }
    }
    mouseup(e: MouseEvent) {
        if (this.drawing && !this.isDragMode) {
            this.drawing = false
            this.curDrawing.fig.append(this.snap.line(this.x, this.y, e.offsetX, e.offsetY))
            this.generalDraw(e)

            let action: Action = {
                id: 0,
                type: ActionTypes.NewDrawing,
                drawing: this.curDrawing
            }

            this.allDrawings.push(this.curDrawing)
            this.actions.push(action)
            
            this.sendAction(action)
            this.curDrawing = new Drawing(this.snap.group())
        }
    } 
    stepBack() {
        if (this.allDrawings.length > 0) {
            let last = this.allDrawings.pop()
            last?.fig.remove()
        }
    }
    drawPackage(msg: Action) {
        switch (msg.type) {
        case ActionTypes.NewDrawing:
            switch(msg.drawing.type) {
            case DrawingTypes.FreeMouse:
                let drawing = new Drawing(this.snap.group())
                for (let i = 0; i < msg.drawing.data.length - 1; i++) {
                    drawing.fig.append(this.snap.line(msg.drawing.data[i].x, msg.drawing.data[i].y, msg.drawing.data[i + 1].x, msg.drawing.data[i + 1].y))
                }
                this.allDrawings.push(drawing)
                msg.drawing = drawing
                this.actions.push(msg)
                console.log(this.actions)
            break;
        }
        break;
    }
    }
    sendAction(action: Action) {
        action.drawing.computePlainData()
        let ts = JSON.stringify({
            type: MsgTypes.Action,
            data: action
        }, (key, val) => { return key == "fig" ? undefined : val })
        this.ws.send(ts)
    }
}