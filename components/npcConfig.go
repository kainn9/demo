package components

import "github.com/yohamta/donburi"

var NpcConfigComponent = donburi.NewComponentType[NpcConfig]()

type NpcConfig struct {
	Name NpcName
}

type NpcName string

func NewNpcConfig(name NpcName) *NpcConfig {

	return &NpcConfig{
		Name: name,
	}
}
