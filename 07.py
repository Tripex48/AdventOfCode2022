import time
from libs import *

from collections import defaultdict


def main():
    print("--- Running Day 7 ---")

    start = time.perf_counter()

    try:
        with open("07.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    directories = defaultdict(int)
    path = []

    for line in data:
        words = line.strip().split()
        if words[1] == 'cd':
            if words[2] == '..':
                path.pop()
            else:
                path.append(words[2])
        elif words[1] == 'ls':
            continue
        else:
            if words[0] != 'dir':
                size = int(words[0])
                for i in range(len(path)+1):
                    new_path = '/'.join(path[:i]).replace('//', '/')
                    if new_path != '':
                        directories[new_path] += size

    answer_p1 = 0
    for value in directories.values():
        if value <= 100000:
            answer_p1 += value

    print(f"Output Part 1: {answer_p1}")

    used_space = directories['/']
    free_space = 70000000 - used_space

    sorted_directories = sorted(
        directories.items(), key=lambda x: x[1],
    )
    sorted_output = dict(sorted_directories)

    answer_p2 = 0
    for v in sorted_output.values():
        if free_space + v >= 30000000:
            answer_p2 = v
            break

    print(f"Output Part 2: {answer_p2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
