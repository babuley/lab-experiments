
class Cell{
    x:number; 
    y:number; 
    key:number;
    constructor(x:number, y:number, key:number) {
        this.x = x;
        this.y = y;
        this.key = key;
    }

    isWall():boolean {
        return this.key === 1;   
    }
    canWalk():boolean {
        return this.key === 0;
    }
    done():boolean {
        return this.key === 3;
    }  
    isStart():boolean {
        return this.key === 2;
    }  
}

function buildCells(maze:number[][]){
    let cells:Cell[] = [];
    maze.forEach( (x, idX) => {
        x.forEach( (y, idY) => {
            cells.push(new Cell(idX, idY, y))
        });
    });
    return cells;
}

function getCell(maze:Cell[], x: number, y:number):Cell {
    return maze.filter( c => { return c.x === x && c.y === y})[0];
}

function getStart(maze:Cell[]):Cell {
    return maze.filter(c => {return c.isStart()})[0];
}

let moveCell:(maze: Cell[], from:Cell)=>Cell;
function moveCellN(maze:Cell[], from:Cell):Cell {
    return getCell(maze, from.x-1,from.y);
} 
function moveCellS(maze:Cell[], from:Cell):Cell {
    return getCell(maze, from.x+1,from.y);
} 
function moveCellE(maze:Cell[], from:Cell):Cell {
    return getCell(maze, from.x,from.y+1);
} 
function moveCellW(maze:Cell[], from:Cell):Cell {
    return getCell(maze, from.x,from.y-1);
} 

const cellMoves:{ [key: string]: typeof moveCell } = {"N": moveCellN, "S": moveCellS, "E":moveCellE,"W":moveCellW };

function move(maze:Cell[], fromCell:Cell, direction:string):Cell {    
    if (typeof(fromCell) === "undefined") return undefined; 
    return cellMoves[direction](maze, fromCell);    
} 

function resolve(cell:Cell):string {
    if (typeof(cell) === "undefined")  {
        return "Dead";
    }
    if (cell.done()) {
        return "Finish";
    }    
    if (cell.isWall()){
        return "Dead"
    } 
    return "Lost";
}

export function mazeRunner(maze:number[][], directions:string[]): string{
    let mazeCells = buildCells(maze); 
    let start:Cell | undefined = getStart(mazeCells);
    for (let s = 0; s< directions.length; s++) {        
        start = move(mazeCells, start, directions[s]);       
        if (typeof(start) === "undefined") return resolve(start); 
        if (start.done() || start.isWall()){
            return resolve(start);            
        }  
    }       
    return resolve(start);    
}

//~~~~~~~~~~~Resolve maze runner~~~~~~~~~~~~~~~~~~~~~~//

let maze = [
    [1,1,1,1,1,1,1],
    [1,0,0,0,0,0,3],
    [1,0,1,0,1,0,1],
    [0,0,1,0,0,0,1],
    [1,0,1,0,1,0,1],
    [1,0,0,0,0,0,1],
    [1,2,1,0,1,0,1]];
    
    let a = mazeRunner(maze, ["N","N","N","N","N","E","E","E","E","E"]);
    console.log(a);
     a = mazeRunner(maze, ["N","N","N","N","N","E","E","S","S","E","E","N","N","E"]);
    console.log(a);
    a =mazeRunner(maze, ["N","N","N","N","N","E","E","E","E","E","W","W"]);
    console.log(a);
    a =mazeRunner(maze, ["N","N","N","W","W"]);
    console.log(a);
    a =mazeRunner(maze, ["N","N","N","N","N","E","E","S","S","S","S","S","S"]);
    console.log(a);
    a =mazeRunner(maze, ["N","E","E","E","E"]);
    console.log(a);
    
    