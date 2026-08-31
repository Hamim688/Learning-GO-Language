package access_modifier

//! Private Variable (Huruf Kecil) - Hanya bisa diakses di dalam package access_modifier
var version = "1.0.0"

//! Public Variable (Huruf Besar) - Bisa diakses dari luar package
var Application = "Golang"

//! Private Function (Huruf Kecil) - Tidak dapat diakses dari luar package
func sayHello(name string) string {
	return "Hello " + name
}

//! Public Function (Huruf Besar) - Dapat diakses dari luar package
func SayGoodBye(name string) string {
	return "Good Bye " + name + " (v" + version + ")"
}
