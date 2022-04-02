package anime

import (
	"AynaAPI/utils/vjson"
	"fmt"
	"testing"
)

func TestSome(t *testing.T) {
	fmt.Println(vjson.MarshalUnescape(Anime{
		AnimeMeta: AnimeMeta{},
		Playlists: nil,
	}))
}
