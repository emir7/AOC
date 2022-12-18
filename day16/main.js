const path = require("path");

const inputParser = require("./input-parser");

const parsedInput = inputParser.parse(path.join(__dirname, "input.txt"));
const numberOfMinutesLeft = 30;

const getOptimalFlow = (valves, visitedValves, numberOfMinutesLeft, currentValveLabel, openedValves) => {
    if(numberOfMinutesLeft < 0) {
        return {
            top: false,
            v: currentPressure
        };
    }


    let optimalFlow = 0;
    const valve = valves[currentValveLabel];

    if(visitedValves.size == Object.keys(valves).length) {
        return {
            top: true,
            v: valve.rate
        };
    }

    visitedValves.add(currentValveLabel);
    let currentMaxFlow = Number.NEGATIVE_INFINITY;
    console.log(typeof valve.rate)
    for(let i = 0; i < valve.nextValves.length; i++) {
        const nextValve = valve.nextValves[i];

        if(visitedValves.has(nextValve.label)) {
            continue;
        }

        const currentFlow = getOptimalFlow(valves, visitedValves, numberOfMinutesLeft - 2, nextValve);
        visitedValves.delete(nextValve.label);

        if(currentFlow.top) {
            return {
                v: currentFlow.v + valve.rate,
                top: true
            };
        }

        if(currentFlow.v > currentMaxFlow) {
            currentMaxFlow = currentFlow.v;
        }
    }


    return {
        top: false,
        v: currentMaxFlow + valve.rate
    };
};

const optimalFlow = getOptimalFlow(parsedInput, new Set([]), numberOfMinutesLeft, "AA", 0);
console.log(JSON.stringify(optimalFlow));