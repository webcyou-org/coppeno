package save

import (
	"encoding/json"
	"fmt"
	"os"
	"webcyou-org/coppeno/lib/load"
)

type Coppeno struct {
	Name string `json:"name"`
	Url string  `json:"url"`
}

func Start(name string, targetPath string) error {
	fmt.Println("save: Start")

	coppeno := load.Start("fileName", "fileGroup")

	// fmt.Println(len(coppeno))
	fmt.Println(coppeno)
	//fmt.Println(name)
	//fmt.Println(targetPath)

	for _, c := range coppeno {
		// fmt.Sprintf("%s: %s", c.Name, c.Url)
		fmt.Println(c.Name)
		fmt.Println(c.Url)
	}

	//a := map[string]interface{}{
	//	"name": name,
	//	"url": targetPath,
	//}

	a := Coppeno{
		Name: name,
		Url: targetPath,
	}
	fmt.Println(a)

	//coppeno = append(coppeno, a)

	f, err := os.Create("data_output.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(a)
	if err != nil {
		return err
	}
	return nil
}