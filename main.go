package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
)

const priceFileName = "price.json"
const configFileName = "config.json"

type Config struct {
	Ratio         float64 `json:"ratio"`
	EmailAddr     string  `json:"email"`
	EmailPassword string  `json:"password"`
}

type BpiRet struct {
	Time time `json:"time"`
	// Disclaimer string `json:"disclaimer"`
	// ChartName string `json:"chartName"`
	Bpi bpi `json:"bpi"`
}
type time struct {
	Updated string `json:"updated"`
}
type bpi struct {
	Usd usd `json:"USD"`
}

type usd struct {
	// Code string `json:"code"`
	// Symbol string `json:"symbol"`
	Rate string `json:"rate"`
	// Description string `json:"description"`
	RateFloat float64 `json:"rate_float"`
}

func main() {

	config := getConfig()
	ratio := config.Ratio
	// eaddr:="wxiaochi@qq.com"

	now, old := getNowAndOld()
	// old := getOldPrice()
	// savePrice(p)
	if exceed(now.Bpi.Usd.RateFloat, old.Bpi.Usd.RateFloat, ratio) {
		fmt.Println("will email")
		email(config.EmailAddr, config.EmailPassword, makeContent(old, now))
	}
}
func getConfig() *Config {
	var c Config
	b := getFileContent(configFileName)
	err := json.Unmarshal(b, &c)
	check(err)
	return &c
}
func getNowAndOld() (now BpiRet, old BpiRet) {

	url := "https://api.coindesk.com/v1/bpi/currentprice.json"
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	err = json.Unmarshal(body, &now)
	check(err)

	old = getOld()

	err = ioutil.WriteFile(priceFileName, body, 0644) // save old
	check(err)

	return
}

func getOld() (old BpiRet) {
	if _, err := os.Stat(priceFileName); os.IsNotExist(err) {
		return old
	}
	b := getFileContent(priceFileName)
	err := json.Unmarshal(b, &old)
	check(err)
	return old
}

func getFileContent(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func exceed(price float64, old float64, ratio float64) bool {
	return math.Abs(price-float64(old)) > ratio/100.0*price
}
func makeContent(old, now BpiRet) string {
	return fmt.Sprintf("价格从 $%s 到 $%s <br>(%s ~ %s)", old.Bpi.Usd.Rate, now.Bpi.Usd.Rate, old.Time.Updated, now.Time.Updated)
}
func email(address string, password string, content string) {
	m := gomail.NewMessage()

	m.SetAddressHeader("From", address /*"发件人地址"*/, "发件人") // 发件人

	m.SetHeader("To",
		m.FormatAddress(address, "收件人")) // 收件人

	m.SetHeader("Subject", "BTC 价格波动 bitcoin") // 主题

	m.SetBody("text/html", content) // 正文

	d := gomail.NewPlainDialer("smtp.qq.com", 465, address, password) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}

	log.Println("done.发送成功")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
