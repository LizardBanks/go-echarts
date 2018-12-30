package goecharts

var headerTpl = `
{{ define "header" }}
<head>
    <meta charset="utf-8">
    <title>{{ .PageTitle }}</title>
{{- range .JSAssets }}
    <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ end }}
`

var routerTpl = `
{{- define "routers" }}
<div class="select" style="margin-left: 10px; margin-top:10px">
{{- if gt .HttpRouters.Len 0}}
    <select onchange="window.location.href=this.options[this.selectedIndex].value">
    {{- range .HttpRouters }}
        <option value="{{ .Url }}">{{ .Text }}</option>
    {{- end }}
    </select>
{{- end -}}
</div>
{{- end }}

`

var baseTpl = `
{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}"
         style="width:{{ .InitOpts.Width }};height:{{ .InitOpts.Height }};"></div>
</div>
<script type="text/javascript">
    let myChart__x__{{ .ChartID }}__x__ = echarts.init(document.getElementById('{{ .ChartID }}'), "{{ .Theme }}");
    let option__x__{{ .ChartID }}__x__ = {
        title: {{ .TitleOpts  }},
        tooltip: {{ .TooltipOpts }},
        legend: {{ .LegendOpts }},
		geo: {{ .GeoOpts }},
	{{- if .ToolboxOpts.Show }}
		toolbox: {{ .ToolboxOpts }},
	{{- end }}
    {{- if gt .DataZoomOptsList.Len 0 }}
        dataZoom:{{ .DataZoomOptsList }},
    {{- end }}
    {{- if gt .VisualMapOptsList.Len 0 }}
        visualMap:{{ .VisualMapOptsList }},
    {{- end }}
    {{- if .HasXYAxis }}
        xAxis: {{ .XAxisOpts }},
        yAxis: {{ .YAxisOpts }},
    {{- end }}
        series: [
        {{ range .Series }}
        {{- . }},
        {{ end -}}
        ],
    {{- if eq .Theme "white" }}
        color: {{ .ColorList }}
    {{- end }}
    };
    myChart__x__{{ .ChartID }}__x__.setOption(option__x__{{ .ChartID }}__x__);
</script>
{{ end }}`

var chartTpl = `
{{- define "chart" }}
<!DOCTYPE html>
<html>
{{- template "header" . }}
<body>
{{- template "routers" . }}
{{- template "base" . }}
<style>
    .container {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .item {
        margin: auto;
    }
</style>
</body>
</html>
{{ end }}
`

var pageTpl = `
{{- define "page" }}
<!DOCTYPE html>
<html>
{{- template "header" . }}
<body>
{{- template "routers" . }}
{{- range .Charts }}
    {{ template "base" .}}
{{- end }}
<style>
    .container {
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .item {
        margin: auto;
    }
</style>
</body>
</html>
{{ end }}
`
