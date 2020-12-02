学习笔记
二.异常处理
    goroutine正常处理方式
    Error vs Exception
        业务逻辑不要用panic，业务逻辑一般用error返回异常
    Error Type
        预定义error
            成为API公共部分，创建依赖
        error type
            和预定义error差不多，但携带更多上下文
        opaque error
            Assert errors for behavior
    如何优雅的处理error
        参考bufio.NewScanner (Scan), sql rows
    wrap error
        error只处理一次
        推荐github.com/pkg/errors
        在最上层打印错误日志
        在自身应用代码中，使用errors.New或者erros.Errorf返回
        如果调用其他包内的函数，简单直接返回
