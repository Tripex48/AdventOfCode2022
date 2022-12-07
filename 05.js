const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('05.txt', 'utf-8');
  return input.trim().split('\n\n');
}

async function puzzle() {
  console.log('--- Running Day 5 ---');

  const start = performance.now();

  const [stackInput, moves] = await readInput();

  let input = stackInput.split('\n');
  input = input.reverse();

  const stackLen = Number(
    input[0]
      .split(/(\s+)/)
      .filter((e) => e.trim().length > 0)
      .pop()
  );

  const stacksP1 = [];
  const stacksP2 = [];
  for (let i = 0; i < stackLen; i += 1) {
    stacksP1[i] = [];
    stacksP2[i] = [];
  }

  for (let i = 1; i < input.length; i += 1) {
    const line = input[i];
    let index = 0;
    for (let j = 1; j <= line.length; j += 4) {
      if (line.slice(j, j + 1) !== '') {
        if (line.slice(j, j + 1) !== ' ') {
          stacksP1[index].push(line.slice(j, j + 1));
          stacksP2[index].push(line.slice(j, j + 1));
        }
        index += 1;
      }
    }
  }

  moves.split('\n').forEach((line) => {
    const strs = line.split(' ');
    const num = +strs[1];
    const init = +strs[3];
    const dest = +strs[5];

    let letters = '';
    for (let i = 0; i < num; i += 1) {
      const letter = stacksP1[init - 1].pop();
      if (letter) {
        stacksP1[dest - 1].push(letter);
      }

      const letterP2 = stacksP2[init - 1].pop();
      if (letterP2) {
        letters += letterP2;
      }
    }
    letters = letters.split('').reverse().join('');
    stacksP2[dest - 1].push(...letters.split(''));
  });

  let answerP1 = '';
  for (let x = 0; x < stackLen; x += 1) {
    if (stacksP1[x].length > 0) {
      answerP1 += stacksP1[x].pop();
    }
  }

  console.log('Part 1: ' + answerP1);

  let answerP2 = '';
  for (let x = 0; x < stackLen; x += 1) {
    if (stacksP2[x].length > 0) {
      answerP2 += stacksP2[x].pop();
    }
  }

  console.log('Part 2: ' + answerP2);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
