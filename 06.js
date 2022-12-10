const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('06.txt', 'utf-8');
  return input.trim().split('\n');
}

const isUnique = (str) => {
  return new Set(str).size == str.length;
};

const getMarkerStart = (input, startIndex, markerLength) => {
  let answer = 0;
  for (let i = startIndex; i < input.length - markerLength; i++) {
    let marker = '';
    for (let j = 0; j < markerLength; j++) {
      marker += input[i + j];
    }

    if (isUnique(marker)) {
      answer = i + markerLength;
      return answer;
    }
  }
};

async function puzzle() {
  console.log('--- Running Day 6 ---');

  const start = performance.now();

  const input = await readInput();

  let answer = getMarkerStart(input[0], 0, 4);
  console.log('Part 1: ' + answer);

  answer = getMarkerStart(input[0], answer - 1, 14);
  console.log('Part 2: ' + answer);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
