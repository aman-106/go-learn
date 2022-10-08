package main

import (
	"fmt"
	"runtime"
	"bufio"
	"os"
)

func cliApp()  {
	reader :=  bufio.NewReader(os.Stdin);
	fmt.Printf("this os version %v, go vsersin %v",runtime.GOOS  , runtime.Version());

	fmt.Println("name?");
	str , _ := reader.ReadString('\n'); //char delimiter
	fmt.Printf("name is %v \n",str); // str conatins delimiter


}


func argsReader()  {
	args:= os.Args;
	fmt.Println(args);

	// 
	argsForFile := args[1:] 
	// skip filename  using slice
	//  go run . someargs  -> args = [filepath , someargs ]
	fmt.Println(argsForFile);

}


func getHelp(){
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("pass /help")
	}else{
		if args[0]=="/help" {
			fmt.Println("some info here")
		}
	}
}


func userTypesSomething(){
	// scanf uses delimiter \n
}

// bufio.Scanner(os.Stdin , filename);
// real time - input - scanner

// scanner - file read 

// os.Exit(0)


// sscanf
// scanln

// log , log files

// trace logger