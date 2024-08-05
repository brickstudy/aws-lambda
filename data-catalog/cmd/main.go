package main

import (
	"log"

	"github.com/robert-min/aws-lambda/data-catalog/adapter/config"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	_, err = mysql.New(config.DB)
	if err != nil {
		log.Fatal(err)
	}
}
