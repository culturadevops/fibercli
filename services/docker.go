package services

//VarDocker.Create(args[0])
func (t *Srv) CreateDocker(filesVersionName string) {
	origFolderName := "dockers"
	destFileName := "Dockerfile"
	t.CopyTemplate(origFolderName, filesVersionName, destFileName)
}
func (t *Srv) DockerGo() {
	t.CreateDocker("go")
}
func (t *Srv) DockerPhp() {
	t.CreateDocker("php")
}
