import strutils
import std/sets
import timeit

proc solve(input: string): int =
  let data = readFile input
  let lines = data.split('\n')

  var priority = 0
  for i in countup(0, len(lines)-3, 3):
    var cmp1 = toHashSet(lines[i].strip())
    var cmp2 = toHashSet(lines[i + 1].strip())
    var cmp3 = toHashSet(lines[i + 2].strip())
    var dup = cmp1 * cmp2 * cmp3
    var dupChr = dup.pop()
    if dupChr.isLowerAscii():
        priority += ord(dupChr) - 96
    else:
        priority += ord(dupChr) - 64 + 26  
  
  return priority
  

var m = monit("Took")
m.start()
echo solve("../inputs/03.txt")
m.finish()