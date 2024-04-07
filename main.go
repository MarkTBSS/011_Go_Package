package main

import (
	any_name_1 "github.com/MarkTBSS/011_Go_Package/folder_name"
)

func main() {
	//Unable to call here : local function, use in same package only
	//any_name_1.local_function_from_file_name_1()
	any_name_1.Glocal_function_from_file_name_1()

	//Unable to call here : local function, use in same package only
	//any_name_1.local_function_from_file_name_2()
	any_name_1.Glocal_function_from_file_name_2()
}
