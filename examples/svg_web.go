package main

import (
	"image/color"
	_ "image/jpeg"
	"time"
	"math/rand"
	"net/http"
	"log"
	"github.com/esimov/gospline"
)

func main()  {
	var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	var points [][]float64
	var width, height int = 800, 600

	for i := 0; i < 20; i++ {
		x := randInt(0, width, rng)
		y := randInt(0, height, rng)
		point := []float64{float64(x), float64(y)}
		points = append(points, point)
	}

	svg := &spline.SVG{
		Width: width,
		Height: height,
		Title: "BSpline",
		Lines: []spline.Line{},
		Color: color.NRGBA{R:255,G:0,B:0,A:255},
		Description: "Convert straight lines to curves",
		StrokeWidth: 2,
		StrokeLineCap: "round", //butt, round, square
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg.Draw(w, points, false)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func randInt(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}