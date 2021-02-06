package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func run() error {
	var limit int
	profile := os.Getenv("BRAVE_PROFILE")
	flag.StringVar(&profile, "profile", profile, "Brave profile directory")
	flag.IntVar(&limit, "limit", 0, "Limit n results")
	flag.Parse()

	query := strings.Join(flag.Args(), " ")
	entries, err := queryHistory(profile, query, query, limit)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
	return nil
}
