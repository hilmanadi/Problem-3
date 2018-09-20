package main

import "io/ioutil"
import "fmt"

func main() {
    files1,_ := ioutil.ReadDir("/mnt/c/Users/ASUS/go/src/problem2/source/folder1")
	files2,_ := ioutil.ReadDir("/mnt/c/Users/ASUS/go/src/problem2/target/folder1")
	files3,_ := ioutil.ReadDir("/mnt/c/Users/ASUS/go/src/problem2/source/folder1/subfolder1")
	files4,_ := ioutil.ReadDir("/mnt/c/Users/ASUS/go/src/problem2/target/folder1/subfolder1")
	files5,_ := ioutil.ReadDir("/mnt/c/Users/ASUS/go/src/problem2/source/folder2")
	files6,_ := ioutil.ReadDir("/mnt/c/Users/ASUS/go/src/problem2/target/folder2")
	fmt.Println("======folder1======")
	if len(files1) == len(files2){
		fmt.Println("SAME")
	}else if len(files1) > len(files2){
		fmt.Println("NEW")
	}else{
		fmt.Println("DELETED")
	}
	fmt.Println("======subfolder1======")
	if len(files3) == len(files4){
		fmt.Println("SAME")
	}else if len(files3) > len(files4){
		fmt.Println("NEW")
	}else{
		fmt.Println("DELETED")
	}
	fmt.Println("======folder2======")
	if len(files5) == len(files6){
		fmt.Println("SAME")
	}else if len(files5) > len(files6){
		fmt.Println("NEW")
	}else{
		fmt.Println("DELETED")
	}
	// if len(files1)>len(files2){
	// 	fmt.Println("File on Target ")
	// }
}