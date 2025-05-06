package user

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

func GenerateQRCode(content string, filePath string) error {
	// 创建二维码
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return fmt.Errorf("生成二维码失败: %v", err)
	}

	// 调整二维码大小
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return fmt.Errorf("调整二维码大小失败: %v", err)
	}

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 保存为PNG格式
	err = png.Encode(file, qrCode)
	if err != nil {
		return fmt.Errorf("保存二维码失败: %v", err)
	}

	fmt.Printf("二维码已保存为: %s", filePath)
	return nil
}
