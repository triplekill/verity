package memory

type Memory struct{}

const name = "memory"

func (self *Memory) Name() string {
	return name
}
