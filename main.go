package main

import (
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"fmt"
	"math"
	"flag"
)

func getLines() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

type Pair struct {
	x, y float64
}

func parse(lines []string) []Pair {
	pairs := []Pair{}
	for _, line := range lines {
		split := strings.Split(line, "\t")
		if len(split) != 2 {
			continue
		}
		x, _ := strconv.ParseFloat(split[0], 64)
		y, _ := strconv.ParseFloat(split[1], 64)
		pairs = append(pairs, Pair{x, y})
	}
	return pairs
}

func getSumsOfPowersOfX(pairs []Pair, maxPow int) []float64 {
	sums := []float64{}
	for i := 0; i <= maxPow; i++ {
		sum := float64(0)
		for _, pair := range pairs {
			sum += math.Pow(pair.x, float64(i))
		}
		sums = append(sums, sum)
	}
	return sums
}

func getSumsOfPowersOfXTimesY(pairs []Pair, maxPow int) []float64 {
	sums := []float64{}
	for i := 0; i <= maxPow; i++ {
		sum := float64(0)
		for _, pair := range pairs {
			sum += pair.y * math.Pow(pair.x, float64(i))
		}
		sums = append(sums, sum)
	}
	return sums
}

func printMatrix(rows [][]float64) {
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if j != 0 {
				fmt.Print("\t")
			}
			fmt.Print(rows[i][j])
		}
		fmt.Println()
	}
}

func toScientific(x float64) string {
	var prefix string
	if x < 0 {
		x = -x
		prefix = "-"
	}
	e := math.Floor(math.Log10(x))
	if e*e < 10 {
		return fmt.Sprintf("%v%v", prefix, x)
	}
	c := x / math.Pow(10, e)
	return fmt.Sprintf("%v%v*10^(%v)", prefix, c, int(e))
}

func main() {
	degptr := flag.Int("degree", 2, "Degree of the polynomial")
	printMat := flag.Bool("matrix", false, "Print matrix instead of solutions")
	eq := flag.Bool("equation", false, "Print as equation")

	flag.Parse()

	deg := *degptr
	lines := getLines()
	pairs := parse(lines)
	xPows := getSumsOfPowersOfX(pairs, 2*deg)
	yxPows := getSumsOfPowersOfXTimesY(pairs, deg)

	rows := [][]float64{}
	for i := 0; i <= deg; i++ {
		row := []float64{}
		for j := i; j <= deg+i; j++ {
			row = append(row, xPows[j])
		}
		row = append(row, yxPows[i])
		rows = append(rows, row)
	}

	if *printMat {
		printMatrix(rows)
		return
	}

	solved := gaussJordan(rows)

	if *eq {
		for i, variable := range solved {
			if i > 0 {
				fmt.Print(" + ")
			}
			fmt.Printf("(%v)*x^%v", toScientific(variable), i)
		}
		fmt.Println()
		return
	}
	for i, variable := range solved {
		if i > 0 {
			fmt.Print("\t")
		}
		fmt.Print(variable)
	}
	fmt.Println()
}
