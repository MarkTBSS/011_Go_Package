//Same Folder Same Package
//package any_name_2

package any_name_1

import "fmt"

func local_function_from_file_name_2() {
	fmt.Println("Hello from local_function_from_file_name_2")
}

func Glocal_function_from_file_name_2() {
	local_function_from_file_name_2()
	fmt.Println("Hello from Glocal_function_from_file_name_2")
}
