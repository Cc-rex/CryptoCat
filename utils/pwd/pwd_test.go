package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("chen0211"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$Z42xE9Ar2cWS6MCKUtAvJOFqeh8/kF4DwHOhEhhHihAWr.HgzGrJa", "chen0211"))
}
