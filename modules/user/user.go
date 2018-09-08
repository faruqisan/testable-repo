package user

type NSQ interface {
	Publish(data string) error
}

type Redis interface {
	Set(key string, data interface{})
	Get(key string)
}

type Module struct {
	nsq   NSQ
	redis Redis
}

type Data struct {
	Name string
	Age  int
}

func New(nsq NSQ, redis Redis) *Module {
	return &Module{
		nsq:   nsq,
		redis: redis,
	}
}

func (m *Module) Get() ([]Data, error) {

	// assume we publish data when user's Data called

	err := m.nsq.Publish("please notify another service to generate something !")
	if err != nil {
		return nil, err
	}

	return []Data{
		{"John", 20},
		{"Dow", 31},
	}, err
}

func (m *Module) Create(name string, age int) error {

	var user Data

	user.Name = name
	user.Age = age

	// assume when create we store it to redis and publish something to nsq
	m.redis.Set("key", user)
	return m.nsq.Publish("some data")
}
