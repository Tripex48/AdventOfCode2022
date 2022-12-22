import strutils
import std/sets
import timeit

proc solve(input: string): int =
  let data = readFile input

  var priority = 0
  for line in data.split("\n"):
      var line = line.strip()
      var cmp1 = toHashSet(line[0..<int(len(line) / 2)])
      for chr in line[int(len(line) / 2)..<len(line)]:
        if chr in cmp1:
          if chr.isLowerAscii():
            priority += ord(chr) - 96
            break
          else:
            priority += ord(chr) - 64 + 26
            break
  
  return priority
  

var m = monit("Took")
m.start()
echo solve("../inputs/03.txt")
m.finish()