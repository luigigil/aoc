const fs = require('fs');

// const input = fs.readFileSync('./dummy.day3', 'utf8').split('\n');
 const input = fs.readFileSync('./input.day3', 'utf8').split('\n');

let ans = 0;

function checkSurroundings(line, column) {
  const aux = []

  while (column >= 0 && !isNaN(parseInt(line[--column]))) { }

  column++;

  while (column < line.length && !isNaN(parseInt(line[column]))) {
    aux.push(line[column++])
  }

  const r = parseInt(aux.join(''))
  return r
}

// input.slice(0,10).forEach(line => console.log(line))
input.forEach(line => console.log(line))

for (let i = 0; i < input.length - 1; i++) {
  for (let j = 0; j < input[i].length; j++) {
    const symbol = input[i][j];

    //if (!isNaN(parseInt(symbol)) || symbol == '.') {
    if (symbol != '*') {
      continue
    }

    const arr = []
    console.log('\nsymbol: ', symbol)
    let adjacentsCount = 0;
    let adjacents = [];

    //get left
    if (j > 0) {
      const left = parseInt(input[i][j - 1]);
      if (!isNaN(left)) {
        const surroundings = checkSurroundings(input[i], j - 1)
        console.log('left surroundings: ', surroundings)
        adjacents.push(surroundings)
      }
    }
    //get right
    if (j < input[i].length - 1) {
      const right = parseInt(input[i][j + 1]);
      if (!isNaN(right)) {
        const surroundings = checkSurroundings(input[i], j + 1)
        console.log('right surroundings: ', surroundings)
        adjacents.push(surroundings)
      }
    }

    //get up-left
    let hasUpLeft = false;
    if (i > 0 && j > 0) {
      const upLeft = parseInt(input[i - 1][j - 1]);
      if (!isNaN(upLeft)) {
        const surroundings = checkSurroundings(input[i - 1], j - 1)
        console.log('up-left surroundings: ', surroundings)
        hasUpLeft = true;
        adjacents.push(surroundings)
      }
    }

    //get up
    let hasUp = false;
    if (i > 0) {
      const up = parseInt(input[i - 1][j]);
      if (!isNaN(up)) {
        const surroundings = hasUpLeft ? 0 : checkSurroundings(input[i - 1], j)
        console.log('up surroundings: ', surroundings)
        hasUp = true;
        if (!hasUpLeft) adjacents.push(surroundings)
      }
    }

    //get up-right
    if (i > 0 && j < input[i].length - 1) {
      const upRight = parseInt(input[i - 1][j + 1]);
      if (!isNaN(upRight)) {
        const surroundings = hasUp ? 0 : checkSurroundings(input[i - 1], j + 1)
        console.log('up-right surroundings: ', surroundings)
        if (!hasUp) adjacents.push(surroundings)
      }
    }

    //get down-left
    let hasDownLeft = false;
    if (i < input.length - 1 && j > 0) {
      const downLeft = parseInt(input[i + 1][j - 1]);
      if (!isNaN(downLeft)) {
        const surroundings = checkSurroundings(input[i + 1], j - 1)
        console.log('down-left surroundings: ', surroundings)
        hasDownLeft = true;
        adjacents.push(surroundings)
      }
    }

    //get down
    let hasDown = false;
    if (i < input.length - 1) {
      const down = parseInt(input[i + 1][j]);
      if (!isNaN(down)) {
        const surroundings = hasDownLeft ? 0 : checkSurroundings(input[i + 1], j)
        console.log('down surroundings: ', surroundings)
        hasDown = true;
        if (!hasDownLeft) adjacents.push(surroundings)
      }
    }

    //get down-right
    if (i < input.length - 1 && j < input[i].length - 1) {
      const downRight = parseInt(input[i + 1][j + 1]);
      if (!isNaN(downRight)) {
        const surroundings = hasDown ? 0 : checkSurroundings(input[i + 1], j + 1)
        console.log('down-right surroundings: ', surroundings)
        if (!hasDown) adjacents.push(surroundings)
      }
    }

    if(adjacents.length == 2) {
      console.log('adjacents: ', adjacents)
      const mul = adjacents.reduce((acc, curr) => acc * curr)
      console.log(mul)
      ans += mul
    }
  }
}

console.log(ans)
