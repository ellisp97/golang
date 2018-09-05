package main

import (
	"os"
	"log"
	"fmt"
	"io"
	"math"
	"bytes"
)

type Matrix struct {
	Dimension int
	Data      [][]int
}

func main() {
	matrix, err := createMatrix(os.Stdin)
	checkError(err)

	fmt.Printf("%d\n", matrix.DiagonalDiff())
}

// --------- copied from stack overflow -------
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			fmt.Printf("%d\n", count)
			return count, nil

		case err != nil:
			fmt.Printf("%d\n", count)
			return count, err
		}
	}
}
// ------------------------------------------

func createMatrix(r io.Reader) (*Matrix,error){
	m := new(Matrix)

	// TODO: are Dimensions passed in or not?
	//m.Dimension, _ = lineCounter(r)
	_, err := fmt.Fscanf(r, "%d", &m.Dimension)
	checkError(err)


	m.Data = make([][]int, m.Dimension)
	for i := range m.Data {
		m.Data[i] = make([]int, m.Dimension)
	}

	for y := 0; y < m.Dimension; y++ {
		for x := 0; x < m.Dimension; x++ {
			_, err := fmt.Fscanf(r, "%d", &m.Data[y][x])
			checkError(err)
		}
	}
	return m,nil
}

func (m *Matrix) Diag() int {
	sum := 0
	for i:=0; i < m.Dimension; i++ {
		sum += m.Data[i][i]
	}
	return sum
}

func (m *Matrix) ReverseDiag() int {
	sum := 0
	for i:=0; i < m.Dimension; i++ {
		sum += m.Data[m.Dimension-i-1][i]
	}
	return sum
}


func (m *Matrix) DiagonalDiff() int {
	diff := math.Abs(float64(m.Diag() - m.ReverseDiag()))
	return int(diff)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}