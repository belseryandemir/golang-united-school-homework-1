package solution

import "github.com/kyokomi/emoji"

func GetMessage() string {
	myvar := emoji.Sprint("Hello :world_map:!")
	return myvar
}
