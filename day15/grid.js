module.exports = class Grid {
    constructor(minX, minY, maxX, maxY, sensors) {
        this.minX = minX;
        this.minY = minY;
        this.maxX = maxX;
        this.maxY = maxY;
        this.sensors = sensors;
        this.markedSpots = new Set([]);
        this.calculatedSpots = new Set([]);
        this.calculatedPoints = new Set([]);

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

    part2(searchSpace) {
        for(let i = 0; i < this.sensors.length; i++) {
            const sensor = this.sensors[i];
            const invisiblePoints = this.#getInvisiblePoints(sensor, searchSpace);

            for(const point of invisiblePoints) {
                let isPointInvisible = true;

                for(let j = 0; j < this.sensors.length; j++) {
                    const otherSensor = this.sensors[j];
    
                    if(i == j) {
                        continue;
                    }
    
                    const distanceFromOtherSensor = Math.abs(otherSensor.x - point.x) + Math.abs(otherSensor.y - point.y);

                    if(distanceFromOtherSensor <= otherSensor.visibleDistance) {
                        isPointInvisible = false;
                        break;
                    }
    
                }

                if(isPointInvisible) {
                    console.log(point);
                    console.log(point.x * 4000000 + point.y);
                    return;
                }
            }
   
        }
    }

    #getInvisiblePoints(sensor, searchSpace) {
        const points = [];
        const farDistance = sensor.visibleDistance + 1;

        for(let x = sensor.x - farDistance, d = 0; x <= sensor.x + farDistance; x++, d++) {
            if(x < 0 || x > searchSpace) {
                continue;
            }



            const y1 = sensor.y + d;
            const y2 = sensor.y - d;

            
            if(y1 >= 0 && y1 < searchSpace && !this.calculatedPoints.has(`${x},${y1}`)) {
                //this.calculatedSpots.add(`${x},${y1}`);
                points.push({x, y: y1})
            }


            if(y2 >= 0 && y2 < searchSpace && !this.calculatedPoints.has(`${x},${y2}`)) {
                //this.calculatedSpots.add(`${x},${y2}`);
                points.push({x, y: y2})
            }

        }

        return points;
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