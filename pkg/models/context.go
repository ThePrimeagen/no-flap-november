package models

type Context struct {
	Screen   *Screen2
	Debug    *Debug
	Bird     *Bird
	Terminal *Terminal
	Pipes    *Pipes
	Events   *GameEventer
}

func (c *Context) Hydrate(screen *Screen2, bird *Bird, terminal *Terminal, pipes *Pipes, events *GameEventer, debug *Debug) {
	c.Debug = debug
	c.Screen = screen
	c.Bird = bird
	c.Terminal = terminal
	c.Pipes = pipes
	c.Events = events
}

func Empty() *Context {
	return &Context{}
}
