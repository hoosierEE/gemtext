package g2h

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func GemToHtml(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	inPre := false
	inList := false
	for scanner.Scan() {
		line := scanner.Text()

		// exit verbatim state?
		if inPre {
			if line == "```" { // TODO - should this check strings.prefix(line, "```") instead?
				fmt.Fprintln(out, "</pre>")
				inPre = false
			} else {
				fmt.Fprintln(out, line) // verbatim
			}
			continue // skip any further processing
		}

		// exit unordered list state?
		if inList {
			if !strings.HasPrefix(line, "* ") {
				inList = false
				fmt.Fprintln(out, "</ul>")
			}
		}

		if strings.Trim(line, " \t") == "" {  // treat whitespace as empty line
			fmt.Fprintln(out, "<br>")
		} else if strings.HasPrefix(line, "### ") {
			fmt.Fprintf(out, "<h3>%s</h3>\n", strings.Trim(line[3:], " \t"))
		} else if strings.HasPrefix(line, "## ") {
			fmt.Fprintf(out, "<h2>%s</h2>\n", strings.Trim(line[2:], " \t"))
		} else if strings.HasPrefix(line, "# ") {
			fmt.Fprintf(out, "<h1>%s</h1>\n", strings.Trim(line[1:], " \t"))
		} else if strings.HasPrefix(line, "=>"){
			url, text, _ := strings.Cut(strings.Trim(line[2:], " \t"), " ")
			fmt.Fprintf(out, "<a href=\"%s\">%s</a>\n", url, text)
		} else if strings.HasPrefix(line, ">") {
			fmt.Fprintf(out, "<blockquote>%s</blockquote>\n", line[1:])
		} else if strings.HasPrefix(line, "* ") { // start or continue a list
			if !inList {
				inList = true
				fmt.Fprintln(out, "<ul>")
			}
			fmt.Fprintf(out, "\t<li>%s</li>\n", strings.Trim(line[2:], " \t"))
		} else if strings.HasPrefix(line, "```") {  // ignore metadata after the 3rd "`" character
			inPre = true
			fmt.Fprintln(out, "<pre>")
		} else {
			fmt.Fprintf(out, "<p>%s</p>\n", line)
		}
	}

	// close anything that's still open
	if inList {
		fmt.Fprintln(out, "</ul>")
	}
	if inPre {
		fmt.Fprintln(out, "</pre>")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(out, "Error:", err)
	}
}
