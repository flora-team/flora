package models

var Element = `package {{.}}.models;

import com.codeborne.selenide.SelenideElement;

import java.util.HashMap;

import static com.codeborne.selenide.Selectors.byXpath;
import static com.codeborne.selenide.Selenide.$;

public class Element {
    public LocatePattern locatePattern;
    public String[] locateParams;
    public boolean isBaseElement;
    public Element(LocatePattern locatePattern, boolean isBaseElement) {
        this.locatePattern = locatePattern;
        this.isBaseElement = isBaseElement;
    }
    public Element(LocatePattern locatePattern, boolean isBaseElement, String[] locateParams) {
        if (isBaseElement) {
            System.out.printf("a base element can not have locateParams\n");
        }
        this.locatePattern = locatePattern;
        this.isBaseElement = isBaseElement;
        this.locateParams = locateParams;
    }

    public SelenideElement getElement(String... params) {
        if (locateParams.length != params.length) {
            System.out.printf("param length can not match: %d!=%d\n", locateParams.length, params.length);
            return null;
        } else {
            String xpath = locatePattern.getXpath(locateParams, params);
            System.out.println(xpath);
            return $(byXpath(xpath));
        }
    }
}`
