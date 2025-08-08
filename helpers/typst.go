package helpers

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func RenderTypst(code string) (*discordgo.File, error) {
	cmd := exec.Command("typst", "compile", "-", "-f", "png", "-")

	var outBuf, errBuf bytes.Buffer

	cmd.Stdin = strings.NewReader(code)
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Typst rendering error:\n```\n%s\n```", errBuf.String())
	}

	return &discordgo.File{
		Name:        "render.png",
		ContentType: "image/png",
		Reader:      &outBuf,
	}, nil
}
