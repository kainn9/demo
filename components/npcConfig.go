package components

import "github.com/yohamta/donburi"

type NpcName string

type NpcConfig struct {
	Name NpcName
}

var NpcConfigComponent = donburi.NewComponentType[NpcConfig]()

func NewNpcConfig(name NpcName) *NpcConfig {

	return &NpcConfig{
		Name: name,
	}
}
