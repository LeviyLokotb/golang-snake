package terminal

import (
	"fmt"
	"testing"
)

func TestStyles(t *testing.T) {
	fmt.Println(GetStyle("RED") + "TEXT" + GetStyle("RESET"))
}
