const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('01.txt', 'utf-8');
  return input.split('\n\n');
}

async function puzzle() {
  console.log('--- Running Day 1 ---');
  const start = performance.now();

  const input = await readInput();
  let highest = 0;
  const elfs = [];
  input.forEach((x) => {
    const elf = x
      .split('\n')
      .map(Number)
      .reduce((sum, val) => sum + val);
    elfs.push(elf);
    if (elf > highest) {
      highest = elf;
    }
  });

  console.log('Part 1: ' + highest);

  elfs.sort((a, b) => (a > b ? -1 : 1));
  console.log('Part 2: ' + (elfs[0] + elfs[1] + elfs[2]));

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
