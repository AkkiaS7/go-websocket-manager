# (WIP) go-websocket-manager 
websocket connection manager

# (WIP)暂定流程如下
```
----------------------------------------------------
|   HEADER(5Byte FIXED)   |          BODY          |
----------------------------------------------------
| CMDLength |  DataLength |    CMD    |    Data    | 
|   uint8   |    uint32   | CMDLength | DataLength |
----------------------------------------------------
```
1. 添加对应CMD的路由
2. 按照上述规格构造数据包
3. 通过CMD从路由中查找处理函数
4. 由找到的处理函数对Data进行处理