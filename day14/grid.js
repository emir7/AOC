module.exports = class Grid {
    constructor(startX, endX, startY, endY, rocksCoordinates) {
        this.startX = startX;
        this.endX = endX;
        this.startY = startY;
        this.endY = endY;
        this.rocksCoordinates = rocksCoordinates;

        this.#initGrid();
    }

    #initGrid() {
        const wid = this.endX - this.startX + 1 + 2;
        const hei = this.endY - this.startY + 1;

        this.grid = [];

        for(let i = 0; i < hei; i++) {
            const newRow = [];
            
            for(let j = 0; j < wid; j++) {
                newRow.push(".");
            }

            this.grid.push(newRow);
        }

        this.rocksCoordinates.forEach((row) => {
            for(let i = 0; i < row.length - 1; i++) {
                const p1 = row[i];
                const p2 = row[i + 1];
                const startY = Math.min(p1.y, p2.y);
                const endY = Math.max(p1.y, p2.y);
                const startX = Math.min(p1.x, p2.x);
                const endX = Math.max(p1.x, p2.x);
                
                for(let i = startY; i <= endY; i++) {
                    for(let j = startX; j <= endX; j++) {
                        this.grid[i][j - this.startX + 1] = "#"
                    }
                }
            }
        });

        this.grid[hei - 1][0] = "~";
        this.grid[hei -1][wid - 1] = "~";
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