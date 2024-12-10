const fs = require('fs');

//const input = fs.readFileSync('./dummy.day4', 'utf8').split('\n');
const input = fs.readFileSync('./input.day4', 'utf8').split('\n');

let ans = 0;

input.forEach(line => console.log(line))

function calculatePoints(count) {
  if (count == 1) return 1
  return calculatePoints(count - 1) * 2;
}

const myMap = new Map();

for (let i = 0; i < input.length - 1; i++) {
  const card = parseInt(input[i].split(': ')[0].split(' ').filter(w => w != '')[1]);
  console.log(input[i].split(': ')[0].split(' ').filter(w => w != ''))
  console.log('card: ', card)

  myMap.set(card, (myMap.get(card) ?? 0) + 1);

  const line = input[i].split(': ')[1].split('| ');
  console.log('line: ', line)
  const winners = new Set(line[0].split(' ').filter(w => w != ''));
  console.log('winners: ', winners)
  const numbers = line[1].split(' ').filter(w => w != '');
  console.log('numbers: ', numbers)

  let count = 0;
  for (let i = 0; i < numbers.length; i++) {
    const num = numbers[i];
    if (winners.has(num)) {
      count++;
    }
  }

  if (count == 0) continue

  for (let k = 1; k <= count; k++) {
    const _card = parseInt(k) + card;
    myMap.set(_card, (myMap.get(_card) ?? 0) + (1 * myMap.get(card)));
  }

  console.log(myMap)

  // ans += calculatePoints(count);
}

for (let [key, value] of myMap) {
  ans += value
}

console.log(ans)
