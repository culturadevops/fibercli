package services

//services.Mdl(args[0])
func (t *Srv) CreateCtlDefault(destfolder string, packageName string, DestFileName string) {
	t.CreateDefaultGoFile("controllers", "default", destfolder, DestFileName, packageName)
}
func (t *Srv) CreateCtlCrud(destfolder string, packageName string, DestFileName string) {
	t.CreateDefaultGoFile("controllers", "crud", destfolder, DestFileName, packageName)
}

func (t *Srv) GetCtlCrud(destfolder string, packageName string, DestFileName string) string {
	return t.GetDefaultGoFile("handle", "default", destfolder, DestFileName, packageName)
}
