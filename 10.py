from __future__ import annotations

import time
from libs import *


def main():
    print("--- Running Day 10 ---")

    start = time.perf_counter()

    try:
        with open("10.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    x = 1
    stack = []

    for line in data:
        if "noop" in line:
            stack.append(x)
        else:
            v = int(line.split()[1])
            stack.append(x)
            stack.append(x)
            x += v

    sum = 0
    for i in [20, 60, 100, 140, 180, 220]:
        sum += i * stack[i - 1]

    print(f"Output Part 1: {sum}")

    str = "\n"
    for i in range(0, len(stack), 40):
        for j in range(0, 40):
            if abs(stack[i + j] - j) <= 1:
                str += "#"
            else:
                str += " "
        str += "\n"

    print(f"Output Part 2: {str}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
