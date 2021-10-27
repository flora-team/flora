package pages

var Page = `# {{.Page.PageDetails}}
*** Settings ***
Resource          ./base.robot
*** Variables ***

*** Keywords ***
{{ range $i, $Element := .Elements }}
{{ $Element.ElementName }}
    [Documentation]    {{ $Element.ElementDetails }}
    {{- if gt (len $Element.LocateParams) 0}}
    {{- range $j, $Param := $Element.LocateParams }}    
    # {{ $Param.Param }}: {{ $Param.Comment }}   {{ end }}
    [Arguments]    {{ range $j, $Param := $Element.LocateParams }}{{"${"}}{{ $Param.Param }}{{"}"}}    {{ end }}
    {{- end }}
    [Return]    {{ $Element.LocatePattern.Xpath }}
{{ range $j, $Function := $Element.Functions }}
{{ $Function.Name }}
    [Documentation]    {{ $Function.Comment }}
    {{- if gt (add (len $Element.LocateParams) (len $Function.Params)) 0}}
    {{- range $k, $Param := $Element.LocateParams }}    
    # {{ $Param.Param }}: {{ $Param.Comment }}   {{ end }}
    {{- range $k, $Param := $Function.Params }}    
    # {{ $Param.Param }}: {{ $Param.Comment }}   {{ end }}
    [Arguments]    {{ range $k, $Param := $Element.LocateParams }}{{"${"}}{{ $Param.Param }}{{"}"}}    {{ end }}{{ range $k, $Param := $Function.Params }}{{"${"}}{{ $Param.Param }}{{"}"}}{{ end }}{{ end }}
    {{"${"}}{{$Element.ElementName}}{{"}"}}    {{$.Page.PageName}}.{{$Element.ElementName}}    {{ range $k, $Param := $Element.LocateParams }}{{"${"}}{{ $Param.Param }}{{"}"}}    {{ end }}
    {{$Function.Operation}}    {{"${"}}{{$Element.ElementName}}{{"}"}}    {{ range $k, $Param := $Function.Params }}{{"${"}}{{$Param.Param}}{{"}"}}{{ end }} 
{{ end -}}
{{ end }}`
