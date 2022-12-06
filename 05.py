import time
from libs import *


def main():
    print("--- Running Day 5 ---")

    start = time.perf_counter()

    try:
        with open("05.txt", "r") as f:
            data = f.read()
    except Exception as e:
        raise ValueError("Unable to read file") from e

    stack_input, moves = data.split("\n\n")

    input = stack_input.splitlines()
    input.reverse()
    stack_len = [int(x) for x in input[0].split()].pop()

    stacks_p1 = [[] for _ in range(0, stack_len)]
    stacks_p2 = [[] for _ in range(0, stack_len)]
    for line in input[1:]:
        for i, c in enumerate(line[1::4]):
            if not c.isspace():
                stacks_p1[i].append(c)
                stacks_p2[i].append(c)

    for line in moves.splitlines():
        _, num_s, _, init_s, _, dest_s = line.split()
        num, init, dest = int(num_s), int(init_s), int(dest_s)

        str = []
        for _ in range(num):
            stacks_p1[dest - 1].append(stacks_p1[init - 1].pop())
            str.append(stacks_p2[init - 1].pop())
        str.reverse()
        stacks_p2[dest - 1].extend(str)

    answer_p1 = ""
    for x in range(stack_len):
        if len(stacks_p1[x]) > 0:
            answer_p1 += stacks_p1[x][-1]

    print(f"Output Part 1: {answer_p1}")

    answer_p2 = ""
    for x in range(stack_len):
        if len(stacks_p2[x]) > 0:
            answer_p2 += stacks_p2[x][-1]

    print(f"Output Part 2: {answer_p2}")

    end = time.perf_counter()
    rtime = (end - start) * 1000  # sec -> ms
    print(f"Took {format_runtime(rtime)}\n")


if __name__ == "__main__":
    main()
