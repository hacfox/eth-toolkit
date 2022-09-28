package contract

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hacfox/eth-toolkit/utils/hashutil"
	"github.com/hacfox/eth-toolkit/utils/log"
)

type DefaultBlock string

const (
	Latest   DefaultBlock = "latest"
	Earliest DefaultBlock = "earliest"
	Pending  DefaultBlock = "pending"
)

var (
	client = &http.Client{
		Timeout: time.Second * 30,
	}
)

type responseCommon struct {
	ID      int       `json:"-"`
	JSONRPC string    `json:"jsonrpc"`
	Error   *RPCError `json:"error"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Stack string `json:"stack"`
		Name  string `json:"name"`
	} `json:"data"`
}

func (err RPCError) Error() string {
	return fmt.Sprintf("Code=%d, Message=%s, Data=%v", err.Code, err.Message, err.Data)
}

type ethCallRequest struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Gas      uint64 `json:"gas,omitempty"`
	GasPrice uint64 `json:"gasPrice,omitempty"`
	// Value    string `json:"value,omitempty"`
	Data string `json:"data,omitempty"`
}

type ethCallResponse struct {
	responseCommon
	Result string `json:"result"`
}

type EthRPC struct {
	RPCURL string
}

func Trim0xPrefix(hex string) string {
	hex = strings.TrimSpace(hex)
	return strings.TrimPrefix(hex, "0x")
}

func (rpc *EthRPC) GetTokenDecimals(tokenAddr string) string {
	method := "decimals()"
	data := hashutil.Sha3Sig4Bytes(method)

	result := rpc.EthCall("", tokenAddr, 0, 0, data, Latest)
	if Trim0xPrefix(result) == "" {
		return "0"
	}

	return strconv.Itoa(ParseInt(result))
}

func ParseInt(s string) int {
	s = strings.TrimPrefix(s, "0x")
	val, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}

	return int(val)
}

func (rpc *EthRPC) GetTokenName(tokenAddr string) string {
	method := "name()"
	data := hashutil.Sha3Sig4Bytes(method)

	result := rpc.EthCall("", tokenAddr, 0, 0, data, Latest)
	if Trim0xPrefix(result) == "" {
		return ""
	}

	index := 0
	if len(result) > 66 {
		index = 2 + 64 + 64
	}

	name := HexToString(result[index:])
	return strings.TrimSpace(name)
}

func (rpc *EthRPC) GetTokenSymbol(tokenAddr string) string {
	method := "symbol()"
	data := hashutil.Sha3Sig4Bytes(method)
	result := rpc.EthCall("", tokenAddr, 0, 0, data, Latest)
	if Trim0xPrefix(result) == "" {
		return ""
	}
	index := 0
	if len(result) > 66 {
		index = 2 + 64 + 64
	}

	symbol := HexToString(result[index:])
	return strings.TrimSpace(symbol)
}

func (rpc *EthRPC) GetApprove(tokenAddr string) string {
	method := "approve"
	data := hashutil.Sha3Sig4Bytes(method)
	fmt.Println(data)
	result := rpc.EthCall("", tokenAddr, 0, 0, data, Latest)
	fmt.Println(result)
	if Trim0xPrefix(result) == "" {
		return ""
	}

	index := 0
	if len(result) > 66 {
		index = 2 + 64 + 64
	}

	name := HexToString(result[index:])
	return strings.TrimSpace(name)
}

func HexToString(raw string) string {
	raw = strings.TrimPrefix(raw, "0x")
	bytesArr, err := hex.DecodeString(raw)
	if err != nil {
		panic(err)
	}
	return string(bytes.Trim(bytesArr, "\x00")[:])
}

func Hex0xPrefix(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return hex
	}

	return "0x" + hex
}

func (rpc *EthRPC) EthCall(from, to string, gas, gasPrice uint64, data string, q DefaultBlock) string {
	method := "eth_call"
	param := ethCallRequest{
		From:     from,
		To:       to,
		Gas:      gas,
		GasPrice: gasPrice,
		Data:     Hex0xPrefix(data),
	}

	params := []interface{}{param, q}
	reqParam := generateReqParam(method, interface{}(params))

	resp := ethCallResponse{}
	rpc.request(reqParam, &resp)
	return resp.Result
}

type requestCommon struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      uint        `json:"id"`
}

func generateReqParam(method string, params interface{}) requestCommon {
	return requestCommon{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
}

func (rpc *EthRPC) request(reqParam requestCommon, target interface{}) {
	url := rpc.RPCURL
	reqParamBytes, err := json.Marshal(reqParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for {
		if count > 3 {
			log.Warn("Max retry")
			return
		}
		count++
		bodyBytes, _ := post(url, bytes.NewBuffer(reqParamBytes))
		err = json.Unmarshal(bodyBytes, target)
		if err != nil {
			log.Errorf("Request body: %v", string(reqParamBytes))
			log.Errorf("Response: %v", string(bodyBytes))
		} else {
			return
		}

		time.Sleep(3 * time.Second)
	}
}

func post(url string, body io.Reader) ([]byte, error) {
	resp, err := client.Post(url, "application/json", body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP Request Error, StatusCode = %d. Node URL=%s", resp.StatusCode, url)
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respCommon := ethCallResponse{}
	err = json.Unmarshal(bodyBytes, &respCommon)
	if err != nil {
		return nil, err
	}
	if respCommon.Error != nil && respCommon.Error.Message != "" {
		return nil, respCommon.Error
	}

	resp.Body.Close()
	return bodyBytes, nil
}
