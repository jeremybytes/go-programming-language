// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // numer of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			generateSVG(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "file" {
		gf, _ := os.Create("out.svg")
		generateSVG(gf)
		gf.Close()
		return
	}
	generateSVG(os.Stdout)
}

func generateSVG(out io.Writer) {
	output := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			z := (az + bz + cz + dz) / 4.0

			red := uint8(0xFF)
			green := uint8(0xFF)
			blue := uint8(0xFF)
			switch {
			case z > float64(0.0):
				green = uint8(0xFF - (0xFF * z))
				blue = uint8(0xFF - (0xFF * z))
			case z < float64(0.0):
				red = uint8(0xFF - (0xFF * -z))
				green = uint8(0xFF - (0xFF * -z))
			}

			output += fmt.Sprintf("<polygon style='fill: #%02X%02X%02X' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				red, green, blue, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	output += fmt.Sprintln("</svg>")
	fmt.Fprint(out, output)
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
