package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var id int
	var set int
	start := time.Now()
	endtime := time.Now()
	var startbidprice int
	var initialbidprice int
	var ref int
	var category string
	var prod_name string
	var desc string
	var mrp int
	var base int
	header := []string{
		"id", "set", "start", "endtime", "startbidprice", "initialbidprice", "ref", "category", "product", "description", "mrp", "base",
	}
	csvFile, csverr := os.Create("twelvehourtable.csv")
	if csverr != nil {
		log.Fatalf("failed creating file: %s", csverr)
	}
	csvwriter := csv.NewWriter(csvFile)
	_ = csvwriter.Write(header)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err: ", err)
	}
	db, dberr := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if dberr != nil {
		log.Fatal("dberr: ", dberr)
	}
	fmt.Println("opened the db connection")
	rows, rowserr := db.Query("select * from twelvehourtable")
	if rowserr != nil {
		log.Fatal("rowserr: ", rowserr)
	}
	fmt.Println("Rows are read")
	for rows.Next() {
		err := rows.Scan(&id, &set, &start, &endtime, &startbidprice, &initialbidprice, &ref, &category, &prod_name, &desc, &mrp, &base)
		if err != nil {
			log.Fatal("looping: ", err)
		}
		row := []string{
			strconv.Itoa(id),
			strconv.Itoa(set),
			start.String(),
			endtime.String(),
			strconv.Itoa(startbidprice),
			strconv.Itoa(initialbidprice),
			strconv.Itoa(ref),
			category,
			prod_name,
			desc,
			strconv.Itoa(mrp),
			strconv.Itoa(base),
		}
		_ = csvwriter.Write(row)

	}
	csvwriter.Flush()
	csvFile.Close()
}
