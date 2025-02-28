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
		os.Exit(1) // to stop the code from running
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

func checkdata(data Data, rowz []int) {
	if data.sc_no == nil {
		fmt.Println("no data found")
		return
	}
	for i := 0; i < len(rowz); i++ {
		if data.quiz[i] > 30 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 5)

		}

	}
	for i := range len(rowz) {
		if data.labtest[i] > 60 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 7)
		}

	}
	for i := range len(rowz) {
		if data.midsem[i] > 75 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 6)
		}

	}
	for i := range len(rowz) {
		if data.welabs[i] > 30 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 8)
		}

	}
	for i := range len(rowz) {
		if data.precomp[i] > 195 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 9)
		}

	}
	for i := range len(rowz) {
		if data.compre[i] > 105 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 10)
		}

	}
	for i := range len(rowz) {
		data.faultrow = append(data.faultrow, i)
		data.faultcol = append(data.faultcol, 11)
	}

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
func quizavg(data Data, j int) float64 {
	if j == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < j; i++ {
		sum += data.quiz[i]
	}

	return sum / float64(j)
}
func labavg(data Data, j int) float64 {
	if j == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < j; i++ {
		sum += data.labtest[i]
	}

	return sum / float64(j)
}
func midsemavg(data Data, j int) float64 {
	if j == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < j; i++ {
		sum += data.midsem[i]
	}

	return sum / float64(j)
}
func labtestavg(data Data, j int) float64 {
	if j == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < j; i++ {
		sum += data.labtest[i]
	}

	return sum / float64(j)
}
func compreavg(data Data, j int) float64 {
	if j == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < j; i++ {
		sum += data.compre[i]
	}

	return sum / float64(j)
}
func totalavg(data Data, j int) float64 {
	if j == 0 {
		return 0
	}

	var sum float64
	for i := 0; i < j; i++ {
		sum += data.total[i]
	}

	return sum / float64(j)
}
func topthree(data Data) []float64 {
	a := make([]float64, len(data.total))
	copy(a, data.total)
	for i := 0; i < len(data.total); i++ {
		for j := i + 1; j < len(data.total); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a[:3]

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	f, err := excelize.OpenFile(filePath)
	check(err)
	defer func() {

		if cerr := f.Close(); cerr != nil {
			fmt.Println(cerr)
		}
	}()
	data := Data{}
	rows, err := f.GetRows("Sheet1")
	check(err)
	var rowz = make([]int, len(rows)-1)
	//var totals = make([]float64, len(rows)-1)
	var j float64 = 0.0

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

		rowz[i] = i + 1
		j = float64(i + 1)
	}

	checkdata(data, rowz)

	fmt.Printf("Faulty Rows: %v\n,Faulty Columns: %v\n", data.faultrow, data.faultcol)

}
