package isp

type Document interface {
	Read() string
	Write(content string)
	Print() string
}

type TextDocument struct {
	content string
}

func (d *TextDocument) Read() string {
	return d.content
}

func (d *TextDocument) Write(content string) {
	d.content = content
}

func (d *TextDocument) Print() string {
	return "Printing: " + d.content
}

type ReadOnlyDocument struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}

func (d *ReadOnlyDocument) Write(content string) {
	// Not supported
}

func (d *ReadOnlyDocument) Print() string {
	// Not supported
	return ""
}
