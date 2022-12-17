module.exports = class Grid {
    constructor(minX, minY, maxX, maxY, sensors) {
        this.minX = minX;
        this.minY = minY;
        this.maxX = maxX;
        this.maxY = maxY;
        this.sensors = sensors;
        this.markedSpots = new Set([]);
        this.calculatedSpots = new Set([]);

        this.#init();
    }

    #init() {
        this.wid = this.maxX - this.minX + 1;
        this.hei = this.maxY - this.minY + 1;
    
        this.sensors.forEach((sensor) => {
            this.markedSpots.add(`${sensor.x},${sensor.y}`);
            this.markedSpots.add(`${sensor.bX},${sensor.bY}`);
        });
    }

    part1(y) {

        this.sensors.forEach((sensor) => {
            const yDiff = Math.abs(sensor.y - y);
            this.#analyse(sensor, yDiff, y);
        });

        console.log(this.calculatedSpots.size)
    }

    #analyse(sensor, yDiff, y) {
        for(let i = sensor.x - sensor.visibleDistance; i <= sensor.x + sensor.visibleDistance; i++) {
            const diffX = Math.abs(i - sensor.x);
            if(diffX + yDiff <= sensor.visibleDistance) {
                if(!this.markedSpots.has(`${i},${y}`)) {
                    this.calculatedSpots.add(`${i},${y}`)
                }
            }
        }
    }


    print() {
        for(let i = 0; i < this.grid.length; i++) {
            for(let j = 0; j < this.grid[i].length; j++) {
                process.stdout.write(this.grid[i][j]);
            }
            console.log("\n");
        }
    }

    
}