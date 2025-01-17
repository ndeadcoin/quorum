package chain

//"fmt"

type Molasses struct {
	name     string
	producer Producer
	user     User
}

func NewMolasses(p Producer, u User) *Molasses {
	return &Molasses{name: "Molasses", producer: p, user: u}
}

func (m *Molasses) Name() string {
	return m.name
}

func (m *Molasses) Producer() Producer {
	return m.producer
}

func (m *Molasses) User() User {
	return m.user
}

func (m *Molasses) SetProducer(p Producer) {
	m.producer = p
}

func (m *Molasses) SetUser(u User) {
	m.user = u
}
