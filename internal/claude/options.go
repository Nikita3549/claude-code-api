package claude

type options struct {
	systemPrompt   string
	outputFormat   string
	allowedTools   string
	permissionMode string
	model          string
	mcpConfig      string
}

func defaultOptions() *options {
	return &options{
		systemPrompt:   "",
		outputFormat:   "json",
		allowedTools:   "",
		mcpConfig:      "",
		permissionMode: "bypassPermissions",
		model:          "claude-sonnet-4-6",
	}
}

type Option func(*options)

func WithMcpConfig(c string) Option {
	return func(o *options) { o.mcpConfig = c }
}

func WithSystemPrompt(p string) Option {
	return func(o *options) { o.systemPrompt = p }
}

func WithOutputFormat(f string) Option {
	return func(o *options) { o.outputFormat = f }
}

func WithAllowedTools(t string) Option {
	return func(o *options) { o.allowedTools = t }
}

func WithModel(m string) Option {
	return func(o *options) { o.model = m }
}
