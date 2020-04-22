package av

import "fmt"

func Version() {
	fmt.Println("libavformat: " + avformatVersion())
	fmt.Println("libavcodec: " + avcodecVersion())
}
