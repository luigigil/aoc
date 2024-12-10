const fs = require('fs');

function readInput(dummy) {
  const input = fs.readFileSync(dummy ? './dummy.day8' : './input.day8', 'utf8').split('\n').filter(line => line != '');
  return input;
}

function part1(dummy = false) {
  const input = readInput(dummy);
  const instructions = input[0];

  let ans = 0;

  const map = new Map();

  for (let i = 1; i < input.length; i++) {
    const line = input[i].trim();
    const location = line.split('=')[0].replace(' ', '');
    const destination = line.split('=')[1].replace('(', '').replace(')', '').replace(' ', '').split(', ');

    map.set(location, {
      'L': destination[0],
      'R': destination[1]
    });
    console.log('map: ', map);
  }

  let current = 'AAA';
  let instructionPointer = 0;
  while (current != 'ZZZ') {
    const instruction = instructions[instructionPointer++];

    console.log('current: ', current);
    console.log('instruction: ', instruction);
    console.log('destination: ', map.get(current));
    console.log('destination2: ', map.get(current)[instruction]);

    current = map.get(current)[instruction];

    ans++;
    if (current == undefined) break;
    if (instructionPointer == instructions.length) instructionPointer = 0;
  }

  return ans;
}

function part2(dummy = false) {
  const input = readInput(dummy);
  const instructions = input[0];

  const map = new Map();

  let currents = [];
  for (let i = 1; i < input.length; i++) {
    const line = input[i].trim();
    const location = line.split('=')[0].replace(' ', '');
    const destination = line.split('=')[1].replace('(', '').replace(')', '').replace(' ', '').split(', ');

    if (location.at(-1) == 'A') {
      currents.push(location);
    }

    map.set(location, {
      'L': destination[0],
      'R': destination[1]
    });
  }

  const steps = []
  for(let current of currents) {
    let index = 0;
    let location = current;
    let counter = 0;
    while(location.at(-1) != 'Z') {
      const instruction = instructions[index++];
      location = map.get(location)[instruction];
      counter++;
      if(index == instructions.length) index = 0;
    }
    steps.push(counter);
  }

  let ans = steps.pop();
  for(let step of steps) {
    ans = (ans * step) / gcd(ans, step);
  }

  return ans;
}

function gcd(a, b) {
  return !b ? a : gcd(b, a % b);
}

// console.log('part1: ', part1());
console.log('part2: ', part2());
