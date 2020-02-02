package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// get number of test cases and prepare for pipeline
	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	caseNum := getNum(reader)
	wg.Add(caseNum)

	resChannel := getSquareSum(caseNum, &wg)

	// blocks until wg.Done() is called caseNum times
	wg.Wait()
	outPut(resChannel, 0, caseNum)
}

// getNum: get number from standard input
func getNum(reader *bufio.Reader) int {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	num, _ := strconv.Atoi(s)
	return num
}

// getSquareSum: compute the sum for each test case and put them into a channel
func getSquareSum(caseNum int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int, caseNum)
	counter := 0
	reader := bufio.NewReader(os.Stdin)

	go func() {
		// use label and goto instead of for statement
		// just to fulfill the rule, this is not recommended for production code
	getSquare:
		size := getNum(reader)
		n, _ := reader.ReadString('\n')
		n = strings.TrimSpace(n)
		numsList := strings.Split(n, " ")
		sum := 0

		getSum(&sum, 0, size, numsList)

		out <- sum
		wg.Done()
		counter++
		if counter < caseNum {
			goto getSquare
		}
		close(out)
	}()

	return out
}

// getSum: recursively compute the sum according to rule.
func getSum(sum *int, counter int, size int, numsList []string) {
	if counter >= size {
		return
	}
	num, _ := strconv.Atoi(numsList[counter])
	if num > 0 {
		*sum += num * num
	}
	counter++
	getSum(sum, counter, size, numsList)
}

// outPut: recursively output to standard output.
func outPut(out <-chan int, counter int, num int) {
	if counter >= num {
		return
	}
	fmt.Println(<-out)
	counter++
	outPut(out, counter, num)
}
