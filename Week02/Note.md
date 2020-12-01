
## Tips

- errors.New()返回的是指针,这样每次New出来的Error,哪怕起string完全一致,返回的也是不同的对象

- 不要在请求中做Gorouting, 避免挂掉无法recover

- 启动配置文件错误,需要panic

- init中做为资源的初始化,不成功,则panic

- 其他情况下,用error进行处理

- 不要对返回的error做任何假设,应该处理它们.

- 做为服务端,在野生的routine中出现的panic会影响整个进程退出,那么需要做recover进行兜底.

- 对于创建的routine需要管理它的生命周期

- 业务逻辑不要用panic

- error考虑的是错误,而不是成功
- 立即处理,而不要丢给上层
- error同样也是value


## Sentinel Error
在包级别变量,预定义的Error叫做`Sentinel Error`.这种方式并不灵活.
当想有判断更多的信息的时候,就很麻烦了.例如Error里面带有其他的内容时...

导致了包之间存在了依赖..
但是返回业务的错误码时,可以使用这种模型.

一句话,尽可能避免

## 自定义的Error类型



使用时,使用类型断言
``` Go
err := test()
switch err := err.(type) {
case nil:
    // nothing to do
case *MyError:
fmt.Printf("Error occurred on line:", err.line)
default:
    // unknown error
}
```

使用错误类型的好处:
- 能够包装上下层以提供更多上下文
- 也是尽量避免

## 建议使用 Opaque errors(非透明)

一句话:做为调用者,就只用关心对错与否即可.函数只返回错误而不假定其他内容

优点:灵活,耦合度低


那么要更多上下文怎么搞?

- 用隐藏的interface(小写的)
- 断言一个行为,在一个暴露的方法内,断言error,然后再进行相关的行为即可
- 例如IsTimeout(error) bool

## Handing error

需要枚举,遍历的时候,可以这样做.会很好
![2d6159189c2162f94514a073a2658164.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p716)

即error包裹,带上状态,最后返回的error就是里面的不暴露的error状态

![f0e00f2e1a9455ba3cbe275b2f2f3c36.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p717)

## Wrap errors

场景: 报错,一层层往上层丢,那么最终看到的是最里面的一个error,而由于没有上下文..无法理解哪里错了.

解决1: return fmt.Errorf("xxx:%v", err)
解决1的新问题:
- 没有stack信息,无法追踪
- 层层记录,日志非常长.

错误行为实例:
![4132a38909d5f61906529a8809c5e66d.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p718)

这里被记录了2次日志.

做微服务,报错了
- 要么则把error吞掉,返回一个降级的数据,
- 否则就要把错误往上面抛

日志记录与错误无关的且调试没有帮助的信息应该视为noise.

### 原则
- 错误要被记录
- 应用程序要处理error,保证100%完整性.(例如降级)
- 之后不再报告当前错误.

看github.com/pkg/errors

### 示范

- 最底层,使用errors.Wrap("xxxx") 会保存stack信息
- 那么上层则就使用errors.WithMessage(err, "Upper message")
- 最上层来打印日志errors.Cause(err)
- 那么实际上底层是把error保存起来了
- 例如中间件,一个地方来打日志即可,可以把stack打印出来.

> 思考

- 在应用代码中,使用errors.New或者errors.Errorf返回错误
- 调用本package内的函数,则简单的(直接)返回error即可
- 调用第三方的库(包括标准库)可以考虑使用errors.Wrap或者errors.Wrapf来保存堆栈信息.(例如最底层与数据库交互的时候)
  -直接返回错误而不是直接打日志
- 在程序的顶部使用%+v打印完整的堆栈
- 使用errors.Cause()来获得root error,再来做`Sentinel Error`来判断.

- 如果kit(基础)库,则不应该Wrap.选择Wrap是只能application应该做的,具有可重用性的包,则返回根error.Example:`sql.ErrNoRows`
- 如果函数或方法不打算处理,则可以用Wrap.Example:调用sql报错了,则可以Wrap,把调用的sql和参数带进去(可以参考java中sql异常的那种打印)
- 如果函数处理了这个error,则不能返回错误值,应该返回nil.Example:解降级处理,你返回了降级数据,然后需要return nil.

> 个人理解
- 如果处理降级,则返回nil
- 如果不处理,可以Wrap一定的内容
- kit库不Wrap.应用库可以Wrap
- 实现了error interface的自定义struct,使用断言来获取更加丰富的上下文
- 连接sql那一层,如果没查到最好还是返回一个error,判断error即可.否则可能要判断2次.

### 如何正确使用 Sentinel Error
这样来说,返回的是一个没有依赖的Error信息
![16eb3c8f1d7949c940d5ba783f629878.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p719)
注意是`%w`,使用error.Is来判断
![9d0085d44c574ba73a03dd869f3bb9a5.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p720)

### errors & github.com/pkg/errors


![a80b4ffbb5897855d1c7ba863252ac0d.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p721)


## Panic处理

![88cd18b27c334d59580b718d198fbdd7.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p722)


# References

go.googlesource.com