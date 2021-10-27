package models

var BasePage = `package {{.}}.models;

import com.codeborne.selenide.Condition;
import com.codeborne.selenide.WebDriverRunner;
import com.codeborne.selenide.impl.SelenidePageFactory;

import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public abstract class BasePage<Page> extends SelenidePageFactory {
    protected List<Element> elements;

    protected BasePage() {
        this.page(WebDriverRunner.driver(), this);
    }

    public List<Element> getAllElements(Object object) {
        List<Element> elements = new ArrayList<>();
        Class className = object.getClass();
        for (; className != Object.class; className = className.getSuperclass()) {
            Field[] fields = className.getDeclaredFields();
            for (Field field : fields) {
                if (field.getGenericType().toString().equals("class {{.}}.models.Element")) {
                    try {
                        field.setAccessible(true);
                        Element e = (Element) field.get(object);
                        elements.add(e);
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
            }
        }
        return elements;
    }

    public void checkBaseElements() {
        List<Element> baseElements = elements.stream().filter(e -> e.isBaseElement).collect(Collectors.toList());
        for (Element e : baseElements) {
            e.getElement().should(Condition.exist);
        }
    }

}`
