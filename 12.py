from __future__ import annotations

import heapq
import time
from libs import *
from typing import Generator


def get_neighbors(x: int, y: int) -> Generator[tuple[int, int], None, None]:
    yield x - 1, y
    yield x, y - 1
    yield x + 1, y
    yield x, y + 1


def bfs(
    coords: dict[tuple[int, int], str],
    start: tuple[int, int],
    end: tuple[int, int],
    part=1,
) -> int:
    visited: list[tuple[int, int]] = [end]
    queue: list[tuple[int, tuple[int, int]]] = []
    heapq.heappush(queue, (0, end))

    while queue:
        dist, last_coord = heapq.heappop(queue)
        x, y = last_coord
        val = coords[last_coord]

        for i, j in get_neighbors(x, y):
            if (i, j) not in coords:
                continue
            if (i, j) in visited:
                continue
            if ord(coords[(i, j)]) - ord(val) < -1:
                continue

            if part == 1:
                if (i, j) == start:
                    return dist + 1
            else:
                if coords[(i, j)] == "a":
                    return dist + 1

            visited.append((i, j))
            heapq.heappush(queue, (dist + 1, (i, j)))

    return 0


def main():
    print("--- Running Day 12 ---")

    start = time.perf_counter()

    try:
        with open("12.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    coords: dict[tuple[int, int], str] = {}
    start_point = (0, 0)
    end_point = (-1, -1)
    for row, line in enumerate(data):
        for col, chr in enumerate(line):
            if chr == "S":
                start_point = (row, col)
                coords[(row, col)] = "a"
            elif chr == "E":
                end_point = (row, col)
                coords[(row, col)] = "z"
            else:
                coords[(row, col)] = chr

    print(f"Output Part 1: {bfs(coords, start_point, end_point)}")
    print(f"Output Part 2: {bfs(coords, start_point, end_point, 2)}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
