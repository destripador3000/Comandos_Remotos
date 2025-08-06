package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

const (
	PORT               = ":8080"
	DISCONNECT_MESSAGE = "!DESCONECTAR"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("[NUEVA CONEXIÓN] %s conectado.\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		cmdStr = strings.TrimSpace(cmdStr)
		if cmdStr == DISCONNECT_MESSAGE {
			break
		}

		fmt.Printf("[COMANDO] %s: %s\n", conn.RemoteAddr().String(), cmdStr)

		// Ejecutar el comando
		out, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()
		response := string(out)
		if err != nil {
			response += "\n[ERROR] " + err.Error()
		}

		// Enviar la respuesta
		response += "\n"
		conn.Write([]byte(response))
	}

	fmt.Printf("[DESCONECTADO] %s desconectado.\n", conn.RemoteAddr().String())
}

func main() {
	fmt.Println("[INICIANDO] Servidor Go escuchando en el puerto", PORT)
	ln, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("[ERROR] No se pudo iniciar el servidor:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[ERROR] Conexión fallida:", err)
			continue
		}
		go handleConnection(conn)
	}
}
