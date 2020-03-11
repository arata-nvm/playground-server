package infra

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

// 引数のプログラミングをファイルに保存し、そのファイル名を返す
func SaveProgram(program string) (string, error) {
	programBytes := []byte(program)
	filename := fmt.Sprintf("./sandbox/src/%x.sl", md5.Sum(programBytes))
	err := ioutil.WriteFile(filename, programBytes, 0644)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// 引数のファイルをコンパイルし、そのバイナリ名を返す
func CompileCode(filename string) (string, error) {
	outputBinName := getFileNameWithoutExt(filename)
	outputPath := fmt.Sprintf("./sandbox/bin/%s", outputBinName)
	args := []string{
		"-O",
		"-o", outputPath,
		filename,
	}

	cmd := &Command{
		Name: "./static/visket",
		Args: args,
	}
	result := execute(cmd)
	if result.Error != nil {
		return "", fmt.Errorf(result.Stderr)
	}

	return outputPath, nil
}

// 引数のバイナリに標準入力を与えて実行する
func RunCode(filename string, stdin string) (string, error) {
	cmd := &Command{
		Name:  filename,
		Stdin: stdin,
	}
	result := execute(cmd)
	if result.Error != nil {
		return "", result.Error
	}

	output := fmt.Sprintf("%s\nExited with code %d\n", result.Stdout, result.ExitCode)

	return output, nil
}
