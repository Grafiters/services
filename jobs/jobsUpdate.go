package jobs

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"riskmanagement/lib"
	"riskmanagement/models"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// - [ ] Cronjob update expired date by timestime, Not including this service api, registered on crontab linux
// - search index where documen if expired_date = now , create to transactionExpireds and delete index from transactions

func JobsUpdate() {
	dateNow := GetTimeNow()
	transactions, err := MatchSearch(dateNow)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("\n ========= Start Create Transaction Expired =========  \n")
	var stat bool
	for _, transaction := range transactions {
		fmt.Println(transaction)
		fmt.Println(transaction.Id)

		fmt.Println("\n ========= Start Create Transaction Expired =========  \n")
		fmt.Println("data transaction=> \n", transaction)
		stat, err = CreateTransactionExpired(transaction)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("\n ========= End Create Transaction Expired =========  \n")
	}
	fmt.Printf("\n ========= End CreateTransactionExpired successfull = %s =========  \n", stat)

	fmt.Printf("\n ========= Delete DeleteTransactionOpen successfull = %s =========  \n", stat)
	if stat {
		status, err := DeleteTransactionOpen(dateNow)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(status)
	}
	fmt.Printf("\n ========= End DeleteTransactionOpen successfull = %s =========  \n", stat)

	// tampung ke struct []
	// create ke transaction.IndexTransactionExpired()
	// loop tampungan
	// delete by id loop ke transaction.IndexTransactionOpen()
}

func MatchSearch(dateNow string) (transactions []models.Transaction, err error) {

	dBEHost, err := lib.GetVarEnv("DBEHost")
	if err != nil {
		return nil, err
	}

	dBEUsername, err := lib.GetVarEnv("DBEUsername")
	if err != nil {
		return nil, err
	}

	// client elasticsearch.Client
	var transaction models.Transaction
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	// Instantiate an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			dBEHost,
		},
		Username: dBEUsername,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	// Check for connection errors to the Elasticsearch cluster
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
		filename, function, line := lib.WhereAmI()
		lib.CreateLogErrorToDB(client, filename, function, line, "Elasticsearch connection error", fmt.Sprintf("%v", err))
	}

	// Create a new query string for the Elasticsearch method call

	var query = `
	"query_string": {
		"query" : "` + dateNow + `",
		"fields"  : ["expiredDate"]
	}`

	// Pass the query string to the function and have it return a Reader object
	read := constructQuery(query, 2)

	// Example of an invalid JSON string
	//read = constructQuery("{bad json", 2)

	fmt.Println("read:", read)

	// Instantiate a map interface object for storing returned documents
	var mapResp map[string]interface{}
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("json.NewEncoder() ERROR:", err)
		filename, function, line := lib.WhereAmI()
		lib.CreateLogErrorToDB(client, filename, function, line, "json.NewEncoder() ERROR", fmt.Sprintf("%v", err))
		// Query is a valid JSON object
	} else {
		fmt.Println("json.NewEncoder encoded query:", read, "\n")

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex(transaction.IndexTransactionOpen()),
			client.Search.WithBody(read),
			client.Search.WithTrackTotalHits(true),
		)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)
			filename, function, line := lib.WhereAmI()
			lib.CreateLogErrorToDB(client, filename, function, line, "Elasticsearch Search() API ERROR", fmt.Sprintf("%v", err))

			// If no errors are returned, parse esapi.Response object
		} else {
			fmt.Println("res TYPE:", reflect.TypeOf(res))

			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
				fmt.Println(`&mapResp:`, &mapResp, "\n")
				fmt.Println(`mapResp["hits"]:`, mapResp["hits"])
			}

			hits := mapResp["hits"].(map[string]interface{})["hits"]
			for _, hit := range hits.([]interface{}) {
				log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
				log.Println(strings.Repeat("=>", 37))
				source := hit.(map[string]interface{})["_source"]

				id := hit.(map[string]interface{})["_id"]
				appname := source.(map[string]interface{})["appname"]
				data := source.(map[string]interface{})["data"]
				prefix := source.(map[string]interface{})["prefix"]
				expiredDate := source.(map[string]interface{})["expiredDate"]
				referenceCode := source.(map[string]interface{})["referenceCode"]
				status := source.(map[string]interface{})["status"]
				created := source.(map[string]interface{})["created"]
				lastUpdate := source.(map[string]interface{})["lastUpdate"]

				transactions = append(transactions, models.Transaction{
					Id:            id.(string),
					Appname:       appname.(string),
					Data:          data,
					Prefix:        prefix.(string),
					ExpiredDate:   expiredDate.(string),
					ReferenceCode: referenceCode.(string),
					Status:        status.(string),
					Created:       created.(string),
					LastUpdate:    lastUpdate.(string),
				})
				log.Println(strings.Repeat("=>", 37))
			}
			fmt.Println("transactions", transactions)
		}
	}
	return transactions, err
}

