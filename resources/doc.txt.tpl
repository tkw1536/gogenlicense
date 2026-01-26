{{ .Options.DeclarationName }} contains legal and license information of external software included in this program. 
These notices consist of a list of dependencies along with their license information.
This string is intended to be displayed to the enduser on demand.

The list of go modules, along with their licenses, is listed below. 

{{ range $index, $library := .Libraries }}
{{ $library.Title | asheading }}

The {{ $library.Title }} is licensed under the Terms of the {{ $library.LicenseName }} License. 
{{if $library.LibraryURL }}See also {{$library.LibraryURL}}. {{ end }}

{{ prefix $library.LicenseText " " }}
{{end}}
