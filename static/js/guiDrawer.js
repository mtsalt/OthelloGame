class GuiDrawer {

    constructor(canvas_id, windowSizeX, windowSizeY) {
        this.canvas = document.getElementById(canvas_id);
        this.setWindowSize(windowSizeX, windowSizeY);
        this.context = this.canvas.getContext("2d");
    }
    setWindowSize(x, y) {
        this.canvas.setAttribute("width" , x.toString());
        this.canvas.setAttribute("height", y.toString());
    }
    /**
     * Draw square 
     * @param {number} x x axis value
     * @param {number} y y axis value
     * @param {number} size length of a side
     * @param {string} fillColor filled color info [e.g. "rgba(255,0,0,0.8)"]
     * @param {string} outlineColor outline color info [e.g. "rgba(255,0,0,0.8)"]
     * @param {number} lineWidth line thickness
     */
    square(x, y, size, fillColor, outlineColor, lineWidth) {
        this.context.beginPath();
        this.context.rect(x, y, size, size);
        this.context.fillStyle = fillColor;
        this.context.strokeStyle = outlineColor;
        this.context.lineWidth = lineWidth;
        this.context.fill();
        this.context.stroke();
    }

    squareFrame(x, y, size, outlineColor, lineWidth) {
        this.context.rect(x, y, size, size);
        this.context.fillStyle = fillColor;
        this.context.strokeStyle = outlineColor;
        this.context.lineWidth = lineWidth;
        this.context.stroke();
    }

    rectangle(x, y, sizeX, sizeY, fillColor, outlineColor, lineWidth) {
        this.context.beginPath();
        this.context.rect(x, y, sizeX, sizeY);
        this.context.fillStyle = fillColor;
        this.context.strokeStyle = outlineColor;
        this.context.lineWidth = lineWidth;
        this.context.fill();
        this.context.stroke();
    }

    circle(x, y, radius, fillColor) {
        this.context.beginPath();
        this.context.arc(x, y, radius, 0 * Math.PI / 180, 360 * Math.PI / 180);
        this.context.fillStyle = fillColor;
        this.context.fill();
    }

    resize(scale) {
        this.context.setTransform(scale, 0, 0, scale, 0, 0);
    }
}