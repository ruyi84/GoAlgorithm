package main

import (
	"errors"
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

//初始化word组件
func init() {
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

	if err != nil {
		fmt.Println("CoInitialize", err)
		panic(err)
	}
}

// word文件锁
var wordOpenMutex = sync.Mutex{}

// 记录打开的word文件
var wordPathMap = make(map[string]int)

// 判断是否重复，
// 如果不重复，则放入并返回false，则允许word文件打开.
// 如果重复，  则不放入并返回true,则不允许打开word文件.
func putWordPath(filePathStr string) bool {
	wordOpenMutex.Lock()
	defer wordOpenMutex.Unlock()

	_, exitence := wordPathMap[filePathStr]
	if !exitence {
		wordPathMap[filePathStr] = 1
	}
	fmt.Println("加入文件锁中:")
	return exitence
}

//word转换完成后删除
func deleteWordPath(filepathStr string) {
	wordOpenMutex.Lock()
	defer wordOpenMutex.Unlock()

	delete(wordPathMap, filepathStr)
	fmt.Println("从文件锁中删除:")
}

//
// 创建文件夹
// 注：只能用于创建文件夹，如果路径不存在，则创建必要的路径.
//
func MustCreatDir(pathDir string) error {
	err1 := os.MkdirAll(pathDir, 0755)
	return err1
}

//
//更换文件后缀名
// 绝对路径和相对路径都可以
//	注意:必须保证 路径文件是存在的
// 使用场景:如 E:/Test/words/07.docx 转换成  E:/Test/words/07.pdf时
//
func FileChangeExt(filePathStr string, newExt string) string {

	//获取后缀名
	//注意:获取的后缀名的格式为: ".jpg",".pdf"的形式
	extName := filepath.Ext(filePathStr)

	return filePathStr[:len(filePathStr)-len(extName)] + newExt
}

//将word文档转换为pdf格式
func WordPdf(wordPath string, pdfPath string) error {
	return wordConvert(wordPath, pdfPath, 17)
}

//将Excels文件转为PDF格式
func ExcelPdf(excelPath string, pdfPath string) error {
	return excelsConvert(excelPath, pdfPath)
}

//ppt转pdf
func PPT2Pdf(excelPath string, pdfPath string) error {
	return pptConvertPdf(excelPath, pdfPath, 32)
}

//word文档转换为html
func WordConvertHtml(wordPath string, outPath string) {
	wordConvert(wordPath, outPath, 10)
}

//word文档转换
func wordConvert(wordPath string, outPath string, fileCoding int) error {
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

	defer func() {
		oleutil.CallMethod(wordFile, "Close")
		wordFile.Release()
	}()

	// 打开文件
	document, err := oleutil.CallMethod(wordFile, "Open", wordPath, true, false, true, "1243")
	if err != nil {
		fmt.Println("文件加密", "====err===", err)
		return errors.New("该文件受密码保护，操作失败！")
	}
	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close")
		document.ToIDispatch().Release()
	}()

	result, err := oleutil.CallMethod(document.ToIDispatch(), "SaveAs", outPath, fileCoding)

	fmt.Println(result)
	if err != nil {
		fmt.Println("转换失败", err)
		return errors.New("转换失败!")
	}

	return nil
}

//转pdf
func documentOpenAndSave(dispatch *ole.IDispatch, wordPath, outPath string) error {
	fmt.Println("转pdf开始...")
	// 打开文件
	document, err := oleutil.CallMethod(dispatch, "Open", wordPath, true, false, true, "1243")
	fmt.Println("document:err", document, err)
	if err != nil {
		fmt.Println("文件加密", "====err===", err)
		return errors.New("该文件受密码保护，操作失败！")
	}
	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close")
		document.ToIDispatch().Release()
	}()

	result, err := oleutil.CallMethod(document.ToIDispatch(), "SaveAs", outPath, 17)
	fmt.Println("result:err", result, err)
	if err != nil {
		return errors.New("转换失败!")
	}

	return nil
}

