# imgutils-sdk

[![Go Reference](https://pkg.go.dev/badge/github.com/imgutils-org/imgutils-sdk.svg)](https://pkg.go.dev/github.com/imgutils-org/imgutils-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/imgutils-org/imgutils-sdk)](https://goreportcard.com/report/github.com/imgutils-org/imgutils-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The unified SDK for the imgutils image manipulation library collection. Part of the [imgutils](https://github.com/imgutils-org) organization.

## Overview

imgutils-sdk provides a convenient wrapper around Go's image types and serves as the central hub for the imgutils ecosystem. While you can use individual imgutils packages directly, the SDK provides a unified Image type for fluent operations.

## Installation

```bash
go get github.com/imgutils-org/imgutils-sdk
```

## Quick Start

```go
package main

import (
    "log"

    imgutils "github.com/imgutils-org/imgutils-sdk"
)

func main() {
    // Open an image
    img, err := imgutils.Open("photo.jpg")
    if err != nil {
        log.Fatal(err)
    }

    // Save as different format
    err = img.SavePNG("photo.png")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Usage Examples

### Basic Operations

```go
// Open an image
img, err := imgutils.Open("input.jpg")
if err != nil {
    log.Fatal(err)
}

// Get dimensions
width := img.Width()
height := img.Height()
bounds := img.Bounds()

// Access underlying image.Image
rawImage := img.Image()

// Clone the image
copy := img.Clone()
```

### Saving Images

```go
// Save as JPEG with quality
err := img.SaveJPEG("output.jpg", 85)

// Save as PNG
err := img.SavePNG("output.png")

// Encode to writer
var buf bytes.Buffer
err := img.EncodeJPEG(&buf, 90)
err := img.EncodePNG(&buf)
```

### Working with image.Image

```go
// Create from existing image.Image
rawImg, _, _ := image.Decode(reader)
img := imgutils.New(rawImg)

// Use with other imgutils packages
import "github.com/imgutils-org/imgutils-resize"

resized := resize.ResizeToWidth(img.Image(), 800)
result := imgutils.New(resized)
result.SaveJPEG("resized.jpg", 85)
```

## The imgutils Ecosystem

The SDK is part of a collection of focused image manipulation packages:

| Package | Description |
|---------|-------------|
| [imgutils-thumbnail](https://github.com/imgutils-org/imgutils-thumbnail) | Thumbnail generation |
| [imgutils-resize](https://github.com/imgutils-org/imgutils-resize) | Image resizing with interpolation |
| [imgutils-crop](https://github.com/imgutils-org/imgutils-crop) | Image cropping |
| [imgutils-rotate](https://github.com/imgutils-org/imgutils-rotate) | Rotation and flipping |
| [imgutils-convert](https://github.com/imgutils-org/imgutils-convert) | Format conversion |
| [imgutils-compress](https://github.com/imgutils-org/imgutils-compress) | Image compression |
| [imgutils-filter](https://github.com/imgutils-org/imgutils-filter) | Filters and effects |
| [imgutils-watermark](https://github.com/imgutils-org/imgutils-watermark) | Watermarking |
| [imgutils-merge](https://github.com/imgutils-org/imgutils-merge) | Image merging |
| [imgutils-upscale](https://github.com/imgutils-org/imgutils-upscale) | Image upscaling |

### Using Multiple Packages

```go
import (
    imgutils "github.com/imgutils-org/imgutils-sdk"
    "github.com/imgutils-org/imgutils-resize"
    "github.com/imgutils-org/imgutils-filter"
    "github.com/imgutils-org/imgutils-watermark"
)

func ProcessImage(inputPath, outputPath, logoPath string) error {
    // Load image
    img, err := imgutils.Open(inputPath)
    if err != nil {
        return err
    }

    // Resize
    resized := resize.ResizeToWidth(img.Image(), 1920)

    // Apply filter
    filtered := filter.Brightness(resized, 5)
    filtered = filter.Contrast(filtered, 10)

    // Load watermark
    logo, err := imgutils.Open(logoPath)
    if err != nil {
        return err
    }

    // Apply watermark
    result := watermark.Apply(filtered, logo.Image(), watermark.Options{
        Position: watermark.BottomRight,
        Opacity:  0.5,
    })

    // Save
    final := imgutils.New(result)
    return final.SaveJPEG(outputPath, 85)
}
```

## API Reference

### Types

#### Image

```go
type Image struct {
    // contains unexported fields
}
```

### Functions

| Function | Description |
|----------|-------------|
| `New(img)` | Create Image from image.Image |
| `Open(path)` | Load image from file |

### Image Methods

| Method | Description |
|--------|-------------|
| `Image()` | Get underlying image.Image |
| `Bounds()` | Get image bounds |
| `Width()` | Get image width |
| `Height()` | Get image height |
| `Clone()` | Create a copy |
| `SaveJPEG(path, quality)` | Save as JPEG file |
| `SavePNG(path)` | Save as PNG file |
| `EncodeJPEG(w, quality)` | Encode to writer as JPEG |
| `EncodePNG(w)` | Encode to writer as PNG |

## Requirements

- Go 1.16 or later

## Philosophy

The imgutils packages follow these principles:

1. **Single Responsibility**: Each package does one thing well
2. **Zero Dependencies**: Only stdlib and golang.org/x/image
3. **Simple API**: Easy to use, hard to misuse
4. **Composable**: Packages work together seamlessly

## License

MIT License - see [LICENSE](LICENSE) for details.
