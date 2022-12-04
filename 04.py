import time
from libs import *


def main():
    print("--- Running Day 3 ---")

    start = time.perf_counter()

    try:
        with open("04.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    count_p1 = 0
    count_p2 = 0
    for line in data:
        p1, _, p2 = line.partition(",")
        p1_num1, _, p1_num2 = p1.partition("-")
        p2_num1, _, p2_num2 = p2.partition("-")

        p1_nums = set()
        for i in range(int(p1_num1), int(p1_num2) + 1):
            p1_nums.add(i)

        p2_nums = set()
        for i in range(int(p2_num1), int(p2_num2) + 1):
            p2_nums.add(i)

        if p1_nums.issubset(p2_nums) or p2_nums.issubset(p1_nums):
            count_p1 += 1

        if len(p1_nums.intersection(p2_nums)) > 0:
            count_p2 += 1

    print(f"Output Part 1: {count_p1}")
    print(f"Output Part 2: {count_p2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
