package adapter

type IconsumerUseCase interface {
	Run(topicName string, action string)
}
