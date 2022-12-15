const fs = require('fs');
const { performance } = require('perf_hooks');
const { formatRuntime } = require('./libs');

async function readInput() {
  const input = await fs.promises.readFile('08.txt', 'utf-8');
  return input.trim().split('\n');
}

const checkTop = (grid, i, j) => {
  const num = grid[i][j];
  for (let x = 0; x < i; x++) {
    if (num <= grid[x][j]) {
      return false;
    }
  }
  return true;
};

const checkTopP2 = (grid, i, j) => {
  const num = grid[i][j];
  let score = 0;
  for (let x = i - 1; x >= 0; x--) {
    if (num <= grid[x][j]) {
      score += 1;
      break;
    } else if (num >= grid[x][j]) {
      score += 1;
    }
  }
  return score;
};

const checkBottom = (grid, i, j) => {
  const num = grid[i][j];
  for (let x = i + 1; x < grid.length; x++) {
    if (num <= grid[x][j]) {
      return false;
    }
  }
  return true;
};

const checkBottomP2 = (grid, i, j) => {
  const num = grid[i][j];
  let score = 0;
  for (let x = i + 1; x < grid.length; x++) {
    if (num <= grid[x][j]) {
      score += 1;
      break;
    } else if (num >= grid[x][j]) {
      score += 1;
    }
  }
  return score;
};

const checkLeft = (grid, i, j) => {
  const num = grid[i][j];
  for (let x = 0; x < j; x++) {
    if (num <= grid[i][x]) {
      return false;
    }
  }
  return true;
};

const checkLeftP2 = (grid, i, j) => {
  const num = grid[i][j];
  let score = 0;
  for (let x = j - 1; x >= 0; x--) {
    if (num <= grid[i][x]) {
      score += 1;
      break;
    } else if (num >= grid[i][x]) {
      score += 1;
    }
  }
  return score;
};

const checkRight = (grid, i, j) => {
  const num = grid[i][j];
  for (let x = j + 1; x < grid.length; x++) {
    if (num <= grid[i][x]) {
      return false;
    }
  }
  return true;
};

const checkRightP2 = (grid, i, j) => {
  const num = grid[i][j];
  let score = 0;
  for (let x = j + 1; x < grid.length; x++) {
    if (num <= grid[i][x]) {
      score += 1;
      break;
    } else if (num >= grid[i][x]) {
      score += 1;
    }
  }
  return score;
};

async function puzzle() {
  console.log('--- Running Day 8 ---');

  const start = performance.now();

  const lines = await readInput();
  const grid = [];
  let visible = lines[0].length * 2 + lines.length * 2 - 4;

  lines.forEach((line) => {
    const numbers = line.split('').map((x) => +x);
    grid.push(numbers);
  });

  for (let i = 1; i < grid.length - 1; i++) {
    for (let j = 1; j < grid[0].length - 1; j++) {
      if (
        checkTop(grid, i, j) ||
        checkLeft(grid, i, j) ||
        checkRight(grid, i, j) ||
        checkBottom(grid, i, j)
      ) {
        visible += 1;
      }
    }
  }

  console.log('Part 1: ' + visible);

  let maxScore = 0;
  for (let i = 1; i < grid.length - 1; i++) {
    for (let j = 1; j < grid[0].length - 1; j++) {
      const score =
        checkTopP2(grid, i, j) *
        checkLeftP2(grid, i, j) *
        checkRightP2(grid, i, j) *
        checkBottomP2(grid, i, j);

      if (score > maxScore) {
        maxScore = score;
      }
    }
  }

  console.log('Part 2: ' + maxScore);

  const end = performance.now();
  const rtime = end - start;
  console.log('Took:', formatRuntime(rtime));
}

puzzle();
