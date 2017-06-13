/* bufio 包的使用 */
package main

import (
	"fmt"
	"bufio"
)

func main() {
	readbuf := strings.NewReader("Hello, Dear World")
	reader := bufio.NewReader(readbuf)
}