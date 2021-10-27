package pages

var Base = `*** Settings ***
Library           Selenium2Library    
*** Variables ***

*** Keywords ***
Wait And Click
    [Timeout]    1 minutes
    [Documentation]    用法同Click Element
    [Arguments]    ${locator}    ${timeout}=5s
    Wait Until Page Contains Element    ${locator}    ${timeout}     
    ${count}    Get Element Count    ${locator}
    FOR    ${i}    IN RANGE    ${count}
        ${tlocator}    Set Variable    (${locator})[${i+1}]
        ${status}    Run Keyword And Return Status    Wait Until Element Is Visible    ${locator}    ${timeout}
        Exit For Loop If    ${status}==${true}
    END
    ${locator}    Set Variable If
    ...   ${count}>${1}    ${tlocator}
    ...    ${locator}
    ${status}    Set Variable    ${false}
    # 进行10次尝试，如果失败则放弃
    FOR    ${i}    IN RANGE    9
        sleep    20ms
        ${status}    Run Keyword And Return Status    Click Element    ${locator}
        Exit For Loop If    ${status}
    END
    Run Keyword If    ${status}==${false}    Click Element    ${locator}  

Wait And Input
    [Timeout]    1 minutes
    [Documentation]    用法同Input Text
    [Arguments]    ${locator}    ${text}    ${timeout}=30s
    Wait Until Page Contains Element    ${locator}    ${timeout}      
    FOR    ${i}    IN RANGE    ${5}      
        ${status}    Run Keyword And Return Status    Input Text    ${locator}    ${text}
        Exit For Loop If    ${status}==${true}
        Sleep    200ms
    END
    FOR    ${i}    IN RANGE    ${5}
        Simulate Event    ${locator}    blur
        Sleep    100ms
        ${isFocused}    Run Keyword And Return Status    Element Should Be Focused    ${locator}
        Log    ${locator} is focused ${isFocused}
        Exit For Loop If    ${isFocused}==${FALSE}
    END
    
Wait And Mouseover
    [Timeout]    1 minutes
    [Documentation]    用法同Mouse Over
    [Arguments]    ${locator}    ${timeout}=30s
    Wait Until Page Contains Element    ${locator}    ${timeout}
    Wait Until Element Is Visible    ${locator}    ${timeout}
    # 进行10次尝试，如果失败则放弃
    FOR    ${i}    IN RANGE    9
        sleep    20ms
        ${status}    Run Keyword And Return Status    Mouse Over    ${locator}
        Exit For Loop If    ${status}
    END
    Run Keyword If    ${status}==${false}    Mouse Over    ${locator}  `
