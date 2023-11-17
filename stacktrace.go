package gostacktrace

type StackTrace struct {
	pCounters pCounters
}

func (s *StackTrace) ToString() string {
	frames := s.GetFrames()
	if len(frames) == 0 {
		return "stack trace is empty"
	}
	return frames.ToString()
}

func (s *StackTrace) GetFrames() Frames {
	var frames Frames
	if len(s.pCounters) == 0 {
		return frames
	}

	for _, programCounter := range s.pCounters {
		frame := getFrame(programCounter)
		frames = append(frames, frame)
	}

	return frames
}
