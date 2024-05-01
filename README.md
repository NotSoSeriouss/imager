# Imager

Imager is a Go library for generating image matrices based on a given template image. It allows you to create pixel matrices with customizable settings such as mirroring, fading, and seeding for randomness.

## Features

- Generate procedurally pixel matrices from template images.
- Customize mirroring along the X and Y axes.
- Seed for deterministic randomness.
- Adjustable dark gradient

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
		Fade:    50, // At 50, the lowest pixel in the y
                             // axys will have removed 50 pt from its brightness
		Seed:    12345, // Use 1 instead to have a random seed
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
(Here the image is in half because this template is meant to be mirrored)

To create a template image:

- Use PNG or JPEG format.
- Use the following color scheme:
  - Red (255,0,0): Border
  - Green (0,255,0): Body
  - Yellow (255,255,0): Border/Body
  - Magenta (255,0,255): Border/Empty
  - Cyan (0,255,255): Empty/Body
  - Anything else: Empty
### Example Results
![2024-05-01-021708_3000x2000_scrot](https://github.com/NotSoSeriouss/imager/assets/77798806/ef162374-2210-4381-8aed-b58d6489c81a)
![2024-05-01-021703_3000x2000_scrot](https://github.com/NotSoSeriouss/imager/assets/77798806/9d85c119-225f-42a4-9e24-0f2f9e334ff8)
![2024-05-01-021700_3000x2000_scrot](https://github.com/NotSoSeriouss/imager/assets/77798806/fc566ed8-9141-4d2c-bd9e-5915c97ec86e)
![2024-05-01-021659_3000x2000_scrot](https://github.com/NotSoSeriouss/imager/assets/77798806/e1a4f0ff-78d0-47ea-8157-31659914f1e0)
![2024-05-01-021656_3000x2000_scrot](https://github.com/NotSoSeriouss/imager/assets/77798806/06c4e966-8f53-41a5-85c9-a7e3e56e7083)


## Contributing

Contributions are welcome! Fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
