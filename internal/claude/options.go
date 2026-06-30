package claude

type options struct {
	systemPrompt   string
	outputFormat   string
	maxTurns       int
	tools          string
	permissionMode string
	model          string
}

func defaultOptions() *options {
	return &options{
		systemPrompt:   "",
		outputFormat:   "json",
		maxTurns:       1,
		tools:          "",
		permissionMode: "bypassPermissions",
		model:          "claude-sonnet-4-6",
	}
}

type Option func(*options)

func WithSystemPrompt(p string) Option {
	return func(o *options) { o.systemPrompt = p }
}

func WithOutputFormat(f string) Option {
	return func(o *options) { o.outputFormat = f }
}

func WithMaxTurns(t int) Option {
	return func(o *options) { o.maxTurns = t }
}

func WithAllowedTools(t string) Option {
	return func(o *options) { o.tools = t }
}

func WithModel(m string) Option {
	return func(o *options) { o.model = m }
}
