package main

import (
	"Belajar_Golang/fundamental"
	"Belajar_Golang/fundamental/access_modifier"
	_ "Belajar_Golang/fundamental/blank_identifier"
	"Belajar_Golang/fundamental/initialization"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=======================================================")
		fmt.Println("       🚀 GOLANG FUNDAMENTAL LEARNING APP 🚀          ")
		fmt.Println("=======================================================")
		fmt.Println(" 1. Hello World                 17. Recover Function")
		fmt.Println(" 2. Data Type                   18. Struct & Struct Method")
		fmt.Println(" 3. Variable                    19. Interface")
		fmt.Println(" 4. Constant                    20. Any (Interface Kosong)")
		fmt.Println(" 5. Data Type Conversion        21. Nil Function")
		fmt.Println(" 6. Type Declaration            22. Type Assertions")
		fmt.Println(" 7. Mathematics                 23. Pointer")
		fmt.Println(" 8. Array                       24. Asterisk Operator (*)")
		fmt.Println(" 9. Slice                       25. New Function")
		fmt.Println("10. Map                         26. Pointer Functions")
		fmt.Println("11. If Else                     27. Pointer Method")
		fmt.Println("12. Switch Expression           28. Access Modifier")
		fmt.Println("13. For Loops                   29. Package Initialization")
		fmt.Println("14. Function (All Practice)     30. Error Functions")
		fmt.Println("15. Closure                     31. Custom Error")
		fmt.Println("16. Defer & Panic Function")
		fmt.Println(" 0. Keluar Aplikasi")
		fmt.Println("=======================================================")
		fmt.Print("Pilih nomor materi yang ingin dieksekusi [0-31]: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		pilihan, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("❌ Error: Harap masukkan angka yang valid!")
			continue
		}

		if pilihan == 0 {
			fmt.Println("\nTerima kasih sudah belajar Golang! Sampai jumpa! 👋✨")
			break
		}

		fmt.Printf("\n-------------------------------------------------------\n")
		switch pilihan {
		case 1:
			fmt.Println("===== 1. Executing Hello World =====")
			fundamental.HelloWorld()
		case 2:
			fmt.Println("===== 2. Executing Data Type =====")
			fundamental.DataType()
		case 3:
			fmt.Println("===== 3. Executing Variable =====")
			fundamental.Variable()
		case 4:
			fmt.Println("===== 4. Executing Constant =====")
			fundamental.Constant()
		case 5:
			fmt.Println("===== 5. Executing Data Type Conversion =====")
			fundamental.DataTypeConversion()
		case 6:
			fmt.Println("===== 6. Executing Type Declaration =====")
			fundamental.TypeDeclaration()
		case 7:
			fmt.Println("===== 7. Executing Mathematics =====")
			fundamental.Mathematics()
		case 8:
			fmt.Println("===== 8. Executing Array =====")
			fundamental.Array()
		case 9:
			fmt.Println("===== 9. Executing Slice =====")
			fundamental.Slice()
		case 10:
			fmt.Println("===== 10. Executing Map =====")
			fundamental.Map()
		case 11:
			fmt.Println("===== 11. Executing If Else =====")
			fundamental.IfElse()
		case 12:
			fmt.Println("===== 12. Executing Switch Expression =====")
			fundamental.SwitchExpression()
		case 13:
			fmt.Println("===== 13. Executing For Loops =====")
			fundamental.ForLoops()
		case 14:
			fmt.Println("===== 14. Executing Function (All Practice) =====")
			fundamental.Function()
		case 15:
			fmt.Println("===== 15. Executing Closure =====")
			fundamental.Closure()
		case 16:
			fmt.Println("===== 16. Executing Defer & Panic Function =====")
			fundamental.DeferFunction()
			fmt.Println()
			fundamental.PanicFunction()
		case 17:
			fmt.Println("===== 17. Executing Recover Function =====")
			fundamental.RecoverFunction()
		case 18:
			fmt.Println("===== 18. Executing Struct & Struct Method =====")
			wonYoung := fundamental.Struct()
			wonYoung.SayHello("Sayangku")
		case 19:
			fmt.Println("===== 19. Executing Interface =====")
			fundamental.Interface()
		case 20:
			fmt.Println("===== 20. Executing Any (Interface Kosong) =====")
			fundamental.AnyFunction()
		case 21:
			fmt.Println("===== 21. Executing Nil Function =====")
			fundamental.NilFunction()
		case 22:
			fmt.Println("===== 22. Executing Type Assertions =====")
			fundamental.TypeAssertions()
		case 23:
			fmt.Println("===== 23. Executing Pointer =====")
			fundamental.Pointer()
		case 24:
			fmt.Println("===== 24. Executing Asterisk Operator =====")
			fundamental.Asterisk()
		case 25:
			fmt.Println("===== 25. Executing New Function =====")
			fundamental.NewFunction()
		case 26:
			fmt.Println("===== 26. Executing Pointer Functions =====")
			fundamental.PointerFunctions()
		case 27:
			fmt.Println("===== 27. Executing Pointer Method =====")
			fundamental.PointerMethod()
		case 28:
			fmt.Println("===== 28. Executing Access Modifier =====")
			fmt.Println("SayGoodBye  :", access_modifier.SayGoodBye("Jang Won-Young"))
			fmt.Println("Application :", access_modifier.Application)
		case 29:
			fmt.Println("===== 29. Executing Package Initialization =====")
			fmt.Println("Database Connection :", initialization.GetDatabase())
		case 30:
			fmt.Println("===== 30. Executing Error Functions =====")
			fundamental.ErrorFunctions()
		case 31:
			fmt.Println("===== 31. Executing Custom Error =====")
			fundamental.ErrorCustom()
		default:
			fmt.Println("❌ Nomor materi tidak ditemukan!")
		}
		fmt.Println("-------------------------------------------------------")

		fmt.Print("\nTekan ENTER untuk kembali ke menu utama...")
		reader.ReadString('\n')
	}
}
