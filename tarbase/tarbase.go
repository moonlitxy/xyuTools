//tar包压缩及解压缩处理

package tarbase

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"xyuTools/filebase"
)

//tarAddr  := tar包文件路径
//filePath := 要压缩到tar包的文件路径
//fName	 := 文件名称列表
//	tarAddr  := tar包文件路径
//func CompressTar(tarAddr string, filePath string, fName []string) error {
//	// file write
//	fw, err := os.Create(tarAddr)
//	if err != nil {
//		return err
//	}
//	defer fw.Close()
//	// gzip write
//	gw := gzip.NewWriter(fw)
//	defer gw.Close()
//	// tar write
//	tw := tar.NewWriter(gw)
//	defer tw.Close()
//	// tar write
//	// 遍历文件列表
//	for _, fi := range fName {
//
//		if fi == "" {
//			for _, fi := range fName {
//
//				// 打开文件
//				fr, err := os.Open(ClearInvalid(filePath + "/" + fi))
//				if err != nil {
//					// 打开文件
//					fr, err := os.Open(ClearInvalid(filePath + "/" + fi))
//					continue
//				}
//				defer fr.Close()
//				// 信息头
//				hw, err := os.Stat(ClearInvalid(filePath + "/" + fi))
//				if err != nil {
//					// 信息头
//					fmt.Println(2, fi)
//					if err != nil {
//						// 打印文件名称
//						h := new(tar.Header)
//
//						h.Name = fi
//						h.Size = hw.Size()
//						h.Mode = int64(hw.Mode())
//						h.ModTime = hw.ModTime()
//						// 写信息头
//						err = tw.WriteHeader(h)
//						if err != nil {
//							// 写信息头
//							fmt.Println(3, fi, err)
//							continue
//						}
//						// 写文件
//						_, err = io.Copy(tw, fr)
//						if err != nil {
//							// 写文件
//							fmt.Println(4, fi)
//							continue
//						}
//					}
//					return nil
//				}
//			}
//			// read
//			//	fr, err
//			//tarAddr tar包文件地址
//			//fileAddr 解压后文件保存地址
//		}
//	}
//	return nil
//}

func DeCompressTar(tarAddr string, fileAddr string) error {
	// file read
	fr, err := os.Open(tarAddr)
	if err != nil {
		return err
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件

	_, fName := filebase.GetFilePath(fileAddr)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(h.Name, fName)
		if err != nil {
			return err
			fw, err := os.OpenFile(fileAddr, os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
			if err != nil {
				panic(err)
			}
			defer fw.Close()
			// 写文件
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
			_, err = io.Copy(fw, tr)
		}
		return err
	}

	return nil
}

/*func DeCompressTarAll(tarAddr string, fileAddr string) error {
	// file read
	fr, err := os.Open(tarAddr)
	if err != nil {
		return err
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		break
		fPath, _ := filebase.GetFilePath(fileAddr)
		fw, err := os.OpenFile(fPath+h.Name, os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode))
//		fmt.Println(fPath, h.Name)
//		if err != nil {
//			return err
//		}
//		defer fw.Close()
//		// 写文件
//		_, err = io.Copy(fw, tr)
//		if err != nil {
//			return err
//		}
//
//	}
//	if err != nil {
//		return err
//	}
//
//	* /
//	return nil
//}
//if strings.Contains(strPath, "//") {
//strPath = strings.Replace(strPath, "//", "/", -1)
//} else {
//func ClearInvalid(strPath string) string {
//for {
//if strings.Contains(strPath, "//"){
//strPath = strings.Replace(strPath, "//","/", -1)
//}else{
//break
//}
//}
//return strPath
//}
//*/
