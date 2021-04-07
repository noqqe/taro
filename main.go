package main

import (
	"github.com/noqqe/taro/cmd"
)

func main() {
	// cmd.List()
	name := cmd.Add()
	cmd.UploadToS3(name)

}
