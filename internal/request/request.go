package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sp1at/go-splash/internal/args"
	"github.com/sp1at/go-splash/internal/config"
)

type Request struct {
	endpoint				string
	header					string
	body					[]byte
}

type RenderHtml struct {
	// All API options are included here for quick adjustments to config settings

	Url						string	`json:"url"` 					//required
	//BaseUrl				string	`json:"baseurl"` 				//optional
	//Timeout				string	`json:"timeout"`				// float (optional) default: 30.0 max: 90.0
	//ResourceTimeout		string	`json:"resource_timeout"`		// float (optional)
	Wait					string	`json:"wait"`					// float (optional) required for full-page png/jpeg rendering
	//Proxy					string	`json:"proxy"`					// string (optional) format: [protocol://][user:password@][proxyhost[:port]]
	//Js					string	`json:"js"`						// string (optional)
	//JsSource				string	`json:"js_source"`				// string (optional)
	//Filters				string	`json:"filters"`				// string (optional) CSV
	//AllowedDomains		string	`json:"allowed_domains"`		// string (optional) CSV
	//AllowedContentTypes	string	`json:"allowed_content_types"`	// string (optional) CSV
	//ForbiddenContentTypes	string	`json:"forbidden_content_types"`// string (optional) CSV
	Viewport				string	`json:"viewport"`				// string (optional) default: 1024x768
	Images					string	`json:"images"`					// integer (optional) 1 or 0
	//Headers				string	`json:"headers"`				// JSON object or string (optional)
	//Body					string	`json:"body"`					// string (optional)
	//HttpMethod			string	`json:"http_method"`			// string (optional)
	//HttpMethod			string	`json:"http_method"`			// string (optional)
	//SaveArgs				string	`json:"save_args"`				// JSON object or string (optional)
	//LoadArgs				string	`json:"load_args"`				// JSON object or string (optional)
	//Html5Media				string	`json:"html5_media"`		// integer (optional) 1 or 0
}

type RenderPng struct {
	RenderHtml
	//Width					string	`json:"width"`					// integer (optional) width in pixels
	//Height				string	`json:"height"`					// integer (optional) height in pixels
	RenderAll				string	`json:"render_all"`				// integer (optional) 1 or 0
	ScaleMethod				string	`json:"scale_method"` 			// string (optional) default: raster option: vector
}

type RenderJpeg struct {
	RenderPng
	Quality					string	`json:"quality"`				// integer (optional) 1 to 100 default: 75
}

type RenderHar struct {
	RenderHtml
	RequestBody				string	`json:"request_body"`			// integer 1 or 0
	ResponseBody			string	`json:"response_body"`			// integer 1 or 0
}

type RenderJson struct {
	RenderJpeg
	Html 					string	`json:"html"` 					// integer 1 or 0 default: 0
	Png 					string 	`json:"png"` 					// integer 1 or 0 default: 0
	Jpeg					string  `json:"jpeg"` 					// integer 1 or 0 default: 0
	Iframes					string  `json:"iframes"`		 		// integer 1 or 0 default: 0
	Script					string	`json:"script"`					// integer 1 or 0 default: 0
	Console					string 	`json:"console"`				// integer 1 or 0 default: 0
	History					string  `json:"history"`				// integer 1 or 0 default: 0
	Har 					string 	`json:"har"`					// integer 1 or 0 default: 0
	RequestBody				string	`json:"request_body"`			// integer 1 or 0 default: 0
	ResponseBody			string	`json:"response_body"`			// integer 1 or 0 default: 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildRequestMap(args args.Args) map[string]string {

	var endpointsMap map[string]string

	endpointsMap = make(map[string]string)

	// make it functional - then make it pretty
	if !args.Endpoints.Json &&
		!args.Endpoints.Har  &&
		!args.Endpoints.Html &&
		!args.Endpoints.Jpeg &&
		!args.Endpoints.Png {
		endpointsMap["json"] = "/render.json"
	}

	if args.Endpoints.Json {
		endpointsMap["json"] = "/render.json"
		return endpointsMap
	}

	if args.Endpoints.Har {
		endpointsMap["har"] = "/render.har"
	}

	if args.Endpoints.Html {
		endpointsMap["html"] = "/render.html"
	}

	if args.Endpoints.Jpeg {
		endpointsMap["jpeg"] = "/render.jpeg"
	}

	if args.Endpoints.Png {
		endpointsMap["png"] = "/render.png"
	}

	return endpointsMap
}

func BuildRequestBody(args args.Args, endpointType string) []byte {

	var j []byte

	var renderHtml = RenderHtml{
		Url:                   args.Url,
		//BaseUrl:               "",
		//Timeout:               "90",
		//ResourceTimeout:       "90",
		Wait:                  "3",
		//Proxy:                 "",
		//Js:                    "",
		//JsSource:              "",
		//Filters:               "",
		//AllowedDomains:        "",
		//AllowedContentTypes:   "",
		//ForbiddenContentTypes: "",
		Viewport:              "1200x800",
		Images:                "1",
		//Headers:               "1",
		//Body:                  "1",
		//HttpMethod:            "POST",
		//SaveArgs:              "",
		//LoadArgs:              "",
		//Html5Media:            "0",
	}
	
	var renderPng = RenderPng{
		RenderHtml:  renderHtml,
		//Width:       "1200",
		//Height:      "800",
		RenderAll:   "1",
		ScaleMethod: "raster",
	}
	
	var renderJpeg = RenderJpeg{
		RenderPng: renderPng,
		Quality:   "75",
	}
	
	var renderHar = RenderHar{
		RenderHtml: renderHtml,
		RequestBody:  "1",
		ResponseBody: "1",
	}

	var renderJson = RenderJson{
		RenderJpeg:   renderJpeg,
		Html:         "1",
		Png:          "0",
		Jpeg:         "1",
		Iframes:      "1",
		Script:       "0", // TODO: Add Custom JS Support
		Console:      "1",
		History:      "0",
		Har:          "1",
		RequestBody:  "1",
		ResponseBody: "1",
	}

	if endpointType == "json" {
		j, err := json.Marshal(renderJson)
		check(err)
		
		return j
	}

	if endpointType == "har" {
		j, err := json.Marshal(renderHar)
		check(err)

		return j
	}

	if endpointType == "html" {
		j, err := json.Marshal(renderHtml)
		check(err)

		return j
	}

	if endpointType == "jpeg" {
		j, err := json.Marshal(renderJpeg)
		check(err)

		return j
	}

	if endpointType == "png" {
		j, err := json.Marshal(renderPng)
		check(err)

		return j
	}

	return j
}

func InitRequests(config config.Config, args args.Args) map[string]string {
	requestMap := buildRequestMap(args)

	var responseMap map[string]string
	responseMap = make(map[string]string)

	var baseUrl = "http://" + config.SplashHost + ":" + config.SplashPort
	var jsonHeader = "content-type: application/json"

	if len(requestMap) > 0 {
		for key, value := range requestMap {
			request := Request{}

			request.endpoint = baseUrl + value
			request.header = jsonHeader
			request.body = BuildRequestBody(args, key)

			response := MakeRequest(request)

			StoreRequest(responseMap, key, response)
		}
	}

	return responseMap
}

func MakeRequest(request Request) string {
	req, err := http.NewRequest("POST", request.endpoint, bytes.NewBuffer(request.body) )
	check(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp,err := client.Do(req)
	check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(body)
}

func StoreRequest(responseMap map[string]string, outputType string, response string) {
	responseMap[outputType] = response
}