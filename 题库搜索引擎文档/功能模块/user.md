# modle

## 角色/Role

- kRoot			代码：0
- kAdmin   	代码： 1
- kUser           代码：10

## 具体含义未知/Module

- kMathDepartment  =  1

## 权限/Permissions

- kCompilcation   ->   101
- kItemSearch       ->  102
- kPaperEdit 		 ->  103

## 用户实体/User

- 数据库表名：user
- email：string
- name：string
- role：枚举 Role
- permissions：列表  Permissions
- valid：bool
- deleted：bool
- module：枚举  Module
- mobile：string
- mark：string
- pass_salt：string
- pass_hash：string
- last_login_ip：string
- last_login_time：float
- ctime：float  创建的时间戳
- mtime：float  最近修改的时间戳
- 用户方法
  - display_name:显示用户名或邮箱
  - set_pass：设置密码
  - validate_pass：验证密码
  - save：保存修改，更新mtime
  - change_pass：修改密码
  - strip_secure_info：将pass_hash和pass_salt设置为""，作用不明
  - to_dict：
  - is_admin：
  - is_valid
  - is_valid_pass：判断密码是否符合要求
  - is_valid_email
  - is_valid_role：
  - is_valid_permission：
  - is_valid_module：
  - create：创建用户

## 操作日志/OpLog

- 数据库表名：op_log

- uid：ObjectID

- ctime：float

- type：string

- obj_type：string

- obj_id：ObjectId

- obj_extra：Dict

- 方法

  - aggregate

    ```python
    def aggregate(cls, conds):
            return db.session.db.get_collection(cls.config_collection_name).aggregate(conds)
    ```

## 反馈/Feedback

- 数据库表名：feedback

- Status

  - kNew  =   1
  - kResolved  = 2
  - kIgnored  =   3

- creator_id：objectID

- obj_type：string

- obj_id：ObjectId

- desc：Dict

- reply：string

- reply_id：ObjectID

- status：枚举 status

- ctime：float

- mtime：float

  

# handlers

这一部分负责和用户个人信息相关的全部操作

URL包括

```json
/user
/op_logs
/feedbacks
/admin/adduser

/user/<oid:uid>
/user/<oid:uid>/statistics
/user/<oid:uid>/edit
/user/<oid:uid>/op_logs
/users/statistics

/profile
/profile/statistics
/profile/op_logs
/profile/feedbacks
/profile/edit

/api/user/<oid:uid>
/api/user/<oid:id>/statistics
/api/user/<oid:id>/op_logs


/api/users
/api/users/statistics

/api/op_logs
/api/feedbacks
/api/feedback/<oid:id>
```









- ensure_logined：确保用户处于登录状态

- login_page：

  - 方法：Get
  - URL：/auth/login
  - 作用：获取登录界面

- login

  - 方法：POST

  - url：/auth/login

  - 参数

    ```json
    {
            'type': 'object',
            'properties': {
                'email': {'type': 'string', 'format': 'email'},
                'pass': {'type': 'string', 'minLength': 1},
                'remember': {'type': 'boolean'}
            },
            'required': ['email', 'pass']
        }
    ```

  - 作用：验证用户登录信息并登录，生成对应session

- logout
  - URL：/auth/logout
  - 作用：退出登录（删除session）

- userlist_page：

  - URL

    ```python
    @app.route('/users')
    @app.route('/op_logs')
    @app.route('/feedbacks')
    @app.route('/admin/adduser')
    @app.route('/user/<oid:uid>/edit')
    @app.route('/user/<oid:uid>/statistics')
    @app.route('/user/<oid:uid>/op_logs')
    @app.route('/users/statistics')
    ```

  - 作用：显示base.html页面

































































































