package Command

// コマンドを入れる構造体
type Context struct {
	route   string
	options []string
}

// コンストラクタ
func New(args []string) *Context {
	if len(args) == 1 {
		return nil
	}

	return &Context{route: args[1], options: args[2:]}
}

// Getter: string Route
func (c *Context) GetRouter() string {
	return c.route
}

// Getter: string[] Args
func (c *Context) GetOptions() []string {
	return c.options
}

// Getter: string Arg
func (c *Context) GetOption(index int) string {
	return c.options[index]
}

// Getter: int length.args
func (c *Context) OptionCount() int {
	return len(c.options)
}

// ブランチ名を取得します。
func (c *Context) GetBranchName() string {
	return c.options[0]
}
