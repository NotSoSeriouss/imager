package imager

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
)

type Settings struct {
	MirrorY bool
	MirrorX bool

	Fade uint8

	Seed int64
}

type rgbaColor struct {
	R uint8
	G uint8
	B uint8
	A float32
}

type pixelType int
const (
	body pixelType = iota
	empty
	border

	// Binary types can be either and will
	// get solved randomly by the seed
	bodyBorder
	bodyEmpty
	borderEmpty
)

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, format, err := image.Decode(file)
	fmt.Printf("Image format: %s\n", format)

	return img, err
}

func getPixelType(col color.Color) pixelType {
	r, g, b, a := col.RGBA()
	var yes uint32 = 65535
	if r == yes && g == 0 && b == 0 && a == yes {
		return border
	}
	if r == 0 && g == yes && b == 0 && a == yes {
		return body
	}
	if r == 0 && g == yes && b == yes && a == yes {
		return bodyEmpty
	}
	if r == yes && g == yes && b == 0 && a == yes {
		return bodyBorder
	}
	if r == yes && g == 0 && b == yes && a == yes {
		return borderEmpty
	}
	return empty
}

func getColor(col rgbaColor, rng *rand.Rand) rgbaColor {
	r := col.R
	g := col.G
	b := col.B
	a := col.A
	r += uint8(rng.Float32() * 25)
	g += uint8(rng.Float32() * 25)
	b += uint8(rng.Float32() * 25)
	return rgbaColor{uint8(r), uint8(g), uint8(b), float32(a / 255)}
}

// Generate a matrix of rgba colors with the imager algorithm.
// It takes an image path as input for the template.
func Generate(path string, col rgbaColor, set Settings) [][]rgbaColor {
	if set.Seed == 1 {
		set.Seed = rand.Int63()
	}
	rng := rand.New(rand.NewSource(set.Seed))

	img, err := loadImage(path)
	if err != nil {
		panic(err)
	}

	var w, h int
	{
		bounds := img.Bounds()
		w = bounds.Dx()
		h = bounds.Dy()
	}

	// Turn the image into a matrix of types
	var matrix [][]pixelType = make([][]pixelType, w)
	for x := 0; x < w; x++ {
		matrix[x] = make([]pixelType, h)
		for y := 0; y < h; y++ {
			matrix[x][y] = getPixelType(img.At(x, y))
		}
	}

	// Resolve the binary types with randomness
	var matCopy [][]pixelType = make([][]pixelType, w)
	for x := 0; x < w; x++ {
		matCopy[x] = make([]pixelType, h)
		for y := 0; y < h; y++ {
			switch(matrix[x][y]) {
			case bodyBorder:
				if rng.Float32() > .5 {
					matCopy[x][y] = body
				} else {
					matCopy[x][y] = border
				}
				break
			case bodyEmpty:
				if rng.Float32() > .7 {
					matCopy[x][y] = body
				} else {
					matCopy[x][y] = empty
				}
				break
			case borderEmpty:
				if rng.Float32() > .5 {
					matCopy[x][y] = border
				} else {
					matCopy[x][y] = empty
				}
				break
			default:
				matCopy[x][y] = matrix[x][y]
			}
		}
	}

	// Insert missing borders
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if (x != 0 && x != w) && (y != 0 && y != h) && matCopy[x][y] == body {
				if matCopy[x-1][y] == empty {
					matCopy[x-1][y] = border
				}
				if matCopy[x+1][y] == empty {
					matCopy[x+1][y] = border
				}
				if matCopy[x][y-1] == empty {
					matCopy[x][y-1] = border
				}
				if matCopy[x][y+1] == empty {
					matCopy[x][y+1] = border
				}
			}
		}
	}

	// Turn pixels into static rgb colors
	var finalMat [][]rgbaColor = make([][]rgbaColor, w)
	for x := 0; x < w; x++ {
		finalMat[x] = make([]rgbaColor, h)
		for y := 0; y < h; y++ {
			if matCopy[x][y] == body {
				color := getColor(col, rng)
				if set.Fade != 0 {
					color.R = uint8(max(0, float32(color.R) - ((float32(y) / float32(h)) * float32(set.Fade))))
					color.G = uint8(max(0, float32(color.G) - ((float32(y) / float32(h)) * float32(set.Fade))))
					color.B = uint8(max(0, float32(color.B) - ((float32(y) / float32(h)) * float32(set.Fade))))
				}
					finalMat[x][y] = color
			}
			if matCopy[x][y] == border {
				color := getColor(col, rng)
				finalMat[x][y] = rgbaColor{min(color.R - 128, 0), min(color.G - 128, 0), min(color.B - 128, 0), color.A}
			}
			if matCopy[x][y] == empty {
				finalMat[x][y] = rgbaColor{255, 255, 255, 0}
			}
		}
	}

	if set.MirrorY {
		if (w % 2) != 0 {
			fmt.Println("Image cannot be mirrored because the height is odd")
		}
		for x := 0; x < w / 2; x++ {
			for y := 0; y < h; y++ {
				finalMat[w - x - 1][y] = finalMat[x][y]
			}
		}
	}

	if set.MirrorX {
		if (h % 2) != 0 {
			fmt.Println("Image cannot be mirrored because the width is odd")
		}
		for x := 0; x < w; x++ {
			for y := 0; y < h / 2; y++ {
				finalMat[x][h - y - 1] = finalMat[x][y]
			}
		}
	}

	return finalMat
}

func Rcolor(alpha float32) rgbaColor {
	return rgbaColor{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), alpha}
}

func Ncolor(r, g, b uint8, alpha float32) rgbaColor {
	return rgbaColor{R: r, G: g, B: b, A: alpha}
}
