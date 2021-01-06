{{ if gt (len .Args) 0 }}
{{ $x := (cslice "test") }}
{{ $y := (index .CmdArgs 0) }}
{{ if in $x $y}}
{{deleteTrigger 3}}
{{else}}
{{end}}
{{end}}