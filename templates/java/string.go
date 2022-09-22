package java

const stringConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor . }}[] {{ constantName . "In" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.In -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.NotIn -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.Pattern }}
		com.google.re2j.Pattern {{ constantName . "Pattern" }} = com.google.re2j.Pattern.compile({{ javaStringEscape $r.GetPattern }});
{{- end -}}`

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.Const }}
			build.buf.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetConst }}");
{{- end -}}
{{- if $r.In }}
			build.buf.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			build.buf.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
{{- if $r.Len }}
			build.buf.pgv.StringValidation.length("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			build.buf.pgv.StringValidation.minLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			build.buf.pgv.StringValidation.maxLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.LenBytes }}
			build.buf.pgv.StringValidation.lenBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLenBytes }});
{{- end -}}
{{- if $r.MinBytes }}
			build.buf.pgv.StringValidation.minBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinBytes }});
{{- end -}}
{{- if $r.MaxBytes }}
			build.buf.pgv.StringValidation.maxBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxBytes }});
{{- end -}}
{{- if $r.Pattern }}
			build.buf.pgv.StringValidation.pattern("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Pattern" }});
{{- end -}}
{{- if $r.Prefix }}
			build.buf.pgv.StringValidation.prefix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetPrefix }}");
{{- end -}}
{{- if $r.Contains }}
			build.buf.pgv.StringValidation.contains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetContains }}");
{{- end -}}
{{- if $r.NotContains }}
			build.buf.pgv.StringValidation.notContains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetNotContains }}");
{{- end -}}
{{- if $r.Suffix }}
			build.buf.pgv.StringValidation.suffix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetSuffix }}");
{{- end -}}
{{- if $r.GetEmail }}
			build.buf.pgv.StringValidation.email("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetAddress }}
			build.buf.pgv.StringValidation.address("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetHostname }}
			build.buf.pgv.StringValidation.hostName("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIp }}
			build.buf.pgv.StringValidation.ip("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv4 }}
			build.buf.pgv.StringValidation.ipv4("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv6 }}
			build.buf.pgv.StringValidation.ipv6("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUri }}
			build.buf.pgv.StringValidation.uri("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUriRef }}
			build.buf.pgv.StringValidation.uriRef("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUuid }}
			build.buf.pgv.StringValidation.uuid("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
