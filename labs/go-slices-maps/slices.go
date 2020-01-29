package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	im :=make([][]uint8,dy)
	for i:=0;i<dy; i++{
		im[i]=make([]uint8,dx)
	}
	
	for y,_ := range im{
	for x,_:= range im[y]{
		im[y][x]=uint8(x+y)
		}
		}
	return im
}

func main() {
	pic.Show(Pic)
}
