package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"getCaptcha/model"
	pb "getCaptcha/proto"
	"github.com/afocus/captcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	cap1 := captcha.New()
	err := cap1.SetFont("handler/comic.ttf")
	if err != nil {
		fmt.Println("font err", err)
		return err
	}
	cap1.SetSize(128, 64)
	cap1.SetDisturbance(captcha.NORMAL)
	img, str := cap1.Create(4, captcha.NUM)
	//存储验证码到redis
	err = model.SaveCaptchaCode(str, req.Uuid)
	if err != nil {
		fmt.Println("getCaptcha save err:", err)
		return err
	}
	imgJson, err := json.Marshal(img)
	if err != nil {
		fmt.Println("json:", err)
		return err
	}
	rsp.Img = imgJson
	return nil
}
