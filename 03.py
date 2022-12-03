import time
from libs import *


def main():
    print("--- Running Day 3 ---")

    start = time.perf_counter()

    try:
        with open("03.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    priority = 0
    for line in data:
        line = line.strip()
        cmp1 = set(line[: int(len(line) / 2)])
        cmp2 = set(line[int(len(line) / 2) :])
        (dup,) = cmp1 & cmp2
        if dup.islower():
            priority += ord(dup) - 96
        else:
            priority += ord(dup) - 64 + 26

    print(f"Output Part 1: {priority}")

    priority = 0
    for i in range(0, len(data), 3):
        cmp1 = set(data[i].strip())
        cmp2 = set(data[i + 1].strip())
        cmp3 = set(data[i + 2].strip())
        (dup,) = cmp1 & cmp2 & cmp3
        if dup.islower():
            priority += ord(dup) - 96
        else:
            priority += ord(dup) - 64 + 26

    print(f"Output Part 2: {priority}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
