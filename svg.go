package pak

import (
	"fmt"
	"strings"
)

const (
	svgtop = `<?xml version="1.0"?>
<svg`
	svginitfmt = `%s width="%f%s" height="%f%s"`
	svgns      = `
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink"`
	svgnsinkscape = `
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"`
	vbfmt = `viewBox="%f %f %f %f"`

	emptyclose = "/>\n"
)

func svgStart(w float64, h float64, unit string, plain bool) string {
	s := fmt.Sprintf(svginitfmt, svgtop, w, unit, h, unit) + " " +
		fmt.Sprintf(vbfmt, 0.0, 0.0, w, h) + svgns
	if plain == false {
		s += svgnsinkscape
	}
	s += ">\n"
	return s
}

func svgEnd(s string) string {
	return s + "\n</svg>\n"
}

func svgGroupStart(ss ...string) string {
	gs := ""
	for _, s := range ss {
		gs += s + " "
	}
	gs = strings.TrimSpace(gs)
	return fmt.Sprintf("<g %s>", gs)
}

func svgGroupEnd(g string) string {
	return g + "\n</g>\n"
}

func svgRect(x float64, y float64, w float64, h float64, s string) string {
	return fmt.Sprintf(`
<rect x="%f" y="%f" width="%f" height="%f" style="%s" />`, x, y, w, h, s)
}

func svgText(x float64, y float64, txt string, s string) string {
	return fmt.Sprintf(`
<text x="%f" y="%f" style="%s" >
%s
</text>`, x, y, s, txt)
}
