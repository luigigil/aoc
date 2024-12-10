const {
  Worker, isMainThread, parentPort, workerData,
} = require('node:worker_threads');
const fs = require('fs');

class Node {
  constructor(src, dest, maps) {
    this.src = src;
    this.dest = dest;
    this.maps = maps;
  }
}

function getNodeByDest(dest, nodes) {
  return nodes.filter(node => node.dest == dest).at(0);
}

function getNodeBySrc(src, nodes) {
  return nodes.filter(node => node.src == src).at(0);
}

function makeSeedMap(src, srcValue, map, nodes) {
  map[src] = srcValue;

  const node = getNodeBySrc(src, nodes);

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

  return makeSeedMap(node.dest, parseInt(destValue), map, nodes);
}

function makeNode(line) {
  const [src, dest] = line.split(' ')[0].split('-to-');

  return new Node(src, dest, []);
}

if (isMainThread) {
  //const input = fs.readFileSync('./dummy.day5', 'utf8').split('\n');
  const input = fs.readFileSync('./input.day5', 'utf8').split('\n');

  let ans = 0;

  const seeds = input[0].split(':')[1].split(' ').filter(w => w != '')

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

  console.log('made nodes')

  let minLocation = Infinity;
  let workerCount = 0;
  let promises = [];
  for (let i = 0; i < seeds.length - 1; i += 2) {
    const initialSeed = seeds[i];
    const seedRange = seeds[i + 1];

    console.log(`starting worker ${workerCount}`)

    promises.push(new Promise((resolve, reject) => {
      const worker = new Worker(__filename, {
        workerData: {
          initialSeed,
          seedRange,
          nodes,
          id: workerCount++
        }
      });

      worker.on('message', (msg) => {
        resolve(Math.min(minLocation, msg))
      });
      worker.on('error', (err) => reject(err));
      worker.on('exit', (code) => {
        if (code !== 0)
          reject(`Worker stopped with exit code ${code}`);
      });
    }));
  }

  Promise.all(promises).then((values) => {
    console.log('all workers done')
    console.log('values: ', values);
    console.log('final: ', Math.min(...values));
  });

} else {
  const { initialSeed, seedRange, nodes, id } = workerData;

  let minLocation = Infinity;

  console.log(`worker ${id}: processing ${seedRange} seeds`)
  for (let j = 0; j < seedRange; j++) {
    const seed = parseInt(initialSeed) + j;
    const seedMap = new Map();
    makeSeedMap('seed', seed, seedMap, nodes)
    minLocation = Math.min(minLocation, seedMap['location'])
  }
  console.log(`worker ${id} finished: `, minLocation)
  parentPort.postMessage(minLocation);
}
