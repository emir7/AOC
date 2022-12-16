const fs = require("fs");

const Grid = require("./grid");
const Sand = require("./sand");

const parseInput = (input) => {
    return input.split("\n").map(line => {
        return line.trim().split(" -> ").map(el => {
            const [x, y] = el.trim().split(",");

            return {
                x: Number(x.trim()),
                y: Number(y.trim())
            }
        });
    });
};

const initGridData = (parsedInput) => {
    let minX = Number.MAX_VALUE;
    let maxX = Number.MIN_VALUE;
    let maxY = Number.MIN_VALUE;

    parsedInput.forEach(row => {
        row.forEach((element) => {
            if (element.x < minX) {
                minX = element.x;
            }
    
            if (element.x > maxX) {
                maxX = element.x;
            }

            if (element.y > maxY) {
                maxY = element.y;

            }
        });
    });

    return new Grid(minX, maxX, 0, maxY, parsedInput);
};

const data = fs.readFileSync("input.txt", {encoding: "utf-8"});
const parsedInput = parseInput(data);
const grid = initGridData(parsedInput);
const sand = new Sand(grid, 500, 0);

let c = 0;

while(sand.fall()) {
    c++;
}

grid.print();
console.log(c);
