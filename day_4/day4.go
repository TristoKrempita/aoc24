package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CrossSearchFromFile(filename string) int {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return CrossSearch(string(f))
}

func WordSearchFromFile(filename string) int {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return WordSearch(string(f))
}

func CrossSearch(crossword string) (totalX_mas int) {
	crosswordSlice := stringToSlice(crossword)
	// printCrosswordSlice(crosswordSlice)
	// row
	for j := 0; j < len(crosswordSlice); j++ {
		// column
		for i := 0; i < len(crosswordSlice[j]); i++ {
			if crosswordSlice[i][j] == rune('A') {
				if checkFirstDiagonal(crosswordSlice, i, j) &&
					checkSecondDiagonal(crosswordSlice, i, j) {
					totalX_mas++
				}
			}
		}
	}
	return
}

func WordSearch(crossword string) (totalXmas int) {
	crosswordSlice := stringToSlice(crossword)
	// printCrosswordSlice(crosswordSlice)
	// row
	for j := 0; j < len(crosswordSlice); j++ {
		// column
		for i := 0; i < len(crosswordSlice[j]); i++ {
			if crosswordSlice[j][i] == rune('X') {
				if i+3 < len(crosswordSlice[j]) {
					if crosswordSlice[j][i+1] == rune('M') {
						totalXmas += findInLineForwards(crosswordSlice[j][i : i+4])
					}
				}
				if i-3 >= 0 {
					if crosswordSlice[j][i-1] == rune('M') {
						totalXmas += findInLineBackwards(crosswordSlice[j][i-3 : i+1])
					}
				}
				if j+3 < len(crosswordSlice) {
					if crosswordSlice[j+1][i] == rune('M') {
						totalXmas += findInColumnDownwards(extractColumnDown(crosswordSlice, j, i))
					}
				}

				if j-3 >= 0 {
					if crosswordSlice[j-1][i] == rune('M') {
						totalXmas += findInColumnUpwards(extractColumnUp(crosswordSlice, j, i))
					}
				}

				if i+3 < len(crosswordSlice[j]) && j+3 < len(crosswordSlice) {
					if crosswordSlice[j+1][i+1] == rune('M') {
						totalXmas += findInDiagonal(extractDiagonalDownRight(crosswordSlice, j, i))
					}
				}

				if i-3 >= 0 && j-3 >= 0 {
					if crosswordSlice[j-1][i-1] == rune('M') {
						totalXmas += findInDiagonal(extractDiagonalTopLeft(crosswordSlice, j, i))
					}
				}

				if i+3 < len(crosswordSlice[j]) && j-3 >= 0 {
					if crosswordSlice[j-1][i+1] == rune('M') {
						totalXmas += findInDiagonal(extractDiagonalTopRight(crosswordSlice, j, i))
					}
				}

				if i-3 >= 0 && j+3 < len(crosswordSlice) {
					if crosswordSlice[j+1][i-1] == rune('M') {
						totalXmas += findInDiagonal(extractDiagonalDownLeft(crosswordSlice, j, i))
					}
				}
			}
		}
	}
	return
}

func stringToSlice(cs string) (finalSlice [][]rune) {
	scanner := bufio.NewScanner(strings.NewReader(cs))
	for scanner.Scan() {
		lineSlice := []rune{}
		for _, ch := range scanner.Text() {
			lineSlice = append(lineSlice, ch)
		}
		finalSlice = append(finalSlice, lineSlice)
	}

	return finalSlice
}

func printCrosswordSlice(crosswordSlice [][]rune) {
	fmt.Print("//")
	for i := range 10 {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for j := 0; j < len(crosswordSlice); j++ {
		fmt.Print(j, " ")
		for i := 0; i < len(crosswordSlice[j]); i++ {
			fmt.Print(string(crosswordSlice[j][i]), " ")
		}
		fmt.Println("|")
	}
	fmt.Println("-----------------------")
}

func findInLineForwards(line []rune) int {
	if strings.Compare(string(line), "XMAS") == 0 {
		return 1
	}
	return 0
}

func findInLineBackwards(line []rune) int {
	if strings.Compare(string(line), "SAMX") == 0 {
		return 1
	}
	return 0
}

func findInColumnDownwards(line []rune) int {
	if strings.Compare(string(line), "XMAS") == 0 {
		return 1
	}
	return 0
}

func findInColumnUpwards(line []rune) int {
	if strings.Compare(string(line), "XMAS") == 0 {
		return 1
	}
	return 0
}

func findInDiagonal(line []rune) int {
	if strings.Compare(string(line), "XMAS") == 0 {
		return 1
	}
	return 0
}

func extractColumnDown(crosswordSlice [][]rune, row, col int) []rune {
	column := []rune{}
	for i := range 4 {
		column = append(column, crosswordSlice[row+i][col])
	}
	return column
}

func extractColumnUp(crosswordSlice [][]rune, row, col int) []rune {
	column := []rune{}
	for i := range 4 {
		column = append(column, crosswordSlice[row-i][col])
	}
	return column
}

func extractDiagonalDownRight(crosswordSlice [][]rune, row, col int) (xmasSlice []rune) {
	for i := range 4 {
		xmasSlice = append(xmasSlice, crosswordSlice[i+row][i+col])
	}
	return
}

func extractDiagonalTopLeft(crosswordSlice [][]rune, row, col int) (xmasSlice []rune) {
	for i := range 4 {
		xmasSlice = append(xmasSlice, crosswordSlice[row-i][col-i])
	}
	return
}

func extractDiagonalTopRight(crosswordSlice [][]rune, row, col int) (xmasSlice []rune) {
	for i := range 4 {
		xmasSlice = append(xmasSlice, crosswordSlice[row-i][col+i])
	}
	return
}

func extractDiagonalDownLeft(crosswordSlice [][]rune, row, col int) (xmasSlice []rune) {
	for i := range 4 {
		xmasSlice = append(xmasSlice, crosswordSlice[row+i][col-i])
	}
	return
}

func checkFirstDiagonal(crosswordSlice [][]rune, row, col int) bool {
	var currentDiagonal []rune
	if row+1 < len(crosswordSlice) && row-1 >= 0 && col+1 < len(crosswordSlice[row]) && col-1 >= 0 {
		currentDiagonal = append(currentDiagonal, crosswordSlice[row-1][col-1])
		currentDiagonal = append(currentDiagonal, crosswordSlice[row][col])
		currentDiagonal = append(currentDiagonal, crosswordSlice[row+1][col+1])
	}
	if strings.Compare(string(currentDiagonal), "MAS") == 0 ||
		strings.Compare(string(currentDiagonal), "SAM") == 0 {
		return true
	}
	return false
}

func checkSecondDiagonal(crosswordSlice [][]rune, row, col int) bool {
	var currentDiagonal []rune
	if row+1 < len(crosswordSlice) && row-1 >= 0 && col+1 < len(crosswordSlice[row]) && col-1 >= 0 {
		currentDiagonal = append(currentDiagonal, crosswordSlice[row+1][col-1])
		currentDiagonal = append(currentDiagonal, crosswordSlice[row][col])
		currentDiagonal = append(currentDiagonal, crosswordSlice[row-1][col+1])
	}
	if strings.Compare(string(currentDiagonal), "MAS") == 0 ||
		strings.Compare(string(currentDiagonal), "SAM") == 0 {
		return true
	}
	return false
}
