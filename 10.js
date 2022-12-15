const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('10.txt', 'utf-8');
  return input.trim().split('\n');
}

async function puzzle() {
  console.log('--- Running Day 10 ---');

  const start = performance.now();

  const input = await readInput();

  let x = 1;
  const stack = [];

  input.forEach((line) => {
    if (line.startsWith('noop')) {
      stack.push(x);
    } else {
      const val = line.split(' ')[1];
      stack.push(x);
      stack.push(x);
      x += +val;
    }
  });

  const sum = [20, 60, 100, 140, 180, 220]
    .map((i) => i * stack[i - 1])
    .reduce((a, b) => a + b);

  console.log('Part 1: ' + sum);

  let str = '\n';
  for (let i = 0; i < stack.length; i += 40) {
    for (let j = 0; j < 40; j++) {
      if (Math.abs(stack[i + j] - j) <= 1) {
        str += '#';
      } else {
        str += ' ';
      }
    }
    str += '\n';
  }

  console.log('Part 2: ' + str);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
