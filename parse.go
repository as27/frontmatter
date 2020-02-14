// Package frontmatter provides a simple way to extract additional informations out of a text.
//
// It is easy to use it. First you need a delimiter, which defines
// the beginning and the end of the frontmatter section.
// The syntax for adding a value:
//
//   key = some value
//
// You can use every string as key and every string as value.
package frontmatter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Parse extrakts frontmatter information out of a io.Reader. The function
// returns the frontmatter as a simple map[string]string. The body part
// of the text is written into an io.Writer. That makes that
// function useful, for using it inside a handler.
func Parse(delim string, r io.Reader, w io.Writer) (frontmatter map[string]string, err error) {
	frontmatter = make(map[string]string)
	scanner := bufio.NewScanner(r)
	scanFrontmatter := false
	for scanner.Scan() {
		l := scanner.Text()

		if scanFrontmatter {
			if l == delim {
				scanFrontmatter = false
				continue
			}

			key, val := splitLine(l)
			if key != "" {
				frontmatter[key] = val
			}
		} else {
			if l == delim {
				scanFrontmatter = true
				continue
			}
			_, err = w.Write([]byte(l + "\n"))
			if err != nil {
				return frontmatter, fmt.Errorf("Parse: Cannot write into Writer: %w", err)
			}
		}

	}
	return frontmatter, nil
}

func splitLine(l string) (key, val string) {
	el := strings.Split(l, "=")
	switch len(el) {
	case 0:
		return "", ""
	case 1:
		key = strings.Trim(el[0], " ")
	default:
		key = strings.Trim(el[0], " ")
		val = strings.Trim(el[1], " ")
	}
	return key, val
}
