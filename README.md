# go2dgame

Experiments with [go-gl](https://github.com/go-gl) and [Go](http://golang.org) game development.

### introduction
After watching the [talk from GopherCon](https://www.youtube.com/watch?v=aiv1JOfMjm0)
about improvements with GC in Go 1.5, I become interested in making games in Go.

This project will contain several experiments using go-gl to make 2D games. Since
I am new to OpenGL, go-gl, and game development in general, this may not
be a great resource for your own learning but maybe you'll find helpful code
examples.

I'm grateful to the people who made go-gl and OpenGL. I am impressed with the
amount of work people have contributed to these efforts over the years.

### getting started
- `go get -u github.com/buwilliams/go2dgame`

### learning

- go-gl package has an examples folder which was helpful in getting my first glfw
window created.
- the glfw code reads easily so I was able to begin to understand the input
  subsystem
- because go-gl is [calling C code within Go](https://golang.org/cmd/cgo/) to do opengl and glfw, the
  documentation/tutorials for those libraries helps in writing/understanding the Go code.
- I found [learnopengl.com](http://learnopengl.com) as a helpful resource for
  initially [understanding why glfw is needed](http://learnopengl.com/#!Getting-started/Creating-a-window).
- Now following a [different tutorial](http://www.opengl-tutorial.org) which uses OpenGL 3.3 (and is a bit easier
  to follow)
