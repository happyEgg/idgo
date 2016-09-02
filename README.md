# idgo
基于beego开发的id生成器。

步骤：
1、安装redis
2、cd idgo   bee run

生成的id是自增的。id每次从redis中获取，性能比从MySQL获取提高很多