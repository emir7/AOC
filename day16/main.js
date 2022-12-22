const path = require("path");

const inputParser = require("./input-parser");

const parsedInput = inputParser.parse(path.join(__dirname, "input.txt"));

let resultMap = {};

const getOptimalFlow2 = (valves, openedValves, numberOfMinutesLeft, currentFlow, currentValveLabel, previousLabel, eleWait) => {
    if(resultMap[currentValveLabel] >= currentFlow) {
        return 0;
    }

    if(numberOfMinutesLeft === 1) {
        if(eleWait) {
            return getOptimalFlow2(parsedInput, openedValves, 26, currentFlow, "AA", "", false);
        }

        return currentFlow;
    }

    const currentValve = valves[currentValveLabel];
    const possibleSolutions = [];

    if(currentValve.rate !== 0 && !openedValves.has(currentValveLabel)) {
        // open
        const option1 = getOptimalFlow2(valves, new Set(openedValves).add(currentValveLabel), numberOfMinutesLeft - 1, currentFlow + currentValve.rate * (numberOfMinutesLeft - 1), currentValveLabel, currentValveLabel, eleWait);

        resultMap[currentValveLabel] = Math.max(resultMap[currentValveLabel], option1);

        possibleSolutions.push(
            option1
        );
    }


    for(let i = 0; i < currentValve.nextValves.length; i++) {
        const nextValveLabel = currentValve.nextValves[i];

        if(nextValveLabel === previousLabel) {
            continue;
        }

        // move
        const option2 = getOptimalFlow2(valves, openedValves, numberOfMinutesLeft - 1, currentFlow, nextValveLabel, currentValveLabel, eleWait);
        resultMap[nextValveLabel] = Math.max(resultMap[nextValveLabel], option2);

        possibleSolutions.push(
            option2
        );

    }

    return Math.max(...possibleSolutions)

};

const getNextPossibleStates = (parsedInput, currentState) => {
    const states = [];
    const nextNodes = parsedInput[currentState.position].nextValves;

    if(currentState.numberOfMinutes === 0) {
        return [];
    }


    // open
    if(parsedInput[currentState.position].rate && !currentState.openedValves.has(currentState.position)) {
        const calculatedFlow = currentState.flow + (parsedInput[currentState.position].rate) * (currentState.numberOfMinutes - 1);
        if(currentState.memo[(currentState.numberOfMinutes + currentState.position)] >= calculatedFlow) {

        } else {
            states.push({
                position: currentState.position,
                numberOfMinutes: currentState.numberOfMinutes - 1,
                flow: calculatedFlow,
                prevState: currentState.position,
                openedValves: new Set(currentState.openedValves).add(currentState.position),
                memo: currentState.memo
            });
            currentState.memo[(currentState.numberOfMinutes + currentState.position)] = calculatedFlow;
        }

    }


    // move
    for(const nextNode of nextNodes) {
        if(currentState.memo[(nextNode + currentState.position)] >= currentState.flow || currentState.prevState === nextNode) {
            continue;
        }

        states.push({
            position: nextNode,
            numberOfMinutes: currentState.numberOfMinutes - 1,
            flow: currentState.flow,
            prevState: currentState.position,
            openedValves: new Set(currentState.openedValves),
            memo: currentState.memo
        })
    }



    return states;
};

const part1 = (parsedInput, numberOfMinutes, start) => {
    const queue = [{
        position: start,
        numberOfMinutes,
        flow: 0,
        prevState: "",
        openedValves: new Set([]),
        memo: {}
    }];

    let max = 0;

    while(queue.length) {
        const head = queue.shift();

        if(!head) {
            return max;
        }

        const nextStates = getNextPossibleStates(parsedInput, head);

        queue.push(...nextStates)

        if(head.flow > max) {
            max = head.flow;
        }

    }

    return max;
};

const memo = {};
const part2 = (hm, t, openedValves, pos, flow) => {
    if(t === 0) {
        return 0;
    }

    if(memo[(pos)] > flow) {
        return 0;
    }

    memo[(pos)] = flow;

    let max = 0;

    for(const nextPos of hm[pos].nextValves) {
        max = Math.max(max, part2(hm, t - 1, new Set(openedValves), nextPos, flow));
    }

    if(hm[pos].rate > 0 && !openedValves.has(pos)) {
        max = Math.max(max, part2(hm, t - 1, new Set(openedValves).add(pos), pos, (t - 1) * hm[pos].rate ));
    }

    return max;
};


const optimalFlow = getOptimalFlow2(parsedInput, new Set([]), 26, 0, "AA", "", true);
console.log(optimalFlow)
