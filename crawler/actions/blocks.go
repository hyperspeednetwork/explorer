package actions

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// object structure, object method, object call database
// Todo change blockInfo struct

func GetBlock(config conf.Config, log zap.Logger, ) error {
	// get height form database
	// get lastblock height form database
	// check height .< get info  .== sleep .empty sleep
	//time.Sleep(config.Param.BlockInterval) //??
	var block models.BlockInfo
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	for {
		lastBlockHeight, publicHeight := block.GetAimHeightAndBlockHeight()
		//check the height difference again
		if publicHeight > lastBlockHeight {
			for publicHeight > lastBlockHeight {
				lastBlockHeight = lastBlockHeight + 1
				url := config.Remote.Lcd + "/blocks/" + strconv.Itoa(lastBlockHeight)
				resp, err := c.Get(url)
				if err != nil {
					lastBlockHeight = lastBlockHeight - 1
					log.Error("get block info failed", zap.String("url", url))
					log.Sync()
					time.Sleep(time.Second * config.Param.BlockInterval) // sleep 2s ,avoid frequent requests
					continue
				} else {
					// deal with resp
					jsonStr, _ := ioutil.ReadAll(resp.Body)
					err = json.Unmarshal(jsonStr, &block)
					// store block
					intHeight, _ := strconv.Atoi(block.Block.Header.Height)
					block.IntHeight = intHeight
					block.SetInfo(log)
					resp.Body.Close()
				}
				//time.Sleep(time.Millisecond* config.Param.BlockInterval)
				time.Sleep(time.Millisecond* 1)

			}

		}
		time.Sleep(time.Second * config.Param.BlockInterval) // sleep 2s ,todo sleep average time -1s
	}
}
