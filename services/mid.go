package services

func (t *Srv) CreateMidDefault(destino string, DestFileName string, packageName string) {
	t.CreateDefaultGoFile("middlewares", "default", destino, DestFileName, packageName)
	t.AddMidToMain(DestFileName)
}
func (t *Srv) CreateMidDefaultToRoutes(destino string, DestFileName string, DestRoutes string, packageName string) {
	t.CreateDefaultGoFile("middlewares", "default", destino, DestFileName, packageName)
	t.AddMidToRts(DestFileName, packageName+"/"+DestRoutes)
}

func (t *Srv) CopyMidCrud(destfolder string, packageName string, DestFileName string) string {
	return t.GetDefaultGoFile("middlewares", "default", destfolder, DestFileName, packageName)
}
