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



Todos:
- Cleanup map(fix ladder and platform placement/visibility)
- Delay for re-interacting interaction markers
- Dodge/Roll
- Support/refactor for multiple attacks for player
- Improve enemy "ai"
- Better callbacks system(perhaps refactor to some sort of event-emitter+queue-processor...this could help
    with proper/better division of responsibilities. For example, preventing chatHandler from triggering
    core sim callbacks despite being a "client system" which is a bit of a violation of concerns).
