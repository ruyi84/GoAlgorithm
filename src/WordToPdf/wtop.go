package main

import (
	"errors"
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"time"
)

func wordConvert1(wordPath string, outPath string, fileCoding int) error {
	now := time.Now()

	//判断文件是否在队列中
	flag := putWordPath(wordPath)
	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在被使用,请稍后再试")

	}
	//从文件锁删除
	defer deleteWordPath(wordPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		defer func() {
			ole.CoUninitialize()
			fmt.Println("word转pdf卸载...")
		}()
	*/

	wordApplication, err := oleutil.CreateObject("Word.Application")
	fmt.Println("wordApplication:err", err)
	if err != nil {
		return err
	}

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("wordAppDispatch:err", err)
	if err != nil {
		return err
	}
	defer func() {
		_, err := oleutil.CallMethod(wordAppDispatch, "Quit")
		fmt.Println("====quit,err:", err)

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "Documents").ToIDispatch()

	//	defer func() {
	//		oleutil.CallMethod(wordFile, "Close")
	//		wordFile.Release()
	//	}()

	fmt.Println("Open:", time.Since(now))
	// 打开文件
	document, err := oleutil.CallMethod(wordFile, "Open", wordPath, true, false, true, "1243")
	if err != nil {
		fmt.Println("文件加密", "====err===", err)
		return errors.New("该文件受密码保护，操作失败！")
	}
	//	defer func() {
	//		oleutil.CallMethod(document.ToIDispatch(), "Close")
	//		document.ToIDispatch().Release()
	//	}()

	result, err := oleutil.CallMethod(document.ToIDispatch(), "SaveAs", outPath, fileCoding)

	fmt.Println(result)
	if err != nil {
		fmt.Println("转换失败", err)
		return errors.New("转换失败!")
	}

	fmt.Println("End:  ", time.Since(now))
	return nil
}
