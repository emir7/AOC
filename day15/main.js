const utils = require("./utils");
const Grid = require("./grid");

const parsedData = utils.parse("input.txt");
const {maxX, maxY, minX, minY} = utils.findMinsMaxs(parsedData);
const grid = new Grid(minX, minY, maxX, maxY, parsedData);

//grid.part1(2000000);
grid.part2(4000000)