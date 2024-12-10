const {
  Worker, isMainThread, parentPort, workerData,
} = require('node:worker_threads');
const Progress = require('progress');
const fs = require('fs');

const input = fs.readFileSync('./dummy.day5', 'utf8').split('\n');
//const input = fs.readFileSync('./input.day5', 'utf8').split('\n');

let ans = 0;

for (let i = 0; i < input.length - 1; i++) {
  const line = input[i];
  console.log(line)
}

console.log(ans)
