package gogenlicense

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"
	"time"
	"unicode"
)

// cspell:words gocode asheading gogenlicense godoc

// generate generates gocode with the provided library and options.
func generate(libraries []Library, opts Options) (string, error) {
	fmtContext := struct {
		Options

		Libraries []Library
		Time      time.Time

		NoticeString string
		DocString    string
	}{
		Libraries: libraries,
		Options:   opts,

		Time: time.Now(),
	}

	var buffer bytes.Buffer
	if err := templates.ExecuteTemplate(&buffer, "notices.txt.tpl", fmtContext); err != nil {
		return "", fmt.Errorf("failed to execute 'notices.txt.tpl': %w", err)
	}
	fmtContext.NoticeString = buffer.String()

	buffer.Reset()
	if err := templates.ExecuteTemplate(&buffer, "doc.txt.tpl", fmtContext); err != nil {
		return "", fmt.Errorf("failed to execute 'doc.txt.tpl': %w", err)
	}
	fmtContext.DocString = buffer.String()

	buffer.Reset()
	if err := templates.ExecuteTemplate(&buffer, "notices.go.tpl", fmtContext); err != nil {
		return "", fmt.Errorf("failed to execute 'notices.go.tpl': %w", err)
	}

	return buffer.String(), nil
}

//go:embed resources/*.tpl
var templateFS embed.FS
var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"fdate": func(t time.Time, layout string) (string, error) {
		return t.UTC().Format(layout), nil
	},

	// prefix adds a prefix to every line of a string.
	// prefix removes trailing newlines.
	"prefix": func(s, prefix string) (string, error) {
		buffer := &bytes.Buffer{}
		for _, l := range strings.Split(s, "\n") {
			buffer.WriteString(prefix)
			buffer.WriteString(l)
			buffer.WriteString("\n")
		}
		return strings.TrimRight(buffer.String(), "\n"), nil
	},

	// asheading turns a string into a heading for godoc
	"asheading": func(s string) (string, error) {
		buffer := &bytes.Buffer{}

		lastValid := true
		for _, r := range s {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				if !lastValid {
					continue
				}
				lastValid = false
				buffer.WriteRune(' ')
				continue
			}
			lastValid = true
			buffer.WriteRune(r)
		}

		return buffer.String(), nil
	},
}).ParseFS(templateFS, "resources/*.tpl"))
