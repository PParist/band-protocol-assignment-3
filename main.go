package main

import (
	"Transaction/handler"
	"Transaction/helpers"
	"Transaction/services"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	transactionService := services.NewTransactionService("https://mock-node-wgqbnxruha-as.a.run.app")
	transactionHandler := handler.NewTransactionHandler(transactionService)

	fmt.Println("BroadcastTransaction....")
	payload, err := helpers.InputData()
	if err != nil {
		log.Fatal(err)
	}
	result, err := transactionHandler.BroadcastTransaction(*payload)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You have tx_hash: ", result)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Input Your Hash: ")
		scanner.Scan()
		if scanner.Text() != "" {
			hash := scanner.Text()
			for {
				res, err := transactionHandler.TransactionStatusMonitoring(hash)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("tx_status: ", res)

				switch res {
				case "CONFIRMED":
					fmt.Println("Transaction Confirmed !")
					return
				case "FAILED":
					fmt.Println("Transaction failed to process !")
					return
				case "PENDING":
					fmt.Println("Transaction is awaiting processing...")
				case "DNE":
					fmt.Println("Transaction does not exist.")
					return
				default:
					return
				}

				fmt.Print("Type 'yes' to continue: ")

				for scanner.Scan() {

					if err := scanner.Err(); err != nil {
						log.Fatal(err)
					}

					if scanner.Text() == "yes" {
						break
					} else {
						log.Fatal("Invalid input.")
					}
				}

			}
		} else {
			log.Fatal("Invalid Hash.")
		}

	}

}
