const fs = require("fs");
const { performance } = require("perf_hooks");
const { formatRuntime } = require("./libs");

async function readInput() {
  const input = await fs.promises.readFile("04.txt", "utf-8");
  return input.trim().split("\n");
}

function contains(self_start, self_end, other_start, other_end) {
  return self_start <= other_start && self_end >= other_end;
}

function overlap(self_start, self_end, other_start, other_end) {
  if (self_end < other_start || other_end < self_start) {
    return false;
  }
  return true;
}

async function puzzle() {
  console.log("--- Running Day 4 ---");

  const start = performance.now();

  const input = await readInput();

  let count_p1 = 0;
  let count_p2 = 0;
  input.forEach((line) => {
    const [p1Str, p2Str] = line.split(",");
    const p1Nums = p1Str.split("-").map((x) => +x);
    const p2Nums = p2Str.split("-").map((x) => +x);

    if (
      contains(p1Nums[0], p1Nums[1], p2Nums[0], p2Nums[1]) ||
      contains(p2Nums[0], p2Nums[1], p1Nums[0], p1Nums[1])
    ) {
      count_p1 += 1;
    }

    if (overlap(p1Nums[0], p1Nums[1], p2Nums[0], p2Nums[1])) {
      count_p2 += 1;
    }
  });

  console.log("Part 1: " + count_p1);
  console.log("Part 2: " + count_p2);

  const end = performance.now();
  const rtime = end - start;
  console.log("Took:", formatRuntime(rtime));
}

puzzle();
