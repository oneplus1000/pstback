package pstback

import "path/filepath"

//Config config
type Config struct {
	Items []Item `json:"folders"`
}

//Item item
type Item struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Max  int    `json:"max"`
}

//SrcBase base of Src
func (i Item) SrcBase() string {
	return filepath.Base(i.Src)
}
