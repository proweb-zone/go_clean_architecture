package entitie

type ListenerEntitie struct {
	Id     int    `json:"Id, omitempty"`
	Name   string `json:"Name, omitempty"`
	Host   string `json:"Host, omitempty"`
	Port   string `json:"Port, omitempty"`
	Deelay string `json:"Deelay, omitempty"`
	Status string `json:"Status, omitempty"`
}

// type Settings struct {
// 	Host string `json:"Host, omitempty"`
// 	Port string `json:"port, omitempty"`
// }
