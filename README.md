# pinduoduo-sdk
拼多多的golang版本的sdk

```
    package main
    
    import (
    	"fmt"
    
    	"github.com/dcsunny/pinduoduo-sdk/common"
    	"github.com/dcsunny/pinduoduo-sdk/ddk"
    )
    
    func main() {
    	api := ddk.NewDuoduoKe(common.NewContext("client_id", "client_secret"))
    
    	result, err := api.GoodsSearch(nil)
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(result)
    }
 
```