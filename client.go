package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	conexao, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}
	defer conexao.Close()

	for {
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("texto a ser enviado: ")
		texto, erro2 := leitor.ReadString('\n')
		if erro2 != nil {
			fmt.Println(erro2)
			os.Exit(3)
		}

		fmt.Fprintf(conexao, texto+"\n")

		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}

		fmt.Print("Resposta do servidor: " + mensagem)
	}
}
