module.exports = class Sensor {
    constructor(x, y, bX, bY) {
        this.x = x;
        this.y = y;
        this.bX = bX;
        this.bY = bY;
    
        this.visibleDistance = null;
        this.#setVisibleDistance();
    }

    #setVisibleDistance() {
        this.visibleDistance = Math.abs(this.x - this.bX) + Math.abs(this.y - this.bY);
    }
}