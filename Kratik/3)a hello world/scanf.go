package main
    import ("bufio"
        "fmt"
        "os"
    )
    func main(){


        fmt.Println("Enter text:1 ")
        text2 := "QWINIX"
        fmt.Scanln(text2)
        fmt.Println(text2)

        fmt.Print("Enter text:2 ")
        var input string
        fmt.Scanln(&input)
        fmt.Print(input)


        reader := bufio.NewReader(os.Stdin)

        // by default, bufio.Scanner scans newline-separated lines

        fmt.Print("Enter text:3 ")
        text,err := reader.ReadString('\n')
        fmt.Println(text)
        fmt.Println(err)


    }