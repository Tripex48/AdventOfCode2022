import strutils
import std/tables
import timeit

proc play(data: string, results: Table[string, int]): int =
  var total_score = 0

  for line in data.split('\n'):
      total_score += results[line.strip()]
  return total_score

proc solve(input: string): int =
  let data = readFile input

  var results = {
    "A X": 4,
    "A Y": 8,
    "A Z": 3,
    "B X": 1,
    "B Y": 5,
    "B Z": 9,
    "C X": 7,
    "C Y": 2,
    "C Z": 6,
  }.toTable

  return play(data, results)
  

var m = monit("Took")
m.start()
echo solve("../inputs/02.txt")
m.finish()