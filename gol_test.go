package main

import "testing"

func TestGameOfLifeChange(t *testing.T) {
    var inpGOL [][]int = [][]int{{0,0,0},{1,1,1},{0,0,0}}
    var expOutput [][]int = [][]int{{0,1,0},{0,1,0},{0,1,0}}
    outGOL := gameOfLife(inpGOL)
    
    if len(expOutput) != len(outGOL) {
       t.Errorf("Length of the output is not correct")
    }
    
    for row := range(outGOL) {
      for column := range(outGOL[0]) {
        if outGOL[row][column] != expOutput[row][column] {
          t.Errorf("Output of GOL is not as expected at row: %d and column %d", row, column)
         }
       }
    } 
}

func TestGameOfLifeNoChange(t *testing.T) {
    var inpGOL [][]int = [][]int{{1,1},{1,1}}
    var expOutput [][]int = [][]int{{1,1},{1,1}}
    outGOL := gameOfLife(inpGOL)
    
    if len(expOutput) != len(outGOL) {
       t.Errorf("Length of the output is not correct")
    }
    
    for row := range(outGOL) {
      for column := range(outGOL[0]) {
        if outGOL[row][column] != expOutput[row][column] {
          t.Errorf("Output of GOL is not as expected at row: %d and column %d", row, column)
         }
       }
    } 
}
