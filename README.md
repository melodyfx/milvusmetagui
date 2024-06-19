## 简介

milvusmetagui是一款用来对milvus的元数据进行解析的工具，milvus的元数据存储在etcd上，而且经过了序列化，通过etcd-manager这样的工具来查看是一堆二进制乱码，因此开发了这个工具对value进行反序列化解析。



## 使用

### 1.解析database

通过etcd-manager工具搜索`database/db-info`可以列出milvus中的数据库。

![](pic\db01.png)

取其中一个，例如`by-dev/meta/root-coord/database/db-info/1`

![](pic\db02.png)