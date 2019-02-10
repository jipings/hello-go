// mandelbrot 函数生成一个 png 格式的 Mandelbrot 分型图
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main()  {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024,1024
	)
	img := image.NewRGBA(image.Rect(0,0,width,height))
	for py := 0; py <height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin)+xmin
			z := complex(x, y)
			// 点(px, py) 表示复数值 z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n< iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

/*
练习
3.5
用 image.NewRGBA 函数和 color.RGBA 类型或 color.YCbCr 类型实现一个 Mandelbrot 
集的全彩图
3.6
超采样 (supersampling) 通过对几个临近像素颜色值取样并取均值，是一种减少锯齿化的方法。
最简单的做法是将每个像素分成4个”子像素“。给出实现方式
3.7
3.8
3.9
*/