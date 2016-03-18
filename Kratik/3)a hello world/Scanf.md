## Reading Input

Input is Read from the Keyboard or standard input,which is os.Stdin
Simplest way to scan is


    package main
    import ("bufio"
        "fmt"
        "os"
    )
    func main(){


        fmt.Println("Enter text:2 ")
        text2 := "QWINIX"
        fmt.Scanln(text2)
        fmt.Println(text2)

        fmt.Print("Enter text:3 ")
        var input string
        fmt.Scanln(&input)
        fmt.Print(input)


        reader := bufio.NewReader(os.Stdin)

        // by default, bufio.Scanner scans newline-separated lines

        fmt.Print("Enter text:1 ")
        text,err := reader.ReadString('\n')
        fmt.Println(text)
        fmt.Println(err)


    }



* Scan scans text read from standard input, storing successive space-separated values into successive arguments. Newlines count as space.

* Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.