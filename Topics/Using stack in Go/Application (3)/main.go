package main

import (
    "errors"
    "fmt"
)

// first you need to change a stack
// to store the coordinates in []int variable
type XYStack struct {
    storage [][]int
}

func (s *XYStack) Push(value ?) {
    s.storage = append(s.storage, value)
}

func (s *XYStack) Pop() (?, error) {
    last := len(s.storage) - 1
    if last <= -1 {
        return ?, errors.New("stack is empty")
    }

    value := s.storage[last]
    s.storage = s.storage[:last]

    return value, nil
}

func main() {
    // don't worry about the input data
    // the readLayout() function will do it for you (don't touch it!)
    layout, start, finish := readLayout()

    fmt.Println(ifPathExists(layout, start, finish))
}

func ifPathExists(layout [][]int, start []int, finish []int) bool {
    var solver ?

    solver.Push(?)

    for {
        tile, err := solver.Pop() // get the last tile to check

        if err != nil {
            return false // the journey is over
        }

        x := tile[0]
        y := tile[1]

        if x == finish[0] && y == finish[1] {
            return true
        }

        // if the lower path is clear
        if y+1 < len(layout)-1 && layout[y+1][x] != 8 {
            // we're moving down
            solver.Push([]int{x, y + 1}) // add the tile to solver
        }
        // if the upper path is clear
        if y-1 >= 0 && layout[y-1][x] != 8 {
            // we're moving up
            ?
        }
        // if the right path is clear x+1
        ?

        // if the left path is clear x-1
        ?

        // we need to mark the tiles where we were
        layout[y][x] = 8 // let it be the wall
    }
}