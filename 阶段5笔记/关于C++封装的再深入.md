## 1、引言

对于前些内容的学习做出反思并加以深入学习

## 2. 封装的概念

### 2.1 定义

封装是一种将数据和操作数据的方法绑定在一起的编程概念。它通过将数据成员和成员函数封装在类中，使得外部代码无法直接访问对象的内部数据，只能通过类提供的公共接口进行访问。

### 2.2 原理

在C++中，封装通过访问控制修饰符实现，常用的有`private`、`protected`和`public`。其中，`private`成员只能在类的内部访问，`protected`成员可以在类的内部和派生类中访问，而`public`成员可以被任何地方访问。


cppclass Student {
private:
    int id;
    std::string name;

public:
    void setId(int newId) {
        id = newId;
    }
    int getId() const {
        return id;
    }
    void setName(const std::string& newName) {
        name = newName;
    }
    std::string getName() const {
        return name;
    }
};


在上述代码中，`id`和`name`是私有成员，只能通过公共的成员函数`setId`、`getId`、`setName`和`getName`来进行访问。

### 2.3 优势

#### 2.3.1 数据隐藏

封装提供了数据隐藏的机制，将对象的内部实现细节对外部隐藏起来。这样，对象的使用者无需关心对象的内部结构，只需要通过公共接口来与对象交互，从而降低了程序的复杂性。

#### 2.3.2 安全性

通过将数据成员设为私有，封装提高了代码的安全性。外部代码无法直接修改对象的内部状态，只能通过类提供的接口来进行操作。这有助于防止不合法的访问和修改，提高了程序的健壮性。

#### 2.3.3 可维护性

封装使得对象的实现细节与外部接口分离开来，当需要修改对象的内部实现时，不会影响到外部代码。这提高了代码的可维护性，使得程序更容易理解和修改。

## 3. 封装的实践

### 3.1 类的设计

在进行封装时，首先需要设计一个类，明确定义类的属性和方法。属性通常设为私有，而方法则作为公共接口对外提供服务。


cppclass Rectangle {
private:
    double length;
    double width;

public:
    Rectangle(double len, double wid) : length(len), width(wid) {}
    double getLength() const {
        return length;
    }
    double getWidth() const {
        return width;
    }
    double area() const {
        return length * width;
    }
};


上述代码定义了一个`Rectangle`类，其中`length`和`width`是私有成员，而`getLength`、`getWidth`和`area`是公共成员函数。

### 3.2 封装的使用

使用封装后的类时，可以通过对象的公共接口来访问和修改对象的状态，而无需直接操作其私有成员。


cppint main() {
    Rectangle myRectangle(5.0, 3.0);
    double length = myRectangle.getLength();
    double width = myRectangle.getWidth();
    double area = myRectangle.area();
    return 0;
}


在上述代码中，通过`getLength`、`getWidth`和`area`等公共接口，实现了对`myRectangle`对象的访问和操作。
