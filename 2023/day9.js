const fs = require('fs');

function readInput(dummy) {
  const input = fs.readFileSync(dummy ? './dummy.day9' : './input.day9', 'utf8').split('\n').filter(line => line != '');

  return input;
}

function makeDifferenceArray(input) {
  let differenceArray = [];

  let allZeros = true;
  for (let i = 1; i < input.length; i++) {
    const diff = input[i] - input[i - 1];
    differenceArray.push(diff);
    if (diff != 0) {
      allZeros = false;
    }
  }

  if (allZeros) {
    differenceArray.push(0);
  }

  return { differenceArray, allZeros };
}

function solvePart1(input) {
  const copy = [...input];

  const { differenceArray, allZeros } = makeDifferenceArray(input);

  const newInput = allZeros ? differenceArray : solvePart1(differenceArray);

  copy.push(newInput.at(-1) + input.at(-1));

  return copy;
}

function solvePart2(input) {
  const copy = [...input];

  const { differenceArray, allZeros } = makeDifferenceArray(input);

  const newInput = allZeros ? differenceArray : solvePart2(differenceArray);

  copy.unshift(input.at(0) - newInput.at(0));

  return copy;
}

function part1(dummy = false) {
  const inputs = readInput(dummy).map(line => line.split(' ').map(word => parseInt(word)));

  let ans = 0;

  for (let input of inputs) {
    const output = solvePart1(input);
    ans += output.at(-1);
  }

  return ans;
}

function part2(dummy = false) {
  const inputs = readInput(dummy).map(line => line.split(' ').map(word => parseInt(word)));

  let ans = 0;

  for (let input of inputs) {
    const output = solvePart2(input);
    ans += output.at(0);
  }

  return ans;
}

// console.log('part1: ', part1());
console.log('part2: ', part2());
