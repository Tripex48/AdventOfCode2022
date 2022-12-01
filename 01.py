import time
from libs import *


def main():
    print("--- Running Day 1 ---")

    start = time.perf_counter()

    try:
        with open("01.txt", "r") as f:
            data = f.read()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    elf_input = data.split("\n\n")
    highest_elf = 0
    elfs = []
    for input in elf_input:
        calories = sum([int(x) for x in input.split("\n") if x != ""])
        elfs.append(calories)
        if calories > highest_elf:
            highest_elf = calories

    elfs = sorted(elfs, reverse=True)
    three_elfs = sum(elfs[:3])

    end = time.perf_counter()

    print(f"Output Part 1: {highest_elf}")
    print(f"Output Part 2: {three_elfs}")
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
