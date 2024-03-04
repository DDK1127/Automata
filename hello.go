package main
	import "fmt"

	func main(){
		fmt.Printf("Hello world!, 你好，世界,こんにちはせかい\n")

		ddk, wall, bell := "董承恩", "少華", "鍾哥"
		eric := "eric"
		fmt.Println(ddk, wall, bell, eric)

		const Pi float64 = 3.1415926
		const c complex64 = 5+5i
		const prefix = "ggelq_..."

		var is_active bool // Default is False.
		is_active = true
		fmt.Println("now is active?\nANS is = ", prefix, is_active)
	}