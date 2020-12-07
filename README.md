# bitree
### 描述
bitree用途是bitmap，当存储部分数据时更加节省内存。
### 应用场景
1. 当不知道要存储数据的范围时，但是知道大概有多少数据，比如存储40亿数据中的某一段大小为1百万的数据
2. 将ip转换成uint32,存储某一网段数据
### 对比
当使用未压缩的bitmap存储数据时，如果要存储数据范围时0-40亿时，数据量为100百万，需要提前分配大概
500MB的内存。使用bitree，如果数据比较满足局部性则需要大概10M大小的内存，最小需要100KB的内存，
最坏数据很离散需要500MB的内存。

#### TODO
1. 是否当叶子节点的bitmap的数据大小为0时，删除该叶子节点。
这个需要加锁，考虑到一般的应用场景不会用到，所以先不加该功能。也可以加配置选项解决。