package grid

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"strings"
)

// Grid represents a 2D grid of runes. It is adapted from the
// image.Gray type in the standard library.
type Grid struct {
	// Pix holds the grid values. The character at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix []rune
	// Stride is the Pix stride (in bytes) between vertically adjacent characters.
	Stride int
	// Rect is the grid's bounds.
	Rect image.Rectangle
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (g *Grid) PixOffset(x, y int) int {
	return (y-g.Rect.Min.Y)*g.Stride + (x-g.Rect.Min.X)*1
}

func (g *Grid) At(x, y int) rune {
	p := image.Pt(x, y)
	if !p.In(g.Rect) {
		return 0
	}
	return g.Pix[g.PixOffset(x, y)]
}

// String prints the part Grid that is delineated by g.Rect.
func (g *Grid) String() string {
	var s strings.Builder
	for y := g.Rect.Min.Y; y < g.Rect.Max.Y; y++ {
		line := g.Pix[g.PixOffset(g.Rect.Min.X, y):g.PixOffset(g.Rect.Max.X, y)]
		fmt.Fprintf(&s, "%s\n", string(line))
	}
	str := s.String()
	return str[:len(str)-1]
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (g *Grid) SubImage(r image.Rectangle) *Grid {
	r = r.Intersect(g.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Grid{}
	}
	i := g.PixOffset(r.Min.X, r.Min.Y)
	return &Grid{
		Pix:    g.Pix[i:],
		Stride: g.Stride,
		Rect:   r,
	}
}

func ReadGrid(r io.Reader) (*Grid, error) {
	var g Grid
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	firstLine := scanner.Text()
	g.Stride = len(firstLine)
	g.Rect.Max.X = g.Stride
	g.Rect.Max.Y++
	g.Pix = []rune(firstLine)
	for scanner.Scan() {
		g.Rect.Max.Y++
		line := scanner.Text()
		if len(line) != g.Stride {
			return nil, fmt.Errorf("wrong input in line %d: got %d chars want %d", g.Rect.Max.Y, len(line), g.Stride)
		}
		g.Pix = append(g.Pix, []rune(line)...)
	}
	return &g, nil
}
