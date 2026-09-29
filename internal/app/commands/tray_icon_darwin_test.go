//go:build darwin && cgo

package commands

import (
	"bytes"
	"image/png"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrayIconsAreValidTemplateImages(t *testing.T) {
	icons := map[string][]byte{
		"play": trayPlayIconPNG,
		"idle": trayIdleIconPNG,
	}

	for name, raw := range icons {
		t.Run(name, func(t *testing.T) {
			require.NotEmpty(t, raw)

			img, err := png.Decode(bytes.NewReader(raw))
			require.NoError(t, err)

			// A template icon carries its shape in the alpha channel: some pixels
			// must be opaque (the glyph) and some fully transparent (the background).
			var opaque, transparent int
			bounds := img.Bounds()
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					_, _, _, a := img.At(x, y).RGBA()
					switch a {
					case 0:
						transparent++
					case 0xFFFF:
						opaque++
					}
				}
			}
			assert.Positive(t, opaque, "glyph shape should have opaque pixels")
			assert.Positive(t, transparent, "background should be transparent")
		})
	}
}
