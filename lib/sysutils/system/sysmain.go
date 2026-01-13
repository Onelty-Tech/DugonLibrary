package system

import (
	"encoding/binary"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

/*
Migrar pronto todos los errores a algo como esto
*/
type Error struct {
	Type  string
	Error error
}

// itob convierte uint64 a []byte
func Itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// converte un []byte a uint64
func Btoi(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, fmt.Errorf("Error: clave []byte no válida para btoi") // o maneja el error como prefieras
	}
	return binary.BigEndian.Uint64(b), nil
}

/*
ejecuta comandos en cualquier OS
*/
func RunCommandMatcher(cmd ...string) (string, error) {
	if len(cmd) == 0 {
		return "", fmt.Errorf("Error: No command provided")
	}

	var command *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		fullCmd := strings.Join(cmd, " ")
		command = exec.Command("cmd", "/C", fullCmd)
	default:
		command = exec.Command(cmd[0], cmd[1:]...)
	}
	output, err := command.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("Error: Could not execute command: %w", err)
	}
	return string(output), nil
}
