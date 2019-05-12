package parser

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"regexp"
	"strconv"
)

const (
	alphaPat = "(?:0|1)?(?:.[0-9]+)?"
)

var (
	rgbaPat = regexp.MustCompile(fmt.Sprintf(`rgba\(\s*(%s)\s*,\s*(%[1]s)\s*,\s*(%[1]s)\s*,\s*(%s)\s*\)`, funcParam, alphaPat))
)

func parseRGBA(line string, lnum int) []*models.Colour {
	var colours []*models.Colour
	matches := rgbaPat.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		colour := new(models.Colour)
		colour.ColStart = match[0] + 1
		colour.ColEnd = match[1]
		colour.Lnum = lnum
		r, err := strToDec(line[match[2]:match[3]])
		g, err := strToDec(line[match[4]:match[5]])
		b, err := strToDec(line[match[6]:match[7]])
		alpha, err := strconv.ParseFloat(line[match[8]:match[9]], 64)
		if err != nil {
			continue
		}
		colour.Hex = rgbToHex(setAlpha(r, g, b, alpha))
		colours = append(colours, colour)
	}
	return colours
}
