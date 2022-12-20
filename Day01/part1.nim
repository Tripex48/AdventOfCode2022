import strutils
import sugar
import math
import timeit

proc solve(input: string): int =
  var highest = 0
  let data = readFile input
  for group in data.split("\n\n"):
    var calories = collect(for x in group.split("\n"): parseInt(x))
    var sum = sum(calories)
    if sum > highest:
      highest = sum

  return highest

var m = monit("Took")
m.start()
echo solve("../inputs/01.txt")
m.finish()