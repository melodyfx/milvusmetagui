## 简介

milvusmetagui是一款用来对milvus的元数据进行解析的工具，milvus的元数据存储在etcd上，而且经过了序列化，通过etcd-manager这样的工具来查看是一堆二进制乱码，因此开发了这个工具对value进行反序列化解析。



## 使用

### 1.解析database

通过etcd-manager工具搜索`database/db-info`可以列出milvus中的数据库。

![](pic\db01.png)

取其中一个，例如`by-dev/meta/root-coord/database/db-info/1`

<img src="pic\db02.png" style="zoom:50%;" />

### 2.解析collection

通过etcd-manager工具搜索`database/collection-info`可以列出milvus中的collection。

![](pic\col01.png)

取其中一个，例如`by-dev/meta/root-coord/database/collection-info/1/449952137045880999`

<img src="pic\col02.png" style="zoom:50%;" />



### 3.解析fields

通过etcd-manager工具搜索`root-coord/fields`可以列出milvus中的field。

![](pic\fields01.png)

取其中一个，例如`by-dev/meta/root-coord/fields/449952137045880999/102`

<img src="pic\fields02.png" style="zoom:50%;" />

### 4.解析field-index

通过etcd-manager工具搜索`field-index`可以列出milvus中的field-index。

![](pic\field-index01.png)

取其中一个，例如`by-dev/meta/field-index/449952137045880999/449952137045881004`

<img src="pic\field-index02.png" style="zoom:50%;" />

### 5.解析segment-index

通过etcd-manager工具搜索`segment-index`可以列出milvus中的segment-index。

![](pic\seg-index01.png)

取其中一个，例如`by-dev/meta/segment-index/449952137045880999/449952137045881000/449952137046086601/449952137047289214`

<img src="pic\seg-index02.png" style="zoom:50%;" />

### 6.解析partition

通过etcd-manager工具搜索`root-coord/partitions`可以列出milvus中的partition。

![](pic\partition01.png)

取其中一个，例如`by-dev/meta/segment-index/449952137045880999/449952137045881000/449952137046086601/449952137047289214`

<img src="pic\partition02.png" style="zoom:50%;" />