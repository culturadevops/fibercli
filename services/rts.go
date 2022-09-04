package services

func (t *Srv) CreateRtsDefault(destino string, DestFileName string, packageName string) {
	t.CreateDefaultGoFile("routes", "default", destino, DestFileName, packageName)
	t.AddRts(DestFileName)
}
func (t *Srv) CreateRtsApi(destino string, DestFileName string, packageName string) {
	t.CreateDefaultGoFile("routes", "api", destino, DestFileName, packageName)
	t.AddRts(DestFileName)
}
