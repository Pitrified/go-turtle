# go-turtle

[Turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics)
in Go.

## Turtle

A minimal Turtle agent, moving on a cartesian plane.

The orientation is in degrees.
`Right` rotates clockwise, `Left` counter-clockwise.

Use it to simulate the movements of the turtle without the drawing overhead.

```go
// create a new turtle
t := turtle.New()

// move it just like you expect
t.Forward(5)
t.Left(45)
t.Forward(5)
t.Right(45)
t.Backward(5)

// get the X, Y, Deg data as needed
fmt.Println(t.X, t.Y, t.Deg)
// 3.5355339059327378 3.5355339059327373 0

// teleport around
t.SetPos(4, 4)
t.SetHeading(120)

// vaguely nice printing
fmt.Println("T:", t)
// T: (4.000000, 4.000000) ^ 120.000000
```

## TurtleDraw

Has the same interface of `Turtle`.
Each `TurtleDraw` is attached to a `World`.

Create a new world to draw in:

```go
w := turtle.NewWorld(900, 600)
```

When creating a `World` with `NewWorld`,
an uniform image of the requested size `(width, height)`
with `SoftBlack` background is generated.

An existing image can be used as base with:

```go
img := image.NewRGBA(image.Rect(0, 0, 900, 600))
draw.Draw(img, img.Bounds(), &image.Uniform{turtle.Cyan},
    image.Point{0, 0}, draw.Src)
wi := turtle.NewWorldImage(img)
```

Create a `TurtleDraw` attached to the `World`:

```go
// create a turtle attached to w
td := turtle.NewTurtleDraw(w)

// position/orientation
td.SetPos(100, 300)
td.SetHeading(turtle.North + 80)

// line style
td.SetColor(turtle.Blue)
td.SetSize(4)

// start drawing
td.PenDown()

// same interface as Turtle
td.Forward(100)
td.Left(160)
td.Forward(100)
```

Save the current image:

```go
err := w.SaveImage("world.png")
if err != nil {
    fmt.Println("Could not save the image: ", err)
}
```

Close the world (there are two open internal channels).

```go
w.Close()

// this is an error: the turtle tries to send the line
// to the world input channel that has been closed
// td.Forward(50)
```

You can create as many turtles as you want.
When drawing, a turtle sends the line to the world on a channel
and blocks until it is done.

## Samples

A few samples are in the
[samples](./samples/draw)
folder.

![sample world](samples/draw/world.png)

## Constants

A few standard colors:

```go
Black     = color.RGBA{0, 0, 0, 255}
SoftBlack = color.RGBA{10, 10, 10, 255}
White     = color.RGBA{255, 255, 255, 255}

Red   = color.RGBA{255, 0, 0, 255}
Green = color.RGBA{0, 255, 0, 255}
Blue  = color.RGBA{0, 0, 255, 255}

Cyan    = color.RGBA{0, 255, 255, 255}
Magenta = color.RGBA{255, 0, 255, 255}
Yellow  = color.RGBA{255, 255, 0, 255}
```

Cardinal directions:

```go
East  = 0.0
North = 90.0
West  = 180.0
South = 270.0
```

## Implementation notes

### Note on float64

A lot of inputs to the API are actually `float64`, so when using a variable
instead of an untyped const the compiler will complain if the var is `int`.

So this works as the var has type float:

```go
segLen := 150.0
td.Forward(segLen)
```

and this does not:

```go
segLen := 150
td.Forward(segLen)
// cannot use segLen (variable of type int) as float64 value
// in argument to td.Forward (compile)
```

This works magically because in Go constants are
[neat](https://blog.golang.org/constants).

```go
td.Forward(150)
```

### Drawing pixels

When drawing points of odd size, the square is centered on the position.
When drawing points of even size, 
a pixel from the left and lower side of the square is removed.

![drawing pixels of even size](samples/draw/pixels.png)

To draw a single point, just call forward with 0 dist:

```go
td.Forward(0)
```

### Channel closing (and line drawing)

The world draws the `Line` it receives on the `DrawLineCh` channel,
so you can technically draw a line directly with that.

## TODO - Ideas

* Hilbert sample!
* More colors!
