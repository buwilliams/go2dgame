package main

import (
    "fmt"
    "strconv"
    "runtime"
    "log"
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
    // GLFW event handling must run on the main OS thread
    runtime.LockOSThread()
}

func main() {

    if err := glfw.Init(); err != nil {
        log.Fatalln("failed to initialize glfw:", err)
    }
    defer glfw.Terminate()

    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 2)
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    window, err := glfw.CreateWindow(800, 600, "Create Window Example", nil, nil)
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        panic(err)
    }

    window.SetKeyCallback(keypress);
    window.SetMouseButtonCallback(mousepress);

    gl.Viewport(0, 0, 800, 600);

    for !window.ShouldClose() {
        draw()
        window.SwapBuffers()
        glfw.PollEvents()
    }

}

func draw() {
    gl.ClearColor(0.2, 0.3, 0.3, 1.0);
    gl.Clear(gl.COLOR_BUFFER_BIT);
    //gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    // gl.Color4f(1, 1, 1, 1) // I have no idea what this does yet.
}

func keypress(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
    if action == glfw.Release {
        s := strconv.Itoa(scancode)
        if key == glfw.KeyA && mods == glfw.ModShift {
            fmt.Println("Key 'A' pressed!")
        } else if key == glfw.KeyA {
            fmt.Println("Key 'a' pressed!")
        } else {
            fmt.Println("Key '" + s + "' pressed!" + s)
        }
    }
}

func mousepress(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
    fmt.Println("Mouse pressed!")
}
