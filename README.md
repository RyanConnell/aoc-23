# aoc-23
Solutions for AdventOfCode 2023: https://adventofcode.com/2023

## Validation and Timing
Solutions can be validated and timed by running `go test` in the directory for each day.
```
~/golang/src/github.com/RyanConnell/aoc-23/09 (master) » go test
|=====================================================|
| Day 09 | Part 1 (sample)     | took 9.735µs         |
| Day 09 | Part 1 (final)      | took 986.957µs       |
| Day 09 | Part 2 (sample)     | took 4.667µs         |
| Day 09 | Part 2 (final)      | took 843.214µs       |
|=====================================================|
PASS
ok      github.com/RyanConnell/aoc-23/09        0.004s
```
Alternatively you can do this for _all_ days at once by running `go test ./... -v` and then filtering the output a little.
```
~/golang/src/github.com/RyanConnell/aoc-23 (master) » go test ./... -v | grep "|"
|=====================================================|
| Day 01 | Part 2 (sample)     | took 7.579µs         |
| Day 01 | Part 2 (final)      | took 1.927923ms      |
|=====================================================|                                                               
|=====================================================|                                                               
| Day 02 | Part 1 (sample)     | took 32.767µs        |
| Day 02 | Part 1 (final)      | took 920.148µs       |
...
```
