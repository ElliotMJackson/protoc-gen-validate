package java

const repeatedConstTpl = `{{ renderConstants (.Elem "" "") }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.GetMinItems }}
			build.buf.pgv.RepeatedValidation.minItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinItems }});
{{- end -}}
{{- if $r.GetMaxItems }}
			build.buf.pgv.RepeatedValidation.maxItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxItems }});
{{- end -}}
{{- if $r.GetUnique }}
			build.buf.pgv.RepeatedValidation.unique("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end }}
			build.buf.pgv.RepeatedValidation.forEach({{ accessor . }}, item -> {
				{{ render (.Elem "item" "") }}
			});
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
