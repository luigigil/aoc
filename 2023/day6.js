// const {
//   Worker, isMainThread, parentPort, workerData,
// } = require('node:worker_threads');
// const Progress = require('progress');
const fs = require('fs');

function readInput(dummy) {
  const input = fs.readFileSync(dummy ? './dummy.day6' : './input.day6', 'utf8').split('\n');
  return input;
}

function readInputP2(dummy) {
  const input = fs.readFileSync(dummy ? './dummy.day6.part2' : './input.day6.part2', 'utf8').split('\n');
  return input;
}

function part1(dummy = false) {
  const input = readInput(dummy);

  const times = input[0].split(':')[1].split(' ').filter(x => x != '');
  const distances = input[1].split(':')[1].split(' ').filter(x => x != '');

  const ways = [];
  for (let j = 0; j < times.length; j++) {
    const time = times[j];

    let dists = []
    for (let i = 0; i < time; i++) {
      const runtime = time - i;

      let dist = runtime * i;

      dists.push(dist)
    }
    const diffWays = dists.filter(x => x > distances[j]).length

    ways.push(diffWays)
  }

  console.log(ways)

  const ans = ways.reduce((a, b) => a * b, 1)

  return ans;
}

function part2(dummy = false) {
  const input = readInputP2(dummy);

  const times = input[0].split(':')[1].split(' ').filter(x => x != '');
  const distances = input[1].split(':')[1].split(' ').filter(x => x != '');

  const ways = [];
  for (let j = 0; j < times.length; j++) {
    const time = times[j];

    let dists = []
    for (let i = 0; i < time; i++) {
      const runtime = time - i;

      let dist = runtime * i;

      dists.push(dist)
    }
    const diffWays = dists.filter(x => x > distances[j]).length

    ways.push(diffWays)
  }

  console.log(ways)

  const ans = ways.reduce((a, b) => a * b, 1)

  return ans;
}

//console.log('part1: ', part1());
console.log('part2: ', part2());