func DeleteTransactionOpen(dateNow string) (status bool, err error) {
	// client elasticsearch.Client
	var transaction models.Transaction
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	// Instantiate an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("DBEHost"),
		},
		Username: os.Getenv("DBEUsername"),
		Password: os.Getenv("DBEPassword"),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	// Check for connection errors to the Elasticsearch cluster
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	// Create a new query string for the Elasticsearch method call

	var query = `
	"query_string": {
		"query" : "` + dateNow + `",
		"fields"  : ["expiredDate"]
	}`

	// Pass the query string to the function and have it return a Reader object
	read := constructQuery(query, 2)

	// Example of an invalid JSON string
	//read = constructQuery("{bad json", 2)

	fmt.Println("read:", read)
	reqDelete := esapi.DeleteByQueryRequest{
		Index: []string{transaction.IndexTransactionOpen()},
		Body:  read,
	}

	resDelete, err := reqDelete.Do(ctx, client)
	defer resDelete.Body.Close()
	fmt.Println(resDelete.Body)

	if resDelete.IsError() {
		return false, fmt.Errorf("Delete: response: %s", resDelete.String())
	}
	// old
	// Instantiate a map interface object for storing returned documents
	// var mapResp map[string]interface{}
	// var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	// if err := json.NewEncoder(&buf).Encode(read); err != nil {
	// 	log.Fatalf("json.NewEncoder() ERROR:", err)

	// 	// Query is a valid JSON object
	// } else {
	// 	fmt.Println("json.NewEncoder encoded query:", read, "\n")

	// 	// Pass the JSON query to the Golang client's Search() method
	// 	res, err := client.Search(
	// 		client.Search.WithContext(ctx),
	// 		client.Search.WithIndex(transaction.IndexTransactionOpen()),
	// 		client.Search.WithBody(read),
	// 		client.Search.WithTrackTotalHits(true),
	// 	)

	// 	// Check for any errors returned by API call to Elasticsearch
	// 	if err != nil {
	// 		log.Fatalf("Elasticsearch Search() API ERROR:", err)

	// 		// If no errors are returned, parse esapi.Response object
	// 	} else {
	// 		fmt.Println("res TYPE:", reflect.TypeOf(res))

	// 		// Close the result body when the function call is complete
	// 		defer res.Body.Close()

	// 		// Decode the JSON response and using a pointer
	// 		if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
	// 			fmt.Println(`&mapResp:`, &mapResp, "\n")
	// 			fmt.Println(`mapResp["failures"]:`, mapResp["failures"])
	// 		}

	// 		failures := mapResp["failures"]
	// 		// .(map[string]interface{})
	// 		fmt.Println(failures)

	// 		timeout := mapResp["timed_out"]
	// 		if timeout.(string) == "true" {
	// 			fmt.Println("timeout=>", timeout)
	// 			return false, err
	// 		}

	// 		if failures == nil {
	// 			fmt.Println("failures=>", failures)
	// 			return true, err
	// 		}
	// 		return false, err

	// 	}
	// }
	return true, err
}

func CreateTransactionExpired(transaction models.Transaction) (status bool, err error) {
	var tranExpired models.TransactionExpired
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Instantiate an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("DBEHost"),
		},
		Username: os.Getenv("DBEUsername"),
		Password: os.Getenv("DBEPassword"),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	// Check for connection errors to the Elasticsearch cluster
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}
	bdy, err := json.Marshal(transaction)
	if err != nil {
		return false, err
	}

	// res, err := p.elastic.client.Create()
	req := esapi.CreateRequest{
		Index:      tranExpired.IndexTransactionExpired(),
		DocumentID: lib.UUID(false),
		Body:       bytes.NewReader(bdy),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := req.Do(ctx, client)
	fmt.Println("response => ", res)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return false, err
	}

	return true, err
}
