package utils

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ReadFile(filePath string) string {
	buf := ReadFileBuf(filePath)
	return string(buf)
}

func ReadFileBuf(filePath string) []byte {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}

	return buf
}

func WriteFile(filePath string, content string) {
	var d1 = []byte(content)
	err2 := ioutil.WriteFile(filePath, d1, 0666) //写入文件(字节数组)
	check(err2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GetAllFiles(dirPth string, ext string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth+PthSep+fi.Name(), ext)
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), "."+ext)
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table, ext)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func MkDir(dir string) {
	if !CheckFileIsExist(dir) {
		os.Mkdir(dir, os.ModePerm)
	}
}

func ReadCheckpoints(file string) []string {
	content := ReadFile(file)

	myExp := regexp.MustCompile(`<<<TC[\S\s]*steps:[^\n]*\n*([\S\s]*)\n+expects:`)
	arr := myExp.FindStringSubmatch(content)

	str := ""
	if len(arr) > 1 {
		checkpoints := arr[1]
		str = RemoveBlankLine(checkpoints)
	}

	ret := GenCheckpointArr(str)

	return ret
}

func ReadExpect(file string) [][]string {
	content := ReadFile(file)

	myExp := regexp.MustCompile(`<<<TC[\S\s]*expects:[^\n]*\n*([\S\s]*)\n+TC;`)
	arr := myExp.FindStringSubmatch(content)

	str := ""
	if len(arr) > 1 {
		expects := arr[1]

		if strings.Index(expects, "@file") > -1 {
			str = ReadFile(ScriptToExpectName(file))
		} else {
			str = RemoveBlankLine(expects)
		}
	}

	ret := GenExpectArr(str)

	return ret
}

func ReadLog(logFile string) [][]string {
	str := ReadFile(logFile)

	ret := GenExpectArr(str)
	return ret
}

func GenCheckpointArr(str string) []string {
	ret := make([]string, 0)
	for _, line := range strings.Split(str, "\n") {
		line := strings.TrimSpace(line)

		if strings.Index(line, "@") == 0 {
			ret = append(ret, line)
		}
	}

	return ret
}

func GenExpectArr(str string) [][]string {
	ret := make([][]string, 0)
	indx := -1
	for _, line := range strings.Split(str, "\n") {
		line := strings.TrimSpace(line)

		if line == "#" {
			ret = append(ret, make([]string, 0))
			indx++
		} else {
			if len(line) > 0 {
				ret[indx] = append(ret[indx], line)
			}
		}
	}

	return ret
}
