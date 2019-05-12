package parser

import (
	"github.com/rrethy/hexokinase/internal/colour"
	"testing"
)

func TestParseRGB(t *testing.T) {
	var tests = []struct {
		line string
		lnum int
		want colours
	}{
		{"rgb(0,0,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 10, Lnum: 0, Hex: "#000000"},
		}},
		{"rgb(176,253,35)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#b0fd23"},
		}},

		// test red value
		{"rgb(35,0,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 11, Lnum: 0, Hex: "#230000"},
		}},
		{"rgb(176,0,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#b00000"},
		}},
		{"rgb(215,0,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#d70000"},
		}},
		{"rgb(253,0,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#fd0000"},
		}},
		{"rgb(255,0,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#ff0000"},
		}},

		// test green value
		{"rgb(0,35,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 11, Lnum: 0, Hex: "#002300"},
		}},
		{"rgb(0,176,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#00b000"},
		}},
		{"rgb(0,215,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#00d700"},
		}},
		{"rgb(0,253,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#00fd00"},
		}},
		{"rgb(0,255,0)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#00ff00"},
		}},

		// test blue value
		{"rgb(0,0,35)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 11, Lnum: 0, Hex: "#000023"},
		}},
		{"rgb(0,0,176)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#0000b0"},
		}},
		{"rgb(0,0,215)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#0000d7"},
		}},
		{"rgb(0,0,253)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#0000fd"},
		}},
		{"rgb(0,0,255)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#0000ff"},
		}},

		// test multiple values
		{"rgb(0,0,255)rgb(176,253,35)  rgb(176,253,35)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 12, Lnum: 0, Hex: "#0000ff"},
			&colour.Colour{ColStart: 13, ColEnd: 27, Lnum: 0, Hex: "#b0fd23"},
			&colour.Colour{ColStart: 30, ColEnd: 44, Lnum: 0, Hex: "#b0fd23"},
		}},

		// tests invalid values
		{"rgb(256,0,0)", 0, []*colour.Colour{}},
		{"rgb(0,0,256)", 0, []*colour.Colour{}},
		{"rgb(0,0,256)", 0, []*colour.Colour{}},
		{"rgb(1000,1000,1000)", 0, []*colour.Colour{}},

		// test handling of whitespace
		{" rgb(0,0,0) ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 11, Lnum: 0, Hex: "#000000"},
		}},
		{" rgb(0,0,0) rgb(0,0,0) ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 11, Lnum: 0, Hex: "#000000"},
			&colour.Colour{ColStart: 13, ColEnd: 22, Lnum: 0, Hex: "#000000"},
		}},
		{"rgb( 0 , 0 , 0 )", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 16, Lnum: 0, Hex: "#000000"},
		}},
		{"rgb(  0  ,  0  ,  0  )", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 22, Lnum: 0, Hex: "#000000"},
		}},
		{"rgb(	0	,	0	,	0	)", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 16, Lnum: 0, Hex: "#000000"},
		}},
	}
	for i, test := range tests {
		if got := parseRGB(test.line, test.lnum); !areSameColours(got, test.want) {
			t.Errorf("test number: %d\n\tgot:    %v\n\twanted: %v", i, got, test.want)
		}
	}
}
