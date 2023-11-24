# JavaScript学习笔记

JavaScript 是一种脚本语言，用于在网页上添加交互性。

1. **变量声明：**

   使用 `var`、`let` 或 `const` 关键字声明变量。

   ```
   javascriptvar x = 5; // 使用 var 声明变量
   let y = 10; // 使用 let 声明变量（具有块级作用域）
   const PI = 3.14; // 使用 const 声明常量
   ```

2. **数据类型：**

   JavaScript 中有多种数据类型，包括字符串、数字、布尔、对象、数组等。

   ```
   javascriptvar message = "Hello, World!"; // 字符串
   var count = 42; // 数字
   var isTrue = true; // 布尔
   var person = { name: "John", age: 30 }; // 对象
   var numbers = [1, 2, 3, 4, 5]; // 数组
   ```

3. **条件语句：**

   使用 `if`、`else if` 和 `else` 来进行条件判断。

   ```
   javascriptvar num = 10;
   
   if (num > 0) {
       console.log("Positive");
   } else if (num < 0) {
       console.log("Negative");
   } else {
       console.log("Zero");
   }
   ```

4. **循环语句：**

   使用 `for`、`while` 或 `do-while` 进行循环。

   ```
   javascriptfor (var i = 0; i < 5; i++) {
       console.log(i);
   }
   
   var i = 0;
   while (i < 5) {
       console.log(i);
       i++;
   }
   ```

5. **函数：**

   使用 `function` 关键字定义函数。

   ```
   javascriptfunction add(a, b) {
       return a + b;
   }
   
   var result = add(3, 4);
   console.log(result); // 输出 7
   ```

### JavaScript 数据类型：

1. **字符串（String）：**

   表示文本数据，用单引号或双引号括起来。

   ```
   javascriptvar greeting = "Hello, World!";
   ```

2. **数字（Number）：**

   表示数值，可以是整数或浮点数。

   ```
   javascriptvar num = 42;
   var pi = 3.14;
   ```

3. **布尔（Boolean）：**

   表示逻辑值，只有两个可能的值：`true` 或 `false`。

   ```
   javascriptvar isTrue = true;
   var isFalse = false;
   ```

4. **Null 和 Undefined：**

   `null` 表示一个空对象引用，`undefined` 表示一个未定义的变量。

   ```
   javascriptvar emptyValue = null;
   var undefinedValue;
   ```

5. **对象（Object）：**

   用于存储键值对的集合。

   ```
   javascriptvar person = {
       name: "John",
       age: 30,
       isStudent: false
   };
   ```

6. **数组（Array）：**

   用于存储一组有序的值。

   ```
   javascriptvar colors = ["red", "green", "blue"];
   ```