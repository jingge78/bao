package merchant

import (
	"bytes"
	"fmt"
	"github.com/tuotoo/qrcode"
	"image"
	"image/png"
	"os"
)

// 从PNG文件解析二维码
func DecodeQRCode(filename string) (string, error) {
	// 打开文件
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前工作目录时出错: %v\n", err)
		return "", nil
	}
	fmt.Println(wd)
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("解码图片失败: %v", err)
	}

	// 创建字节缓冲区
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return "", fmt.Errorf("编码图片失败: %v", err)
	}

	// 解析二维码
	qrmatrix, err := qrcode.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", fmt.Errorf("解析二维码失败: %v", err)
	}
	return qrmatrix.Content, nil
}
