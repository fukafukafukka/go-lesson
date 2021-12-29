package bitflyer

// lesson98_bitflyer.goと関数ダブルためコメントアウト

// // go mod bitflyer
// import (
// 	"bytes"
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"time"
// )

// const baseURL = "https://api.bitflyer.com/v1/"

// type APIClient struct {
// 	key        string
// 	secret     string
// 	httpClient *http.Client
// }

// func New(key, secret string) *APIClient {
// 	apiClient := &APIClient{key, secret, &http.Client{}} // &と*の関係を忘れたので調べる
// 	return apiClient
// }

// func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
// 	// ref) section3/13_mold_replace.go
// 	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
// 	log.Println(timestamp)
// 	message := timestamp + method + endpoint + string(body) // bitflyerの仕様

// 	// ref) section14/lesson75_hmac.go
// 	mac := hmac.New(sha256.New, []byte(api.secret))
// 	mac.Write([]byte(message))
// 	sign := hex.EncodeToString(mac.Sum(nil))
// 	return map[string]string{
// 		"ACCESS-KEY":       api.key,
// 		"ACCESS-TIMESTAMP": timestamp,
// 		"ACCESS-SIGN":      sign,
// 		"Content-Type":     "application/json",
// 	}
// }

// func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
// 	baseURL, err := url.Parse(baseURL) // TODO:中で何をやっているのか調べる
// 	if err != nil {
// 		return
// 	}
// 	apiURL, err := url.Parse(urlPath) // TODO:中で何をやっているのか調べる
// 	if err != nil {
// 		return
// 	}
// 	endpoint := baseURL.ResolveReference(apiURL).String()
// 	log.Printf("action=doRequest endpoint=%s", endpoint)

// 	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer((data)))
// 	if err != nil {
// 		return
// 	}

// 	q := req.URL.Query()
// 	for key, value := range query {
// 		q.Add(key, value)
// 	}
// 	req.URL.RawQuery = q.Encode()

// 	for key, value := range api.header(method, req.URL.RequestURI(), data) {
// 		req.Header.Add(key, value)
// 	}

// 	resp, err := api.httpClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close() // なんでresp.Bodyがcloseメソッドを持っているのか不明。直感的にrespが持ってそう。

// 	body, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, nil
// }

// type Balance struct {
// 	CurrentCode string  `json:"currency_code"`
// 	Amount      float64 `json:"amount"`
// 	Available   float64 `json:"available"`
// }

// // https://lightning.bitflyer.com/docs/#%E8%B3%87%E7%94%A3%E6%AE%8B%E9%AB%98%E3%82%92%E5%8F%96%E5%BE%97
// func (api *APIClient) GetBalance() ([]Balance, error) {
// 	url := "me/getbalance"
// 	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
// 	log.Printf("url=%s resp=%s", url, string(resp))
// 	if err != nil {
// 		log.Printf("action=GetBalance err=%s", err.Error())
// 		return nil, err
// 	}

// 	var balance []Balance
// 	err = json.Unmarshal(resp, &balance)
// 	if err != nil {
// 		log.Printf("action=GetBalance err=%s", err.Error())
// 		return nil, err
// 	}
// 	return balance, nil
// }
