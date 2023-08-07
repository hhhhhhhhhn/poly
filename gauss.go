package main

import (
	"math"
)

func cloneMatrix(rows [][]float64) [][]float64 {
	clonedRows := make([][]float64, len(rows))
	for i, row := range rows {
		clonedRows[i] = make([]float64, len(row))
		copy(clonedRows[i], row)
	}
	return clonedRows
}

const EPSILON = 1e-10

func isZero(float float64) bool {
	return math.Abs(float) < EPSILON
}

func swapRows(rows [][]float64, i, j int) {
	rows[i], rows[j] = rows[j], rows[i]
}

func divideRowByFactor(rows [][]float64, i int, scale float64) {
	for j := range rows[i] {
		rows[i][j] /= scale
	}
}

func addRowToRowWithFactor(rows [][]float64, src, dest int, factor float64) {
	for j := range rows[dest] {
		rows[dest][j] += factor * rows[src][j]
	}
}

func gaussJordan(rows [][]float64) []float64 {
	rows = cloneMatrix(rows)

	pivotRow := 0
	pivotCol := 0

	for pivotRow < len(rows) && pivotCol < len(rows[0]) - 1 {
		if isZero(rows[pivotRow][pivotCol]) {
			for swapRow := pivotRow + 1; swapRow < len(rows); swapRow++ {
				if !isZero(rows[swapRow][pivotCol]) {
					swapRows(rows, swapRow, pivotRow)
					break
				}
			}
		}
		if isZero(rows[pivotRow][pivotCol]) {
			pivotCol++
			continue
		}
		for targetRow := 0; targetRow < len(rows); targetRow++ {
			if targetRow == pivotRow {
				continue
			}
			factor := -rows[targetRow][pivotCol]/rows[pivotRow][pivotCol]
			addRowToRowWithFactor(rows, pivotRow, targetRow, factor)
		}


		pivotRow++
		pivotCol++
	}

	for i := 0; i < len(rows); i++ {
		var pivot float64
		for j := 0; j < len(rows[i]); j++ {
			if !isZero(rows[i][j]) {
				pivot = rows[i][j]
				break
			}
		}
		divideRowByFactor(rows, i, pivot)
	}

	solutions := make([]float64, len(rows))
	for i, row := range rows {
		solutions[i] = row[len(row)-1]
	}
	return solutions
}
