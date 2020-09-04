package release

type Version struct {
	latest   string
	previous string
}

func (v Version) Latest() string {
	return v.latest
}

func (v Version) Previous() string {
	return v.previous
}
