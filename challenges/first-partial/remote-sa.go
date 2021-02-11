package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// Point is a structure that represents a point on a 2D plane
type Point struct {
	X, Y float64
}

func main() {
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	// shoelace algorithm
	/*
			poly.append(poly[0])
		    sum1, sum2 = 0, 0
		    for i in range(len(poly) - 1):
		        sum1 += poly[i].x * poly[i + 1].y
		        sum2 += poly[i].y * poly[i + 1].x
		    return fabs(sum1 - sum2) / 2.0
	*/
	return 0.0
}

// Distance between two Points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
	// distance sqrt((point1.x - point2.x) ** 2 + (point1.y - point2.y) ** 2)
	// return math.Sqrt(math.Pow(p.X-q.X, 2) + math.Pow(p.Y-q.Y, 2))
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	// sum distance between first and last
	points = append(points, points[0])
	perimeter := 0.0
	// for each point
	for i := 0; i < len(points)-1; i++ {
		// distance between actual and next (until penultimate)
		distance := Distance(points[i], points[i+1])
		// sum all distances
		perimeter += distance
	}
	return perimeter
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
	response += fmt.Sprintf(" - Area            : %v\n", area)

	// Send response to client
	fmt.Fprintf(w, response)
}
