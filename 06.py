import time
from libs import *
from collections import Counter


def isUniqueChars(string: str) -> bool:
    freq = Counter(string)

    if len(freq) == len(string):
        return True
    else:
        return False


def main():
    print("--- Running Day 6 ---")

    start = time.perf_counter()

    try:
        with open("06.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    input = data[0]
    answer_p1 = 0
    for i in range(0, len(input) - 4):
        marker = "".join(input[i + x] for x in range(4))

        if isUniqueChars(marker):
            answer_p1 = i + 4
            break

    print(f"Output Part 1: {answer_p1}")

    answer_p2 = 0
    for i in range(0, len(input) - 14):
        marker = "".join(input[i + x] for x in range(14))

        if isUniqueChars(marker):
            answer_p2 = i + 14
            break

    print(f"Output Part 2: {answer_p2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
