package table

import (
	"strings"

	"github.com/fatih/color"
	"github.com/giantswarm/columnize"
)

// Colourize makes the given table parts more colourful. It colourizes the
// header of the table by modifying all elements of the first row.
func Colourize(parts [][]string) [][]string {
	var colourized [][]string

	for i, p := range parts {
		var rowParts []string

		for _, rp := range p {
			// In case we deal with the first row we colourize the header parts.
			if i == 0 {
				rp = color.CyanString(rp)
			}
			rowParts = append(rowParts, rp)
		}

		colourized = append(colourized, rowParts)
	}

	return colourized
}

// Format parses the table parts and formats them into a columnized table string
// which is easily printable.
func Format(parts [][]string) string {
	var rows []string

	for _, p := range parts {
		var rowParts []string

		for _, rp := range p {
			rowParts = append(rowParts, rp)
		}

		row := strings.Join(rowParts, " | ")
		rows = append(rows, row)
	}

	formatted := columnize.SimpleFormat(rows)

	return formatted
}
