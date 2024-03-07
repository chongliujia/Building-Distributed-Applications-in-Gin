# To-do List Application

Web Application

## 项目结构
```
todo-list-app/
│
├── cmd/                   # 应用程序的入口点
│   └── server/            # 主服务应用
│       └── main.go        # 主程序启动入口
│
├── internal/              # 私有应用程序和库代码
│   ├── api/               # Web API层，处理HTTP请求
│   │   └── handler.go     # HTTP处理函数（例如添加、获取Todo项）
│   │
│   ├── model/             # 应用程序的数据模型
│   │   └── model.go       # 定义Todo项的结构体等
│   │
│   └── store/             # 数据存储逻辑
│       └── store.go       # 处理数据存储（例如内存中或数据库）
│
├── pkg/                   # 外部应用程序可以使用的库代码
│
├── test/                  # 额外的外部测试（可选）
│
└── go.mod                 # 定义模块的依赖
└── go.sum                 # 锁定依赖模块的版本

```

## 说明

- cmd/ : 包含项目的入口点。服务启动代码位于 cmd/server/main.go
- internal/: 包含私有的应用代码，该目录下的代码还能被项目内部其他代码导入。这里面包括业务逻辑，API处理，数据模型
	- api/: 包含处理HTTP请求的逻辑
	-  model/: 定义了应用数据的结构体和类型 	
	-  store/: 包含数据存取逻辑，比如保存Todo项
- pkg/: 包含可以被外部应用导入的库代码
- test/: 用于存放外部测试代码，这些代码可以是集成测试或者端到端测试
- go.mod 和 go.sum 定义了项目的依赖关系。go.mod 显示了项目的模块路径和依赖模块。go.sum 包含了每个依赖的预期加密哈希，用于确保这些依赖未被篡改
 
## 计划
-  [x] 创建项目
-  [ ] 数据库支持
-  [ ] 