/**Excel文档转换
注：如果pdf文件已存在，则会报错
*/
func excelsConvert(excelPath string, outPath string) error {
	//判断文件是否在队列中
	flag := putWordPath(excelPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在被使用,请稍后再试")

	}

	//从文件锁中清除
	defer deleteWordPath(excelPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		defer ole.CoUninitialize()
	*/

	wordApplication, err := oleutil.CreateObject("Excel.Application")
	fmt.Println("err", err)

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("err", err)

	defer func() {
		_, err := oleutil.CallMethod(wordAppDispatch, "Quit")
		fmt.Println("quit01", err)

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "WorkBooks").ToIDispatch()

	defer func() {

		oleutil.CallMethod(wordFile, "Close")
		wordFile.Release()
	}()

	// 打开文件
	document, err := oleutil.CallMethod(wordFile, "Open", excelPath)
	fmt.Println("Open", err)
	//

	defer func() {
		_, err := oleutil.CallMethod(document.ToIDispatch(), "Close", false)
		fmt.Println("Closeerr::", err)

		document.ToIDispatch().Release()
	}()

	_, err = oleutil.CallMethod(document.ToIDispatch(), "ExportAsFixedFormat", 0, outPath)
	fmt.Println("err", err)
	if err != nil {
		return err
	}

	return nil
}

//ppt文档转换pdf
func pptConvertPdf(wordPath string, outPath string, fileCoding int) error {
	//判断文件是否在队列中
	flag := putWordPath(wordPath)
	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在被使用,请稍后再试")

	}
	//从队列中删除
	defer deleteWordPath(wordPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		defer ole.CoUninitialize()
	*/

	word, err := oleutil.CreateObject("Powerpoint.Application")
	fmt.Println("err:", err)
	if err != nil {
		return err
	}
	doc, err := word.QueryInterface(ole.IID_IDispatch)
	fmt.Println("err", err)
	defer func() {
		_, err2 := oleutil.CallMethod(doc, "Quit")
		fmt.Println("err2", err2)

		doc.Release()
		word.Release()
	}()

	//实例化模板
	document, err := oleutil.GetProperty(doc, "Presentations")
	fmt.Println("err", err)

	//打开
	wordbooks, err := oleutil.CallMethod(document.ToIDispatch(), "Open", wordPath, true, true, false)
	fmt.Println("err", err)
	if err != nil {
		return err
	}
	defer func() {
		_, err2 := oleutil.CallMethod(wordbooks.ToIDispatch(), "Close")
		fmt.Println("err2", err2)
	}()

	_, err = oleutil.CallMethod(wordbooks.ToIDispatch(), "SaveAs", outPath, fileCoding)
	fmt.Println("err", err)
	if err != nil {
		return err
	}

	return nil
}

// 生成pdf的缩略图
func PdfThumbnail(pdfPath string, imgPath string) {

	cmdDo := exec.Command("gm", "convert", pdfPath+"[0]", imgPath)
	err := cmdDo.Run()
	fmt.Println("CMdeiaThumbnail,imgPath", "err:", err)
}

/*
图片压缩
注：只能压缩jpg,png图片
*/
func ImageCompression(imgPath string) {
	fileOrigin, err1 := os.Open(imgPath)
	if err1 != nil {
		fmt.Println("open err", err1)
		return
	}
	fileinfo, err2 := fileOrigin.Stat()
	if err2 != nil {
		fmt.Println("Stat err", err2)
		return
	}
	fmt.Println("fileinfo.Size:", fileinfo.Size(), err2)
	// 文件大小 大于 512K 则 准备压缩
	if fileinfo.Size() < (512 * 1024) {
		return
	}
	fileExt := filepath.Ext(imgPath)
	fileExt = strings.ToLower(fileExt)
	fmt.Println("fileExt:", fileExt)

	var imgsBounds image.Rectangle
	if fileExt == ".jpg" {
		originImg, err2 := jpeg.Decode(fileOrigin)
		defer fileOrigin.Close()
		if err2 != nil {
			fmt.Println("jpeg.Decode err", err2)
			return
		}

		imgsBounds = originImg.Bounds()

	} else if fileExt == ".png" {
		originImg, err2 := png.Decode(fileOrigin)
		defer fileOrigin.Close()
		if err2 != nil {
			fmt.Println("png.Decode err", err2)
			return
		}

		imgsBounds = originImg.Bounds()
	} else {
		return
	}

	// 图片的宽度,图片的高度
	fmt.Println("imgsBounds.Dx():", imgsBounds.Dx(), "imgsBounds.Dy():", imgsBounds.Dy())
	defalutMaxSize := 1500

	//如果宽度或者高度超过预设值，则压缩为原始尺寸的50%
	if imgsBounds.Dx() > defalutMaxSize || imgsBounds.Dy() > defalutMaxSize {
		cmdDo := exec.Command("gm", "convert", "-density", "288", "-geometry", "50%", imgPath, imgPath)
		err := cmdDo.Run()
		fmt.Println("GraphicsMagick convert,imgPath", imgPath, "err:", err)
	} else {
		// 只压缩分辨率
		cmdDo := exec.Command("gm", "convert", "-density", "288", imgPath, imgPath)
		err := cmdDo.Run()
		fmt.Println("GraphicsMagick convert,imgPath", imgPath, "err:", err)
	}
}

/**
	根据标签名称从标签列表中获取标签
	返回值:
		标签的句柄;
		是否出现异常;

 	注:会判断某一个标签是否存在,如果不存在，则会按异常处理.
*/
func getBookMarkIDis(marksIDis *ole.IDispatch, markName string) (*ole.IDispatch, error) {
	//判断书签是否存在
	ExistsItem, err := oleutil.CallMethod(marksIDis, "Exists", markName)

	if err != nil {
		return nil, err
	}

	if ExistsItem.Value() == false {
		// 没找到指定的书签
		return nil, errors.New("标签:" + markName + "不存在")
	}

	//不可以直接根据名字选择Item，会崩溃.
	// 必须先判断是否存在
	//
	markItem, err := oleutil.CallMethod(marksIDis, "Item", markName)
	if err != nil {
		return nil, err
	}

	return markItem.ToIDispatch(), nil
}

//修改书签的文字
func changeBookMarkText(bookMarks *ole.IDispatch, markName, newText string) {

	markItemDis, errItem := getBookMarkIDis(bookMarks, markName)
	if errItem != nil {
		return
	}

	RangeResult, err := oleutil.GetProperty(markItemDis, "Range")
	fmt.Println("RangeResult:", RangeResult, ",err:", err)

	putResult, err := oleutil.PutProperty(RangeResult.ToIDispatch(), "Text", newText)
	fmt.Println("putResult:", putResult, ",err:", err)

	//重新添加书签
	addItem, errItem := oleutil.CallMethod(bookMarks, "Add", markName, RangeResult)
	fmt.Println("addItem:", addItem, ",errItem:", errItem)
}

//文件复制
//src为目标文件  绝对路径
//des为复制后文件  绝对路径
//
func CopyFile(des, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

//文件内容
//fullPath：模板文件
//appendFile：用户上传内容文件
//办公室流程,修改字段
func FileOfficeInsertAndPdf(fullPath, pdfPath, leaderSign, codeNum string) error {
	//判断文件是否被占用
	flag := putWordPath(fullPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在处理,请稍后再试...")
	}
	defer deleteWordPath(fullPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		defer ole.CoUninitialize()
	*/

	wordApplication, err := oleutil.CreateObject("Word.Application")
	fmt.Println("Application", err)
	if err != nil {
		return err
	}

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("Application", err)
	if err != nil {
		return err
	}

	defer func() {
		oleutil.MustCallMethod(wordAppDispatch, "Quit")

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "Documents")

	defer func() {

		oleutil.CallMethod(wordFile.ToIDispatch(), "Close")
		wordFile.ToIDispatch().Release()
	}()

	//打开模板
	document, err := DocmentOpen(wordFile, fullPath)
	if err != nil {
		return err
	}

	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close")
		document.ToIDispatch().Release()
	}()

	selection, selectErr := oleutil.GetProperty(wordAppDispatch, "Selection")

	fmt.Println("selection:", selection, ",selectErr:", selectErr)
	//通过获取标签然后选中,间接实现跳转到指定位置.
	BookMarks, selectErr := oleutil.GetProperty(document.ToIDispatch(), "BookMarks")
	fmt.Println("BookMarks:", BookMarks, ",selectErr:", selectErr)

	//替换发稿人等
	changeBookMarkText(BookMarks.ToIDispatch(), "签批人", leaderSign)
	changeBookMarkText(BookMarks.ToIDispatch(), "编号", codeNum)
	//删除未授予文号的pdf文件
	os.RemoveAll(pdfPath)

	/*
		//转pdf
		//转pdf
		e := documentOpenAndSave(wordFile.ToIDispatch(), fullPath, pdfPath)
		if e != nil {
			return e
		}
	*/

	result, err := oleutil.CallMethod(document.ToIDispatch(), "SaveAs", pdfPath, 17)
	fmt.Println("result:err", result, err)
	if err != nil {
		return errors.New("转换失败!")
	}

	return nil
}

//办公室通过
func FileOfficeInsert(fullPath, leaderSign, codeNum string) error {
	//判断文件是否被占用
	flag := putWordPath(fullPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在处理,请稍后再试...")
	}
	defer deleteWordPath(fullPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		fmt.Println("插入CoInitialize:", err)
		defer ole.CoUninitialize()
	*/

	wordApplication, err := oleutil.CreateObject("Word.Application")
	fmt.Println("Application", err)
	if err != nil {
		return err
	}

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("Application", err)
	if err != nil {
		return err
	}

	defer func() {

		oleutil.MustCallMethod(wordAppDispatch, "Quit")

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "Documents").ToIDispatch()
	defer func() {

		oleutil.CallMethod(wordFile, "Close")
		wordFile.Release()
	}()

	// 打开文件
	document, err := oleutil.CallMethod(wordFile, "Open", fullPath, true, false, true, "1243")
	if err != nil {
		fmt.Println("文件加密", "err:", err)

		return errors.New("该文件受密码保护，操作失败！")
	}

	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close", true)
		document.ToIDispatch().Release()
	}()

	selection, selectErr := oleutil.GetProperty(wordAppDispatch, "Selection")

	fmt.Println("selection:", selection, ",selectErr:", selectErr)
	//通过获取标签然后选中,间接实现跳转到指定位置.
	BookMarks, selectErr := oleutil.GetProperty(document.ToIDispatch(), "BookMarks")
	fmt.Println("BookMarks:", BookMarks, ",selectErr:", selectErr)

	//替换发稿人等
	changeBookMarkText(BookMarks.ToIDispatch(), "签批人", leaderSign)
	changeBookMarkText(BookMarks.ToIDispatch(), "编号", codeNum)

	return nil
}

//文件内容
//fullPath：模板文件
//appendFile：用户上传内容文件
//拟稿发文流程,修改word字段
func FileInserts(fullPath, appendFile, drafter, checkDrafter, deptName, codeType, year, month, day, hour, minute, grade, phoneNum, mainSend, copySend string) error {
	//判断文件是否在队列中
	flag := putWordPath(fullPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在被使用,请稍后再试...")

	}
	defer deleteWordPath(fullPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		// fmt.Println("插入:CoInitialize", err)

		defer func() {
			ole.CoUninitialize()
			fmt.Println("插入卸载...")
		}()
	*/

	wordApplication, _ := oleutil.CreateObject("Word.Application")

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("Application", err)
	if err != nil {
		return err
	}

	defer func() {
		oleutil.MustCallMethod(wordAppDispatch, "Quit")

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "Documents").ToIDispatch()
	defer func() {

		oleutil.CallMethod(wordFile, "Close")
		wordFile.Release()
	}()

	// 打开文件
	document, err := oleutil.CallMethod(wordFile, "Open", fullPath, true, false, true, "1243")
	fmt.Println("open:", err)
	if err != nil {
		fmt.Println("文件加密", "err", err)
		return errors.New("文件受密码保护，操作失败！")
	}
	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close", true)
		document.ToIDispatch().Release()
	}()

	selection, selectErr := oleutil.GetProperty(wordAppDispatch, "Selection")

	fmt.Println("selection:", selection, ",selectErr:", selectErr)
	//通过获取标签然后选中,间接实现跳转到指定位置.
	BookMarks, selectErr := oleutil.GetProperty(document.ToIDispatch(), "BookMarks")
	fmt.Println("BookMarks:", BookMarks, ",selectErr:", selectErr)

	//替换发稿人等

	changeBookMarkText(BookMarks.ToIDispatch(), "发稿人", drafter)
	changeBookMarkText(BookMarks.ToIDispatch(), "核稿人", checkDrafter)
	changeBookMarkText(BookMarks.ToIDispatch(), "拟稿部门", deptName)
	changeBookMarkText(BookMarks.ToIDispatch(), "文号类型", codeType)
	changeBookMarkText(BookMarks.ToIDispatch(), "年份", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "等级", grade)
	changeBookMarkText(BookMarks.ToIDispatch(), "年", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "月", month)
	changeBookMarkText(BookMarks.ToIDispatch(), "日", day)
	changeBookMarkText(BookMarks.ToIDispatch(), "时", hour)
	changeBookMarkText(BookMarks.ToIDispatch(), "分", minute)
	changeBookMarkText(BookMarks.ToIDispatch(), "联系电话", phoneNum)

	changeBookMarkText(BookMarks.ToIDispatch(), "主送", mainSend)
	changeBookMarkText(BookMarks.ToIDispatch(), "抄送", copySend)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发年份", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发月份", month)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发日", day)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发部门", deptName)

	markName := "公文正文"
	markItemIDis, errItem := getBookMarkIDis(BookMarks.ToIDispatch(), markName)
	fmt.Println("errItem", errItem)
	if errItem != nil {
		return errors.New("插入错误")
	}

	// 标签执行选中
	// 使文件在当前标签的位置执行插入
	SelectItem, errItem := oleutil.CallMethod(markItemIDis, "Select")

	fmt.Println("SelectItem:", SelectItem, ",errItem:", errItem)
	//将另一个文件填入到当前文件
	insertFile, errMethod := oleutil.CallMethod(selection.ToIDispatch(), "InsertFile", appendFile)
	fmt.Println("errItem", errMethod)
	if errMethod != nil {

		return errMethod
	}
	fmt.Println("insertFile:", insertFile, ",errMethod:", errMethod)

	resizeWordShapeLineSpanOneHeheight(document)

	return nil
}

