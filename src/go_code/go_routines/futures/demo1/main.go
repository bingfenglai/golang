package main

func main() {

	file := receiverImagesFile()
	resolve(file)

}

func receiverImagesFile() string {
	return "1.png"
}

func resolve(s string) {
	println("对图片 " + s + " 进行特征提取")
}
