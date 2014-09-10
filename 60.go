package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
  "image/color"
)
type Image struct{
  w, h int
    rgba color.RGBA
}
// ColorModelが何やってんのかわからん・・・
func (m Image)rgbaColor(c color.Color) color.Color {
  return c
}
func (m Image)ColorModel() color.Model{
    return color.ModelFunc(m.rgbaColor)
}
func (m Image)Bounds() image.Rectangle{
    return image.Rect(0, 0, m.w, m.h)
}
func (m Image)At(x, y int) color.Color{
    return m.rgba
}
func main() {
    m := Image{}
    m.h = 50
    m.w = 100
    m.rgba = color.RGBA{1, 155, 255, 55}
    pic.ShowImage(m)
}