// 将word中包含图片的行设置为单倍行距
func resizeWordShapeLineSpanOneHeheight(document *ole.VARIANT) {

	staticResult, _ := oleutil.CallMethod(document.ToIDispatch(), "ComputeStatistics", 1)
	lineAccount := int(staticResult.Value().(int32))
	//fmt.Println("========staticResult=====",lineAccount,err)

	contentResult, contentErr := oleutil.GetProperty(document.ToIDispatch(), "Content")
	fmt.Println("content:", contentResult, ",contentErr:", contentErr)

	for lineIndex := 0; lineIndex < lineAccount; lineIndex++ {

		lineRange, _ := getSpecailLine(document, contentResult, lineIndex, lineAccount)

		inlineShapes, _ := oleutil.GetProperty(lineRange.ToIDispatch(), "InlineShapes")

		shapesCount, _ := oleutil.GetProperty(inlineShapes.ToIDispatch(), "Count")
		shapeAccount := int(shapesCount.Value().(int32))

		if shapeAccount > 0 {

			paragraphFormat, _ := oleutil.GetProperty(lineRange.ToIDispatch(), "ParagraphFormat")
			oleutil.PutProperty(paragraphFormat.ToIDispatch(), "LineSpacingRule", 0)
		}

	}
}

func getSpecailLine(documentResult, contentResult *ole.VARIANT, lineIndex, maxLine int) (*ole.VARIANT, error) {

	GotoStart, _ := oleutil.CallMethod(documentResult.ToIDispatch(), "GoTo", 3, 1, lineIndex)

	startResult, _ := oleutil.GetProperty(GotoStart.ToIDispatch(), "Start")

	var endResult *ole.VARIANT
	if lineIndex >= maxLine {
		endResult, _ = oleutil.GetProperty(contentResult.ToIDispatch(), "End")
	} else {
		GotoEnd, _ := oleutil.CallMethod(documentResult.ToIDispatch(), "GoTo", 3, 1, lineIndex+1)

		endResult, _ = oleutil.GetProperty(GotoEnd.ToIDispatch(), "Start")

	}

	RangeResult, contentErr := oleutil.CallMethod(documentResult.ToIDispatch(), "Range", startResult, endResult)
	//fmt.Println("RangeResult:", RangeResult, ",contentErr:", contentErr)

	return RangeResult, contentErr

}

