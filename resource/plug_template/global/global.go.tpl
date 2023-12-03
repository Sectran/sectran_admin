package global

{{- if .HasGlobal }}

import "github.com/Sectran/sectran_admin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}