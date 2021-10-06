package main

import (
	_ "embed"
	"utils/pkg/app"

	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
  stats:=&app.MyStruct{}
  
  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "utils",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(stats)
  app.Run()
}
