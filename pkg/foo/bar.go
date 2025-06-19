package foo

type FooBar struct {
	Content string
}

func New(c string) FooBar {
	return FooBar{
		Content: c,
	}
}
