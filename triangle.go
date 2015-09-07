package main

import (
    "fmt"
    "strconv"
    "runtime"
    "log"
    "github.com/go-gl/gl/v3.3-core/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
    // GLFW event handling must run on the main OS thread
    runtime.LockOSThread()
}

func main() {

    width := 1024
    height := 768
    title := "Simple Triangle"

    if err := glfw.Init(); err != nil {
        log.Fatalln("failed to initialize glfw:", err)
    }
    defer glfw.Terminate()

    glfw.WindowHint(glfw.Samples, 4); // 4x antialiasing
    glfw.WindowHint(glfw.ContextVersionMajor, 3); // We want OpenGL 3.3
    glfw.WindowHint(glfw.ContextVersionMinor, 3);
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True); // To make MacOS happy; should not be needed
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile); //We don't want the old OpenGL
    window, err := glfw.CreateWindow(width, height, title, nil, nil)
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        panic(err)
    }
    window.SetKeyCallback(keypress);
    window.SetMouseButtonCallback(mousepress);
    gl.Viewport(0, 0, 1024, 768);
    for !window.ShouldClose() {
        draw()
        window.SwapBuffers()
        glfw.PollEvents()
    }
}

func draw() {
    // TODO: draw triangle!
}

func keypress(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
    // Pressing escape will close the window
    if key == glfw.KeyEscape {
        window.SetShouldClose(true)
    }
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
