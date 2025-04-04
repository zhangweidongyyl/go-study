# 数据压缩算法

    question：一个数组存了一千个数字，设计一个压缩算法能进行压缩和解压
    answer：

    1、数字相邻较近，差分数组存储，如果数字都较小可以用byte来存，这里知识点为
    在 Go 语言中，数字的存储占用的字节数取决于其类型。以下是常见整数类型及其占用的字节数：
    类型	字节数	取值范围
    uint8	1	0 到 255
    int8	1	-128 到 127
    uint16	2	0 到 65535
    int16	2	-32768 到 32767
    uint32	4	0 到 4294967295
    int32	4	-2147483648 到 2147483647
    uint64	8	0 到 18446744073709551615
    int64	8	-9223372036854775808 到 9223372036854775807
    int	4 或 8	取决于平台（32 位或 64 位）
    uint	4 或 8	取决于平台（32 位或 64 位）
    
    2、数字重复的较多，可以用游程编码来实现，至于游程编码如何实现，详细的可见代码

    3、哈弗曼编码实现
    1. 哈夫曼编码
    哈夫曼编码是一种基于字符出现频率的变长编码方法。对于分布不均匀的数字，出现频率较高的数字可以用较短的编码表示，从而减少总存储空间。
    
    实现步骤
    统计每个数字的出现频率。
    
    根据频率构建哈夫曼树。
    
    为每个数字生成哈夫曼编码。
    
    将原始数组转换为哈夫曼编码的二进制流。
    
    解压时，根据哈夫曼树将二进制流还原为原始数字。

    4、稀疏矩阵：如果数组中大部分数字是 0 或某个默认值，可以使用稀疏矩阵存储来压缩数据。




常见的可以分为几种case：
## 1、数据相邻的较多--使用差分编码解决
    思路：存储相邻数据的差值，这样能少存数据，适合数据相邻较近的case
## 1、数据里面重复较多--使用游程编码解决

    前提条件：这种处理适合重复数据在连续位置的情况
  使用游程编码就能实现，即记录每个数据出现的次数，一般要将给出的数据进行升序排列
  例如：原始数据为
    [1,2,2,3,3,44,44]
   游程编码处理后的数据为
    [[1,1],[2,2],[3,2],[44,2]]
    数据结构为 

    
    type RuneLengthDomain struct {
        Val   int // 具体数字
        Count int // 数字出现次数
    }
    
### 1.1、数据量比较小的case
  直接内存操作

### 1.2、数据量较大，如超过1000万，1个亿这样的case
  使用流式处理，逐块处理，处理完后压缩存到文件中

