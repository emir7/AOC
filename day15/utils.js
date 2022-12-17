const fs = require("fs");

const Sensor = require("./sensor");

module.exports.parse = (path) => {
    const input = fs.readFileSync(path, {encoding: "utf8"});
    
    return input.split("\n").map((line) => {
        line = line.replace("Sensor at", "").trim();
        line = line.replace("closest beacon is at").trim();
    
        const [sensorPart, beaconPart] = line.split(":");
        const [sensorXPart, sensorYPart] = sensorPart.split(",");
        const [beaconXPart, beaconYPart] = beaconPart.split(",");
        const [, sensorXstr] = sensorXPart.split("=");
        const [, sensorYStr] = sensorYPart.split("=");
        const [, beaconXstr] = beaconXPart.split("=");
        const [, beaconYStr] = beaconYPart.split("=");
        const sX = Number(sensorXstr);
        const sY = Number(sensorYStr);
        const bX = Number(beaconXstr);
        const bY = Number(beaconYStr);

        return new Sensor(sX, sY, bX, bY); 
    });
};

module.exports.findMinsMaxs = (sensors) => {
    let maxX = Number.MIN_VALUE;
    let maxY = Number.MIN_VALUE;
    let minX = Number.MAX_VALUE;
    let minY = Number.MAX_VALUE;

    sensors.forEach(sensor => {
        if(sensor.x > maxX) maxX = sensor.x;
        if(sensor.bX > maxX) maxX = sensor.bX;
        if(sensor.y > maxY) maxY = sensor.y;
        if(sensor.bY > maxY) maxY = sensor.bY;

        if(sensor.x < minX) minX = sensor.x;
        if(sensor.bX < minX) minX = sensor.bX;
        if(sensor.y < minY) minY = sensor.y;
        if(sensor.bY < minY) minY = sensor.bY;
    });

    return {maxX, maxY, minX, minY};
};