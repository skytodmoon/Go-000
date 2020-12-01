项目配置说明
数据库:mysql
初始化数据库文件：initUser.sql
默认用户名：root
密码：mysql-password
可以在main.go中修改为自己的数据库和密码

作业说明：
在Dao层和Service层封装Wrap error，方便查询错误日志
Gorm已经在底层封装了gorm.ErrRecordNotFound
在http.go中判断错误类型是否是gorm.ErrRecordNotFound，如果是，则转义为业务层错误，隐藏接口内部错误