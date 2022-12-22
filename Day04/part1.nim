import strutils
import timeit

proc contains(
        self_start: int,
        self_end: int,
        other_start: int,
        other_end: int,
): bool =
    return (self_start <= other_start) and (self_end >= other_end)

proc solve(input: string): int =
  let data = readFile input
  
  var count = 0

  for line in data.split("\n"):
    var parts_str = line.split(",")
    
    var p1_num_str = parts_str[0].split("-")
    var p1_num1 = parseInt(p1_num_str[0])
    var p1_num2 = parseInt(p1_num_str[1])
    
    var p2_num_str = parts_str[1].split("-")
    var p2_num1 = parseInt(p2_num_str[0])
    var p2_num2 = parseInt(p2_num_str[1])

    if contains(
            p1_num1, p1_num2, p2_num1, p2_num2
        ) or contains(p2_num1, p2_num2, p1_num1, p1_num2):
            count += 1

  return count 
  
var m = monit("Took")
m.start()
echo solve("../inputs/04.txt")
m.finish()