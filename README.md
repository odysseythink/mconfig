# mconfig是干什么的
mconfig是一个读取配置文件的库。无需定义struct，可以动态获取json配置文件的各个节点的值。目前只支持json，后面会加入对xml，yaml，ini等主流格式的支持。

# 为什么要写mconfig库
golang传统解析json格式的配置文件是先定义一个struct结构体，然后通过json.Unmarshal解析到这个结构体中。这种方式有一些不灵活的地方：
- 对于配置文件里面的字段必须要有一个struct结构体与之对应
- 配置文件中增加字段，struct里面必须要相应增加字段
mconfig库把整个json文件解析到map[string]interface{}中，然后需要什么字段就是用Get方法直接获取。
这样就不需要定义struct，不管json文件怎么变，代码都不需要变。子结构也支持，直接用形如：key1.key2.key3这种格式获取。
