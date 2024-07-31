package entitie

type ConsumerEntitie struct {
	Id     int    `json:"Id, omitempty"`
	Name   string `json:"Name, omitempty"`
	Host   string `json:"Host, omitempty"`
	Port   string `json:"Port, omitempty"`
	Deelay int    `json:"Deelay, omitempty"`
	Status int    `json:"Status, omitempty"`
}
