# demo(WIP)
A Metroidvania-lite/Platformer game written in Go.


# To Run:
- When using mac: 
    ```bash
        EBITENGINE_GRAPHICS_LIBRARY="opengl" go run .
    ```
- When using windows: 
    ```bash
    go run.
    ```
- Notes: 
    - Using `opengl` option for mac resolves issue where invoking `ebiten.SetFullscreen(true)` results in a lower frame rate then when setting fullscreen via mac UI.
    
    - Also resolves issue where fps decreases when running OBS while game is in fullscreen too.



### Random List of TODOs:
- Scene transitions(fade in/out + sound)
- Better simple AI system
- Movement system cleanup
- Sound system cleanup
- Better entity Factories
- Fix duplicate sprite asset loading for chat sprites
- Organize components
- Proper intro Map(s)
- Delay for re-interacting interaction markers
- Better custom callback system
