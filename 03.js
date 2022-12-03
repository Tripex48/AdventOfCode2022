const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('03.txt', 'utf-8');
  return input.split('\n');
}

async function puzzle() {
  console.log('--- Running Day 3 ---');

  const start = performance.now();

  const input = await readInput();

  let priority = 0;
  input.forEach((line) => {
    const entry = line.trim();
    const cmp1 = new Set(entry.slice(0, entry.length / 2));
    const cmp2 = new Set(entry.slice(entry.length / 2));
    const [dup] = new Set([...cmp1].filter((i) => cmp2.has(i)));
    if (dup == dup.toLowerCase()) {
      priority += dup.charCodeAt(0) - 96;
    } else {
      priority += dup.charCodeAt(0) - 64 + 26;
    }
  });

  console.log('Part 1: ' + priority);

  priority = 0;
  for (let i = 0; i < input.length; i += 3) {
    const cmp1 = new Set(input[i].trim());
    const cmp2 = new Set(input[i + 1].trim());
    const cmp3 = new Set(input[i + 2].trim());
    const [dup] = new Set(
      [...cmp1].filter((i) => cmp2.has(i)).filter((i) => cmp3.has(i))
    );
    if (dup == dup.toLowerCase()) {
      priority += dup.charCodeAt(0) - 96;
    } else {
      priority += dup.charCodeAt(0) - 64 + 26;
    }
  }

  console.log('Part 2: ' + priority);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
