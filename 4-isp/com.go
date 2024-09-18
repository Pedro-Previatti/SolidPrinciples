package isp

type Reader interface {
	Read() string
}

type Writer interface {
	Write(content string)
}

type Printer interface {
	Print() string
}

type TextDocumentISP struct {
	content string
}

func (d *TextDocumentISP) Read() string {
	return d.content
}

func (d *TextDocumentISP) Write(content string) {
	d.content = content
}

func (d *TextDocumentISP) Print() string {
	return "Printing: " + d.content
}

type ReadOnlyDocumentISP struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}
