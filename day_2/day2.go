package day2

import (
	"bufio"
	"io/fs"
	"log"
	"math"
	"strconv"
	"strings"
)

func CountSafeFromFile(fileName string, filesystem fs.FS) (safeLevels int) {
	data, err := fs.ReadFile(filesystem, fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	return CountSafeLevels(string(data))
}

func CountSafeLevels(levels string) (safeLevels int) {
	reader := strings.NewReader(levels)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		anlysisResult := 0
		level := scanner.Text()
		levelSlice := LevelToSlice(level)
		if AnalyzeLevel(levelSlice) == 0 {
			for i := 0; i < len(levelSlice); i++ {
				newSlice := make([]int, len(levelSlice))
				copy(newSlice, levelSlice)
				newSlice = newSlice[: i : i+1]
				if i+1 != len(levelSlice) {
					newSlice = append(newSlice, levelSlice[i+1:]...)
				}
				anlysisResult = AnalyzeLevel(newSlice)
				if anlysisResult == 1 {
					safeLevels++
					break
				}
			}
		} else {
			safeLevels++
		}
	}
	return
}

func AnalyzeLevel(line []int) int {
	var comparissonFunction func(int, int) bool
	if line[0] < line[1] {
		comparissonFunction = lessThan
	} else {
		comparissonFunction = moreThan
	}
	for i := 0; i < len(line); i++ {
		if i+1 == len(line) {
			break
		}
		difference := math.Abs(float64(line[i] - line[i+1]))
		if !comparissonFunction(line[i], line[i+1]) || !(difference < 4) || !(difference > 0) {
			return 0
		}
	}
	return 1
}

func lessThan(a, b int) bool {
	return a < b
}

func moreThan(a, b int) bool {
	return a > b
}

func LevelToSlice(level string) (levelSlice []int) {
	stringSlice := strings.Split(level, " ")
	for _, stringElement := range stringSlice {
		intElement, err := strconv.Atoi(stringElement)
		if err != nil {
			panic(err)
		}
		levelSlice = append(levelSlice, intElement)
	}
	return
}
