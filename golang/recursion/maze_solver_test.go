package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type thing struct {
	x, y int
}

func Test_MazeSolver(t *testing.T) {

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}


	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	start := Point{
		x: 10,
		y: 10,
	}
	end := Point{x: 1, y: 5}
	result := maze_solver(maze, "x", start, end);
	r := drawPath(maze, result)
	e := drawPath(maze, mazeResult)

	printMaze(r)
	printMaze(e)
	assert.Equal(t, r, e)
}

func printMaze(maze []string) {
	for _, v := range maze {
		fmt.Printf("%v\n", v)
	}
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func drawPath(data []string, path []Point) []string {
	for i:=0; i<len(path); i++ {
		x := path[i].x
		y:= path[i].y
			data[y] = replaceAtIndex(data[y], []rune("*")[0], x)
	}

	return data
}
