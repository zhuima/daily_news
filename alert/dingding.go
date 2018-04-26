/*
提供投递信息的入口，使用钉钉来实现, 返回执行后的状态信息
*/

package alert

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"errors"
)

func SendMsg(dingdingUrl, Title, Content string) (status string, err error){
	formt := `
    {
        "msgtype": "markdown",
        "markdown": {
            "title":"%s",
            "text": "%s"
        }
    }`

	body := fmt.Sprintf(formt, Title, Content)
	jsonValue := []byte(body)

	resp, err := http.Post(dingdingUrl, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		logrus.Fatal("notify", err.Error())
		return "faild", errors.New("Send Msg faild!")
	}

	logrus.Info("Send message", resp)
	return "successful!", nil
}