//拟稿发文
func FileInsertsAndPdf(fullPath, appendFile, pdfpath, drafter, checkDrafter, deptName, codeType, year, month, day, hour, minute, grade, phoneNum, mainSend, copySend string) error {
	//判断文件是否在队列中
	flag := putWordPath(fullPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在被使用,请稍后再试...")

	}
	//从文件锁删除
	defer deleteWordPath(fullPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		fmt.Println("插入:CoInitialize", err)

		defer func() {
			ole.CoUninitialize()
			fmt.Println("插入卸载...")
		}()
	*/

	wordApplication, _ := oleutil.CreateObject("Word.Application")

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("Application", err)
	if err != nil {
		return err
	}
	defer func() {
		_, err := oleutil.CallMethod(wordAppDispatch, "Quit")
		fmt.Println("====quit,err:", err)

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "Documents")
	defer func() {

		oleutil.CallMethod(wordFile.ToIDispatch(), "Close")
		wordFile.ToIDispatch().Release()
	}()

	//判断正文是否加密
	/*
		_, err = judgeDocmentOpen(wordFile, appendFile)
		if err != nil {
			return err
		}
	*/

	//打开模板
	document, err := DocmentOpen(wordFile, fullPath)
	if err != nil {
		return err
	}

	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close")
		document.ToIDispatch().Release()
	}()

	selection, selectErr := oleutil.GetProperty(wordAppDispatch, "Selection")

	fmt.Println("selection:", selection, ",selectErr:", selectErr)
	//通过获取标签然后选中,间接实现跳转到指定位置.
	BookMarks, selectErr := oleutil.GetProperty(document.ToIDispatch(), "BookMarks")
	fmt.Println("BookMarks:", BookMarks, ",selectErr:", selectErr)

	//替换发稿人等

	changeBookMarkText(BookMarks.ToIDispatch(), "发稿人", drafter)
	changeBookMarkText(BookMarks.ToIDispatch(), "核稿人", checkDrafter)
	changeBookMarkText(BookMarks.ToIDispatch(), "拟稿部门", deptName)
	changeBookMarkText(BookMarks.ToIDispatch(), "文号类型", codeType)
	changeBookMarkText(BookMarks.ToIDispatch(), "年份", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "等级", grade)
	changeBookMarkText(BookMarks.ToIDispatch(), "年", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "月", month)
	changeBookMarkText(BookMarks.ToIDispatch(), "日", day)
	changeBookMarkText(BookMarks.ToIDispatch(), "时", hour)
	changeBookMarkText(BookMarks.ToIDispatch(), "分", minute)
	changeBookMarkText(BookMarks.ToIDispatch(), "联系电话", phoneNum)

	changeBookMarkText(BookMarks.ToIDispatch(), "主送", mainSend)
	changeBookMarkText(BookMarks.ToIDispatch(), "抄送", copySend)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发年份", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发月份", month)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发日", day)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发部门", deptName)

	markName := "公文正文"
	markItemIDis, errItem := getBookMarkIDis(BookMarks.ToIDispatch(), markName)
	fmt.Println("errItem", errItem)
	if errItem != nil {
		return errors.New("插入错误")
	}

	// 标签执行选中
	// 使文件在当前标签的位置执行插入
	SelectItem, errItem := oleutil.CallMethod(markItemIDis, "Select")

	fmt.Println("SelectItem:", SelectItem, ",errItem:", errItem)
	//将另一个文件填入到当前文件
	insertFile, errMethod := oleutil.CallMethod(selection.ToIDispatch(), "InsertFile", appendFile)
	fmt.Println("errItem", errMethod)
	if errMethod != nil {

		return errMethod
	}
	fmt.Println("insertFile:", insertFile, ",errMethod:", errMethod)

	resizeWordShapeLineSpanOneHeheight(document)
	oleutil.MustCallMethod(document.ToIDispatch(), "Close", true)
	document.ToIDispatch().Release()
	document = oleutil.MustCallMethod(wordFile.ToIDispatch(), "Open", fullPath)

	/*
		//转pdf
		e := documentOpenAndSave(wordFile.ToIDispatch(), fullPath, pdfpath)
		fmt.Println("转pdf:err", e)
		if e != nil {
			return e
		}
	*/

	result, err := oleutil.CallMethod(document.ToIDispatch(), "SaveAs", pdfpath, 17)
	fmt.Println("result:err", result, err)
	if err != nil {
		return errors.New("转换失败!")
	}

	return nil
}

