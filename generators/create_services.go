package generators

import (
	_ "embed"
	"fmt"
	"golang-starter/helpers"
	"log"
	"path/filepath"

	"github.com/iancoleman/strcase"
)

//go:embed templates/delivery/usecase/init.tmpl
var initTemplate string

//go:embed templates/delivery/usecase/create.tmpl
var createTemplate string

//go:embed templates/delivery/usecase/list.tmpl
var listTemplate string

//go:embed templates/delivery/usecase/detail.tmpl
var detailTemplate string

//go:embed templates/delivery/usecase/delete.tmpl
var deleteTemplate string

//go:embed templates/delivery/usecase/update.tmpl
var updateTemplate string

//go:embed templates/domain/usecase/interface.tmpl
var domainUsecaseTemplate string

func GenerateService() {

	type Data struct {
		PackageName string
		ServiceName string
	}
	var packageName string
	var serviceName string

	fmt.Println("Package Name:")
	_, err := fmt.Scanln(&packageName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Serivce Name:")
	_, err = fmt.Scanln(&serviceName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := Data{PackageName: packageName, ServiceName: serviceName}

	if err != nil {
		log.Fatal(err)
		return
	}

	helpers.ProcessTemplate(initTemplate, "init.tmpl", filepath.Join("service/delivery/usecase", strcase.ToSnake(serviceName)+"/init.go"), data)
	helpers.ProcessTemplate(createTemplate, "create.tmpl", filepath.Join("service/delivery/usecase/", strcase.ToSnake(serviceName)+"/create.go"), data)
	helpers.ProcessTemplate(listTemplate, "list.tmpl", filepath.Join("service/delivery/usecase/", strcase.ToSnake(serviceName)+"/list.go"), data)
	helpers.ProcessTemplate(detailTemplate, "detail.tmpl", filepath.Join("service/delivery/usecase/", strcase.ToSnake(serviceName)+"/detail.go"), data)
	helpers.ProcessTemplate(deleteTemplate, "delete.tmpl", filepath.Join("service/delivery/usecase/", strcase.ToSnake(serviceName)+"/delete.go"), data)
	helpers.ProcessTemplate(updateTemplate, "update.tmpl", filepath.Join("service/delivery/usecase/", strcase.ToSnake(serviceName)+"/update.go"), data)
	helpers.ProcessTemplate(domainUsecaseTemplate, "interface.tmpl", filepath.Join("serivce/domain/usecase", strcase.ToSnake(serviceName)+".go"), data)
}
