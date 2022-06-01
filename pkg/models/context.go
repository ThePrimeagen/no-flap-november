package models

type Context struct {
    screen *Screen2
    bird *Bird
    terminal *Terminal
    pipes *Pipes
}

func (c *Context) Hydrate(screen *Screen2, bird *Bird, terminal *Terminal, pipes *Pipes) {
    c.screen = screen
    c.bird = bird
    c.terminal = terminal
    c.pipes = pipes
}

func Empty() *Context {
    return &Context{}
}


