const fs = require("fs");
const { performance } = require("perf_hooks");
const { formatRuntime } = require("./libs");

async function readInput() {
  const input = await fs.promises.readFile("04.txt", "utf-8");
  return input.trim().split("\n");
}

function isSuperset(set, subset) {
  for (const elem of subset) {
    if (!set.has(elem)) {
      return false;
    }
  }
  return true;
}

function intersection(setA, setB) {
  const intersectionSet = new Set();
  for (const elem of setB) {
    if (setA.has(elem)) {
      intersectionSet.add(elem);
    }
  }
  return intersectionSet;
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

    const cmp1 = new Set();
    for (let i = p1Nums[0]; i <= p1Nums[1]; i += 1) {
      cmp1.add(i);
    }

    const cmp2 = new Set();
    for (let i = p2Nums[0]; i <= p2Nums[1]; i += 1) {
      cmp2.add(i);
    }

    if (isSuperset(cmp1, cmp2) || isSuperset(cmp2, cmp1)) {
      count_p1 += 1;
    }

    if (intersection(cmp1, cmp2).size > 0) {
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
