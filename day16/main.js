const path = require("path");

const inputParser = require("./input-parser");

const parsedInput = inputParser.parse(path.join(__dirname, "input.txt"));

let resultMap = {};

const getOptimalFlow2 = (valves, openedValves, numberOfMinutesLeft, currentFlow, currentValveLabel, previousLabel) => {
    if(resultMap[currentValveLabel] >= currentFlow) {
        return 0;
    }

    if(numberOfMinutesLeft === 1) {
        return currentFlow;
    }

    const currentValve = valves[currentValveLabel];
    const possibleSolutions = [];

    // generate possible moves
    // set max to 0
    // iterate over possible moves and check if each recursive result is greater than max
    // max = result
    // return max

    if(currentValve.rate !== 0 && !openedValves.has(currentValveLabel)) {
        // odpres sebe
        const option1 = getOptimalFlow2(valves, new Set(openedValves).add(currentValveLabel), numberOfMinutesLeft - 1, currentFlow + currentValve.rate * (numberOfMinutesLeft - 1), currentValveLabel, currentValveLabel);

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

        // premik
        const option2 = getOptimalFlow2(valves, openedValves, numberOfMinutesLeft - 1, currentFlow, nextValveLabel, currentValveLabel);
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

    if(resultMap[currentState.position] >= currentState.flow) {
        return [];
    }

    if(currentState.numberOfMinutes === 0) {
        return [];
    }

    // open
    if(parsedInput[currentState.position].rate && !currentState.openedValves.has(currentState.position)) {
        const calculatedFlow = currentState.flow + (parsedInput[currentState.position].rate) * (currentState.numberOfMinutes - 1);

        //resultMap[currentState.position] = Math.max(resultMap[currentState.position] ?? 0, calculatedFlow);

        states.push({
            position: currentState.position,
            numberOfMinutes: currentState.numberOfMinutes - 1,
            flow: calculatedFlow,
            prevState: currentState.position,
            openedValves: new Set(currentState.openedValves).add(currentState.position)
        });
    }


    // move
    for(const nextNode of nextNodes) {
        if(currentState.prevState === nextNode) {
            continue
        }

        states.push({
            position: nextNode,
            numberOfMinutes: currentState.numberOfMinutes - 1,
            flow: currentState.flow,
            prevState: currentState.position,
            openedValves: new Set(currentState.openedValves)
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
        openedValves: new Set([])
    }];

    let max = 0;

    while(queue.length) {
        const head = queue.pop();

        if(!head) {
            return max;
        }

        const nextStates = getNextPossibleStates(parsedInput, head);
        resultMap[head.position] = Math.max(resultMap[head.position], head.flow)
        queue.push(...nextStates)

        if(head.flow > max) {
            max = head.flow;
        }

    }

    return max;
};


const part2 = (hm, openedValves, t, turn, flow, e, h, ep, hp) => {
    if(t === 0) {
        return flow;
    }

    const solutions = [];

    let nextTurn = "";

    if(turn === "e") {
        nextTurn = "h";
    } else if(turn === "h") {
        nextTurn = "e"
    }

    if(e === h) {
        // same position
        if(hm[e].rate > 0 && !openedValves.has(e)) {
            // not opened valve, and has rate greater than 0
            if(turn === "e") {
                const o1 = part2(hm, new Set(openedValves).add(e), t - 1 ,"h", flow + (t - 1) * hm[e].rate, e, h, ep, hp);
                solutions.push(o1);
            } else {
                const o2 = part2(hm, new Set(openedValves).add(e), t,"e", flow + (t - 1) * hm[e].rate, e, h, ep, hp);
                solutions.push(o2);
            }
        }
    } else {
        if(hm[e].rate > 0 && !openedValves.has(e) && turn === "e") {
            const o = part2(hm, new Set(openedValves).add(e), t - 1, nextTurn, flow + (t - 1) * hm[e].rate, e, h, ep, hp);
            solutions.push(o);
        }

        if(hm[h].rate > 0 && !openedValves.has(h) && turn === "h") {
            const o = part2(hm, new Set(openedValves).add(h), t, nextTurn, flow + (t - 1) * hm[h].rate, e, h, ep, hp);
            solutions.push(o);
        }
    }

    let comparisonKey = (turn === "e") ? ep : hp;
    let currentValveKey = (turn === "e") ? e : h;

    // move
    for(const valveLabel of hm[currentValveKey].nextValves) {
        if(valveLabel === comparisonKey) {
            continue;
        }

        // (hm, openedValves, t, turn, flow, e, h, ep, hp)
        if(turn === "e") {
            const o1 = part2(hm, openedValves, t - 1, nextTurn, flow, valveLabel, h, e, hp);
            solutions.push(o1);
        } else {
            const o2 = part2(hm, openedValves, t, nextTurn, flow, e, valveLabel, ep, h);
            solutions.push(o2);
        }

    }

    return Math.max(...solutions);

}


//console.time("getOptimalFlow2");
//const optimalFlow = getOptimalFlow2(parsedInput, new Set([]), 30, 0, "AA", "");
//console.timeEnd("getOptimalFlow2");
//resultMap = {};
/*console.time("part1")
const p1 = part1(parsedInput, 30, "AA");
console.log(p1)
console.timeEnd("part1");*/

const p2 = part2(
    parsedInput,
    new Set([]),
    26,
    "h",
    0,
    "AA",
    "AA",
    "",
    ""
);

console.log(p2)