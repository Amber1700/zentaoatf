package utils

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"sync"
)

type Config struct {
	Language string
	Width    int
	Height   int
}

var Conf Config

func InitConfig() {
	// config from yaml
	Conf = getInst()

	// screen size
	InitScreenSize()

	// internationalization
	InitI118(Conf.Language)
	color.Blue(I118Prt.Sprintf("current_config", ""))

	// print config object
	val := reflect.ValueOf(Conf)
	typeOfS := val.Type()
	for i := 0; i < reflect.ValueOf(Conf).NumField(); i++ {
		fmt.Printf("  %s: %v \n", typeOfS.Field(i).Name, val.Field(i).Interface())
	}
}

func Set(param string, val string) {
	buf, _ := ioutil.ReadFile(ConfFile)
	yaml.Unmarshal(buf, &Conf)

	if param == "lang" {
		Conf.Language = val

		data, _ := yaml.Marshal(&Conf)
		ioutil.WriteFile(ConfFile, data, 0666)

		color.Blue(I118Prt.Sprintf("set_config", I118Prt.Sprintf("lang"), I118Prt.Sprintf(Conf.Language)))
	}

}

func getInst() Config {
	var once sync.Once
	once.Do(func() {
		Conf = Config{}
		buf, _ := ioutil.ReadFile(ConfFile)
		yaml.Unmarshal(buf, &Conf)
	})
	return Conf
}
