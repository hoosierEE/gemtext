package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inPre := false
	inList := false
	for scanner.Scan() {
		line := scanner.Text()

		// exit verbatim state?
		if inPre {
			if line == "```" { // TODO - should this check strings.prefix(line, "```") instead?
				fmt.Println("</pre>")
				inPre = false
			} else {
				fmt.Println(line) // verbatim
			}
			continue // skip any further processing
		}

		// exit unordered list state?
		if inList {
			if !strings.HasPrefix(line, "* ") {
				inList = false
				fmt.Println("</ul>")
			}
		}

		if strings.Trim(line, " \t") == "" {  // treat whitespace as empty line
			fmt.Println("<br>")
		} else if strings.HasPrefix(line, "### ") {
			fmt.Printf("<h3>%s</h3>\n", strings.Trim(line[3:], " \t"))
		} else if strings.HasPrefix(line, "## ") {
			fmt.Printf("<h2>%s</h2>\n", strings.Trim(line[2:], " \t"))
		} else if strings.HasPrefix(line, "# ") {
			fmt.Printf("<h1>%s</h1>\n", strings.Trim(line[1:], " \t"))
		} else if strings.HasPrefix(line, "=>"){
			url, text, _ := strings.Cut(strings.Trim(line[2:], " \t"), " ")
			fmt.Printf("<a href=\"%s\">%s</a>\n", url, text)
		} else if strings.HasPrefix(line, ">") {
			fmt.Printf("<blockquote>%s</blockquote>\n", line[1:])
		} else if strings.HasPrefix(line, "* ") { // start or continue a list
			if !inList {
				inList = true
				fmt.Println("<ul>")
			}
			fmt.Printf("\t<li>%s</li>\n", strings.Trim(line[2:], " \t"))
		} else if strings.HasPrefix(line, "```") {  // ignore metadata after the 3rd "`" character
			inPre = true
			fmt.Println("<pre>")
		} else {
			fmt.Printf("<p>%s</p>\n", line)
		}
	}

	// close anything that's still open
	if inList {
		fmt.Println("</ul>")
	}
	if inPre {
		fmt.Println("</pre>")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
