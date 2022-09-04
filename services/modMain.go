package services

func (t *Srv) AddValueToFile(index string, value string, file string) {
	MapForReplace := make(map[string]string)
	MapForReplace[index] = value
	t.UpdateFile(file, MapForReplace)
}
func (t *Srv) AddMidToMain(filename string) {
	index := "func mainMiddlewares(e *echo.Echo) {"
	value := index + "\n\te.Use(middlewares." + t.aUpper(filename) + "())"
	t.AddValueToFile(index, value, "main.go")
}
func (t *Srv) AddMidToRts(filename string, fileDest string) {
	index := "/*middlewares*/"
	value := index + "\n\tr.Use(middlewares." + t.aUpper(filename) + "())"
	t.AddValueToFile(index, value, fileDest+".go")

}
func (t *Srv) AddRts(filename string) {
	index := "func mainRoutes(app *fiber.App) {"
	value := index + "\n\troutes." + t.aUpper(filename) + "(app,new(handlers." + t.aUpper(filename) + "))"
	t.AddValueToFile(index, value, "main.go")
}
