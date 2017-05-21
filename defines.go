package luna

// TODO 变量定义拆分到各个模块

const DefaultTimeout = 0

type BasicAuth struct {
	User     string
	Password string
}

type File struct {
	Name string
	Path string // support absolute path and relative path
}
