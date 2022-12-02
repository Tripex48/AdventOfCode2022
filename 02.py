import time
from libs import *


def main():
    print("--- Running Day 2 ---")

    start = time.perf_counter()

    try:
        with open("02.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    results = {
        "A X": 4,
        "A Y": 8,
        "A Z": 3,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 7,
        "C Y": 2,
        "C Z": 6,
    }

    total_score = play(data, results)

    print(f"Output Part 1: {total_score}")

    results_r2 = {
        "A X": 3,
        "A Y": 4,
        "A Z": 8,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 2,
        "C Y": 6,
        "C Z": 7,
    }

    total_score_r2 = play(data, results_r2)

    print(f"Output Part 2: {total_score_r2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


def play(data, results):
    total_score = 0

    for line in data:
        total_score += results[line.strip()]
    return total_score


if __name__ == "__main__":
    main()
