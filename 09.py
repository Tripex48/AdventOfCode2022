from __future__ import annotations

import time
import enum
from libs import *


class Direction(enum.Enum):
    UP = (0, -1)
    RIGHT = (1, 0)
    DOWN = (0, 1)
    LEFT = (-1, 0)

    def __init__(self, x: int, y: int) -> None:
        self.x, self.y = x, y

    @property
    def _vals(self) -> tuple[Direction, ...]:
        return tuple(type(self).__members__.values())

    @property
    def cw(self) -> Direction:
        vals = self._vals
        return vals[(vals.index(self) + 1) % len(vals)]

    @property
    def ccw(self) -> Direction:
        vals = self._vals
        return vals[(vals.index(self) - 1) % len(vals)]

    @property
    def opposite(self) -> Direction:
        vals = self._vals
        return vals[(vals.index(self) + 2) % len(vals)]

    def apply(self, x: int, y: int, *, n: int = 1) -> tuple[int, int]:
        return self.x * n + x, self.y * n + y


def main():
    print("--- Running Day 9 ---")

    start = time.perf_counter()

    try:
        with open("09.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    D = {
        'R': Direction.RIGHT,
        'U': Direction.UP,
        'L': Direction.LEFT,
        'D': Direction.DOWN,
    }

    head = tail = (0, 0)
    seen: set[tuple[int, int]] = set()
    seen.add(tail)

    for line in data:
        dir_s, n_s = line.split()
        move = D[dir_s]
        n = int(n_s)

        for _ in range(n):
            head = move.apply(*head)
            if abs(head[0] - tail[0]) >= 2 or abs(head[1] - tail[1]) >= 2:
                tail = move.opposite.apply(*head)
                seen.add(tail)

    print(f"Output Part 1: {len(seen)}")

    def fixup(head: tuple[int, int], tail: tuple[int, int]) -> tuple[int, int]:
        hx, hy = head
        tx, ty = tail

        if abs(hy - ty) == 2 and abs(hx - tx) == 2:
            return ((hx + tx) // 2, (hy + ty) // 2)
        if abs(hy - ty) == 2:
            return (hx, (ty + hy) // 2)
        elif abs(hx - tx) == 2:
            return ((tx + hx) // 2, hy)
        else:
            return tail

    snake: list[tuple[int, int]] = [(0, 0)] * 10
    seen = {snake[0]}

    for line in data:
        dir_s, n_s = line.split()
        move = D[dir_s]
        nums = int(n_s)

        for _ in range(nums):
            snake[0] = move.apply(*snake[0])

            prev = snake[0]
            for i in range(1, len(snake)):
                snake[i] = fixup(prev, snake[i])
                prev = snake[i]

            seen.add(snake[-1])

    print(f"Output Part 2: {len(seen)}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
