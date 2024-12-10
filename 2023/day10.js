const fs = require('fs');

function readInput(dummy) {
  const input = fs.readFileSync(dummy ? './dummy.part10' : './input.day10', 'utf8').split('\n').filter(line => line != '');

  return input;
}

const directionsMap = {
  N: [0, -1],
  E: [1, 0],
  S: [0, 1],
  W: [-1, 0],
  X: [0, 0], // X is a special case, it means that the pipe is not connected
};

/**
  * Map of pipes and their connections
 */
const pipes = {
  '|': { 'S': 'N', 'N': 'S' },
  '-': { 'W': 'E', 'E': 'W' },
  'L': { 'E': 'N', 'N': 'E' },
  'J': { 'N': 'W', 'W': 'N' },
  '7': { 'W': 'S', 'S': 'W' },
  'F': { 'S': 'E', 'E': 'S' },
  '.': { 'X': 'X', 'X': 'X' }, // X is a special case, it means that the pipe is not connected
};

function getNewDirection({ direction, pipe }) {
  return pipes[pipe][direction];
}

function findS(input) {
  for (let i = 0; i < input.length; i++) {
    for (let j = 0; j < input[i].length; j++) {
      if (input[i][j] == 'S') {
        return [i, j];
      }
    }
  }
  return [0, 0];
}

function getSrcDirection({ prevPosition, currPosition }) {
  const [prevY, prevX] = prevPosition;
  const [currY, currX] = currPosition;

  if (prevY < currY) {
    return 'N';
  }
  if (prevY > currY) {
    return 'S';
  }
  if (prevX < currX) {
    return 'W';
  }
  if (prevX > currX) {
    return 'E';
  }
}

function calculateCicle({ inputs, initialDirection, x, y }) {
  let direction = initialDirection;

  let [dx, dy] = directionsMap[direction];
  let prevPosition = [y, x];
  let currPosition = [y + dy, x + dx];

  let count = 1;
  while (true) {
    const pipe = inputs[currPosition[0]][currPosition[1]];

    if (pipe == 'S') {
      return count;
    }

    if (pipe == '.') {
      return 0;
    }

    const srcDirection = getSrcDirection({ prevPosition, currPosition });
    direction = getNewDirection({ direction: srcDirection, pipe });

    if (direction == null) {
      return 0;
    }

    [dx, dy] = directionsMap[direction];

    prevPosition[0] = currPosition[0];
    prevPosition[1] = currPosition[1];

    currPosition[0] = currPosition[0] + dy;
    currPosition[1] = currPosition[1] + dx;

    count++;
  }
}

function part1(dummy = false) {
  const inputs = readInput(dummy);

  const [y, x] = findS(inputs);

  const north = calculateCicle({ inputs, initialDirection: 'N', x, y });
  const south = calculateCicle({ inputs, initialDirection: 'S', x, y });
  const east = calculateCicle({ inputs, initialDirection: 'E', x, y });
  const weast = calculateCicle({ inputs, initialDirection: 'W', x, y });

  const ans = Math.max(north, south, east, weast) / 2;

  return ans;
}

function part2(dummy = false) {
  const inputs = readInput(dummy);

  inputs.forEach((line) => {
    console.log(line);
  });

  const [y, x] = findS(inputs);
  console.log('(x, y): ', x, y);

  const north = calculateCicle({ inputs, initialDirection: 'N', x, y });
  const south = calculateCicle({ inputs, initialDirection: 'S', x, y });
  const east = calculateCicle({ inputs, initialDirection: 'E', x, y });
  const weast = calculateCicle({ inputs, initialDirection: 'W', x, y });

  console.log('north: ', north);
  console.log('south: ', south);
  console.log('east: ', east);
  console.log('weast: ', weast);

  const ans = Math.max(north, south, east, weast) / 2;

  return ans;
}

// console.log('part1: ', part1());
console.log('part2: ', part2());
