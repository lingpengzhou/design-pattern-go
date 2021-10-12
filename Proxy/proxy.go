package Proxy

import (
	"debug/dwarf"
	"log"
)

type Image interface {
	display() dwarf.VoidType
}

type RealImage struct {
	fileName string
}

func (realImage *RealImage) LoadFromDisk() {
	log.Println("从磁盘获取图片信息" + realImage.fileName)
}

func (realImage *RealImage) RealImageInit(fileName string) {
	realImage.fileName = fileName
	realImage.LoadFromDisk()
}

func (realImage *RealImage) display() {
	log.Println("从磁盘展示图片信息" + realImage.fileName)
}

type ImageProxy struct {
	fileName  string
	realImage RealImage
}

func (imageProxy *ImageProxy) ImageProxyInit(fileName string) {
	imageProxy.fileName = fileName
}

func (imageProxy *ImageProxy) display() {
	if imageProxy.realImage.fileName == "" {
		imageProxy.realImage.RealImageInit(imageProxy.fileName)
	}
	imageProxy.realImage.display()
}
