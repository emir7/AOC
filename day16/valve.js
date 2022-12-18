module.exports = class Valve {
    constructor(label, rate, nextValves) {
        this.label = label;
        this.rate = rate;
        this.nextValves = nextValves;
    }
}