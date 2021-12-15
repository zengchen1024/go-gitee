package gitee

func (l *LabelHook) GetID() int32 {
	if l == nil {
		return 0
	}

	return l.Id
}

func (l *LabelHook) GetName() string {
	if l == nil {
		return ""
	}

	return l.Name
}

func (l *LabelHook) GetColor() string {
	if l == nil {
		return ""
	}

	return l.Color
}
