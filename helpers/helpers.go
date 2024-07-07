package helpers

import (
	"Transaction/entities"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputData() (*entities.PayLoad, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Symbol: ")
	symbol, _ := reader.ReadString('\n')
	symbol = strings.TrimSpace(symbol)

	if symbol == "" {
		return nil, errors.New("Invalid Symbol")
	}

	fmt.Print("Enter Price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	price, err := strconv.ParseUint(priceStr, 10, 64)
	if err != nil {
		return nil, errors.New("Invalid price")
	}

	fmt.Print("Enter Timestamp: ")
	timestampStr, _ := reader.ReadString('\n')
	timestampStr = strings.TrimSpace(timestampStr)
	timestamp, err := strconv.ParseUint(timestampStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid timestamp:", err)
		return nil, errors.New("Invalid timestamp")
	}

	transaction := entities.PayLoad{
		Symbol:    symbol,
		Price:     price,
		Timestamp: timestamp,
	}

	return &transaction, nil
}
