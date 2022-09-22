package java

const requiredTpl = `{{ $f := .Field }}
	{{- if .Rules.GetRequired }}
		if ({{ hasAccessor . }}) {
			build.buf.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", {{ accessor . }});
		} else {
			build.buf.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", null);
		};
	{{- end -}}
`
