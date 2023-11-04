package main

import (
    _ "embed"
    "os/exec"
    "log"
)

//go:embed files/pre-commit-hook.sh
var commandStr string

func generate_prehook() (string, error) {
    cmd         := exec.Command("/bin/sh", "-c",  "tee -a .git/hooks/pre-commit <<FILE\n" + commandStr + ";")
    stdout, err := cmd.Output()

    if err == nil {
        cmd     = exec.Command("/bin/sh", "-c",  "chmod +x .git/hooks/pre-commit;")
        _, err  = cmd.Output()
    }

    if err == nil {
        output := string(stdout)
        log.Println(output)
        return output, err
    } else {
        log.Fatal(err)
        return "", err
    }
}

func main() {
    generate_prehook()
}