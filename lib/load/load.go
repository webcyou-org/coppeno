package load

import (
	"encoding/json"
	"fmt"
	"os"
)

type Coppeno struct {
	Name string `json:"name"`
	Url string  `json:"url"`
}

func Start(fileName string, fileGroup string) []*Coppeno {
	fmt.Println("load: Start")
	f, _ := os.Open("coppeno.json")
	//if err != nil {
	//	return err
	//}
	defer f.Close()

	var coppeno []*Coppeno
	json.NewDecoder(f).Decode(&coppeno)
	//err = json.NewDecoder(f).Decode(&coppeno)
	//if err != nil {
	//	return err
	//}

	//for _, v := range coppeno {
	//	fmt.Printf("%+v\n", v)
	//}
	//// fmt.Printf("%+v\n", coppeno)
	fmt.Println(fileName)
	fmt.Println(fileGroup)
	//return nil
	return coppeno
}
