
# 博客后端
>使用gin框架进行开发

## 基础业务
+ 标签管理
+ 文章管理

## 目录结构
```shell
blog-service
├─configs             # 配置文件
├─docs                # 文档集合
├─internal            # 内部模块
│  ├─dao              # 数据访问层， 所有与数据相关的操作都会在dao层进行，例如mysql、Elasticsearch
│  ├─middleware       # HTTP中间件
│  ├─model            # 模型层， 用于存放model 对象
│  ├─routers          # 路由相关的逻辑
│  └─service          # 项目核心业务逻辑
├─pkg                 # 项目相关的模块包
├─scripts             # 项目生成的临时文件
├─storage             # 各类构建、安装、分析等操作的脚本
└─third_party         # 第三方的资源工具，如swagger UI
```