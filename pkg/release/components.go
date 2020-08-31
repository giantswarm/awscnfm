package release

type Components struct {
	latest   map[string]string
	previous map[string]string
}

func (c Components) Latest() map[string]string {
	return c.latest
}

func (c Components) Previous() map[string]string {
	return c.previous
}
