import strutils
import sugar
import math
import std/algorithm
import timeit

proc solve(input: string): int =
  let data = readFile input
  var sums: seq[int]
  for group in data.split("\n\n"):
    var calories = collect(for x in group.split("\n"): parseInt(x))
    var sum = sum(calories)
    sums.add sum

  sums = sorted(sums, Descending)
  var answer = sums[0] + sums[1] + sums[2]
  return answer

var m = monit("Took")
m.start()
echo solve("../inputs/01.txt")
m.finish()