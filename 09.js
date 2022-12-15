const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('09.txt', 'utf-8');
  return input.trim().split('\n');
}

const moveSnake = (lines, len) => {
  const visited = new Set();

  const snake = [];
  for (let i = 0; i < len; i++) {
    snake.push([0, 0]);
  }

  lines.forEach((line) => {
    const fields = line.split(' ');
    const x = fields[0];
    const y = +fields[1];

    for (let i = 0; i < y; i++) {
      let dx = 0;
      if (x === 'R') {
        dx = 1;
      } else if (x === 'L') {
        dx = -1;
      }

      let dy = 0;
      if (x === 'U') {
        dy = 1;
      } else if (x === 'D') {
        dy = -1;
      }

      snake[0][0] += dx;
      snake[0][1] += dy;

      for (let j = 0; j < len - 1; j++) {
        const H = snake[j];
        const T = snake[j + 1];

        const _x = H[0] - T[0];
        const _y = H[1] - T[1];

        if (Math.abs(_x) > 1 || Math.abs(_y) > 1) {
          if (_x == 0) {
            if (_y > 0) {
              T[1] += 1;
            } else {
              T[1] += -1;
            }
          } else if (_y == 0) {
            if (_x > 0) {
              T[0] += 1;
            } else {
              T[0] += -1;
            }
          } else {
            if (_x > 0) {
              T[0] += 1;
            } else {
              T[0] += -1;
            }

            if (_y > 0) {
              T[1] += 1;
            } else {
              T[1] += -1;
            }
          }
        }
      }

      visited.add(
        JSON.stringify({
          x: snake[snake.length - 1][0],
          y: snake[snake.length - 1][1],
        })
      );
    }
  });

  return visited;
};

async function puzzle() {
  console.log('--- Running Day 9 ---');

  const start = performance.now();

  const input = await readInput();

  let visited = moveSnake(input, 2);

  console.log('Part 1: ' + visited.size);

  visited = moveSnake(input, 10);

  console.log('Part 2: ' + visited.size);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
