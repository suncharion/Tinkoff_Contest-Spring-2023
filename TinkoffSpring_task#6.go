package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	Max, Score int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	count, max := parseInput(firstLine)

	studs := []*Student{}
	for i := 0; i < count; i++ {
		line, _ := reader.ReadString('\n')
		min, smax := parseInput(line)
		studs = append(studs, &Student{
			Max:   smax,
			Score: min,
		})
	}

	fmt.Println(MaxMedian(max, studs))
}

func parseInput(in string) (int, int) {
	strs := strings.Split(strings.Trim(in, "\r\n"), " ")
	out1, _ := strconv.Atoi(strs[0])
	out2, _ := strconv.Atoi(strs[1])
	return out1, out2
}

func MaxMedian(scoreLimit int, studs []*Student) int {
	medianI := len(studs) / 2
	sum := 0
	max := 0
	for i := range studs {
		if studs[i].Score > max {
			max = studs[i].Score
		}
		sum += studs[i].Score
		scoreLimit -= studs[i].Score
	}

	if scoreLimit != 0 {
		sort.Slice(studs, func(i, j int) bool {
			if studs[i].Max != studs[j].Max {
				return studs[i].Max > studs[j].Max
			}
			return studs[i].Score > studs[j].Score
		})
	} else {
		sort.Slice(studs, func(i, j int) bool {
			return studs[i].Score > studs[j].Score
		})
		return studs[medianI].Score
	}

	changed := true
	for changed && scoreLimit > 0 {
		changed = false
		min := max
		currMinEl := 0
		for i := 0; i <= medianI; i++ {
			if studs[i].Score < min {
				min = studs[i].Score
				currMinEl = i
			}
		}

		if min == max {
			break
		}
		studs[currMinEl].Score++
		scoreLimit--
		changed = true
	}

	for scoreLimit > 0 {
		changed := false
		leftOvers := false

		for i := 0; i < len(studs); i++ {
			if studs[i].Score >= studs[i].Max {
				continue
			}
			if i > 0 {
				if studs[i].Score >= studs[i-1].Score {
					continue
				}
			}

			if i > medianI && !leftOvers {
				continue
			}
			studs[i].Score++
			scoreLimit--
			if scoreLimit == 0 {
				break
			}
			changed = true
		}
		if !changed {
			leftOvers = true
		}
	}

	sort.Slice(studs, func(i, j int) bool {
		return studs[i].Score > studs[j].Score
	})

	return studs[medianI].Score
}
