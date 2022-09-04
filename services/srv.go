package services

import (
	"fmt"
	"strings"

	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

var VarSrv *Srv

type Mvcdirs struct {
	Mdl         string
	FlagCrud    string
	Hde         string
	Type        string
	Mid         string
	Cfg         string
	Libs        string
	Rts         string
	Sev         string
	Projectname string
}

var Dirs Mvcdirs

type Srv struct {
}

func (t *Srv) aUpper(b string) string {
	b = strings.ToLower(b)
	b = strings.Title(b)
	return b
}

func (t *Srv) CreateMapDefault(fileName string, packageName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%PACKAGENAME%"] = strings.ToLower(packageName)
	MapForReplace["%MODELNAME%"] = strings.ToLower(fileName)
	MapForReplace["%EXPORTNAME%"] = t.aUpper(fileName)

	MapForReplace["%PROJECT_NAME%"] = strings.ToLower(Dirs.Projectname)
	MapForReplace["%HANDLER_PACKAGE_NAME%"] = strings.ToLower(Dirs.Hde)
	MapForReplace["%MODEL_PACKAGE_NAME%"] = strings.ToLower(Dirs.Mdl)
	MapForReplace["%SERVICE_PACKAGE_NAME%"] = strings.ToLower(Dirs.Sev)
	MapForReplace["%MID_PACKAGE_NAME%"] = strings.ToLower(Dirs.Mid)
	MapForReplace["%ROUTER_PACKAGE_NAME%"] = strings.ToLower(Dirs.Rts)
	MapForReplace["%LIBS_PACKAGE_NAME%"] = strings.ToLower(Dirs.Libs)
	MapForReplace["%CFG_PACKAGE_NAME%"] = strings.ToLower(Dirs.Cfg)

	return MapForReplace
}
func (t *Srv) GetCopyTempleteName(origFolderName string, filesVersionName string) string {
	return viper.GetString("homedir") + "/cp/" + origFolderName + "/" + filesVersionName + ".stub"
}
func (t *Srv) GetTempleteName(origFolderName string, filesVersionName string) string {
	return viper.GetString("homedir") + "/" + origFolderName + "/" + filesVersionName + ".stub"
}
func (t *Srv) GetTempleteDir(origFolderName string, filesVersionName string) string {

	return viper.GetString("homedir") + "/" + origFolderName + "/" + filesVersionName
}

func (t *Srv) CreateDefaultGoFile(origFolderName string, filesVersionName string, destFolderName string, destFileName string, packageName string) {
	nameStruct := destFileName
	destFileName = destFileName + ".go"
	fmt.Println(t.GetTempleteDir(origFolderName, filesVersionName) + "->>" + destFolderName + "/" + destFileName)
	t.CreateDefaultFile(origFolderName, filesVersionName, destFolderName, nameStruct, destFileName, packageName)
}
func (t *Srv) GetDefaultGoFile(origFolderName string, filesVersionName string, destFolderName string, destFileName string, packageName string) string {
	nameStruct := destFileName
	destFileName = destFileName + ".go"
	fmt.Println(t.GetTempleteDir(origFolderName, filesVersionName) + "->>" + destFolderName + "/" + destFileName)
	return t.GetDefaultFile(origFolderName, filesVersionName, destFolderName, nameStruct, destFileName, packageName)
	//ReplaceTextInFile(templateName string, MapForReplace map[string]string)
}
func (t *Srv) GetDefaultFile(origFolderName string, filesVersionName string, destFolderName string, nameStruct string, destFileName string, packageName string) string {

	MapForReplace := t.CreateMapDefault(nameStruct, packageName)
	return jio.ReplaceTextInFile(t.GetCopyTempleteName(origFolderName, filesVersionName), MapForReplace)
}

func (t *Srv) CreateDefaultFile(origFolderName string, filesVersionName string, destFolderName string, nameStruct string, destFileName string, packageName string) {
	jio.CreateFolder(destFolderName)
	MapForReplace := t.CreateMapDefault(nameStruct, packageName)
	newName := destFolderName + "/" + destFileName
	jio.NewFileforTemplate(newName, t.GetTempleteName(origFolderName, filesVersionName), MapForReplace)
}
func (t *Srv) CopyfileAndRemplace(origFolderName string, filesVersionName string, destFileName string, MapForReplace map[string]string) {
	jio.NewFileforTemplate(destFileName, t.GetTempleteDir(origFolderName, filesVersionName), MapForReplace)

}

func (t *Srv) CopyTemplateDir(origFolderName string, filesVersionName string, destFileName string) {
	fmt.Println(t.GetTempleteDir(origFolderName, filesVersionName) + "->>" + destFileName)
	jio.Copy(t.GetTempleteDir(origFolderName, filesVersionName), destFileName)
}

func (t *Srv) CopyTemplate(origFolderName string, filesVersionName string, destFileName string) {
	jio.Copy(t.GetTempleteName(origFolderName, filesVersionName), destFileName)
}
func (t *Srv) UpdateFile(destFileName string, MapForReplace map[string]string) {
	jio.ChancarFile(destFileName, MapForReplace)
}
