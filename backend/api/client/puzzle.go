package client

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type PuzzleSteps struct {
	ClientId string
	UnSorted []int
	Sorted   []int
}

func GenerateRandomArray(size int) []int {
	arr := []int{}
	for i := 0; i < size; i++ {
		arr = append(arr, (rand.Intn(20)+1)*5)
	}
	return arr
}

func ConvertSliceToString(arr []int) string {
	puzzle := ""
	for indx, element := range arr {
		puzzle = fmt.Sprintf("%s,%s:%s", puzzle, strconv.Itoa(indx+1), strconv.Itoa(element))
	}
	return puzzle[1:]
}

func BinarySort(c *Client, messageType int, arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			//fmt.Printf("Checking positions: %d, %d\n", i, j)
			select {
			case <-c.Pool.SolvePuzzleStop:
				<-c.Pool.SolvePuzzleResume
				break
			case <-time.After(1 * time.Millisecond):
				break
			}
			if arr[j] < arr[i] {
				fmt.Printf("Swapping positions: %d, %d\t", i, j)
				arr[i], arr[j] = arr[j], arr[i]

				c.Puzzle = arr
				c.Pool.PuzzleUpdate <- c

				fmt.Printf("Array Step: %v\n", arr)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}

	c.Pool.PuzzleSolved <- c
}
