# CSS学习笔记

## 一、CSS基础用法

### 1. **基本概念：**

- **选择器（Selector）：** 选择器用于选中HTML中的元素，以便对其应用样式。常见的选择器有类选择器（.class）、ID选择器（#id）、元素选择器（element）等。
- **属性（Property）：** CSS规则由选择器和一组属性-值对组成。属性定义了要应用的样式，比如`color`、`font-size`等。
- **值（Value）：** 属性的值定义了样式的具体表现，例如，`color: blue;`中的"blue"是值。
- **盒模型（Box Model）：** 每个HTML元素都被看作是一个矩形的盒子，包括内容区、内边距、边框和外边距。

### 2. **常用属性：**

- **文本样式：**
  - `color`：文本颜色。
  - `font-family`：字体系列。
  - `font-size`：字体大小。
  - `font-weight`：字体粗细。
- **盒模型：**
  - `width`：元素宽度。
  - `height`：元素高度。
  - `margin`：外边距。
  - `padding`：内边距。
  - `border`：边框。
- **布局：**
  - `display`：指定元素应该生成的框的类型。
  - `position`：设置元素的定位方式。
  - `float`：浮动。
- **背景：**
  - `background-color`：背景颜色。
  - `background-image`：背景图片。
  - `background-position`：背景位置。
- **定位：**
  - `top`、`right`、`bottom`、`left`：相对定位元素的位置。
- **动画与过渡：**
  - `transition`：定义过渡效果。
  - `animation`：定义动画效果。

### 3. **样式的引入方式：**

- **内部样式表：** 在HTML文档头部的`<style>`标签中定义样式。

  ```
  html<head>
    <style>
      body {
        background-color: #f0f0f0;
      }
      h1 {
        color: blue;
      }
    </style>
  </head>
  ```

- **外部样式表：** 将样式定义在一个独立的CSS文件中，然后在HTML中引入。

  ```
  html<head>
    <link rel="stylesheet" type="text/css" href="styles.css">
  </head>
  ```

- **内联样式：** 将样式直接应用到HTML元素。

  ```
  <p style="color: green; font-size: 16px;">This is a paragraph.</p>
  ```



## 二、对于Flex、Grid、响应式设计的原理理解

### 1. **Flex布局的基本原理：**

Flex布局是一种用于设计复杂布局结构的弹性盒子模型。它的基本原理包括以下几点：

- **容器属性：**
  - `display: flex;` 将容器设置为弹性容器。
  - `flex-direction` 定义主轴的方向（row、row-reverse、column、column-reverse）。
  - `justify-content` 控制项目在主轴上的对齐方式。
  - `align-items` 控制项目在交叉轴上的对齐方式。
  - `align-content` 处理多根轴线的对齐方式。
- **项目属性：**
  - `order` 定义项目的排列顺序。
  - `flex-grow` 定义项目的放大比例。
  - `flex-shrink` 定义项目的缩小比例。
  - `flex-basis` 定义项目在主轴上的初始大小。
  - `flex` 是 `flex-grow`, `flex-shrink`, 和 `flex-basis` 的缩写。

### 2. **Grid布局的基本原理：**

Grid布局是一种二维布局系统，通过将布局划分为行和列来实现。基本原理包括：

- **容器属性：**
  - `display: grid;` 将容器设置为网格容器。
  - `grid-template-columns` 和 `grid-template-rows` 定义网格的列和行。
  - `grid-gap` 定义行和列之间的间隔。
  - `grid-template-areas` 定义区域的名称，然后使用 `grid-area` 将项目放置在特定区域。
- **项目属性：**
  - `grid-column` 和 `grid-row` 定义项目所占的列和行。
  - `grid-area` 用于指定项目所占的区域。

### 3. **响应式设计的基本原理：**

响应式设计旨在使网站能够适应不同设备和屏幕尺寸，基本原理包括：

- **媒体查询（Media Queries）：**
  - 使用 `@media` 查询，根据不同的条件（如屏幕宽度）应用不同的CSS规则。
- **弹性网格与相对单位：**
  - 使用相对单位（如百分比、em、rem）而非绝对单位，以确保元素相对于其容器或父元素的大小而变化。
  - 使用弹性网格布局，如Flex布局和Grid布局，以适应不同屏幕尺寸。
- **图片和多媒体的响应式处理：**
  - 使用 `max-width: 100%;` 保证图片和多媒体在小屏幕上不会超出其容器。
- **流式布局（Fluid Layout）：**
  - 使用百分比布局或弹性盒子模型，使布局具有弹性，能够适应不同屏幕尺寸。
- **隐藏、显示与重新排列：**
  - 使用 `display: none;` 或 `visibility: hidden;` 根据需要隐藏某些元素。
  - 使用 `order` 属性重新排列Flex项目的顺序。