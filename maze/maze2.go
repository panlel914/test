package main

import (
	"fmt"
	"os"
)

func readMaze2(filename string) [][]int{
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)

	for i := range maze{
		maze[i] = make([]int, col)
		for j := range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func main() {
	m := readMaze2("maze/maze.in")
	walk(m,point{0,0}, point{5,4})
}

type point struct {
	i int
	j int
}

func (p point)compare(po point) bool{
	if p.i == po.i && p.j == po.j{
		return true
	}

	return false
}

func (p point) add(r point) point{
	return point{i:p.i+r.i,j:p.j+r.j}
}

var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1},
}

func walk(maze [][]int, start, end point){
	steps := make([][]int, len(maze))
	for i := range steps{
		steps[i] = make([]int, len(maze[i]))
	}

	q := []point{start}
	// index := 1
	for len(q) >0 {
		cur := q[0]
		q = q[1:]
		if cur.compare(end){
			break
		}
		for _,v := range dirs{
			next := cur.add(v)
			if next.i>=0 && next.j >=0 && next.i < len(maze) && next.j < len(maze[next.i]) && maze[next.i][next.j] == 0 && !next.compare(start){
				if steps[next.i][next.j] ==0 {
					q = append(q, next)
					st := steps[cur.i][cur.j] + 1
					steps[next.i][next.j] = st
				}
			}
		}
	}

	for _,r:= range steps{
		for _,v:= range r{
			fmt.Printf("%3d ", v)
		}
		fmt.Println()
	}
}
