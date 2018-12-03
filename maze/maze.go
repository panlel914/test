package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
)

var list []image.Point
var already  = make(map[string] image.Point)
var his []image.Point

func main() {
	maze :=readMaze("maze/maze.in")
	for _,r:= range maze{
		for _,v:= range r{
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	check(maze, image.Point{0,0})
	already["0,0"] = image.Point{0,0}
	for {
		if len(list) > 0 {
			ok := check(maze, list[0])
			if ok{
				fmt.Println("success")
				break
			}
		} else{
			break
		}

		list = append(list[:0], list[0+1:]...)
		fmt.Println(his)
	}
}

func add(p image.Point, mx int, my int) (suc bool){
	if _,ok := already[strconv.Itoa(p.X)+","+strconv.Itoa(p.Y)];!ok {
		if p.X == mx && p.Y == my{
			return true
		}
		list = append(list, p)
		his = append(his, p)
	}

	return
}

func addAlready(p image.Point){
	if _,ok := already[strconv.Itoa(p.X)+","+strconv.Itoa(p.Y)];!ok{
		already[strconv.Itoa(p.X)+","+strconv.Itoa(p.Y)] = p
	}
}

func check(maze [][]int, p image.Point)(ok bool){
	mx := len(maze)-1
	my := len(maze[0])-1

	if p.X == mx && p.Y == my {
		return true
	}
	// 上
	if p.X >0 && maze[p.X-1][p.Y] ==0{
		ok = add(image.Point{X:p.X-1, Y:p.Y}, mx, my)
		if ok{
			return ok
		}
	}
	addAlready(image.Point{X:p.X-1, Y:p.Y})
	//右
	if p.Y < my && maze[p.X][p.Y+1] == 0{
		ok = add(image.Point{X:p.X, Y:p.Y+1}, mx, my)
		if ok{
			return ok
		}
	}
	addAlready(image.Point{X:p.X, Y:p.Y+1})
	//下
	if p.X < mx && maze[p.X+1][p.Y] == 0{
		ok = add(image.Point{X:p.X+1, Y:p.Y}, mx, my)
		if ok{
			return ok
		}
	}
	addAlready(image.Point{X:p.X+1, Y:p.Y})
	//左
	if p.Y > 0 && maze[p.X][p.Y-1] == 0{
		ok = add(image.Point{X:p.X, Y:p.Y-1}, mx, my)
		if ok{
			return ok
		}
	}
	addAlready(image.Point{X:p.X, Y:p.Y-1})

	return
}

func readMaze(filename string) [][]int{
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
