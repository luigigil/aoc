const fs = require('fs');

//const input = fs.readFileSync('./dummy.day5', 'utf8').split('\n');
const input = fs.readFileSync('./input.day5', 'utf8').split('\n');

let ans = 0;

class Node {
  constructor(src, dest, maps) {
    this.src = src;
    this.dest = dest;
    this.maps = maps;
  }
}

const seeds = input[0].split(':')[1].split(' ').filter(w => w != '')

function makeNode(line) {
  const [src, dest] = line.split(' ')[0].split('-to-');

  return new Node(src, dest, []);
}

const nodes = []
for (let i = 1; i < input.length - 1; i++) {
  const line = input[i];

  if (line == '') {
    const node = makeNode(input[++i]);
    nodes.push(node)

    continue;
  }

  const [dest, src, range] = line.split(' ');
  nodes[nodes.length - 1].maps.push([parseInt(dest), parseInt(src), parseInt(range)]);
}

function getNodeByDest(dest) {
  return nodes.filter(node => node.dest == dest).at(0);
}

function getNodeBySrc(src) {
  return nodes.filter(node => node.src == src).at(0);
}

function makeSeedMap(src, srcValue, map) {
  map[src] = srcValue;

  const node = getNodeBySrc(src);

  if (node == undefined) {
    return
  }

  let destValue = srcValue;

  const _map = node.maps.filter(m => {
    const [dest, src, range] = m;

    const c1 = (src + range) >= srcValue
    const c2 = srcValue >= src

    const f = c1 && c2;
    return f
  });

  if (_map.length > 0) {
    const [dest, src, range] = _map[0];
    destValue = dest + (srcValue - src);
  }

  return makeSeedMap(node.dest, parseInt(destValue), map);
}

const seedMaps = []
for (let i = 0; i < seeds.length; i++) {
  const seed = seeds[i];
  const seedMap = new Map();
  makeSeedMap('seed', parseInt(seed), seedMap)
  console.log(seedMap)
  seedMaps.push(seedMap)
}

console.log(seedMaps.map(map => map['location']))

