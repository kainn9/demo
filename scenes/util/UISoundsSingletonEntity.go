package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	"github.com/kainn9/demo/tags"
	"github.com/yohamta/donburi"
)

func AddUISoundsSingletonEntity(scene *coldBrew.Scene) {
	singletonSoundEntity := AddUISoundsSingletonEntityWithoutContext(scene)

	// Initialize AudioContext.
	components.AudioContextComponent.SetValue(singletonSoundEntity, *components.NewAudioContext())

}

func AddUISoundsSingletonEntityWithoutContext(scene *coldBrew.Scene) *donburi.Entry {
	// Initialize UI Singleton SpritesMap.
	singletonSoundEntity := scene.AddEntity(
		components.SoundsMapComponent,
		components.AudioContextComponent,
		tags.UISingletonSoundsTag,
	)

	sounds := make(map[string]*components.Sound)

	sounds[UIGlobals.CHAT_BOX_NEW_SOUND_NAME] = components.NewSound(-1, 1)

	components.SoundsMapComponent.SetValue(singletonSoundEntity, sounds)

	return singletonSoundEntity
}
