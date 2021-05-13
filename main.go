package main

import (
	"bufio"
	"fmt"
	"net/http"
    "strconv"
    "os"
)

var clear map[string]func() //create a map for storing clear funcs


func banner() {
	asciiget, err := http.Get("https://gist.githubusercontent.com/belbomemo/b5e7dad10fa567a5fe8a/raw/4ed0c8a82a8d1b836e2de16a597afca714a36606/gistfile1.txt")

	if err != nil {
        panic(err)
    }
	defer asciiget.Body.Close()

	ascii := bufio.NewScanner(asciiget.Body)
    for i := 0; ascii.Scan() && i < 5; i++ {
        fmt.Println(ascii.Text())
    }

	
    if err := ascii.Err(); err != nil {
        panic(err)
    }

}
func main() {
	
	banner();
	var i string
	fmt.Println("\n   Go Lang Data Organizator :D   \n ")
	fmt.Println("[1] Send Data          [2] average ")
    fmt.Println("[3] median             [4] exit ")
    fmt.Print("Operation:")
    fmt.Scanf("%s", &i)
	switch i {
    case "1":
        fmt.Println("Option 1 selected")
        data_save();    
    case "2":
        fmt.Println("Option 1 selected")
    case "3":
        fmt.Println("Option 1 selected")
    case "4":
        exit()
    }
}

func data_save() {
    fmt.Print("\033[H\033[2J")
    banner();

    var onenum string
    fmt.Print("\n1 number:")
    fmt.Scanf("\n%s", &onenum)
    one, _ := strconv.ParseFloat(onenum, 64)
    fmt.Print("\033[H\033[2J")
    banner();
    
    var twonum string
    fmt.Print("\n2 number:")
    fmt.Scanf("\n%s", &twonum)
    two, _ := strconv.ParseFloat(twonum, 64)
    fmt.Print("\033[H\033[2J")
    banner();

    var treenum string
    fmt.Print("\n3 number:")
    fmt.Scanf("\n%s", &treenum)
    tree, _ := strconv.ParseFloat(treenum, 64)
    fmt.Print("\033[H\033[2J")
    banner();

    var fournum string
    fmt.Print("\n4 number:")
    fmt.Scanf("\n%s", &fournum)
    four, _ := strconv.ParseFloat(fournum, 64)
    fmt.Print("\033[H\033[2J")
    banner();

    var  fivenum string
    fmt.Print("\n5 number:")
    fmt.Scanf("\n%s", &fivenum)
    five, _ := strconv.ParseFloat(fivenum, 64)
    fmt.Print("\033[H\033[2J")
    banner();

    data := [5]float64{one, two, tree, four, five}
    fmt.Println("\nsaved number", onenum, twonum, treenum, fournum, fivenum, "\n", "")
    fmt.Println(data)
    main()
}

func average() {
	banner();
    main()
}
func median() {
}

func exit() {
	banner();
    fmt.Println("\n   Good Bye gohper :)   \n ")
    os.Exit(1)
}
