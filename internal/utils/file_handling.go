package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadNumbersFromFile lê números de um arquivo, um por linha
func ReadNumbersFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("erro ao converter número: %v", err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// WriteNumbersToFile escreve números em um arquivo, um por linha
func WriteNumbersToFile(filePath string, numbers []int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

