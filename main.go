package rottenOrange

// defining directions for rottenOrange
var dir [][]int = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func rottenOrange(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var queue [][]int
	freshCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				// counting total number of fresh oranges
				freshCount++
			}
			if grid[i][j] == 2 {
				// loading rottenOrange to queue for initial time
				queue = append(queue, []int{i, j})
			}
		}
	}
	// no fresh oranges found
	if freshCount == 0 {
		return 0
	}
	time := 0
	// lets say each iteration of queue represents 1 unit of time in the grid
	for len(queue) > 0 {
		for _, node := range queue {
			// check all 4 directions for rottenOrange
			for _, d := range dir {
				i, j := node[0]+d[0], node[1]+d[1]
				if isValidNeighbor(i, j, grid) {
					// add new infected oranges to queue for next unit of time
					queue = append(queue, []int{i, j})
					freshCount--

					// update grid to mark as rottenOrange
					grid[i][j] = 2
				}
			}
			// dequeue the rottenOrange
			queue = queue[1:]
		}
		// increment time
		time++
		//fmt.Println(time)
	}
	// all oranges are rotten after time
	if freshCount > 0 {
		return -1
	}

	// index of time is the number of unit time it takes to rotten all oranges
	return time - 1
}

// helper function to check if the given node is valid
func isValidNeighbor(x, y int, grid [][]int) bool {
	rows, cols := len(grid), len(grid[0])
	if x < 0 || x >= rows || y < 0 || y >= cols {
		return false
	}
	return grid[x][y] == 1
}
