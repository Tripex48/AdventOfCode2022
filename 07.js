const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('07.txt', 'utf-8');
  return input.trim().split('\n');
}

const smallestDirectory = (free_space, vals) => {
  for (const v of vals) {
    if (free_space + v >= 30000000) {
      return v;
    }
  }
};

async function puzzle() {
  console.log('--- Running Day 7 ---');

  const start = performance.now();

  const input = await readInput();

  const directories = new Map();
  const path = [];

  input.forEach((line) => {
    const words = line.trim().split(' ');
    if (words[1] === 'cd') {
      if (words[2] === '..') {
        path.pop();
      } else {
        path.push(words[2]);
      }
    } else if (words[1] !== 'ls') {
      if (words[0] !== 'dir') {
        const size = +words[0];
        for (let i = 0; i < path.length + 1; i++) {
          const newPath = path.slice(0, i).join('/').replace('//', '/').trim();
          const origSize = directories.get(newPath);
          if (origSize) {
            directories.set(newPath, origSize + size);
          } else {
            directories.set(newPath, size);
          }
        }
      }
    }
  });

  let answer = 0;
  for (const value of directories.values()) {
    if (value <= 100000) {
      answer += value;
    }
  }

  console.log('Part 1: ' + answer);

  const used_space = directories.get('/');
  const free_space = 70000000 - used_space;

  const vals = [...directories.values()];
  vals.sort();

  answer = smallestDirectory(free_space, vals);

  console.log('Part 2: ' + answer);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
