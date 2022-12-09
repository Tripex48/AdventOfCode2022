import time
from libs import *


def main():
    print("--- Running Day 8 ---")

    start = time.perf_counter()

    try:
        with open("08.txt", "r") as f:
            data = f.readlines()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    # Part 1
    def checkTopP1(grid: list[list[int]], i: int, j: int) -> bool:
        num = grid[i][j]
        for x in range(0, i):
            if not num > grid[x][j]:
                return False
        return True

    def checkBottomP1(grid: list[list[int]], i: int, j: int) -> bool:
        num = grid[i][j]
        for x in range(i + 1, len(grid)):
            if not num > grid[x][j]:
                return False
        return True

    def checkLeftP1(grid: list[list[int]], i: int, j: int) -> bool:
        num = grid[i][j]
        for x in range(0, j):
            if not num > grid[i][x]:
                return False
        return True

    def checkRightP1(grid: list[list[int]], i: int, j: int) -> bool:
        num = grid[i][j]
        for x in range(j + 1, len(grid)):
            if not num > grid[i][x]:
                return False
        return True

    lines = data
    grid = []
    scenicScores = []
    visible = len(lines[0]) * 2 + (len(lines) * 2) - 4
    for line in lines:
        numbers = [int(x) for x in line.strip()]
        grid.append(numbers)
        scores = [0 for _ in line]
        scenicScores.append(scores)

    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[0]) - 1):
            if (
                checkTopP1(grid, i, j)
                or checkLeftP1(grid, i, j)
                or checkRightP1(grid, i, j)
                or checkBottomP1(grid, i, j)
            ):
                visible += 1

    print(f"Output Part 1: {visible}")

    # Part 2
    def checkTopP2(grid: list[list[int]], i: int, j: int) -> int:
        num = grid[i][j]
        score = 0
        for x in reversed(range(0, i)):
            if num <= grid[x][j]:
                score += 1
                break
            elif num >= grid[x][j]:
                score += 1
        return score

    def checkBottomP2(grid: list[list[int]], i: int, j: int) -> int:
        num = grid[i][j]
        score = 0
        for x in range(i + 1, len(grid)):
            if num <= grid[x][j]:
                score += 1
                break
            elif num >= grid[x][j]:
                score += 1
        return score

    def checkLeftP2(grid: list[list[int]], i: int, j: int) -> int:
        num = grid[i][j]
        score = 0
        for x in reversed(range(0, j)):
            if num <= grid[i][x]:
                score += 1
                break
            elif num >= grid[i][x]:
                score += 1
        return score

    def checkRightP2(grid: list[list[int]], i: int, j: int) -> int:
        num = grid[i][j]
        score = 0
        for x in range(j + 1, len(grid)):
            if num <= grid[i][x]:
                score += 1
                break
            elif num >= grid[i][x]:
                score += 1
        return score

    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[0]) - 1):
            score = (
                checkTopP2(grid, i, j)
                * checkLeftP2(grid, i, j)
                * checkRightP2(grid, i, j)
                * checkBottomP2(grid, i, j)
            )
            scenicScores[i][j] = score

    answer_p2 = max([i for lst in scenicScores for i in lst])

    print(f"Output Part 2: {answer_p2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
