package Proxy

import (
	"log"
	"testing"
)

func Test_Proxy(t *testing.T) {
	image := ImageProxy{fileName: "test_10mb.jpg"}
	image.display()
	log.Println()
	image.display()
}
