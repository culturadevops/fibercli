package services

func (t *Srv) CreateLibDefault(destfolder string, packageName string, DestFileName string) {
	t.CreateDefaultGoFile("libs", "default", destfolder, DestFileName, packageName)
}
