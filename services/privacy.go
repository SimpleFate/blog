package services

import (
	"blog/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	IP_SEARCH_URL_FORMAT = "http://ip.360.cn/IPQuery/ipquery?ip=%s"
	IP_URL               = "http://ip.360.cn/IPShare/info"
)

type Privacy struct {
	Ip       string
	Location string
	Dt       int64
}

func (p *Privacy) String() string {
	return fmt.Sprintf("ip : %s\nlocation : %s\ndt : %d\ntime : %s",
		p.Ip,
		p.Location,
		p.Dt,
		utils.DtToString(p.Dt))
}

func GetPrivacy(r *http.Request) *Privacy {
	addr := r.RemoteAddr
	//[::1]:53275
	ip, _, _ := net.SplitHostPort(addr)

	if ip == "::1" {
		ip = GetLocalIp()
	}
	if ip == "" {
		return nil
	}
	requst := fmt.Sprintf(IP_SEARCH_URL_FORMAT, ip)
	resp, err := http.Get(requst)
	if err != nil {
		return nil
	}
	resJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	res := struct {
		Data string `json:"data"`
	}{}
	err = json.Unmarshal(resJson, &res)
	if err != nil {
		return nil
	}

	location := res.Data
	dt := time.Now().Unix()

	privacy := Privacy{
		Ip:       ip,
		Location: location,
		Dt:       dt,
	}

	return &privacy
}

func GetLocalIp() string {
	resp, err := http.Get(IP_URL)
	if err != nil {
		return ""
	}
	resJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	res := struct {
		Ip string `json:"ip"`
	}{}

	err = json.Unmarshal(resJson, &res)

	if err != nil {
		return ""
	}
	return res.Ip
}
