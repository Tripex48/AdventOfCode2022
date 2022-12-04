import time
from libs import *


def contains(self_start, self_end, other_start, other_end):
    return (self_start <= other_start) and (self_end >= other_end)


def overlap(self_start, self_end, other_start, other_end):
    if self_end < other_start or other_end < self_start:
        return False
    return True


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

        if contains(
            int(p1_num1), int(p1_num2), int(p2_num1), int(p2_num2)
        ) or contains(int(p2_num1), int(p2_num2), int(p1_num1), int(p1_num2)):
            count_p1 += 1

        if overlap(int(p1_num1), int(p1_num2), int(p2_num1), int(p2_num2)):
            count_p2 += 1

    print(f"Output Part 1: {count_p1}")
    print(f"Output Part 2: {count_p2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
