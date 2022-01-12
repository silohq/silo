package silo

type Tree struct {
	T []*Node `json:"tree"`
}

type Doc struct {
	Parent string  `json:"parent"`
	Type   string  `json:"type"`
	Nodes  []*Node `json:"nodes"`
}

type Node struct {
	Parent string  `json:"parent"`
	Type   string  `json:"type"`
	Kind   string  `json:"kind"`
	Unique bool    `json:"unique"`
	Label  string  `json:"label"`
	Nodes  []*Node `json:"nodes"`
}
