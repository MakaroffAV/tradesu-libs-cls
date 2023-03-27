package cls

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMe(t *testing.T) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}
