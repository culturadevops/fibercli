package services

func (t *Srv) CreateMdlDefault(destfolder string, packageName string, DestFileName string) {
	t.CreateMdl(destfolder, DestFileName, packageName, "default")
}
func (t *Srv) CreateMdlCrud(destfolder string, packageName string, DestFileName string) {
	t.CreateMdl(destfolder, DestFileName, packageName, "crud")
}
func (t *Srv) CreateMdl(destfolder string, DestFileName string, packageName string, verFileName string) {
	t.CreateDefaultGoFile("models", verFileName, destfolder, DestFileName, packageName)
}
