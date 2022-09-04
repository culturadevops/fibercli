package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

func (t *Srv) CopyInitConfig(Config string) {
	origFolderName := "init/cleacode/" + Config
	t.CopyTemplate(origFolderName, "main", "main.go")
	origFolderName = origFolderName + "/libs"
	jio.CreateFolder("libs")
	t.CopyTemplate(origFolderName, "libs", "libs/libs.go")
	t.CopyTemplate(origFolderName, "mysql", "libs/mysql.go")
}

func (t *Srv) Copydir(base string, origFolderName string) {
	fmt.Println("Nombre de carpeta inicial " + origFolderName)
	files, err := ioutil.ReadDir(viper.GetString("homedir") + "/" + base + "/" + origFolderName)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		path := ""
		if origFolderName == "" {
			path = f.Name()
		} else {
			path = origFolderName + "/" + f.Name()
			//fmt.Println(path)
		}
		if f.IsDir() {
			jio.CreateFolder(path)
			t.Copydir(base, path)
		} else {
			t.CopyTemplateDir(base, path, path)
		}
	}
}
func (t *Srv) CopydirAndRemplaceMark(base string, origFolderName string, nameStruct string) {
	fmt.Println("Nombre de carpeta inicial " + origFolderName)
	files, err := ioutil.ReadDir(viper.GetString("homedir") + "/" + base + "/" + origFolderName)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		path := ""
		if origFolderName == "" {
			path = f.Name()
		} else {
			path = origFolderName + "/" + f.Name()
			//fmt.Println(path)
		}
		if f.IsDir() {
			jio.CreateFolder(path)
			t.CopydirAndRemplaceMark(base, path, nameStruct)
		} else {
			MapForReplace := make(map[string]string)
			MapForReplace["%FIBERCLI%"] = strings.ToLower(nameStruct)
			t.CopyfileAndRemplace(base, path, path, MapForReplace)
		}
	}
}

func (t *Srv) InitServerlessCopyAll(Config string) {
	origFolderName := ""
	base := "init/lambda"
	t.CopydirAndRemplaceMark(base, origFolderName, Config)

}

func (t *Srv) InitCleanCodeCopyAll(Config string) {
	origFolderName := ""
	base := "init/cleancode"
	t.CopydirAndRemplaceMark(base, origFolderName, Config)
	//t.Copydir(base, origFolderName)
}
func (t *Srv) InitMVCCopyAll(Config string) {
	origFolderName := ""
	base := "init/mvc"
	t.CopydirAndRemplaceMark(base, origFolderName, Config)
}
