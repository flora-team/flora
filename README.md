# Flora

[![GitHub release](https://img.shields.io/github/v/tag/flora-team/flora.svg?label=flora)](https://github.com/Flora-team/flora/releases)

flora是一个用于web自动化测试中PO模式下元素管理的工具，支持以yaml配置文件的方式管理元素，并自动生成可直接执行的程序代码。

## 为什么使用flora?

实现PO模式时，由于页面元素数量巨大，需要编写大量的重复性代码，各种参数、注释，写得多了以后会觉得很麻烦。而且一个文件可能内容很长，也给查找元素对应代码或修改元素定位路径带来了麻烦。flora用配置文件的方式代替写代码，然后自动生成可直接粘贴到项目中的代码，从而降低编写页面类代码的难度和复杂度，并提升元素变更后的维护效率。不同团队若使用不同编程语言，也可共同维护一份配置文件，因为flora会帮助您生成各种语言的代码。没有代码能力的测试人员也可以使用flora来分担自动化测试人员的工作。

## 内容列表

- [Flora](#Flora)
  - [安装](#安装)
  - [使用说明](#使用说明)
    - [页面](#页面)
    - [元素](元素)
    - [生成代码](#生成代码)
  - [示例](#示例)
  - [维护者](#维护者)
  - [如何贡献](#如何贡献)
  - [使用许可](#使用许可)

## 安装

这个项目使用 [golang](https://golang.org/) 。如果你本地安装了golang，可以直接执行命令安装:

```sh
$ go install github.com/Flora-team/flora
```

否则，可以直接下载[可执行文件](https://github.com/Flora-team/flora/releases)。

## 使用说明
### 页面

你可以通过flora new page来创建一个页面。你需要在命令最后指定将此页面保存的位置。

```sh
$ flora new page [path]
```

例如执行:

```sh
$ flora new page ./page
```

将会有命令行交互指引你生成一个新页面。

页面的存储形式为一个文件夹和一个同名yaml文件。

你也可以不通过flora new page命令，直接手动新建一个文件夹和yaml文件来创建一个页面。

页面只有两个属性: pageName和pageDetails。它们将会影响生成代码的页面相关名称和注释。

### 元素

你可以通过flora new element来创建一个元素。你需要在命令最后指定将此页面保存的位置。

```sh
$ flora new element [path]
```

例如执行:

```sh
$ flora new element ./page/defaultPage
```

将会有命令行交互指引你生成一个新元素。

元素的存储形式为一个yaml文件。

你也可以不通过flora new element命令，直接手动新建一个yaml文件来创建一个页面。

元素有以下属性:

elementName: 元素名称。此属性必须与文件名相同。

elementDetails: 对此元素的描述。将生成注释。

isBaseElement: 是否是当前页面的基本元素。将用来生成判断页面是否缺失元素的代码。

locateParams: 用于元素定位的参数。

param: 参数。可以用在元素定位和函数中。param有三个属性，param(参数名)，comment(参数注释)和type(参数类型)。

locatePattern: 元素定位路径。当前只支持xpath。其中可以有参数，用${param}的形式直接写在字符串中。用到的参数必须在locateParams中事先定义。

functions: 此元素挂载的函数。函数可以有多个，每个函数可定义name(函数名称)、params(参数列表)、comment(函数注释)和operation(函数行为)。

元素必须在某页面的文件夹下。一个页面下可以有多级目录，但所有元素都会直接挂载到此页面上。多级目录只是为了人工分类。

### 生成代码

通过flora generate命令来生成代码。

有以下参数需要指定:

--language: 需要生成的代码的语言。当前只支持java和robot。其中，生成的java代码使用selenide库。

--package: 需要生成的代码的包名。会影响代码中与包有关的部分(如java)，且会影响输出目录。

--source: 源文件目录。即你的页面、元素所在目录。

--target: 输出目录。生成的代码保存的目录。默认为当前目录。

--operation: 更改默认的函数行为。例如，元素的function的operation因为要输出java代码设定为click，需要改为输出robot时，可以使用 -o click="Click Element"来使其符合robot的代码。-o可以有多个，如 -o click="Click Element" -o setValue="Input Element"。

## 示例
拉取代码:
```sh
$ git clone https://github.com/Flora-team/flora-examples.git
```

java:
```sh
$ flora generate -l java -p com.baidu.test -s baidu
```
robot:
```sh
$ flora generate -l robot -p . -s baidu -o click="Wait And Click" -o setValue="Wait And Input" -t robot
```
## 维护者

[@Naughtz](https://github.com/naughtz)。

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/Flora-team/flora/issues/new) 或者提交一个 Pull Request。


## 使用许可

[MIT](LICENSE) © Naughtz