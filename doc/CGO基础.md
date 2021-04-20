# CGO基础



## 1.准备工作

### 1.1 新增的工具链

在需要进行CGO的开发工作前，除了已经有了Golang的常规开发环境以外，还需要一些C/C++的构建工具链。在macOS/OSX和Linux下，是要安装GCC，在windows下是需要安装MinGW工具（本文档所有实践优先在macOS下测试通过）。

macOS/OSx下面安装gcc（基于Homebrew）

```shell
brew install gcc
```

windows下面安装MinGW可以参考[MinGW官网](http://mingw-w64.org/)

### 1.2 环境变量

在CGO开发和程序执行的过程中，有一些环境变量是会影响程序行为的，我先把主要的一些列举出来，后面的会逐渐补充。



<b>CGO_ENABLED</b>

该环境变量主要用于标识是否开启CGO支持，默认值为1，为1时则开启了CGO的编译支持，需要使用CGO的编译支持时必须开启。



<b>LD_LIBRARY_PATH</b>

该环境变量主要用于指定查找共享库（动态链接库）时除了默认路径之外的其他路径。如果需要加载的动态库在默认路径之外的话，需要对该环境变量进行赋值。



<b>CGO_CFLAGS</b>

该环境变量指定了CGO去读取头文件的路径，同C中的CFLAGS。



<b>CGO_LDFLAGS</b>

该环境变量指定了CGO去读取库文件的路径，同C中的LDFLAGS。



### 1.3 import "C" 语句



如果要在Go中使用CGO的话，需要在import地点加上一行`import "C"`，必须是独立一行，不可以和golang自身引用的包写在一起。

例如：

```go
package main

/*
#include<stdio.h>

void print(int num){
	printf("number is:%d\n",num);
}
*/
import "C"

func main() {
	num := 13
	C.print(C.int(num))
}
```

上例中的`C.int(num)`就是一次将go的num变量（类型int）转为c语言int类型变量的过程，并传递给了print函数的参数。



---

下一节[类型转换](./类型转换.md)