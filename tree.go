package silo

type Tree struct {
	T []Doc `json:"tree"`
}

type Doc struct {
	Type  string  `json:"type"`
	Label string  `json:"label,omitempty"`
	Nodes []*Node `json:"nodes"`
}

type Node struct {
	Parent string `json:"parent"`
	Type   string `json:"type"`
	Kind   string `json:"kind"`
	Unique string `json:"unique"`
	Label  string `json:"label"`
}
