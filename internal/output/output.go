package output

import (
	"os"
	"encoding/base64"
	"encoding/json"
	"net/url"

	"github.com/sp1at/go-splash/internal/args"
	"github.com/tidwall/gjson"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type responseJson struct {
	Har 	string
	Html 	string
	Jpeg 	string
	Png 	string
}

func getHostFromUrl(argUrl string) string {
	u, err := url.Parse(argUrl)
	check(err)

	return u.Host
}

func createOutputDirectory(host string) {
	//fmt.Printf("%s", "outside")
	//if _, err := os.Stat("./out"); os.IsNotExist(err) {
		os.MkdirAll("./out/" + host, 0750)
		//fmt.Printf("%s", "Made it")
	//} else {
	//	if
	//}
}

func SaveOutput(response map[string]string, args args.Args) {

	host := getHostFromUrl(args.Url)

	createOutputDirectory(host)

	for key, value := range response {

		if key == "json" {
			if !args.Endpoints.Json &&
				!args.Endpoints.Har &&
				!args.Endpoints.Html &&
				!args.Endpoints.Jpeg &&
				!args.Endpoints.Png {


				var responseJson responseJson
				json.Unmarshal([]byte(value), &responseJson)

				png, err := base64.StdEncoding.DecodeString(responseJson.Png)
				check(err)

				jpg, err := base64.StdEncoding.DecodeString(responseJson.Jpeg)
				check(err)

				writeToFile(host, "html", responseJson.Html)
				writeToFile(host, "har", gjson.Get(value, "har").String())
				writeToFile(host, "png", string(png))
				writeToFile(host, "jpg", string(jpg))
			}
		}
		writeToFile(host, key, value)
	}
}

func writeToFile(host string, key string, value string) {
	if value != "" {
		var fName = "./out" + "/" + host + "/" + host + "." + key
		f, err := os.Create(fName)
		check(err)

		f.WriteString(value)
		check(err)

		f.Sync()
	}
}
