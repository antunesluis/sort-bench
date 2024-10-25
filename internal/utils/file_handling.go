package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (f *FileHandler) ReadNumbers(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, scanner.Err()
}

func (f *FileHandler) WriteNumbers(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		if _, err := writer.WriteString(fmt.Sprintf("%d\n", num)); err != nil {
			return err
		}
	}

	return writer.Flush()
}

