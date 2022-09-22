package java

const boolTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			build.buf.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetConst }});
{{- end }}`
