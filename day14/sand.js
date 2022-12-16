module.exports = class Sand {
    constructor(grid, startX, startY) {
        this.grid = grid.grid;
        this.startXPos = startX - grid.startX + 1;
        this.currentX = this.startXPos;
        this.currentY = startY;
    }

    fall() {
        this.currentX = this.startXPos;
        this.currentY = 0;
        let nextMove = this.getNextMove();
 
        if(nextMove == null) {
            return false;
        }


        while(nextMove) {
            if(nextMove.e) {
                return false;
            }
            
            this.currentX = nextMove.x;
            this.currentY = nextMove.y;

            nextMove = this.getNextMove();
        }


        this.grid[this.currentY][this.currentX] = "o";
        return true;
    }

    getNextMove() {
        // try down
        const nextY = this.currentY + 1;
        const canMoveDown = nextY <= this.grid.length - 1;
        
        if(canMoveDown && this.grid[nextY][this.currentX] == ".") {
            return {
                x: this.currentX,
                y: nextY,
                e: false
            };
        }

        if(canMoveDown && this.grid[nextY][this.currentX] == "~") {
            return {e: true};
        }


        // try left down
        const leftX = this.currentX - 1;
        const canMoveDownLeft = leftX >= 0 && nextY <= this.grid.length - 1;

        if(canMoveDownLeft && this.grid[nextY][leftX] == ".") {
            return {
                x: leftX,
                y: nextY,
                e: false
            };
        }

        if(canMoveDownLeft && this.grid[nextY][leftX] == "~") {
            return {e: true};
        }

        // try right down
        const rightX = this.currentX + 1;
        const canMoveDownRight = rightX <= this.grid[0].length - 1 && nextY <= this.grid.length - 1;

        if(canMoveDownRight && this.grid[nextY][rightX] == ".") {
            return {
                x: rightX,
                y: nextY,
                e: false
            };
        }

        if(canMoveDownRight && this.grid[nextY][rightX] == "~") {
            return {e: true};
        }
    
        return null;
    }
}