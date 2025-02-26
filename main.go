package main

import (
	"fmt"
	"os"
	"strconv"

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

type Data struct {
	sc_no    []int
	class_no []int
	Implide  []int
	campusid []string
	quiz     []float64
	midsem   []float64
	labtest  []float64
	welabs   []float64
	precomp  []float64
	compre   []float64
	total    []float64
	faultrow []int
	faultcol []int
}

func parseInt(value string) (int, error) {
	if value == "" {
		return 0, nil
	}
	return strconv.Atoi(value)
}
func parseFloat(value string) (float64, error) {
	if value == "" {
		return 0.0, nil
	}
	return strconv.ParseFloat(value, 64)
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
	data := Data{}
	rows, err := f.GetRows("Sheet1")
	check(err)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if isRowEmpty(row) {
			continue
		}
		scNo, err1 := parseInt(row[0])
		check(err1)
		classNo, err2 := parseInt(row[1])
		check(err2)
		implide, err3 := parseInt(row[2])
		check(err3)

		campusID := row[3]

		quiz, err4 := parseFloat(row[4])
		check(err4)

		midsem, err5 := parseFloat(row[5])
		check(err5)

		labTest, err6 := parseFloat(row[6])
		check(err6)

		weLabs, err7 := parseFloat(row[7])
		check(err7)
		precomp, err8 := parseFloat(row[8])
		check(err8)

		compre, err9 := parseFloat(row[9])
		check(err9)

		total, err10 := parseFloat(row[10])
		check(err10)

		data.sc_no = append(data.sc_no, scNo)
		data.class_no = append(data.class_no, classNo)
		data.Implide = append(data.Implide, implide)
		data.campusid = append(data.campusid, campusID)
		data.quiz = append(data.quiz, quiz)
		data.midsem = append(data.midsem, midsem)
		data.labtest = append(data.labtest, labTest)
		data.welabs = append(data.welabs, weLabs)
		data.precomp = append(data.precomp, precomp)
		data.compre = append(data.compre, compre)
		data.total = append(data.total, total)

	}
	fmt.Printf("Faulty Rows: %v\n", data.faultrow)
	fmt.Printf("Faulty Columns: %v\n", data.faultcol)

}
