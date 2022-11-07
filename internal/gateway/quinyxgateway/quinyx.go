package quinyxgateway

type Client interface {
	CallWSDL() error
}

type Concrete struct {
}

func (c *Concrete) CallWSDL() error {
	// some logics
	return nil
}

type Mock struct {
}

func (m *Mock) CallWSDL() error {
	// some logics
	return nil
}
