package main

import (
	"fmt"
	"os"
	"strconv"

	"strings"

	"sort"

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
	//for i := range len(rowz)
	for i := 0; i < len(rowz); i++ {
		if data.labtest[i] > 60 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 7)
		}

	}
	//for i := range len(rowz)
	for i := 0; i < len(rowz); i++ {
		if data.midsem[i] > 75 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 6)
		}

	}
	//for i := range len(rowz)
	for i := 0; i < len(rowz); i++ {
		if data.welabs[i] > 30 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 8)
		}

	}
	//for i := range len(rowz)
	for i := 0; i < len(rowz); i++ {
		if data.precomp[i] > 195 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 9)
		}

	}
	//for i := range len(rowz)
	for i := 0; i < len(rowz); i++ {
		if data.compre[i] > 105 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 10)
		}

	}
	//for i := range len(rowz)
	for i := 0; i < len(rowz); i++ {
		if data.total[i] > 300 {
			data.faultrow = append(data.faultrow, i)
			data.faultcol = append(data.faultcol, 11)
		}

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

type Student struct {
	Emplid int
	Total  float64
}

func topThreeStudents(data Data) {

	students := make([]Student, len(data.total))
	for i := 0; i < len(data.total); i++ {
		students[i] = Student{Emplid: data.Implide[i], Total: data.total[i]}
	}
	sort.Slice(students, func(i, j int) bool {
		return students[i].Total > students[j].Total
	})

	fmt.Println("Top 3 Students Based on Total Marks:")
	rankings := []string{"1st", "2nd", "3rd"}
	for i := 0; i < 3; i++ {
		fmt.Printf("%s Place: Emplid: %d, Marks: %.2f\n", rankings[i], students[i].Emplid, students[i].Total)
	}
}
func branchavg(data Data) float64 {
	var code string
	var count float64 = 0
	fmt.Println("enter the branch code")
	fmt.Scanln(&code)
	var sum float64 = 0.0
	switch code {
	case "A7":

		for i := 0; i < len(data.total); i++ {
			if data.campusid[5] == "7" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "A8":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "8" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "AA":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "A" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "AD":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "D" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "A2":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "2" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "A3":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "3" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "A1":

		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "1" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "A4":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "4" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "A5":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[5] == "2" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "B1":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[4] == "B" && data.campusid[5] == "1" {
				sum = sum + data.total[i]
				count++
			}

		}

	case "B5":
		var sum float64 = 0.0
		for i := 1; i <= len(data.total); i++ {
			if data.campusid[4] == "B" && data.campusid[5] == "5" {
				sum = sum + data.total[i]
				count++
			}

		}
	default:
		fmt.Println("Please enter a valid code.")
		return 0.0
	}
	if count == 0 {
		fmt.Println("No matching data found.")
		return 0.0
	}

	return sum / float64(count)

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
	rows, err := f.GetRows("CSF111_202425_01_GradeBook")
	check(err)
	//var rowz = make([]int, len(rows)-1)
	var rowz []int
	//var totals = make([]float64, len(rows)-1)
	for i, row := range rows {
		{
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
			rowz = append(rowz, i)

		}

		checkdata(data, rowz)

	}

	var choice int
	fmt.Printf("Faulty Rows index: %v,Faulty Columns index: %v\n", data.faultrow, data.faultcol)
	fmt.Println("1 total average")
	fmt.Println("2 compre avg\n3 pre-compre avg\n4 midesm average\n5 labtest average\n6 lab averages\n7 quiz average ")
	fmt.Println("8 the top three students")
	fmt.Println("9 brach wise averages")
	fmt.Println("enter the index of what u want to analys")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Printf("Total Average: %.2f\n", totalavg(data, len(data.total)))
	case 2:
		fmt.Printf("Compre Average: %.2f\n", compreavg(data, len(data.total)))
	case 3:
		fmt.Printf("Pre-Compre Average: %.2f\n", quizavg(data, len(data.total)))
	case 4:
		fmt.Printf("Midsem Average: %.2f\n", midsemavg(data, len(data.total)))
	case 5:
		fmt.Printf("Lab Test Average: %.2f\n", labtestavg(data, len(data.total)))
	case 6:
		fmt.Printf("Lab Average: %.2f\n", labavg(data, len(data.total)))
	case 7:
		fmt.Printf("Quiz Average: %.2f\n", quizavg(data, len(data.total)))
	case 8:
		topThreeStudents(data)
	case 9:
		fmt.Println("Branch-wise Average:", branchavg(data))
	default:
		fmt.Println("Invalid option.")
	}
}
