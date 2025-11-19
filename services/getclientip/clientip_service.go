package getclientip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	models "riskmanagement/models/getclientip"

	"gitlab.com/golang-package-library/logger"
)

type ClientIPDefinition interface {
	ClientIP() (response models.ClientIP, err error)
}

type ClientIpService struct {
	logger logger.Logger
}

func NewClientIPService(
	logger logger.Logger,
) ClientIPDefinition {
	return ClientIpService{
		logger: logger,
	}
}

type Response struct {
	IP string `json:"ip"`
}

// type Response struct {
// 	DNS struct {
// 		Geo string `json:"geo"`
// 		IP  string `json:"ip"`
// 	} `json:"dns"`
// }

type IPAddressInfo struct {
	Query    string  `json:"query"`
	Country  string  `json:"country"`
	City     string  `json:"city"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone"`
	ISP      string  `json:"isp"`
}

func getIPAddress() (IPAddressInfo, error) {
	response, err := http.Get("https://api.ipify.org?format=json")
	// response, err := http.Get("http://edns.ip-api.com/json")
	if err != nil {
		return IPAddressInfo{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return IPAddressInfo{}, err
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return IPAddressInfo{}, err
	}

	if data.IP != "" {
		infoURL := fmt.Sprintf("http://ip-api.com/json/%s", data.IP)
		infoResponse, err := http.Get(infoURL)
		if err != nil {
			return IPAddressInfo{}, err
		}

		defer infoResponse.Body.Close()

		infoBody, err := ioutil.ReadAll(infoResponse.Body)
		if err != nil {
			return IPAddressInfo{}, err
		}

		var infoData IPAddressInfo
		err = json.Unmarshal(infoBody, &infoData)
		if err != nil {
			return IPAddressInfo{}, err
		}

		ipInfo := IPAddressInfo{
			Query:    infoData.Query,
			Country:  infoData.Country,
			City:     infoData.City,
			Lat:      infoData.Lat,
			Lon:      infoData.Lon,
			Timezone: infoData.Timezone,
			ISP:      infoData.ISP,
		}

		return ipInfo, err
	}

	// ipInfo := IPAddressInfo{
	// 	IP:  data.DNS.IP,
	// 	Geo: data.DNS.Geo,
	// }

	return IPAddressInfo{}, nil
}

// ClientIP implements ClientIPDefinition
func (cs ClientIpService) ClientIP() (response models.ClientIP, err error) {
	ipInfo, err := getIPAddress()

	if err != nil {
		cs.logger.Zap.Info("Error fetching IP:", err)
		return response, err
	}

	response = models.ClientIP{
		IP:       ipInfo.Query,
		Country:  ipInfo.Country,
		City:     ipInfo.City,
		Lat:      ipInfo.Lat,
		Lon:      ipInfo.Lon,
		Timezone: ipInfo.Timezone,
		ISP:      ipInfo.ISP,
	}

	return response, err
}
