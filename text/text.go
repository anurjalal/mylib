package text

type Name struct {
	MyName string
}

func (name Name) name() string  {
	return "this is v1.0.0"
}