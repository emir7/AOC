const path = require("path");

const inputParser = require("./input-parser");

const parsedInput = inputParser.parse(path.join(__dirname, "input.txt"));

const resultMap = {};

const getOptimalFlow2 = (valves, openedValves, numberOfMinutesLeft, currentFlow, currentValveLabel, previousLabel) => {
    if(resultMap[currentValveLabel] >= currentFlow) {
        return 0;
    }

    if(numberOfMinutesLeft === 1) {
        return currentFlow;
    }

    const currentValve = valves[currentValveLabel];
    const possibleSolutions = [];

    if(currentValve.rate !== 0 && !openedValves.has(currentValveLabel)) {
        // odpres sebe
        openedValves.add(currentValveLabel);
        const option1 = getOptimalFlow2(valves, openedValves, numberOfMinutesLeft - 1, currentFlow + currentValve.rate * (numberOfMinutesLeft - 1), currentValveLabel, "");

        resultMap[currentValveLabel] = Math.max(resultMap[currentValveLabel], option1);

        possibleSolutions.push(
            option1
        );

        openedValves.delete(currentValveLabel);
    }


    for(let i = 0; i < currentValve.nextValves.length; i++) {
        const nextValveLabel = currentValve.nextValves[i];

        if(nextValveLabel === previousLabel) {
            continue;
        }

        // premik
        const option2 = getOptimalFlow2(valves, openedValves, numberOfMinutesLeft - 1, currentFlow, nextValveLabel, currentValveLabel);
        resultMap[nextValveLabel] = Math.max(resultMap[nextValveLabel], option2);

        possibleSolutions.push(
            option2
        );

    }

    return Math.max(...possibleSolutions)

};

const optimalFlow = getOptimalFlow2(parsedInput, new Set([]), 30, 0, "AA", "");
console.log(JSON.stringify(optimalFlow));