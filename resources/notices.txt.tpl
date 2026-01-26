The following go packages are imported:
{{ range $index, $library := .Libraries }}- {{ if $library.Path }}{{$library.Path}}{{ else }}{{ $library.Title}}{{end}} ({{ $library.LicenseName }}{{if $library.LibraryURL }}; see {{$library.LibraryURL}}{{end}})
{{end}}
================================================================================

{{ range $index, $library := .Libraries }}
================================================================================
{{$library.Title}}
Licensed under the Terms of the {{ $library.LicenseName }} License{{if $library.LibraryURL }}, see also {{$library.LibraryURL}}{{end}}. {{ if $library.NoticeURL }}See also the notice file at {{$library.NoticeURL}}. {{end}}

{{ $library.LicenseText }}{{ if $library.NoticeText }}

{{ $library.NoticeText }}
{{ end }}
================================================================================
{{end}}