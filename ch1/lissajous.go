/*
 * @Date: 2024-01-08 14:40:15
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 17:21:50
 * @FilePath: \go_learn\入门\lissajous.go
 * @Description:
 */
// Lissajous generates GIF animations of random Lissajous figures.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White,
	color.RGBA{0x58, 0xEF, 0x7D, 0xFF},
	color.RGBA{0x94,0x00,0x00, 0xFF},
	color.RGBA{0x44, 0x4B, 0xB5, 0xFF}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// 使用管道符输出到指定文件中,windows下`powershell`会转换字符导致生成的gif无法打开,可以切换到`cmd`或者程序中使用create创建文件输出
func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 练习题添加多个颜色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%3+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
