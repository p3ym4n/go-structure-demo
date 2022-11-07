package riderprofilegateway

type Client interface {
	CreateRider(name, email, phone string) error
}

type Concrete struct {
}

func (c *Concrete) CreateRider(name, email, phone string) error {
	// some logics
	return nil
}

type Mock struct {
}

func (m *Mock) CreateRider(name, email, phone string) error {
	// some logics
	return nil
}
