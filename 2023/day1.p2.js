const fs = require('fs');

const input = fs.readFileSync('./input.day1', 'utf8').split('\n');

let ans = 0;

const stringsAsArrays = [
  ['o', 'n', 'e'],
  ['t', 'w', 'o'],
  ['t', 'h', 'r', 'e', 'e'],
  ['f', 'o', 'u', 'r'],
  ['f', 'i', 'v', 'e'],
  ['s', 'i', 'x'],
  ['s', 'e', 'v', 'e', 'n'],
  ['e', 'i', 'g', 'h', 't'],
  ['n', 'i', 'n', 'e']
]

const digitMap = new Map()
digitMap.set('one', 1);
digitMap.set('two', 2);
digitMap.set('three', 3);
digitMap.set('four', 4);
digitMap.set('five', 5);
digitMap.set('six', 6);
digitMap.set('seven', 7);
digitMap.set('eight', 8);
digitMap.set('nine', 9);

function parser(word, pointer, direction) {
  let parsed;
  let digit;
  let temp = [];

  while (true) {
    const char = word[direction == 'left' ? pointer++ : pointer--]
    parsed = parseInt(char)

    if (!isNaN(parsed)) {
      digit = parsed;
      break;
    }

    direction == 'left' ? temp.push(char) : temp.unshift(char);

    while (temp.length > 0 &&
      !stringsAsArrays
        .map(arr => (direction == 'left' ? arr.slice(0, temp.length) : arr.slice(arr.length - temp.length, arr.length)).join(''))
        .includes(temp.join(''))
    ) {
      direction == 'left' ? temp.shift(char) : temp.pop(char);
    }

    if (temp.length == 0) continue

    const joined = temp.join('');

    if (!digitMap.has(joined)) continue

    digit = digitMap.get(joined);
    break;
  }

  return digit;
}

for (let i = 0; i < input.length - 1; i++) {
  const word = input[i];

  let left = 0;
  let right = word.length - 1;

  let leftDigit;
  let rightDigit;

  // walk on left
  leftDigit = parser(word, left, 'left');

  // walk on right
  rightDigit = parser(word, right, 'right');

  const joined = `${leftDigit}${rightDigit}`;
  ans += parseInt(joined);
}

console.log('final: ', ans);