//审核通过拟稿发文
func PassedFileInsertsAndPdf(fullPath, appendFile, pdfpath, drafter, checkDrafter, deptName, codeType, year, month, day, hour, minute, grade, phoneNum, mainSend, copySend, leaderSign, codeNum string) error {
	//判断文件是否在队列中
	flag := putWordPath(fullPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在被使用,请稍后再试...")

	}
	//从文件锁删除
	defer deleteWordPath(fullPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		fmt.Println("插入:CoInitialize", err)

		defer func() {
			ole.CoUninitialize()
			fmt.Println("插入卸载...")
		}()
	*/

	wordApplication, _ := oleutil.CreateObject("Word.Application")

	wordAppDispatch, err := wordApplication.QueryInterface(ole.IID_IDispatch)
	fmt.Println("Application", err)
	if err != nil {
		return err
	}
	defer func() {
		_, err := oleutil.CallMethod(wordAppDispatch, "Quit")
		fmt.Println("====quit,err:", err)

		wordAppDispatch.Release()
		wordApplication.Release()
	}()

	wordFile := oleutil.MustGetProperty(wordAppDispatch, "Documents")

	defer func() {
		oleutil.CallMethod(wordFile.ToIDispatch(), "Close")
		wordFile.ToIDispatch().Release()
	}()

	//判断正文是否加密
	/*
		_, err = judgeDocmentOpen(wordFile, appendFile)
		if err != nil {
			return err
		}
	*/

	//打开模板
	document, err := DocmentOpen(wordFile, fullPath)
	if err != nil {
		return err
	}

	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close")
		document.ToIDispatch().Release()
	}()

	selection, selectErr := oleutil.GetProperty(wordAppDispatch, "Selection")

	fmt.Println("selection:", selection, ",selectErr:", selectErr)
	//通过获取标签然后选中,间接实现跳转到指定位置.
	BookMarks, selectErr := oleutil.GetProperty(document.ToIDispatch(), "BookMarks")
	fmt.Println("BookMarks:", BookMarks, ",selectErr:", selectErr)

	//替换发稿人等

	changeBookMarkText(BookMarks.ToIDispatch(), "发稿人", drafter)
	changeBookMarkText(BookMarks.ToIDispatch(), "核稿人", checkDrafter)
	changeBookMarkText(BookMarks.ToIDispatch(), "拟稿部门", deptName)
	changeBookMarkText(BookMarks.ToIDispatch(), "文号类型", codeType)
	changeBookMarkText(BookMarks.ToIDispatch(), "年份", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "等级", grade)
	changeBookMarkText(BookMarks.ToIDispatch(), "年", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "月", month)
	changeBookMarkText(BookMarks.ToIDispatch(), "日", day)
	changeBookMarkText(BookMarks.ToIDispatch(), "时", hour)
	changeBookMarkText(BookMarks.ToIDispatch(), "分", minute)
	changeBookMarkText(BookMarks.ToIDispatch(), "联系电话", phoneNum)

	changeBookMarkText(BookMarks.ToIDispatch(), "主送", mainSend)
	changeBookMarkText(BookMarks.ToIDispatch(), "抄送", copySend)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发年份", year)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发月份", month)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发日", day)
	changeBookMarkText(BookMarks.ToIDispatch(), "印发部门", deptName)
	//替换发稿人等
	changeBookMarkText(BookMarks.ToIDispatch(), "签批人", leaderSign)
	changeBookMarkText(BookMarks.ToIDispatch(), "编号", codeNum)

	markName := "公文正文"
	markItemIDis, errItem := getBookMarkIDis(BookMarks.ToIDispatch(), markName)
	fmt.Println("errItem", errItem)
	if errItem != nil {
		return errors.New("插入错误")
	}

	// 标签执行选中
	// 使文件在当前标签的位置执行插入
	SelectItem, errItem := oleutil.CallMethod(markItemIDis, "Select")

	fmt.Println("SelectItem:", SelectItem, ",errItem:", errItem)
	//将另一个文件填入到当前文件
	insertFile, errMethod := oleutil.CallMethod(selection.ToIDispatch(), "InsertFile", appendFile)
	fmt.Println("errItem", errMethod)
	if errMethod != nil {

		return errMethod
	}
	fmt.Println("insertFile:", insertFile, ",errMethod:", errMethod)

	resizeWordShapeLineSpanOneHeheight(document)
	oleutil.MustCallMethod(document.ToIDispatch(), "Close", true)
	document.ToIDispatch().Release()
	document = oleutil.MustCallMethod(wordFile.ToIDispatch(), "Open", fullPath)

	/*
		//转pdf
		e := documentOpenAndSave(wordFile.ToIDispatch(), fullPath, pdfpath)
		fmt.Println("转pdf", e)
		if e != nil {
			return e
		}
	*/

	result, err := oleutil.CallMethod(document.ToIDispatch(), "SaveAs", pdfpath, 17)
	fmt.Println("result:err", result, err)
	if err != nil {
		return errors.New("转换失败!")
	}

	return nil
}

