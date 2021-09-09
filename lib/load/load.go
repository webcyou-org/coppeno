package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Coppeno struct {
	Name string `json:"name"`
	Url string  `json:"url"`
}

func Start(fileName string, fileGroup string) []Coppeno {
	fmt.Println("load: Start")
	//f, _ := os.Open("coppeno.json")
	////if err != nil {
	////	return err
	////}
	//defer f.Close()

	//var coppeno []*Coppeno
	//json.NewDecoder(f).Decode(&coppeno)
	//err = json.NewDecoder(f).Decode(&coppeno)
	//if err != nil {
	//	return err
	//}

	raw, err := ioutil.ReadFile("coppeno.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var coppeno []Coppeno
	json.Unmarshal(raw, &coppeno)

	//for _, v := range coppeno {
	//	fmt.Printf("%+v\n", v)
	//}
	//// fmt.Printf("%+v\n", coppeno)
	//fmt.Println(fileName)
	//fmt.Println(fileGroup)
	//return nil
	return coppeno
}
