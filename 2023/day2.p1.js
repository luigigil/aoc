const fs = require('fs');

const input = fs.readFileSync('./input.day2', 'utf8').split('\n');

let ans = 0;

for (let i = 0; i < input.length - 1; i++) {
  let isPossible = true
  const line = input[i];
  const gameId = line.split(':')[0].split(' ')[1];
  const gameSets = line.split(':')[1].split(';')

  console.log('gameId: ', gameId);

  const minMap = new Map();
  minMap.set('red', 0);
  minMap.set('green', 0);
  minMap.set('blue', 0);

  for (let j = 0; j < gameSets.length; j++) {
    const set = gameSets[j];
    console.log('set: ', set);
    const itemSplit = set.split(',');
    console.log('itemSplit: ', itemSplit);
    for (let k = 0; k < itemSplit.length; k++) {
      const item = itemSplit[k];
      console.log('item: ', item);
      const cubeColor = item.split(' ')[2];
      const cubeCount = item.split(' ')[1];
      console.log('cubeColor: |', cubeColor, '|  ', 'cubeCount: ', cubeCount);

      if (minMap.get(cubeColor) > parseInt(cubeCount)) continue

      minMap.set(cubeColor, parseInt(cubeCount));
    }
  }

  console.log('minMap: ', minMap);

  if (isPossible) {
    ans += minMap.get('red') * minMap.get('green') * minMap.get('blue');
  }
}

console.log(ans)
