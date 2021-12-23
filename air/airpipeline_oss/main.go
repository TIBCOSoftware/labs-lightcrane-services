// main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	folder := "./"
	flogoApp, _ := FromFile("pipeline.json")
	log.Println(flogoApp)
	categories, _ := ioutil.ReadDir(folder)
	imports := make(map[string]interface{})
	for _, category := range categories {
		if category.IsDir() {
			log.Println("Category : ", category.Name())
			templates, _ := ioutil.ReadDir(folder + "/" + category.Name())
			for _, fileInfo := range templates {
				log.Println(fileInfo.Name())
				desc, _ := FromFile(folder + category.Name() + "/" + fileInfo.Name() + "/" + fileInfo.Name() + ".json")
				for _, anImport := range desc["imports"].([]interface{}) {
					if nil == imports[anImport.(string)] && "github.com/TIBCOSoftware/labs-air-contrib/activity/log" != imports[anImport.(string)] {
						log.Println(anImport.(string))
						flogoApp["imports"] = append(flogoApp["imports"].([]interface{}), anImport.(string))
					}
					imports[anImport.(string)] = ""
				}
			}
		}
	}

	flogoAppBytes, _ := json.Marshal(flogoApp)
	flogoAppString := string(flogoAppBytes)
	log.Println(flogoAppString)

	f, err := os.OpenFile("flogo-generic.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer f.Close()

	if err != nil {
		panic(err)
		return
	}

	f.WriteString(flogoAppString)
}

func FromFile(filename string) (map[string]interface{}, error) {
	log.Println(":::::::::", filename)
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(fileContent, &result)

	if nil != err {
		return nil, err
	}

	//log.Debug("[BasePipelineComponent:buildFromFile] FlogoTemplate : filename = ", filename, ", template = ", result)
	return result, nil
}
