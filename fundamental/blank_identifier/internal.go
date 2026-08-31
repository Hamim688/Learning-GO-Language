package blank_identifier

import "fmt"

//! Function init() dipanggil otomatis meskipun dipanggil via Blank Identifier (_)
func init() {
	fmt.Println("Init: Executing Blank Identifier Package (Internal Package)")
}