//打开文档 判断是否加密 执行close
func judgeDocmentOpen(wordFile *ole.VARIANT, fullPath string) (*ole.VARIANT, error) {
	//fmt.Println("open document")
	document, err := oleutil.CallMethod(wordFile.ToIDispatch(), "Open", fullPath, true, false, true, "1243")
	if err != nil {
		return nil, errors.New("该文件受密码保护,操作失败!")
	}
	defer func() {
		oleutil.CallMethod(document.ToIDispatch(), "Close")
		document.ToIDispatch().Release()
	}()
	return document, nil
}

//打开文档
func DocmentOpen(wordFile *ole.VARIANT, fullPath string) (*ole.VARIANT, error) {

	document, err := oleutil.CallMethod(wordFile.ToIDispatch(), "Open", fullPath, true, false, true, "1243")
	if err != nil {
		return nil, err
	}
	return document, nil
}

//Excel转Thml
//excelPath Excel路径
//htmlPath 路径
func ExcelToHtml(excelPath, htmlPath string) error {
	//判断文件是否被占用
	flag := putWordPath(excelPath)

	if flag {
		fmt.Println("文件正在使用:")
		return errors.New("文件正在处理,请稍后再试...")
	}

	//删除锁
	defer deleteWordPath(excelPath)

	/*
		err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)

		if err != nil {
			fmt.Println("======CoInitialize=======", err)
			return err
		}
		fmt.Println("CoInitialize", err)
		defer ole.CoUninitialize()
	*/

	unknown, _ := oleutil.CreateObject("Excel.Application")
	excel, _ := unknown.QueryInterface(ole.IID_IDispatch)

	defer func() {
		_, e := oleutil.CallMethod(excel, "Quit")
		fmt.Println("quit", e)

		excel.Release()
		unknown.Release()

	}()

	workbooksProperty, _ := oleutil.GetProperty(excel, "Workbooks")
	workbooks := workbooksProperty.ToIDispatch()
	defer func() {
		oleutil.CallMethod(workbooks, "Close")
		workbooks.Release()
	}()

	workbook, err := oleutil.CallMethod(workbooks, "Open", excelPath)
	//	fmt.Println("open", err)

	if err != nil {
		return errors.New("打开Excel错误")
	}
	defer func() {
		oleutil.CallMethod(workbook.ToIDispatch(), "Close", false)

		workbook.ToIDispatch().Release()
	}()

	sheets := oleutil.MustGetProperty(excel, "Sheets").ToIDispatch()
	sheetCount := (int)(oleutil.MustGetProperty(sheets, "Count").Val)
	fmt.Println("sheet count=", sheetCount)
	defer sheets.Release()

	// 第一个sheet
	worksheet := oleutil.MustGetProperty(workbook.ToIDispatch(), "Worksheets", 1).ToIDispatch()
	defer func() {

		worksheet.Release()
	}()

	_, err = oleutil.CallMethod(worksheet, "SaveAs", htmlPath, 44)
	fmt.Println("saveAs:", err)

	return nil
}
