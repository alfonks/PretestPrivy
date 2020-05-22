package main

import(
	"fmt"
)


/*contoh sample input :
1. {[{[]}]}
2. iasundfasn
3. {[]
4. 65123!#$1
5. {[[asfdmnk]]}
*/

func checkString(kata string){ //Membuat function untuk mengecek string
	squareleft, squereright, curlyleft, curlyright := 0,0,0,0

	for i := range kata {
		if string(kata[i]) != "[" && string(kata[i]) != "]" && string(kata[i]) != "{" && string(kata[i]) != "}"   {
			fmt.Println("nil returned") //jika bukan karakter yang ditentukan akan mengembalikan nilai nil
			return
		} else { //jika iya, akan dihitung jumlah kurungnya
			if string(kata[i]) == "[" {
				squareleft++
			} else if string(kata[i]) == "]" {
				squereright++
			} else if string(kata[i]) == "{" {
				curlyleft++
			} else if string(kata[i]) == "}" {
				curlyright++
			}
		} 
	}

	if squareleft != squereright || curlyleft != curlyright {
		fmt.Println("false") //jika jumlah kurung buka dan kurung tutup tidak sama maka akan mengembalikan nilai false
	} else {
		fmt.Println("true") //jika jumlah kurung buka dan tutup sama maka akan mengembalikan nilai true
	}
}

func main() {
	var kata string
	for {
		fmt.Print("Masukan kata: ")
		fmt.Scanf("%s\n", &kata)

		if kata == "exit" { //jika mengetikan exit pada masukan kata maka akan keluar dari loop
			fmt.Println("Goodbye!")
			break
		}

		checkString(kata) //mengirim string yang baru saja di scan kedalam fuction sebagai parameter
		fmt.Println("")
	}
	
}