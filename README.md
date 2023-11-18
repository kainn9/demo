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
- Delay for re-interacting interaction markers
- Emphasis font options for dialogue
- Dodge/Roll
- "Better" enemy "ai"
