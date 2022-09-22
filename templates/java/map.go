package java

const mapConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		{{ renderConstants (.Key "key" "Key") }}
		{{ renderConstants (.Elem "value" "Value") }}
{{- end -}}
`

const mapTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.GetMinPairs }}
			build.buf.pgv.MapValidation.min("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinPairs }});
{{- end -}}
{{- if $r.GetMaxPairs }}
			build.buf.pgv.MapValidation.max("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxPairs }});
{{- end -}}
{{- if $r.GetNoSparse }}
			build.buf.pgv.MapValidation.noSparse("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
			build.buf.pgv.MapValidation.validateParts({{ accessor . }}.keySet(), key -> {
				{{ render (.Key "key" "Key") }}
			});
			build.buf.pgv.MapValidation.validateParts({{ accessor . }}.values(), value -> {
				{{ render (.Elem "value" "Value") }}
			});
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
