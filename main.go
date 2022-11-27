package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

type Point struct {
	x int
	y int
	surroundingColour color.RGBA
}

var (
	pointsAmount = 10
	height = 2400
	width = 2400
	minOffset = 10

	RED = color.RGBA{255, 0, 0, 255}
	MAGENTA = color.RGBA{255, 0, 255, 255}
	GREEN = color.RGBA{0, 255, 0, 255}
	CYAN = color.RGBA{0, 255, 255, 255}
	BLUE = color.RGBA{0, 0, 255, 255}
	YELLOW = color.RGBA{255, 255, 0, 255}
	WHITE = color.RGBA{255, 255, 255, 255}
	VIOLET = color.RGBA{51, 0, 154, 255}
	ORANGE = color.RGBA{255, 175, 0, 255}
	BROWN = color.RGBA{70, 63, 23, 255}
	EMPTY_PIXEL = color.RGBA{}

	COLOURS []color.RGBA = []color.RGBA{RED, MAGENTA, GREEN, CYAN, BLUE, YELLOW, WHITE, VIOLET, ORANGE, BROWN}
)

func init() {
	rand.Seed(time.Now().Unix())
}

func getShortestDistanceColour(vectors []Point, point Point) color.RGBA {
	distances := make([]float64, pointsAmount)
	for i := 0; i < len(vectors); i++ {
		distances[i] = math.Sqrt(
			math.Pow(float64(vectors[i].x) - float64(point.x), 2) +
			math.Pow(float64(vectors[i].y) - float64(point.y), 2),
		)
	}

	min := distances[0]
	minIndex := 0
	for x := 1; x < len(distances); x++ {
		if distances[x] < min {
			min = distances[x]
			minIndex = x
		}
	}
	return vectors[minIndex].surroundingColour
}


func paintblueprintImage(img *image.RGBA, vectors []Point) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if img.At(x, y) == EMPTY_PIXEL {
				img.Set(x, y, getShortestDistanceColour(vectors, Point{y: x, x: y}))
			}
		}
	}
}

func generateRandomPoints(points *image.RGBA) []Point {
	var vectors []Point
	for x := 0; x < pointsAmount; x++ {
		h := rand.Intn(height - 2 * minOffset) + minOffset
		w := rand.Intn(width - 2 * minOffset) + minOffset

		points.Set(h, w, color.Black)
		vectors = append(vectors, Point{surroundingColour: COLOURS[x], x: w, y: h})
	}

	return vectors
}

func main() {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{height, width}})
	vectors := generateRandomPoints(img)
	paintblueprintImage(img, vectors)
	
	f, _ := os.Create("voronoi.png")
	defer f.Close()

	png.Encode(f, img)
}