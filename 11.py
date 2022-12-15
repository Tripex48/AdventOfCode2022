from __future__ import annotations

import time
from libs import *
import copy
from typing import Callable


def main():
    print("--- Running Day 11 ---")

    start = time.perf_counter()

    try:
        with open("11.txt", "r") as f:
            data = f.read()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    monkeys: list[tuple[list[int], Callable[[int], int], int, int, int]] = []

    for monkey in data.strip().split("\n\n"):
        lines = monkey.splitlines()
        nums = list([int(x) for x in lines[1].split(": ")[1].split(", ")])
        calc = eval("lambda old:" + lines[2].split("=")[1])
        div = int(lines[3].split()[-1])
        pos = int(lines[4].split()[-1])
        neg = int(lines[5].split()[-1])
        monkeys.append((nums, calc, div, pos, neg))

    counts = solve(monkeys)
    print(f"Output Part 1: {counts[-1] * counts[-2]}")

    counts = solve(monkeys, 10000, 2)
    print(f"Output Part 2: {counts[-1] * counts[-2]}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


def solve(monkeys, iter=20, part=1):
    counts = [0] * len(monkeys)
    monkeys_temp = copy.deepcopy(monkeys)

    mod = 1
    if part == 2:
        for input in monkeys_temp:
            mod *= input[2]

    for _ in range(iter):
        for index, input in enumerate(monkeys_temp):
            for item in input[0]:
                item = input[1](item)

                if part == 1:
                    item //= 3
                else:
                    item %= mod

                if item % input[2] == 0:
                    monkeys_temp[input[3]][0].append(item)
                else:
                    monkeys_temp[input[4]][0].append(item)

            counts[index] += len(input[0])

            (_, calc, div, pos, neg) = input
            monkeys_temp[index] = ([], calc, div, pos, neg)

    counts.sort()
    return counts


if __name__ == "__main__":
    main()
