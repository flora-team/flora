package models

var LocatePattern = `package {{.}}.models;

public class LocatePattern {
    private String xpath;
    public LocatePattern(String xpath) {
        this.xpath = xpath;
    }

    public String getXpath() {
        return xpath;
    }
    public String getXpath(String[] locateParams, String[] params) {
        String tmpXpath = xpath;
        for (int i=0; i<params.length; i++) {
            tmpXpath = tmpXpath.replace("${"+locateParams[i]+"}", params[i]);
        }
        return tmpXpath;
    }
}`
