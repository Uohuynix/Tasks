<<<<<<< HEAD
# Markdown



<img src="https://img-blog.csdnimg.cn/84b60bcc127844ddacf39a6d9175af0c.png" alt="在这里插入图片描述" style="zoom:150%;" />

## 1.Markdown的基础用法

Markdown是一种轻量级标记语言，用于简单而易读的文本格式化。以下是Markdown的基本语法规范：

**标题**

使用`#`符号表示标题，数量表示标题级别，例如：

```
markdown# 一级标题
## 二级标题
### 三级标题
```

**段落**

段落之间用空行分隔，可以在段落中使用换行符。

**强调**

使用`*`或`_`包裹文本表示斜体，使用`**`或`__`包裹文本表示粗体。

```
markdown*斜体* 或 _斜体_
**粗体** 或 __粗体__
```

**列表**

有序列表使用数字和英文句点，无序列表使用`*`、`+`或`-`：

```
markdown1. 有序列表项
2. 另一个项

* 无序列表项
* 另一个项
```

**链接**

使用`[]`和`()`表示链接，中括号内是链接文本，小括号内是链接地址：

```
markdown[文本](http://www.example.com)
```

**图片**

与链接类似，前面加上`！`表示图片：

```
markdown![文本](http://www.example.com/image.jpg)
```

**引用**

使用`>`表示引用：

```
markdown> 这是引用的文本。
```

**代码**

使用反引号 ` 表示行内代码，使用三个反引号 ``` 表示多行代码块：

行内代码： `这是行内代码`

多行代码块：

\```语言 代码... ```

**分割线**

使用三个或更多的`*`、`-`或`_`表示分割线：

```
markdown***
```

以上是Markdown的基本语法规范，它被广泛用于编写文档、博客文章和README文件等。



## 2.Markdown的高级用法

Markdown支持创建简单的表格，使用`|`来分隔单元格，使用`-`来分隔表头和表格内容：

```
markdown| 列1   | 列2   | 列3   |
| ----- | ----- | ----- |
| 内容1  | 内容2  | 内容3  |
```

**任务列表**

任务列表可以用于跟踪任务完成情况。使用`- [ ]`表示未完成任务，`- [x]`表示已完成任务：

```
markdown- [ ] 未完成任务1
- [x] 已完成任务2
- [ ] 未完成任务3
```

**代码块**

指定代码块的语言可以提高可读性，并且一些Markdown编辑器会根据语言进行语法高亮显示：

\```python def hello_world(): print("Hello, world!") ```

**脚注**

可以使用脚注来提供额外的注释或解释。在文本中使用`[^1]`，然后在文档底部添加脚注内容：

```
markdown这是一个文本[^1]。

[^1]: 这是脚注的内容。
```

**删除线**

使用双波浪线`~~`可以表示删除线：

```
markdown~~这是删除线~~
```

**表情符号**

一些平台支持使用表情符号，例如 `:smile:`：

```
markdown:smile:
```

这只是Markdown的一些高级用法，具体支持程度可能因不同的Markdown编辑器或平台而有所不同。使用这些高级功能可以让文档更具表现力和可读性。



## 3.Markdown的快捷键

一级标题：快捷键Ctrl+1或者#

二级标题：快捷键Ctrl+2或者##

三级标题：快捷键Ctrl+3或者###

四级标题：快捷键Ctrl+4或者####

五级标题：快捷键Ctrl+5或者#####

六级标题：快捷键Ctrl+6或者######

段落：Ctrl+0（段落文字输入时候不需要按快捷键，typora会默认为段落文字）

提升标题等级：Ctrl+=

降低标题等级：Ctrl+-

跳转到文首:Ctrl+Home（当前功能还不明白如何操作，不知道该如何使用）

跳转到文末：Ctrl+End

应用内窗口切换：Ctrl+Tab

源代码模式：Ctrl+/

清除样式：Ctrl+\

减少缩进：Ctrl+[

增加缩进：Ctrl+]

偏好设置：Ctrl+逗号

常见快捷方式：

全选：Ctrl+A

加粗：Ctrl+B或者四个*号

复制：Ctrl+C

选中当前词：Ctrl+D

选中当前格式文本：Ctrl+E

查找：Ctrl+F、F3（查找下一个）、Shift+F3（查找上一个）

替换：Ctrl+H

斜体：Ctrl+I

跳转到所选内容：Ctrl+J

超链接：Ctrl + K

选中当前行：Ctrl + L

新建文件：Ctrl + N

打开：Ctrl + O

快速打开（快速打开之前编辑过的历史文件）：Ctrl + P

保存：Ctrl + S

插入表格：Ctrl + T

下划线：Ctrl + U

粘贴：Ctrl + V

关闭：Ctrl + W

剪切：Ctrl + X

重做：Ctrl + Y

撤销：Ctrl + Z ！！！！！

Ctrl + shift + 1 ：大纲

Ctrl + shift + 2 ：文档列表

Ctrl + shift + 3 ：文件树

Ctrl + shift + 9 ：实际大小

Ctrl + shift + C ：复制为MarkDown格式

Ctrl + shift + D ：删除当前词

Ctrl + shift + E ：使用上一次设置导出

Ctrl + shift + F ：搜索

Ctrl + shift + I ：插入图片

Ctrl + shift + K ：插入代码块！！！！！（这个更常用）

Ctrl + shift + L ：显示/隐藏侧边栏

Ctrl + shift + M ：插入公式

Ctrl + shift + N ：新建窗口

Ctrl + shift + Q ：引用

Ctrl + shift + S ：另存为

Ctrl + shift + T ：重新打开关闭的文件

Ctrl + shift + V ：粘贴为纯文本格式

Ctrl + shift + X ：任务列表

Ctrl + shift + [ ：有序列表

Ctrl + shift + ] ：无序列表

Ctrl + shift + = ：放大

Ctrl + shift + - ：缩小

Ctrl + shift + ` ：代码！！！！

Alt + shift + P ：打印

Alt + shift + 5 ：删除线

Alt + Ctrl + P ：选择段落或块

Alt + Ctrl + shift + P ：删除块

Alt + Ctrl + shift + L ：删除当前行或句

Alt + Ctrl + shift + E ：删除当前格式文本

shift + F10 ：快捷操作弹窗

shift + F12 ：开发者工具

F3 ：查找

F8 ：专注模式

F9 ：打字机模式

F11 ：全屏

win + 句号 ：表情与符号（很有意思的一个快捷键）

Home ：跳转到行首

End ：跳转到行尾
>>>>>>> 4b4657df34e534b4b13342503142a821087934ac
