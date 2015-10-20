package main

import (
	"log"
	_ "time"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	_ "github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

func drawTimeAxis(canvas gxui.Canvas) {
	p := gxui.Polygon{
		gxui.PolygonVertex{
			Position: math.Point{
				X: 0,
				Y: 400,
			},
			RoundedRadius: 0,
		},
		gxui.PolygonVertex{
			Position: math.Point{
				X: 400,
				Y: 400,
			},
			RoundedRadius: 0,
		},
	}

	canvas.DrawLines(p, gxui.CreatePen(1.0, gxui.White))
}

func drawPlot(canvas gxui.Canvas) {
	drawTimeAxis(canvas)
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	canvas := driver.CreateCanvas(math.Size{W: 500, H: 300})

	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)
	layout.SetDirection(gxui.TopToBottom)

	buttonsLayout := theme.CreateLinearLayout()
	buttonsLayout.SetSizeMode(gxui.Fill)
	buttonsLayout.SetDirection(gxui.LeftToRight)

	button := func(name string, action func()) gxui.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action() })
		return b
	}

	okayButton := button("Okay", func() { log.Println("Okay") })
	buttonsLayout.AddChild(okayButton)
	cancelButton := button("Cancel", func() { log.Println("Cancel") })
	buttonsLayout.AddChild(cancelButton)

	drawPlot(canvas)
	canvas.Complete()

	image := theme.CreateImage()
	image.SetCanvas(canvas)

	window := theme.CreateWindow(800, 600, "bview")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))
	layout.AddChild(buttonsLayout)
	layout.AddChild(image)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})

	window.OnResize(func() { log.Println(layout.Children().String()) })

}

func main() {
	gl.StartDriver(appMain)
}
