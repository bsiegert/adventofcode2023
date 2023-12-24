package grid

import (
	"image"
	"strconv"
	"unicode"
)

func (g *Grid) HasControlCharacter() bool {
	for y := g.Rect.Min.Y; y < g.Rect.Max.Y; y++ {
		for x := g.Rect.Min.X; x < g.Rect.Max.X; x++ {
			r := g.At(x, y)
			if !unicode.IsDigit(r) && r != '.' && r != 0 {
				return true
			}
		}
	}
	return false
}

func nextPoint(p image.Point, bbox image.Rectangle) image.Point {
	q := p
	q.X++
	if q.X >= bbox.Max.X {
		q.X = bbox.Min.X
		q.Y++
	}
	return q
}

// FindNextNumber searches for a number (a string of consecutive digits)
// from the point at start. It returns the number, the rectangle conaining
// the digits, and whether a number was found.
func (g *Grid) FindNextNumber(start image.Point) (int, image.Rectangle, bool) {
	p := start
	var num []rune
	for p.In(g.Rect) {
		if r := g.At(p.X, p.Y); unicode.IsDigit(r) {
			// Found the start of a number
			num = append(num, r)
			for x := p.X + 1; x <= g.Rect.Max.X; x++ {
				r := g.At(x, p.Y)
				if !unicode.IsDigit(r) || x == g.Rect.Max.X {
					// Found the end of the number
					n, err := strconv.Atoi(string(num))
					if err != nil {
						panic(err)
					}
					return n, image.Rect(p.X, p.Y, x, p.Y+1), true
				}
				num = append(num, r)
			}
		}
		p = nextPoint(p, g.Rect)
	}
	return 0, image.Rectangle{}, false
}

// HasControlCharacterNearby checks if there is a control character *adjacent*
// to the rectangle r.
func (g *Grid) HasControlCharacterNearby(r image.Rectangle) bool {
	return g.SubImage(r.Inset(-1)).HasControlCharacter()
}

// NumbersWithCC returns all the numbers in g that have a control character nearby.
func NumbersWithCC(g *Grid) []int {
	var (
		nums []int
		p    = g.Rect.Min
	)
	for {
		n, r, ok := g.FindNextNumber(p)
		if !ok {
			break
		}
		if g.HasControlCharacterNearby(r) {
			nums = append(nums, n)
		}
		p = nextPoint(image.Pt(r.Max.X, r.Max.Y-1), g.Rect)
	}
	return nums
}

func Sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
