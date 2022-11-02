package tool

import (
	"fmt"
	_ "github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

func ChangeIndexToAxis(intIndexX int, intIndexY int) string {
	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	intIndexY = intIndexY + 1
	resultY := ""
	for true {
		if intIndexY <= 26 {
			resultY = resultY + arr[intIndexY-1]
			break
		}
		mo := intIndexY % 26
		resultY = arr[mo-1] + resultY
		shang := intIndexY / 26
		if shang <= 26 {
			resultY = arr[shang-1] + resultY
			break
		}
		intIndexY = shang
	}
	return resultY + strconv.Itoa(intIndexX+1)
}

func Exec(Cmd string) error {
	cmd := exec.Command("/bin/bash", "-c", Cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
		return err
	}
	defer stdout.Close() // 保证关闭输出流
	err = cmd.Start()
	if err != nil { // 运行命令
		log.Fatal(err)
		return err
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil { // 读取输出结果
		log.Fatal(err)
		return err
	} else {
		fmt.Println(string(opBytes))
	}
	return nil
}
