`_test.go`为后缀的源文件是go test测试的一部分。
在这些文件中有三种类型的函数：
 1. 测试函数，以Test函数名为前缀的函数
 2. 基准测试函数，以Benchmark函数名为前缀的函数，用于衡量函数的性能。go test命令会多次运行基准函数以计算一个平均的执行时间
 3. 示例函数，以Example函数名为前缀的函数

## 扩展
Benchmark https://en.wikipedia.org/wiki/Benchmark_(computing)

