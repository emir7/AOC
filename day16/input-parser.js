const fs = require("fs");

const Valve = require("./valve");

module.exports.parse = (path) => {
    const data = fs.readFileSync(path, {encoding: "utf-8"});
    return data.split("\n").map((line) => {
        const [part1, part2] = line.split(";");
        const splittedPart1 = part1.split(" ");
        const label = splittedPart1[1];
        const ratePart = splittedPart1[splittedPart1.length - 1];
        const [,rate] = ratePart.split("=");
        
        let relevantPart2 = part2.replace("tunnels lead to valves", "");
        if(relevantPart2.indexOf("tunnel") >= 0) {
            relevantPart2 = part2.replace("tunnel leads to valve", "")
        }

        const splittedRelevantPart2 = relevantPart2.split(", ");
        const nextValves = splittedRelevantPart2.flatMap(v => {
            if(v == "s") {
                return [];
            }

            return v.trim();
        });

        return new Valve(label, Number(rate.trim()), nextValves);
    }).reduce((acc, currentValue) => {
        acc[currentValue.label] = currentValue;

        return acc;
    }, {})


};