const fs = require('fs');

const input = fs.readFileSync('./input.day1', 'utf8').split('\n');

let ans = 0;

for (let i = 0; i < input.length - 1; i++) {
  const word = input[i];

  let left = 0;
  let right = word.length - 1;

  let leftDigit;
  let rightDigit;

  // walk on left
  while (true) {
    leftDigit = parseInt(word[left++])
    if (!isNaN(leftDigit)) {
      break;
    }
  }

  while (true) {
    rightDigit = parseInt(word[right--])
    if (!isNaN(rightDigit)) {
      break;
    }
  }

  ans += parseInt(`${leftDigit}${rightDigit}`);
}

console.log('final: ', ans);
