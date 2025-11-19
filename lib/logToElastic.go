package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"riskmanagement/models"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type logActivity struct {
	Uri       string `json:"uri"`
	Agent     string `json:"agent"`
	Ipaddress string `json:"ipaddress"`
	Request   string `json:"request"`
	Response  string `json:"response"`
}

type logError struct {
	Month         string `json:"month"`
	Year          string `json:"year"`
	Filename      string `json:"filename"`
	Function      string `json:"function"`
	Line          string `json:"line"`
	MessageCustom string `json:"messageCustom"`
	MessageSystem string `json:"messageSystem"`
}

func WhereAmI(depthList ...int) (string, string, string) {

	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}

	function, file, line, _ := runtime.Caller(depth)
	fl := chopPath(file)
	fn := runtime.FuncForPC(function).Name()
	ln := fmt.Sprintf("%d", line)

	return fl, fn, ln
}

func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

func CreateLogErrorToDB(db *elasticsearch.Client, filename, function, line, messageCustom, messageSystem string) {
	year := GetTimeNow("year")
	month := strings.ToLower(GetTimeNow("month"))
	model := models.LogError{}
	logs := logError{
		Month:         month,
		Year:          year,
		Filename:      filename,
		Function:      function,
		Line:          line,
		MessageCustom: messageCustom,
		MessageSystem: messageSystem,
	}
	body, err := json.Marshal(logs)
	if err != nil {
		fmt.Println(err.Error())
	}

	req := esapi.CreateRequest{
		Index:      model.IndexLogError(month),
		DocumentID: UUID(false),
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := req.Do(ctx, db)
	if err != nil {
		fmt.Println("LogErrorToElastic =>", err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println("LogErrorToElastic=> ", res.String())
	}

}

func CreateLogActivityToDB(db *elasticsearch.Client, uri string, agent string, ipaddress string, reqBody string, respBody string) {
	model := models.LogAccess{}
	logs := logActivity{
		Uri:       uri,
		Agent:     agent,
		Ipaddress: ipaddress,
		Request:   reqBody,
		Response:  respBody,
	}
	body, err := json.Marshal(logs)
	if err != nil {
		fmt.Println(err.Error())
	}

	month := strings.ToLower(GetTimeNow("month"))
	req := esapi.CreateRequest{
		Index:      model.IndexLogAccess(month),
		DocumentID: UUID(false),
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := req.Do(ctx, db)
	if err != nil {
		fmt.Println("LogActivityToElastic =>", err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println("LogActivityToElastic ", res.String())
	}
	// create log activity
	// fmt.Println("==>", db)
	// fmt.Println("uri==>", uri)
	// fmt.Println("agent==>", agent)
	// fmt.Println("ipaddress==>", ipaddress)
	// fmt.Println("request==>", reqBody)
	// fmt.Println("response==>", respBody)
}
