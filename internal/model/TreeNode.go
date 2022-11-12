package model

type TreeNode struct {
	Id         int64             `json:"id"`
	Label      string            `json:"label"`
	Depth      int               `json:"depth"`
	State      string            `d:"closed" json:"state"`
	Attributes map[string]string `json:"attributes"`
	Children   []TreeNode        `json:"children"`
}
