# Markdown



<img src="https://img-blog.csdnimg.cn/84b60bcc127844ddacf39a6d9175af0c.png" alt="�刻����亙��餈?" style="zoom:150%;" />

## 1.Markdown的基础用法

Markdown是一种轻量级标记语言，用于简单而易读的文本格式化。以下是Markdown的基本语法规范：

**标题**

使用`#`符号表示标题，数量表示标题级别，例如

```
markdown# 一级标题
## 二级标题
### 三级标题
```

**段落**

段落之间用空行分隔，可以在段落中使用换行符

**强调**

使用`*`或`_`包裹文本表示斜体，使用`**`或`__`包裹文本表示粗体

```
markdown*斜体*  _斜体_
**粗体**  __粗体__
```

**列表**

有序列表使用数字和英文句点，无序列表使用`*`、`+`或`-`

```
markdown1. 有序列表
2. 另一个项

* 无序列表
* 另一个项
```

**链接**

使用`[]`和`()`表示链接，中括号内是链接文本，小括号内是链接地址

```
markdown[文本](http://www.example.com)
```

**图片**

与链接类似，前面加上`！`表示图片

```
markdown![文本](http://www.example.com/image.jpg)
```

**引用**

使用`>`表示引用

```
markdown> 这是引用的文本
```

**代码**

使用反引` `表示行内代码，使用三个反引号 ``` 表示多行代码块：

行内代码 `这是行内代码`

多行代码块：

\```语言 代码... ```

**分割**

使用三个或更多的`*`、`-`或`_`表示分割线：

```
markdown***
```

以上是Markdown的基本语法规范，它被广泛用于编写文档、博客文章和README文件等


## 高级用法：
表格：
markdown
| 列1 | 列2 |
|-----|-----|
| 内容1 | 内容2 |
任务列表：
markdown
- [x] 已完成的任务
- [ ] 未完成的任务
脚注：
markdown
文本[^1]

[^1]: 脚注内容
删除线：
~~删除线~~


## 快捷键用法：
Typora:

粗体：Ctrl + B
斜体：Ctrl + I
引用：Ctrl + Q
代码块：Ctrl + Shift + K
有序列表：Ctrl + Shift + O
无序列表：Ctrl + Shift + U
标题：Ctrl + 1~6
插入链接：Ctrl + K
插入图片：Ctrl + Shift + I
表格：Ctrl + T
段落对齐：Ctrl + R（右对齐）、Ctrl + E（居中对齐）、Ctrl + L（左对齐）
VS Code with Markdown All in One 插件:

预览：Ctrl + Shift + V
切换大纲：Ctrl + Shift + E
插入代码块：Ctrl + Shift + I
加粗：Ctrl + B
斜体：Ctrl + I
插入链接：Ctrl + Shift + L
插入图片：Ctrl + Shift + I
插入表格：Ctrl + Alt + T
插入标题：Ctrl + 1~6
