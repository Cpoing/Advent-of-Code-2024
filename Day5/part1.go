package main

import (
	"fmt"
	"strings"
  "os"
)

func isValidUpdate(rules map[int][]int, update []int) bool {
	pos := make(map[int]int)
	for i, page := range update {
		pos[page] = i
	}

	for page, dependencies := range rules {
		if p, ok := pos[page]; ok {
			for _, dep := range dependencies {
				if d, ok := pos[dep]; ok && d <= p {
					return false
				}
			}
		}
	}

	return true
}


func main() {
	orderingRulesBytes, err := os.ReadFile("orders.txt")
  if err != nil {
    fmt.Println("error", err)
    return
  }

  orderingRules := string(orderingRulesBytes)

	updatesBytes, err := os.ReadFile("./updates.txt")
  if err != nil {
    fmt.Println("error", err)
    return
  }

  updates := string(updatesBytes)

	rules := make(map[int][]int)
	for _, line := range strings.Split(orderingRules, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d|%d", &x, &y)
		rules[x] = append(rules[x], y)
	}

	var sum int
	for _, updateLine := range strings.Split(updates, "\n") {
		updateParts := strings.Split(updateLine, ",")
		update := make([]int, len(updateParts))
		for i, part := range updateParts {
			fmt.Sscanf(part, "%d", &update[i])
		}

		if isValidUpdate(rules, update) {
			middle := update[len(update)/2]
			sum += middle
		}
	}

	fmt.Println("Sum: ", sum)
}

