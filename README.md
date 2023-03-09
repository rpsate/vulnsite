# VulnSite  API说明文档

1. ### 登陆
   
   1. URL：/v1/auth/login
   
   2. 提交方式：GET
   
   3. 参数格式：JSON
   
   4. 提交参数：
      
      1. Username    用户名    string    必填
      2. Password    密码    string    必填
   
   5. 返回数据：
      
      1. token    登陆凭证    string

2. ### 管理员管理
   
   1. ##### 获取管理员列表
      
      1. URL：/v1/Admin
      2. 提交方式：GET
      3. 携带token：是
      4. 提交参数：无
      5. 返回数据：
         1. AdminInfoList    管理员信息数组 array
            1. ID    用户id    int
            2. Username    用户名   string
            3. Name    用户真实姓名   string
            4. Grade    权限    int（1：超级管理员，2：普通管理员）
   
   2. ##### 添加管理员
      
      1. URL：/v1/Admin
      2. 提交方式：POST
      3. 参数格式：JSON
      4. 携带token：是
      5. 参数格式：JSON
      6. 提交参数：
         1. Username    用户名    string    必填
         2. Password    密码     string    必填
         3. Name    真实姓名    string    必填
         4. Grade    权限    string    可选（default：2）
      7. 返回数据：空
   
   3. ##### 修改管理员信息
      
      1. URL：/v1/Admin
      2. 提交方式：PUT
      3. 参数格式：JSON
      4. 携带token：是
      5. 参数格式：JSON
      6. 提交参数：
         1. ID    需要修改管理员的id    int    必填
         2. name    修改后的真实姓名    string 可选
         3. password    修改后的新密码    string    可选
         4. Grade    修改后的权限    int    可选
      7. 返回数据：空
   
   4. ##### 删除管理员
      
      1. URL：/v1/Admin/:id
      2. 提交方式：DELETE
      3. 携带token：是
      4. 参数格式：Params（将url中的:id替换成参数id即可）
      5. 提交参数：
         1. id    要删除管理员的id    string    必填
      6. 返回数据：空

3. ### 我的账号信息
   
   1. ##### 查看我的信息
      
      1. URL：/v1/Admin/My
      2. 提交方式：GET
      3. 携带token：是
      4. 提交参数：无
      5. 返回数据：
         1. AdminInfo    用户信息    array
            1. ID    用户ID    int
            2. Username    用户名    string
            3. Name    用户真实姓名    string
            4. Grade    用户权限    int
   
   2. ##### 修改我的信息
      
      1. URL：/v1/Admin/My
      2. 提交方式：PUT
      3. 参数格式：JSON
      4. 携带token：是
      5. 提交参数：
         1. Password    修改后的新密码    string    可选
         2. Name    修改后的姓名    string   可选
      6. 返回数据：
         1. AdminInfo    用户信息    array
            1. ID    用户ID    int
            2. Username    用户名    string
            3. Name    用户真实姓名    string
            4. Grade    用户权限    int

4. ### 靶机管理
   
   1. ##### 添加靶机
      
      1. URL：/v1/Machine
      2. 提交方式：POST
      3. 参数格式：JSON
      4. 携带token：是
      5. 提交参数：
         1. Title    靶机标题    string    必填
         2. Desc    靶机描述    string    可选
         3. DownAddr    网盘下载地址    string    可选
         4. DownPwd     网盘提取密码    string    可选
         5. Difficulty     靶机难度    int     可选（default：1）
      6. 返回数据：无
   
   2. ##### 删除靶机
      
      1. URL：/v1/Machine/:id
      2. 提交方式：DELETE
      3. 参数格式：Params
      4. 携带token：是
      5. 提交参数：
         1. Id    要删除靶机的ID    int    必填
      6. 返回数据：无
   
   3. ##### 修改靶机
      
      1. URL：/v1/Machine
      2. 提交方式：PUT
      3. 参数格式：JSON
      4. 携带token：是
      5. 提交参数：
         1. ID    要修改的靶机id    int    必填
         2. Title    修改后靶机标题    string    可选
         3. Desc    修改后靶机描述    string    可选
         4. DownAddr    修改后网盘下载地址    string    可选
         5. DownPwd     修改网盘提取密码    string    可选
         6. Difficulty     修改靶机难度    int     可选
      6. 返回数据：无

5. ### 查询靶机
   
   1. ##### 查询靶机
      
      1. URL：/v1/Machine/:pageNum/:pageSize
      2. 携带token：否
      3. 提交方式：GET
      4. 参数格式：Params
         1. pageNum    页数    string    必选
         2. pageSize    每页数据条数    string   必选
      5. 返回数据：
         1. MachinesList    靶机数据    array
            1. ID    靶机数据id    int
            2. CreatedAt    创建时间    string
            3. UpdatedAt    修改时间    string
            4. DeletedAt    删除时间    string
            5. Title    靶机标题    string
            6. Author    靶机作者    string
            7. Desc    靶机描述    string
            8. DownAddr    靶机网盘地址    string
            9. DownPwd    靶机网盘下载密码    string
            10. Difficulty    靶机难度    int
         2. Total    靶机数据总条数    int
   
   2. ##### 查询靶机+难度
      
      1. URL：/v1/Machine/:pageNum/:pageSize/:Difficulty
      2. 提交参数
         1. pageNum    页数    string    必选
         2. pageSize    每页数据条数    string   必选
         3. Difficulty    靶机难度    int
      3. 其他与查询靶机相同
   
   3. ##### 靶机搜索
      
      1. URL：/v1/Machine/:pageNum/:pageSize/:Keys
      2. 提交参数
         1. pageNum    页数    string    必选
         2. pageSize    每页数据条数    string   必选
         3. Keys    靶机难度    string    必填（多个关键词用空格隔开，可搜索：标题，表述，作者）
      3. 其他与查询靶机相同

### 错误状态码和错误信息

```go
const(
   SUCCESS = 200
   FAIL = 400
   PERMISSION_DENIED = 401
   USERNAME_PASSWORD_ERR = 402
   LOGIN_SUCCESS = 403
   PARAMETER_ERROR = 404
   USERNAME_EXIST = 405
   USERNAME_NOT_EXIST = 406
   RECORD_NOT_EXIST = 407
   SYSTEM_ERR = 422
)

var MSG = map[int]string {
   SUCCESS: "成功",
   FAIL: "失败",
   PERMISSION_DENIED: "权限不足",
   USERNAME_PASSWORD_ERR: "账号或密码错误",
   LOGIN_SUCCESS: "登陆成功",
   PARAMETER_ERROR: "参数错误",
   USERNAME_EXIST: "用户已经存在",
   USERNAME_NOT_EXIST: "用户不存在",
   RECORD_NOT_EXIST: "该记录不存在",
   SYSTEM_ERR: "系统错误",
}
```