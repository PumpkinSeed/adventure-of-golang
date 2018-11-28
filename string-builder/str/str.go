package str

type Builder string

func (b *Builder) Write(str string) {
	*b += Builder(str)
}

func (b *Builder) WriteByte(str []byte) {
	*b += Builder(str)
}

func (b *Builder) String() string {
	return string(*b)
}
