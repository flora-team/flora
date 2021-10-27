package pages

var Page = `package {{.PackageName}}.pages;

import {{.PackageName}}.models.BasePage;
import {{.PackageName}}.models.Element;
import {{.PackageName}}.models.LocatePattern;


/**
 * {{.Page.PageDetails}}
 */
public class {{.Page.PageName}} extends BasePage<{{.Page.PageName}}> {
    {{ range $i, $Element := .Elements }}
    /**
     * {{.ElementDetails}}
     */
    private Element {{.ElementName}} = new Element(
            new LocatePattern({{.LocatePattern.Xpath}}),
            {{.IsBaseElement}}{{if gt (len .LocateParams) 0}},
            // {{ range $j, $Param := .LocateParams }}{{ $Param.Comment }}{{repeat $j (len $Element.LocateParams) ", "}}{{ end }}
            new String[] { {{ range $j, $Param := .LocateParams }}"{{ $Param.Param }}"{{repeat $j (len $Element.LocateParams) ", "}}{{ end }} } {{end}}
    );
    {{ end }}
    public {{.Page.PageName}}() {
        super.elements = getAllElements(this);
    }
    {{ range $i, $Element := .Elements -}}
    {{ range $j, $Function := $Element.Functions }}
    /**
     * {{$Function.Comment}}
     {{ range $k, $Param := $Element.LocateParams }}* @param {{$Param.Param}} {{$Param.Comment}} 
     {{ end -}}
     {{ range $k, $Param := $Function.Params }}* @param {{$Param.Param}} {{$Param.Comment}} 
     {{ end -}}
     * @return {{$.Page.PageName}}
     */
    public {{$.Page.PageName}} {{$Function.Name}}({{ range $k, $Param := $Element.LocateParams }}String {{$Param.Param}}{{repeat $k (add (len $Element.LocateParams) (len $Function.Params)) ", "}}{{ end }}{{ range $k, $Param := $Function.Params }}{{$Param.Type}} {{$Param.Param}}{{repeat $k (len $Function.Params) ", "}}{{- end -}}) {
        {{$Element.ElementName}}.getElement({{ range $k, $Param := $Element.LocateParams }}{{$Param.Param}}{{repeat $k (len $Element.LocateParams) ", "}}{{ end }}).{{$Function.Operation}}({{ range $k, $Param := $Function.Params }}{{$Param.Param}}{{repeat $k (len $Function.Params) ", "}}{{- end -}});
        return this;
    }
    {{ end }}
    {{- end }}

}`
