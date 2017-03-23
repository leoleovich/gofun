package main

import (
	"os"
	"github.com/gocarina/gocsv"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

/*
Data example
"String","String","String","String","String","String","String","String","String","String","String","String","String","String","String","String"
 */

const chunks = 76
const pass = "pass"

type Data struct {
	Val1,Val2,Val3,Val4,Val5,Val6,Val7,Val8,Val9,Val10,Val11,Val12,Val13,Val14,Val15,Val16 string
}

func execute(data []*Data, thread int) {
	fmt.Printf("+%d:%s Starting with chunk size %d\n", thread, time.Now().String(), len(data))
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:3306)/test_oleg?autocommit=false", pass))
	if err != nil {
		fmt.Println("Unable to connect to mysql", err.Error())
		return
	}
	defer db.Close()

	for transaction:=0; transaction*2000 < len(data) ; transaction++ {
		//fmt.Printf("Transaction %d\n", transaction)
		tx, err := db.Begin()
		if err != nil {
			fmt.Println("Unable to start transaction", err.Error())
			return
		}

		query := ""
		for i:=0; i<2000; i++ {
			tmp_data := &Data{}
			if transaction*2000+i < len(data) {
				tmp_data = data[transaction*2000+i]
			} else {
				break
			}

			query = fmt.Sprintf("insert into data values('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');\n",
				tmp_data.Val1,
				tmp_data.Val2,
				tmp_data.Val3,
				tmp_data.Val4,
				tmp_data.Val5,
				tmp_data.Val6,
				tmp_data.Val7,
				tmp_data.Val8,
				tmp_data.Val9,
				tmp_data.Val10,
				tmp_data.Val11,
				tmp_data.Val12,
				tmp_data.Val13,
				tmp_data.Val14,
				tmp_data.Val15,
				tmp_data.Val16)

			_, err = tx.Exec(query)
			if err != nil {
				tx.Rollback()
				fmt.Println(err)
				fmt.Println(query)
				return
			}
		}

		err = tx.Commit()
		if err != nil {
			fmt.Println(err)
			fmt.Println(query)
			return
		}
	}
	fmt.Printf("+%d:%s Finished\n", thread, time.Now().String())
}

func cleanDB() error {
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:3306)/test_oleg", pass))
	if err != nil {
		fmt.Println("Unable to connect to mysql")
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Unable to start transaction")
		return err
	}
	tx.Exec("TRUNCATE data")
	return tx.Commit()
}

func main() {
	file, err := os.Open("/tmp/huge.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []*Data{}
	if err := gocsv.UnmarshalFile(file, &data); err != nil { // Load clients from file
		panic(err)
	}

	err = cleanDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	chunkSize := len(data)/chunks
	fmt.Println("Queries:", len(data))
	for thread:=0 ; thread < chunks; thread++ {
		if chunks-thread == 1 {
			fmt.Printf("Thread %d: %d - %d\n", thread, thread*chunkSize, len(data)-1)
			go execute(data[thread*chunkSize:len(data)-1], thread)
		} else {
			fmt.Printf("Thread %d: %d - %d\n", thread, thread*chunkSize, (thread+1)*chunkSize-1)
			go execute(data[thread*chunkSize:(thread+1)*chunkSize-1], thread)
		}
	}
	wg.Wait()
}
