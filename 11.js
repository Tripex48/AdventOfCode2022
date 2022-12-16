const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('11.txt', 'utf-8');
  return input.trim().split('\n\n');
}

const solve = (monkeys, iter, part) => {
  const counts = [];
  let mod = 1n;
  for (let i = 0; i < monkeys.length; i++) {
    counts.push(0);
    if (part === 2) {
      mod *= BigInt(monkeys[i][2]);
    }
  }

  for (let i = 0; i < iter; i++) {
    for (let j = 0; j < monkeys.length; j++) {
      const input = monkeys[j];
      input[0].forEach((item) => {
        let val = input[1](item);

        if (part == 1) {
          val = val / 3n;
        } else {
          val = val % mod;
        }

        if (val % BigInt(input[2]) === 0n) {
          monkeys[input[3]][0].push(val);
        } else {
          monkeys[input[4]][0].push(val);
        }
      });

      counts[j] += input[0].length;
      monkeys[j][0] = [];
    }
  }

  const answers = counts.sort((a, b) => a - b);
  return answers.pop() * answers.pop();
};

async function puzzle() {
  console.log('--- Running Day 11 ---');

  const start = performance.now();

  const input = await readInput();

  let monkeys = [];
  input.forEach((monkey) => {
    const lines = monkey.split('\n');
    const nums = lines[1]
      .split(': ')[1]
      .split(', ')
      .map((x) => BigInt(+x));
    const test = lines[2].split('=')[1].match(/([0-9])+/g);
    let calc = null;
    if (test) {
      const newnum = test[0] + 'n';
      calc = eval(
        '(old) => ' + lines[2].split('=')[1].replace(/([0-9])+/g, newnum)
      );
    } else {
      calc = eval('(old) => ' + lines[2].split('=')[1]);
    }
    const div = lines[3].split(' ').pop();
    const pos = lines[4].split(' ').pop();
    const neg = lines[5].split(' ').pop();

    monkeys.push([nums, calc, div, pos, neg]);
  });

  console.log('Part 1: ' + solve(monkeys, 20, 1));

  monkeys = [];
  input.forEach((monkey) => {
    const lines = monkey.split('\n');
    const nums = lines[1]
      .split(': ')[1]
      .split(', ')
      .map((x) => BigInt(+x));
    const test = lines[2].split('=')[1].match(/([0-9])+/g);
    let calc = null;
    if (test) {
      const newnum = test[0] + 'n';
      calc = eval(
        '(old) => ' + lines[2].split('=')[1].replace(/([0-9])+/g, newnum)
      );
    } else {
      calc = eval('(old) => ' + lines[2].split('=')[1]);
    }
    const div = lines[3].split(' ').pop();
    const pos = lines[4].split(' ').pop();
    const neg = lines[5].split(' ').pop();

    monkeys.push([nums, calc, div, pos, neg]);
  });

  console.log('Part 2: ' + solve(monkeys, 10000, 2));

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
