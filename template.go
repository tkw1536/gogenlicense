package gogenlicense

import (
	"bytes"
	"strings"
	"text/template"
	"time"
	"unicode"

	"github.com/pkg/errors"
)

// generate generates gocode with the provided library and options.
func generate(libraries []Library, opts Options) (string, error) {
	tplCtx := tplContext{
		Libraries: libraries,
		Options:   opts,

		Time: time.Now(),
	}

	notice := &bytes.Buffer{}
	if err := noticeTpl.Execute(notice, noticeContext(tplCtx)); err != nil {
		return "", errors.Wrap(err, "Unable to execute noticeString")
	}

	doc := &bytes.Buffer{}
	if err := docTpl.Execute(doc, docContext(tplCtx)); err != nil {
		return "", errors.Wrap(err, "Unable to execute noticeComment")
	}

	master := &bytes.Buffer{}
	masterCtx := &masterContext{
		tplContext: tplCtx,

		DocString:    doc.String(),
		NoticeString: notice.String(),
	}
	if err := masterTemplate.Execute(master, masterCtx); err != nil {
		return "", errors.Wrap(err, "Unable to execute master")
	}

	return master.String(), nil
}

var masterTemplate = maketemplate("template.go", `
package {{ .Options.OutPackage }}

// ===========================================================================================================
// This file was generated automatically at {{ fdate .Time "02-01-2006 15:04:05" }} using gogenlicense.
// Do not edit manually, as changes may be overwritten.
// ===========================================================================================================

{{ prefix .DocString "// " }}
var {{ .Options.DeclarationName }} string

func init() {
	{{ .Options.DeclarationName }} = {{ .NoticeString | printf "%q" }}
}
`)

// masterContext is the context for masterTemplate
type masterContext struct {
	tplContext

	NoticeString string
	DocString    string
}

// docTpl is the template that will show up in godoc.
var docTpl = maketemplate("doc",
	`{{ .Options.DeclarationName }} contains legal and license information of external software included in this program. 
These notices consist of a list of dependencies along with their license information.
This string is intended to be displayed to the enduser on demand.

Even though the value of this variable is fixed at compile time it is omitted from this documentation.
Instead the list of go modules, along with their licenses, is listed below. 

{{ range $index, $library := .Libraries }}
{{ $library.Title | asheading }}

The {{ $library.Title }} is licensed under the Terms of the {{ $library.LicenseName }} License. 
{{if $library.LibraryURL }}See also {{$library.LibraryURL}}. {{ end }}

{{ prefix $library.LicenseText " " }}
{{end}}

Generation

This variable and the associated documentation have been automatically generated using the 'gogenlicense' tool. 
It was last updated at {{ fdate .Time "02-01-2006 15:04:05" }}. 
`)

// docContext is the context for docTpl
type docContext tplContext

// noticeTpl is the template that will be displayed to the user
var noticeTpl = maketemplate("notice",
	`The following go packages are imported:
{{ range $index, $library := .Libraries }}- {{ if $library.Path }}{{$library.Path}}{{ else }}{{ $library.Title}}{{end}} ({{ $library.LicenseName }}{{if $library.LibraryURL }}; see {{$library.LibraryURL}}{{end}})
{{end}}
================================================================================

{{ range $index, $library := .Libraries }}
================================================================================
{{$library.Title}}
Licensed under the Terms of the {{ $library.LicenseName }} License{{if $library.LibraryURL }}, see also {{$library.LibraryURL}}. {{end}}

{{ $library.LicenseText }}
================================================================================
{{end}}`)

// noticeContext is the context for noticeTpl
type noticeContext tplContext

// templateFunctions are functions used during templating
var templateFunctions = template.FuncMap{
	"fdate": func(t time.Time, layout string) (string, error) {
		return t.Format(layout), nil
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
}

// maketemplate is a utility function that makes a template fro a string
func maketemplate(name, content string) *template.Template {
	return template.Must(template.New(name).Funcs(templateFunctions).Parse(content))
}

type tplContext struct {
	Options
	Libraries []Library
	Time      time.Time
}
