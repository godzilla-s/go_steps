/* bufio 包实现有缓冲的IO读写 */
package main

import (
	"fmt"
	"bufio"
)

func main() {
	readbuf := strings.NewReader("Hello, Dear World")
	reader := bufio.NewReader(readbuf)

	
}
