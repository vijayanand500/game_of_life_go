package main

import (
	// "fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
	"log"
)

const scale=2
const WIDTH=160
const HEIGHT=120
var black color.RGBA=color.RGBA{75,139,190,255}//95,95,95
var white color.RGBA=color.RGBA{255,232,115,255}//233,233,233
var grid [WIDTH][HEIGHT]int=[WIDTH][HEIGHT]int{}
var buffer [WIDTH][HEIGHT]int=[WIDTH][HEIGHT]int{}
var counter int=0

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	var neighbors [][]int = [][]int{{-1, -1}, {-1, 0}, {-1, 1},
                          { 0, -1}, { 0, 1},
                          { 1, -1}, { 1, 0}, { 1, 1},}
        
	for row:= 0; row< WIDTH; row++{
		for col:= 0; col< HEIGHT; col++{
			live_neighbors := 0
			for _, neighbor := range neighbors{
				r := row + neighbor[0]
				c := col + neighbor[1]
				
				if (r < len(grid) && r >= 0) && (c < len(grid[0]) && c >= 0) && Abs(grid[r][c]) == 1{
					live_neighbors += 1
				}
				
			}
			if grid[row][col] == 1 && (live_neighbors < 2 || live_neighbors >3){
				grid[row][col] = -1
			}
			if grid[row][col] == 0 && (live_neighbors == 3){
				grid[row][col] = 2
			}
		}
	}
    for row:= 0; row < WIDTH; row++{
        for col:= 0; col < HEIGHT; col++ {
            if grid[row][col] > 0{
               grid[row][col] = 1 
            } else{
                grid[row][col] = 0
            }
        }
	}
	return nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(white)
	for x:=0;x<WIDTH;x++ {
		for y:=0;y<HEIGHT;y++ {
			if grid[x][y]>0{
				for x1:=0;x1<scale;x1++ {
					for y1:=0;y1<scale;y1++ {
						screen.Set((x*scale)+x1,(y*scale)+y1,black)
					}
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	for x:=1;x<WIDTH-1;x++ {
		for y:=1;y<HEIGHT-1;y++ {
			if(rand.Float32()<0.4){
				grid[x][y]=1
			}
		}
	}
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}