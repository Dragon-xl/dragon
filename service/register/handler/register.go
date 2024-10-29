package handler

import (
	"context"
	register "register/proto"
	"register/utils"

	"encoding/json"

	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	"os"
	"register/model"
	"strings"
	"time"
)

type Register struct{}

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
func _main(args []*string, phone string, smsCode string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}
	send := fmt.Sprintf("{\"code\":\"%s\"}", smsCode)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("爱家租房"),
		TemplateCode:  tea.String("SMS_474790402"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(send),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		resp, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		fmt.Println(resp.Body.Message)
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func (e *Register) SendSms(ctx context.Context, req *register.CallRequest, rsp *register.CallResponse) error {
	uuid := req.Uuid
	capCode := req.CapCode
	//验证图片验证码 在model封装好
	if !model.CheckCaptchaCode(uuid, capCode) {
		//校验失败 返回ctx
		fmt.Println("校验失败 captchaCode: ", capCode)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return fmt.Errorf("校验失败 captchaCode: %s", capCode)
	}
	//发送短信
	rand.Seed(time.Now().UnixNano())
	smsCode := fmt.Sprintf("%04d", rand.Int31n(1000000)) //0-9999

	err := _main(tea.StringSlice(os.Args[1:]), req.Phone, smsCode)
	if err != nil {
		//短信发送失败 返回ctx
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return err
	}
	//短信发送成功
	err = model.SaveSmsCode(req.Phone, smsCode)
	//存redis失败
	if err != nil {
		fmt.Println("redis存储手机号失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}
func (e *Register) Register(ctx context.Context, req *register.RegRequest, rsp *register.CallResponse) error {

	//校验短信验证码
	if !model.CheckSmsCode(req.Mobile, req.SmsCode) {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return fmt.Errorf("model.CheckSmsCode(req.Mobile, req.SmsCode) failed")
	}
	//校验通过 存入MYSQL
	err := model.SaveRegisterUser(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}
