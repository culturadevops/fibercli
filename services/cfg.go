package services

//services.Mdl(args[0])
func (t *Srv) CreateConfigMysqlToml() {
	t.CreateConfig("mysql", "mysql", "toml", "")
}
func (t *Srv) CreateConfigYml(destFileName string) {
	t.CreateConfig("yml", destFileName, "yml", "")
}

func (t *Srv) CreateConfig(origFileName string, destFileName string, ext string, packageName string) {
	t.CreateDefaultFile("configs", origFileName, "configs", destFileName, destFileName+"."+ext, packageName)
}
