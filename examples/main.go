// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/Tsubasa-2005/sumaregi-go"
// )

// func main() {
// 	envVari, err := sumaregi.LoadEnv(false)
// 	if err != nil {
// 		log.Fatalf("Failed to load environment variables: %v", err)
// 	}
// 	config := sumaregi.NewConfig(envVari)
// 	scopes := []string{sumaregi.TransactionsRead}
// 	client, err := sumaregi.NewClient(config, scopes, envVari)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	ctx := context.Background()
// 	now := time.Now()
// 	transactionDateTimeFrom := now.AddDate(0, 0, -30)
// 	transactionDateTimeTo := now.AddDate(0, 0, -1)

// 	transactions, err := client.GetTransactions(ctx, sumaregi.GetTransactionsOpts{
// 		TransactionDateTimeFrom: transactionDateTimeFrom.Format("2006-01-02T15:04:05-07:00"),
// 		TransactionDateTimeTo:   transactionDateTimeTo.Format("2006-01-02T15:04:05-07:00"),
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to get products: %v", err)
// 	}

//		sumaregi.PrintResponse(*transactions)
//		for _, transaction := range *transactions {
//			transactionDetail, err := client.GetTransactionDetail(ctx, sumaregi.GetTransactionDetailOpts{}, transaction.TransactionHeadID)
//			if err != nil {
//				log.Printf("Failed to get transaction details for ID %s: %v", transaction.TransactionHeadID, err)
//				break
//			}
//			sumaregi.PrintResponse(*transactionDetail)
//		}
//	}
package main

import (
	"context"
	"log"
	"time"

	"github.com/generalworksinc/sumaregi-go"
)

func main() {
	development := false
	envVari, err := sumaregi.LoadEnv(development)
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	config := sumaregi.NewConfig(envVari)
	scopes := []string{sumaregi.TransactionsRead}
	client, err := sumaregi.NewClient(config, scopes, envVari)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	now := time.Now()
	transactionDateTimeFrom := now.AddDate(0, 0, -1)
	transactionDateTimeTo := now

	transactions, err := client.GetTransactions(ctx, sumaregi.GetTransactionsOpts{
		Fields:                  "transactionHeadId,transactionDateTime,transactionHeadDivision,cancelDivision,unitNonDiscountsubtotal,subtotal,total,cashTotal,creditTotal,deposit,depositCash,depositCredit,change,storeId,storeCode,terminalId,customerId,customerCode,terminalTranDateTime,taxRate,taxRounding,updDateTime,transactionUuid",
		TransactionDateTimeFrom: sumaregi.FormatToISO8601(transactionDateTimeFrom),
		TransactionDateTimeTo:   sumaregi.FormatToISO8601(transactionDateTimeTo),
	})
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	// fmt.Println("len: ", len(*transactions))
	sumaregi.PrintResponse(*transactions)
	// for _, transaction := range *transactions {
	// 	transactionDetail, err := client.GetTransactionDetail(ctx, sumaregi.GetTransactionDetailOpts{}, transaction.TransactionHeadID)
	// 	if err != nil {
	// 		log.Printf("Failed to get transaction details for ID %s: %v", transaction.TransactionHeadID, err)
	// 		break
	// 	}
	// 	sumaregi.PrintResponse(*transactionDetail)
	// }
}
