package recursion

type Point struct {
	x, y int
}

var dir = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func deleteLastElement(slice []Point) []Point {
	slice = slice[:len(slice)-1]
	return slice
}

func walk(maze []string, wall string, curr, end Point, seen [][]bool, path *[]Point) bool {
	if curr.x < 0 || curr.y >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	if seen[curr.y][curr.x] {
		return false
	}

	if curr.y == end.y && curr.x == end.x {
		*path = append(*path, curr)
		return true
	}

	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for i := 0; i < len(dir); i++ {
		x := dir[i][0]
		y := dir[i][1]
		var newPoint = Point{x: curr.x + x, y: curr.y + y}
		if walk(maze, wall, newPoint, end, seen, path) {
			return true
		}
	}

	// Remove the last element: pop
	*path = append((*path)[:len(*path)-1], (*path)[len(*path):]...)

	return false
}

func maze_solver(maze []string, wall string, start, end Point) []Point {
	var seen = make([][]bool, len(maze))
	path := []Point{}
	walk(maze, wall, start, end, seen, &path)
	return path
}
