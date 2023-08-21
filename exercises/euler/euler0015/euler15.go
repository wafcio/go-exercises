package main

func Euler15(size int) int {
	grid := prepareGrid(size)

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[size][size]
}

func prepareGrid(size int) [][]int {
	var grid [][]int

	for position := 0; position <= size; position += 1 {
		grid = append(grid, make([]int, size+1))
		grid[position][0] = 1
		grid[0][position] = 1
	}

	return grid
}
