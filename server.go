package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println("Servidor aguardando conexões...")

	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}
	defer ln.Close()

	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer conexao.Close()

	fmt.Println("Conexão aceita...")

	for {
		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}

		fmt.Print("Mensagem recebida:", string(mensagem))

		novamensagem := strings.ToUpper(mensagem)

		conexao.Write([]byte(novamensagem + "\n"))
	}
}
