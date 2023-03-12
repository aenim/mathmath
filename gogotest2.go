package main
 
import (
	"os"
	"os/exec"
	"io/ioutil"
	"regexp"
	"fmt"
)
 
func main() {
	var s []string
	s = listDirByReadDir(".")
	for i := 0; i < len(s); i++ {

		matched, err := regexp.MatchString("host", s[i])
		if err != nil {
			fmt.Println(err, matched)
		} else {
			if matched != false {
				cmd := exec.Command("echo", "working")
				//cmd := exec.Command("ansible-inventory", "-i",  s[i], "-y", "--list",">", "inventory.yaml") // input command there
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
				cmd.Run()
			} else {
				continue
			}
		}
	} 
	
	
	}

func echocmd() {
    cmd := exec.Command("df", "-h")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func listDirByReadDir(path string) []string {
	var arrfiles []string
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			//fmt.Printf("[%s]\n", val.Name())
			//fmt.Printf("IS_Dir")
			continue
		} else {
			arrfiles = append(arrfiles, val.Name())
			//fmt.Println(val.Name())
			//fmt.Println("IS_NOT_Dir")	
		}
	}
	return arrfiles
}