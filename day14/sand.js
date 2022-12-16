module.exports = class Sand {
    constructor(gridWrapper, startX, startY) {
        this.gridWrapper = gridWrapper;
        this.grid = this.gridWrapper.grid;
        this.startX = startX;
        
        this.setPos();
        this.currentY = startY;
    }

    setPos() {
        this.startXPos = this.startX - this.gridWrapper.startX;
        this.currentX = this.startXPos;
        this.currentY = 0;
    }

    fall() {
        let nextMove = this.getNextMove();
 
        if(nextMove == null) {
            return false;
        }

        while(nextMove) {
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
            };
        }


        // try left down
        const leftX = this.currentX - 1;
        const canMoveDownLeft = leftX >= 0 && nextY <= this.grid.length - 1;

        if(canMoveDownLeft && this.grid[nextY][leftX] == ".") {
            return {
                x: leftX,
                y: nextY,
            };
        }

        if(leftX < 0) {
            this.gridWrapper.extendGrid();
            this.setPos();
            return this.getNextMove();
        }


        // try right down
        const rightX = this.currentX + 1;
        const canMoveDownRight = rightX <= this.grid[0].length - 1 && nextY <= this.grid.length - 1;

        if(canMoveDownRight && this.grid[nextY][rightX] == ".") {
            return {
                x: rightX,
                y: nextY,
            };
        }

        if(rightX == this.grid[0].length) {
            this.gridWrapper.extendGrid();
            this.setPos();

            return this.getNextMove();
        }

    
        return null;
    }
}