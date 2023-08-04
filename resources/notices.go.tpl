package {{ .Options.OutPackage }}

// spellchecker:disable

// ===========================================================================================================
// This file was generated automatically at {{ fdate .Time "02-01-2006 15:04:05" }} using gogenlicense.
// Do not edit manually, as changes may be overwritten.
// ===========================================================================================================

{{ prefix .DocString "// " }}
var {{ .Options.DeclarationName }} string

func init() {
	{{ .Options.DeclarationName }} = {{ .NoticeString | printf "%q" }}
}
