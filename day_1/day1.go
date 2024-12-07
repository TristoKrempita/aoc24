package day1

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
)

type RepeatedElement struct {
	Repetitions int
	Amount      int
}

func SumSimilarity() (score int) {
	left, right, err := SlicesFromFile("input", os.DirFS("."))
	cachedIDs := make(map[int]int)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, leftElement := range left {
		if cachedIDs[leftElement] != 0 {
			score += cachedIDs[leftElement]
		} else {
			cachedIDs[leftElement] = leftElement * CalculateRepetitions(leftElement, right)
			score += cachedIDs[leftElement]
		}
	}
	return
}

func CalculateRepetitions(elem int, targetSlice []int) (totalCount int) {
	// TODO: sorted slice can be searched using binary search
	for _, e := range targetSlice {
		if elem == e {
			totalCount++
		}
	}
	return
}

func SumAllDifferences() int {
	// Open file and return 2 slices
	left, right, err := SlicesFromFile("input", os.DirFS("."))
	if err != nil {
		log.Fatal(err.Error())
	}

	SortSlice(left)
	SortSlice(right)

	var finalSum float64

	for i := 0; i < len(left); i++ {
		finalSum += math.Abs(float64(left[i] - right[i]))
	}
	return int(finalSum)
}

func SlicesFromFile(fileName string, fs fs.FS) ([]int, []int, error) {
	left, right := []int{}, []int{}
	buff := make([]byte, 14)
	f, err := fs.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		count, err := f.Read(buff)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return left, right, nil
			}
			return nil, nil, err
		}
		leftNumber, rightNumber := sliceLine(buff[:count])
		left = append(left, leftNumber)
		right = append(right, rightNumber)

	}
}

func sliceLine(line []byte) (leftNum, rightNum int) {
	base := 10_000
	for _, num := range line[:5] {
		leftNum += base * int(num-'0')
		base /= 10
	}

	base = 10_000
	for _, num := range line[8:13] {
		rightNum += base * int(num-'0')
		base /= 10
	}

	return
}

func SortSlice(sliceName []int) {
	var temp int
	for j := 0; j < len(sliceName); j++ {
		for i := 0; i < len(sliceName); i++ {
			if i+1 == len(sliceName) {
				break
			}
			if sliceName[i+1] < sliceName[i] {
				temp = sliceName[i+1]
				sliceName[i+1] = sliceName[i]
				sliceName[i] = temp
			}
		}
	}
}
