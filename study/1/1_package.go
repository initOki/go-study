// 모든 Go 프로그램은 패키지로 구성되어 있음
package main

// import 문은 현재 프로그램에서 사용할 패키지를 지정
// 여기서는 "fmt"와 "math/rand" 패키지를 사용
import (
	"fmt"
	"math/rand"
)

func main() {
	// 패키지 내부의 함수를 사용하려면 패키지 이름을 함수 이름 앞에 붙여야 함
	// 여기서는 "fmt.Println"과 "rand.Intn"을 사용
	// rand.Intn은 0과 50 사이의 난수를 생성
	fmt.Println("내가 좋아하는 숫자는????", rand.Intn(50))
}
