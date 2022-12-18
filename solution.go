package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	mystr := emoji.Sprint("Hello :world_map:!")
	return mystr
}
