合并重合数组：
问题：
给定一个区间列表，合并所有重叠的区间。例如：

输入：[[1, 3], [2, 6], [8, 10], [15, 18]]

输出：[[1, 6], [8, 10], [15, 18]]

解法：首先将数组按start进行升序排列，然后如何判断是重合数组？上一个数组的End 大于下一个数组的start