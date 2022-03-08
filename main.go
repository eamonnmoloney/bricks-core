package main

import (
	"bricks-core/internals/controllers"
	"fmt"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"strings"
)

type Import struct {
	name string
	dep  string
}

type Component interface {
	write() string
}

type Button struct {
	Classes []string `yaml:"classes"`
}

type Page struct {
	imports       []Import
	PreComponents []interface{} `yaml:"components"`
	Components    []Component   `yaml:"-"`
}

func (b Button) write() string {
	return "<Button>" + strings.Join(b.Classes, ",") + "<Button>"
}

func (i Import) write() string {
	return "import " + i.name + " from " + i.dep
}

func mapToString(slice []Component) string {
	result := ""
	for _, item := range slice {
		result += item.write() + "\n"
	}
	return result
}

func importsToString(slice []Import) string {
	result := ""
	for _, item := range slice {
		result += item.write() + "\n"
	}
	return result
}

func (p Page) write() string {
	return importsToString(p.imports) +
		`export default function Home() {
	return (<div>` + "\n" +
		mapToString(p.Components) +
		`</div>)
}`
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/applications", controllers.ReadApplications)
	r.POST("/applications", controllers.CreateApplication)

	return r
}

func main() {
	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Println("=======================================")
	log.Println("Runinng gin-lambda server in " + addr)
	log.Println("=======================================")
	if mode == "production" {
		log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
	} else {
		log.Fatal(http.ListenAndServe(addr, routerEngine()))
	}

	//r.GET("/components", controllers.ReadComponents)
	//r.POST("/components", controllers.CreateComponent)
	//
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//convertYmlToJs()

}

func convertYmlToJs() {
	// consume yml
	// product jsx
	//path, _ := filepath.Abs("hello-world.yml")
	//input, _ := ioutil.ReadFile(path)
	var page Page
	//test := map[string]string{}
	s := `components: 
  - classes: hello`
	err := yaml.Unmarshal([]byte(s), &page)
	if err != nil {
		return
	}

	var arr []Component
	for _, entry := range page.PreComponents {
		switch i := entry.(type) {
		case string:
			log.Printf("i is a %+v\n", i)
		case map[string]interface{}:
			log.Printf("i is a %+v\n", i)
			button := Button{}
			for _, v := range i {
				switch j := v.(type) {
				case string:
					button.Classes = strings.Split(j, ",")
				}
			}
			arr = append(arr, button)
		}
	}

	page.Components = arr
	page.imports = append(page.imports, Import{name: "Head", dep: "next/Head"})

	marshal := page.write()

	fmt.Println(marshal)
}
