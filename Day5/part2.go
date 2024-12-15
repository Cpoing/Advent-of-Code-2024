package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	orderingRulesBytes, err := os.ReadFile("orders.txt")
	if err != nil {
		fmt.Println("Error reading orders.txt:", err)
		return
	}
	orderingRules := string(orderingRulesBytes)

	updatesBytes, err := os.ReadFile("updates.txt")
	if err != nil {
		fmt.Println("Error reading updates.txt:", err)
		return
	}
	updates := string(updatesBytes)

	lines := strings.Split(strings.TrimSpace(orderingRules+"\n\n"+updates), "\n")
	output := 0

	rules := make(map[int][]int)
	updatesList := [][]int{}

	isAddRule := true

	for _, line := range lines {
		if len(line) == 0 {
			isAddRule = false
			continue
		}

		if isAddRule {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])

			rules[num1] = append(rules[num1], num2)
		} else {
			currentUpdate := []int{}
			parts := strings.Split(line, ",")

			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				currentUpdate = append(currentUpdate, num)
			}

			updatesList = append(updatesList, currentUpdate)
		}
	}

	for _, currentUpdate := range updatesList {
		isValid := true
		for i := len(currentUpdate) - 1; i > 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}
		}

		if !isValid {
			ordered := []int{}
			remaining := map[int]bool{}
			dependencies := make(map[int][]int)

			for _, num := range currentUpdate {
				remaining[num] = true

				for _, dep := range rules[num] {
					dependencies[dep] = append(dependencies[dep], num)
				}
			}

			for len(remaining) > 0 {
				for num := range remaining {
					if len(dependencies[num]) == 0 {
						ordered = append(ordered, num)
						delete(remaining, num)

						for key, val := range dependencies {
							newList := []int{}

							for _, n := range val {
								if n != num {
									newList = append(newList, n)
								}
							}

							dependencies[key] = newList
						}
					}
				}
			}

			middle := ordered[len(ordered)/2]
			output += middle
		}
	}

	fmt.Println("Sum : ", output)
}

