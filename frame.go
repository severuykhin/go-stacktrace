package gostacktrace

import "fmt"

type Frame struct {
	File string `json:"file"`
	Func string `json:"func"`
	Line uint   `json:"line"`
}

func (f Frame) ToString() string {
	return fmt.Sprintf("%s %s %d", f.File, f.Func, f.Line)
}

type Frames []Frame

func (f Frames) ToString() string {
	if len(f) == 0 {
		return ""
	}

	var res string

	for i, frame := range f {
		if i == 0 {
			res = fmt.Sprintf("%d.%s", i, frame.ToString())
		} else {
			res = fmt.Sprintf("%s | %d.%s", res, i, frame.ToString())
		}
	}

	return res
}
