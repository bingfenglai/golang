package interfaces

// 灵根接口，实现了灵根接口的凡人即可修炼
type SpiritualRootAble interface {

	// 生成灵根
	GenSpiritualRootNames()

	// 获取生成的灵根
	SpiritualRoot() string

	// 修行方法
	Practice()
}
