package silo

type Tree struct {
	Doc Document `json:"document"`
}

type Document struct {
	Typ   string  `json:"type,omitempty"`
	Label string  `json:"label,omitempty"`
	Nodes []*Node `json:"nodes,omitempty"`
}

type Node struct {
	Parent string `json:"parent"`
	Kind   string `json:"kind,omitempty"`
	Unique string `json:"unique,omitempty"`
	Typ    string `json:"typ,omitempty"`
	Label  string `json:"label,omitempty"`
}
