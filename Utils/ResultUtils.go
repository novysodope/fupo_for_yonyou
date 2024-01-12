package Utils

import "os"

func SaveResultToFile(result, filename string) error {
	//追加不覆盖
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将结果追加写入文件
	_, err = file.WriteString(result + "\n\n")
	if err != nil {
		return err
	}

	return nil
}
