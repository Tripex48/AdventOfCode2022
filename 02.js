const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('02.txt', 'utf-8');
  return input.split('\n');
}

async function puzzle() {
  console.log('--- Running Day 2 ---');

  const start = performance.now();

  const input = await readInput();

  results = {
    'A X': 4,
    'A Y': 8,
    'A Z': 3,
    'B X': 1,
    'B Y': 5,
    'B Z': 9,
    'C X': 7,
    'C Y': 2,
    'C Z': 6,
  };

  total_score = play(input, results);
  console.log('Part 1: ' + total_score);

  results_r2 = {
    'A X': 3,
    'A Y': 4,
    'A Z': 8,
    'B X': 1,
    'B Y': 5,
    'B Z': 9,
    'C X': 2,
    'C Y': 6,
    'C Z': 7,
  };

  total_score_r2 = play(input, results_r2);
  console.log('Part 2: ' + total_score_r2);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

function play(input, results) {
  let total_score = 0;

  input.forEach((x) => {
    total_score += results[x.trim()];
  });

  return total_score;
}

puzzle();
