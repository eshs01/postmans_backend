package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/xuri/excelize/v2"
)

// fuction to check errors
func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}

}
func isRowEmpty(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

type data struct {
	sc_no    []int
	class_no []int
}

func main() {

	dat, err := os.ReadFile("/tmp/dat")
	check(err)

	f, err := excelize.OpenFile(string(dat))
	check(err)
	defer func() {

		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("Sheet1")
	check(err)

	for _, row := range rows { // this code was taken to skip the empty rows in aspreadsheet

		if isRowEmpty(row) {
			continue
		}

	}
}
