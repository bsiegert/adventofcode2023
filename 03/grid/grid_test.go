package grid

import (
	"image"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	_ "embed"
)

func TestReadGrid(t *testing.T) {
	input := strings.NewReader("....\n....\n....")
	grid, err := ReadGrid(input)
	if err != nil {
		t.Fatal(err)
	}
	wantRect := image.Rect(0, 0, 4, 3)
	if grid.Rect != wantRect {
		t.Errorf("unexpected Rect; got %v want %v", grid.Rect, wantRect)
	}
	t.Logf("\n%s", grid)
}

func TestControlCharacter(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{".....", false},
		{".123.", false},
		{"....\n...*", true},
	}
	for i, tc := range tests {
		g, err := ReadGrid(strings.NewReader(tc.input))
		if err != nil {
			t.Fatal(err)
		}
		if got := g.HasControlCharacter(); got != tc.want {
			t.Errorf("[%d] HasControlCharacter = %t, want %t\n%s", i, got, tc.want, g)
		}
	}
}

func TestFindNextNumber(t *testing.T) {
	tests := []struct {
		input    string
		wantOK   bool
		wantRect image.Rectangle
		wantNum  int
	}{
		{
			input:  ".....",
			wantOK: false,
		}, {
			input:    ".123.",
			wantOK:   true,
			wantRect: image.Rect(1, 0, 4, 1),
			wantNum:  123,
		}, {
			input:    ".....\n.123.",
			wantOK:   true,
			wantRect: image.Rect(1, 1, 4, 2),
			wantNum:  123,
		}, {
			input:  "....\n...*",
			wantOK: false,
		},
	}
	for i, tc := range tests {
		g, err := ReadGrid(strings.NewReader(tc.input))
		if err != nil {
			t.Fatal(err)
		}
		num, rect, ok := g.FindNextNumber(image.Pt(0, 0))
		if ok != tc.wantOK {
			t.Errorf("[%d] FindNextNumber: ok = %t, want %t", i, ok, tc.wantOK)
		}
		if !ok {
			continue
		}
		if num != tc.wantNum {
			t.Errorf("[%d] FindNextNumber: num = %v, want %v", i, num, tc.wantNum)
		}
		if rect != tc.wantRect {
			t.Errorf("[%d] FindNextNumber: rect = %v, want %v", i, rect, tc.wantRect)
		}
	}
}

func TestHasControlCharacterNearby(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"..1.\n...*", true},
		{".123.", false},
	}
	for i, tc := range tests {
		g, err := ReadGrid(strings.NewReader(tc.input))
		if err != nil {
			t.Fatal(err)
		}
		_, rect, ok := g.FindNextNumber(image.Pt(0, 0))
		if !ok {
			t.Fatalf("[%d] grid contains no number\n%s", i, g)
		}
		if got := g.HasControlCharacterNearby(rect); got != tc.want {
			t.Errorf("[%d] HasControlCharacterNearby = %t, want %t\n%s", i, got, tc.want, g)
		}
	}
}

//go:embed testdata.txt
var testdata string

func TestNumbersWithCC(t *testing.T) {
	g, err := ReadGrid(strings.NewReader(testdata))
	if err != nil {
		t.Fatal(err)
	}
	got, want := Sum(NumbersWithCC(g)), 4361
	if got != want {
		t.Errorf("unexpected sum value; got %v want %v", got, want)
	}
}

//go:embed testdata2.txt
var testdata2 string

func TestNumbersWithCC_Edge(t *testing.T) {
	g, err := ReadGrid(strings.NewReader(testdata2))
	if err != nil {
		t.Fatal(err)
	}
	got := NumbersWithCC(g)
	want := []int{346, 12}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected diff:\n%v", diff)
	}
}
