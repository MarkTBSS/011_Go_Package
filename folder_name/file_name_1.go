package any_name_1

import "fmt"

func local_function_from_file_name_1() {
	fmt.Println("Hello from local_function_from_file_name_1")
}

func Glocal_function_from_file_name_1() {
	local_function_from_file_name_1()
	fmt.Println("Hello from Glocal_function_from_file_name_1")
}
