# Imager

Imager is a Go library for generating image matrices based on a given template image. It allows you to create pixel matrices with customizable settings such as mirroring, fading, and seeding for randomness.

## Features

- Generate procedurally pixel matrices from template images.
- Customize mirroring along the X and Y axes.
- Seed for deterministic randomness.

## Installation

To install Imager, use `go get`:

```bash
go get github.com/NotSoSeriouss/imager
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/NotSoSeriouss/imager"
)

func main() {
	// Define settings
	settings := imager.Settings{
		MirrorY: true,
		MirrorX: false,
		Fade:    0.5,
		Seed:    12345,
	}

	// Define color
	color := imager.Rcolor(1.0) // Random color with full opacity

	// Generate pixel matrix
	pixelMatrix := imager.Generate("path/to/template/image.png", color, settings)

	// Do something with pixelMatrix...
	fmt.Println(pixelMatrix)
}
```

## Template Image

![template](https://imgur.com/EVb2SZp.png)
`Here the image is in half because this template is meant to be merrored`

To create a template image:

- Use PNG or JPEG format.
- Use the following color scheme:
  - Red (255,0,0): Border
  - Green (0,255,0): Body
  - Yellow (255,255,0): Border/Body
  - Magenta (255,0,255): Border/Empty
  - Cyan (0,255,255): Empty/Body
  - Anything else: Empty

## Contributing

Contributions are welcome! Fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
