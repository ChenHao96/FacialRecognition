/**
描述

获取session相关状态和结果

可能的status：INQUEUE(队列中), SUCC(成功) 和FAILED(失败)
当status是SUCC时，返回结果中还包含session对应的结果
所有session都将在计算完成72小时之后过期，并被自动清除。
status返回值为SUCC仅表示成功取得运行结果，实际任务成功与否请根据result内容判断
*/
package info
