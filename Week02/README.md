# Week02 作业题目：

## 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

sql.ErrNoRows 是标准库的函数，dao层需要对其做wrap包装。但不应该包含具体的数据操作的信息（e.g.: sql scripts）. <br>
实际上，和具体数据库相关的信息应该要对上层隐藏。<br>
但就dao这一层来说，个人认为可以另起一些log，把出现错误的语句log起来。 当然，也可以在开发中打一些debug的log。<br>
比如在使用orm的时候，可以查看orm生成的sql语句是否正确，是否有逻辑/性能问题等等。<br>
<br>
大致代码如下：<br>


```go
//service UserService
func GetUser(id int) User {
    b := UserBiz{}
    user, err := b.GetUser(id)
    if err!=nil{
        fmt.Printf("%+v\r\n", err)
        //handle the error and return
        //for example: return 404 not found / an empty response body.
	//that based on the business/real scenario.
        return User{}
    }
    return user
}

//biz UserBiz
func GetUser(id int) (User, error) {
    d := UserDao{}
    return d.GetUser(id)
}

//dao UerDao
func GetUser(id int) (User, error) {
    var u User
    if id >= 1 && id <= 10 {
        u = User{accountID: id, firstName: "Hello", lastName: "world"}
	return u, nil
    }
    u = User{}
    var message = fmt.Sprintf("dao: user not found. (id:%d)", id)
    return u, errors.Wrapf(sql.ErrNoRows, message)
}
```
