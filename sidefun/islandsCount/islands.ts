function countIslands(image){
    const adjacency = toAdjacency(image);
    return traverse(adjacency);
  }

function traverse(graph) {
    const visited: string[] = [];
    let islandCount = 0;
    for (let node in graph) 
    {
        let islandSize = explore(node, graph, visited);
        if (islandSize > 0) {
            islandCount++;
        }
    }
    return islandCount;
}

function explore(node: string, graph, visited: string[]) {
    let queue: string[] = [];
    queue.push(node);
    let nodeCount = 0;
    while (queue.length > 0) {
        let current: string = queue.shift();
        if (visited.indexOf(current) < 0 ) 
        {
            visited.push(current);
            nodeCount++;
            if (graph[current]) {
                for (let n of graph[current])
                {
                    queue.push(n);
                }
            }
        }  
    }
    return nodeCount;
}

function toAdjacency(grid: number[][]) {
    
    let graph = {};

    for (let i = 0; i < grid.length; i++ ) {
        for (let j = 0; j < grid[i].length; j++) {
            
            let cell = grid[i][j];
            if (cell == 1) {
                let key = `${j},${i}`;
                let edges = graph[key];
                if (!edges) {
                    edges = [];
                }
                
                let neightbours = getNeighboursOf(j,i, grid);
                graph[key] = neightbours
            }

        }
    }
    return graph;
}

function getNeighboursOf(x:number, y: number, grid: number[][]) 
{
    let neightbours:string[] = [];
    for (let i = -1; i <= 1; i++) {
        for (let j = -1; j <= 1; j++) {
            let m = y + i;
            let n = x + j;
            if ( !(i == 0 && j == 0)){
                if (m >=0 && n >=0 && m < grid.length && n < grid[m].length) 
                {
                    let neighbour = grid[m][n];
                    if (neighbour == 1) 
                    {
                        neightbours.push(`${n},${m}`);
                    }
                }
            }
        }
    }
    return neightbours
}


  let image = [
    [0,0,0,0],
    [0,1,1,0],
    [0,1,1,0],
    [0,0,0,0],
  ];

  let r = countIslands(image);
  console.log(r)

  var image2 = [
    [0,0,0,0,0,0,0,0,0,0],
    [0,0,1,1,0,0,0,0,0,0],
    [0,0,1,1,0,0,0,0,0,0],
    [0,0,0,0,0,0,0,0,1,0],
    [0,0,0,0,0,1,1,1,0,0],
    [0,0,0,0,0,0,0,0,0,0],
  ];
  let r2 = countIslands(image2);
  console.log(r2)

  let corner = [ 
    [ 1, 1, 0, 1 ], 
  [ 0, 0, 0, 1 ], 
  [ 1, 0, 0, 0 ],
   [ 1, 0, 1, 1 ] ]
   let r3 = countIslands(corner);
   console.log(r3)