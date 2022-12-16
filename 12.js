const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

class Queue {
  constructor() {
    this.elements = {};
    this.head = 0;
    this.tail = 0;
  }
  enqueue(element) {
    this.elements[this.tail] = element;
    this.tail++;
  }
  dequeue() {
    const item = this.elements[this.head];
    delete this.elements[this.head];
    this.head++;
    return item;
  }
  peek() {
    return this.elements[this.head];
  }
  get length() {
    return this.tail - this.head;
  }
  get isEmpty() {
    return this.length === 0;
  }
}

async function readInput() {
  const input = await fs.promises.readFile('12.txt', 'utf-8');
  return input.trim().split('\n');
}

const getNeighbours = (x, y) => {
  const neighbours = [];
  neighbours.push({ x: x - 1, y: y });
  neighbours.push({ x: x, y: y - 1 });
  neighbours.push({ x: x + 1, y: y });
  neighbours.push({ x: x, y: y + 1 });
  return neighbours;
};

const bfs = (coords, start, end, part) => {
  const visited = new Set();
  visited.add(JSON.stringify(end));
  const queue = new Queue();
  queue.enqueue({ dist: 0, coord: end });

  while (!queue.isEmpty) {
    const item = queue.dequeue();
    const dist = item.dist;
    const x = item.coord.x;
    const y = item.coord.y;

    const val = coords.get(JSON.stringify({ x: x, y: y }));

    const neighbours = getNeighbours(x, y);
    for (let i = 0; i < neighbours.length; i++) {
      const coord = neighbours[i];
      if (coords.get(JSON.stringify(coord)) === undefined) {
        continue;
      }
      if (visited.has(JSON.stringify(coord))) {
        continue;
      }
      if (
        coords.get(JSON.stringify(coord)).charCodeAt() - val.charCodeAt() <
        -1
      ) {
        continue;
      }
      if (part === 1) {
        if (coord.x === start.x && coord.y === start.y) {
          return dist + 1;
        }
      } else {
        if (coords.get(JSON.stringify(coord)) == 'a') {
          return dist + 1;
        }
      }
      visited.add(JSON.stringify(coord));
      queue.enqueue({ dist: dist + 1, coord: coord });
    }
  }
  return 0;
};

async function puzzle() {
  console.log('--- Running Day 12 ---');

  const start = performance.now();

  const input = await readInput();

  const coords = new Map();
  let startPoint = { x: 0, y: 0 };
  let endPoint = { x: -1, y: -1 };

  for (let row = 0; row < input.length; row++) {
    const line = input[row];
    for (let col = 0; col < line.length; col++) {
      const chr = line[col];
      if (chr === 'S') {
        startPoint = { x: row, y: col };
        coords.set(JSON.stringify(startPoint), 'a');
      } else if (chr === 'E') {
        endPoint = { x: row, y: col };
        coords.set(JSON.stringify(endPoint), 'z');
      } else {
        coords.set(JSON.stringify({ x: row, y: col }), chr);
      }
    }
  }

  console.log('Part 1: ' + bfs(coords, startPoint, endPoint, 1));
  console.log('Part 2: ' + bfs(coords, startPoint, endPoint, 2));

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
