# HTML学习笔记

### 常用HTML标签：

1. **`<div>`（Division）:**

   - 用于定义文档中的区块或容器，常用于组织和布局页面结构。

   ```
   html<div>
       <!-- 内容放在这里 -->
   </div>
   ```

2. **`<p>`（Paragraph）:**

   - 用于定义段落。

   ```
   html<p>This is a paragraph.</p>
   ```

3. **`<a>`（Anchor）:**

   - 用于创建超链接，将文本或图像链接到其他页面或资源。

   ```
   html<a href="https://www.example.com">Visit our website</a>
   ```

4. **`<button>`:**

   - 用于创建按钮。

   ```
   html<button type="button">Click me</button>
   ```

### 块级元素与内联元素：

1. **块级元素（Block-level Elements）:**
   - 占据页面的一整行或一整块，独自一行显示，宽度默认为父元素的100%。
   - 例如：`<div>`, `<p>`, `<h1>` to `<h6>`, `<ul>`, `<ol>`, `<li>`, 等。
2. **内联元素（Inline Elements）:**
   - 只占据其本身的宽度，不会导致换行，与相邻的内联元素在同一行显示。
   - 例如：`<a>`, `<span>`, `<strong>`, `<em>`, `<img>`, 等。

### HTML标签的样式：

HTML本身并没有样式，样式通常通过CSS来设置。但是，有一些HTML标签本身带有一些默认样式。

1. **样式属性：**

   - 有一些HTML标签支持内联样式，可以通过`style`属性设置。

   ```
   html<p style="color: blue; font-size: 16px;">This is a styled paragraph.</p>
   ```

2. **CSS样式表：**

   - 最常见的做法是使用外部或内部CSS样式表，通过选择器为HTML元素应用样式。

   ```
   <head>
       <style>
           p {
               color: red;
               font-size: 18px;
           }
       </style>
   </head>
   <body>
       <p>This is a styled paragraph.</p>
   </body>
   ```