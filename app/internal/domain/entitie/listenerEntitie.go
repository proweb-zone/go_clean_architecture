package entitie

type ListenerEntitie struct {
	Name string `json:"Name"`
	Settings Settings `json:"Settings"`
}

type Settings struct {
	Host string `json:"Host"`
	Port string `json:"port"`
}
