package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/swexbe/zettleIT/api"
)

func main() {

	api.GetAuthkey()

	url := "https://purchase.izettle.com/purchases/v2?limit=10&&descending=true"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer eyJraWQiOiIxNTc0NDM1MDM3MTIzIiwidHlwIjoiSldUIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJpWmV0dGxlIiwiaWF0IjoxNTc0NTE1NjY4LCJhdWQiOiJBUEkiLCJleHAiOjE1NzQ1MjI4NjgsInN1YiI6ImQ1ODZmYWVhLTlhYWUtNDVlMi05M2VhLTNmNTQ4MDJmMGMzMyIsImNsaWVudF9pZCI6ImU5Njc3NDZmLWVjYjMtNDI0My05M2MzLTk1YTYxYTI4MjNmZiIsInNjb3BlIjpbIkFMTDpJTlRFUk5BTCIsIlJFQUQ6Q1VTVE9NRVIiLCJSRUFEOkZJTkFOQ0UiLCJSRUFEOk9OTElORVBBWU1FTlQiLCJSRUFEOlBST0RVQ1QiLCJSRUFEOlBVUkNIQVNFIiwiUkVBRDpVU0VSSU5GTyIsIldSSVRFOkNVU1RPTUVSIiwiV1JJVEU6RklOQU5DRSIsIldSSVRFOk9OTElORVBBWU1FTlQiLCJXUklURTpQUk9EVUNUIiwiV1JJVEU6UFVSQ0hBU0UiLCJXUklURTpSRUZVTkQiLCJXUklURTpVU0VSSU5GTyJdLCJyZW5ld2VkIjp0cnVlLCJ1c2VyIjp7InVzZXJUeXBlIjoiVVNFUiIsInV1aWQiOiJkNTg2ZmFlYS05YWFlLTQ1ZTItOTNlYS0zZjU0ODAyZjBjMzMiLCJvcmdVdWlkIjoiNWVkNWEwMmYtMGFjNC00ZjQ1LThhNGItZTAwZjE1N2Y3ZTEzIiwidXNlclJvbGUiOiJPV05FUiJ9fQ.xFpZe2Yae3u4GGgkBhi07VsbdsG18WXHfhPr6ZAH6jmtFnTe3X6X0aojB-PgV6usVc1go4_7IPzz0PEC9kgT_mYZWnL-dg0kOU1w2gYWnToWgO09kWe9jWOCCGYysjeyiKTuDXOl00jgAgYcG72LBTSKyfGeDmWi7ADaKoH6L_lmzrJHcikL5iF_g12x-5STHfteOQ4PEQ96vYVvi-gFZ5Y_Nd4tHizvIK_Ln3yG-zwJn1IjHSl2qqy_fnLomuHSm46OdqZRnUEXImXt9GrCLtR6EwCsgJcXMoP8_GjEi1mAGrrL18hsBkwA0fa6WK8A2xjYAb2K6CNogclkUGkmEg")
	resp, _ := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
