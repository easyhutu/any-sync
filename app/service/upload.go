package service

import (
	"bytes"
	"fmt"
	"github.com/easyhutu/any-sync/app/utils/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

func (svr *AnySyncSvr) Upload(ctx *gin.Context) {
	fromBuvid := middleware.WithDevBuvid(ctx)
	fromDev := svr.devs.WithDevice(fromBuvid)
	file, _ := ctx.FormFile("file")
	totalChunks, ok := ctx.GetPostForm("totalChunks")
	fp := fmt.Sprintf("%s/%s/%s", svr.config.ShareFilesPrefix, fromDev.Md5, file.Filename)

	if ok {
		totalChunksInt, _ := strconv.Atoi(totalChunks)
		if totalChunksInt > 1 {
			chunkNumber, _ := strconv.Atoi(ctx.PostForm("chunkNumber"))
			identifier := ctx.PostForm("identifier")

			svr.addMuFileHs(identifier, chunkNumber, file)
			if totalChunksInt > chunkNumber {
				println("split upload...", totalChunksInt, chunkNumber)
				ctx.JSON(200, fmt.Sprintf("%d-%d", chunkNumber, totalChunksInt))
			} else {
				totalSize, _ := strconv.ParseInt(ctx.PostForm("totalSize"), 10, 64)
				println("merge block upload... ", math.MaxInt)
				svr.saveBlockFile(identifier, fp)
				ctx.JSON(200, fromDev.AddFile(file.Filename, fp, totalSize).Desc)
			}
			return
		}
	}

	log.Println(file.Filename, file.Size)
	ctx.SaveUploadedFile(file, fp)

	ctx.JSON(200, fromDev.AddFile(file.Filename, fp, file.Size).Desc)
}

func (svr *AnySyncSvr) Download(ctx *gin.Context) {
	params := &struct {
		Filename string `form:"filename"`
		FromMD5  string `form:"from_md5"`
	}{}
	ctx.ShouldBind(params)
	fmd5 := ctx.Param("fmd5")
	filename := ctx.Param("filename")
	if fmd5 != "" && filename != "" {
		params.FromMD5 = fmd5
		params.Filename = filename
	}
	fmt.Printf("dl params: %+v", params)

	fromDev := svr.devs.WithDevice(middleware.WithDevBuvid(ctx))
	if fromDev == nil {
		ctx.JSON(400, "FILE NOT EXIST")
		return
	}
	println(ctx.GetHeader("User-Agent"), ctx.MustGet("User-Agent"), "fromdev...")
	fromDev.SplitSync()
	for _, sf := range fromDev.SyncFile {
		if sf.FromMd5 == params.FromMD5 {
			for _, ss := range sf.Details {
				if ss.Desc == params.Filename {
					ctx.File(ss.Content)
					return
				}
			}
		}
	}
	ctx.JSON(200, "FILE NOT EXIST")
}

func (svr *AnySyncSvr) addMuFileHs(identifier string, chunkNumber int, file *multipart.FileHeader) {
	_, ok := svr.muFileHs[identifier]
	svr.lock.Lock()
	if !ok {
		svr.muFileHs[identifier] = map[int]*bytes.Buffer{}
	}
	open, _ := file.Open()
	defer open.Close()
	mb := &bytes.Buffer{}
	io.Copy(mb, open)
	svr.muFileHs[identifier][chunkNumber] = mb
	svr.lock.Unlock()
}

func (svr *AnySyncSvr) saveBlockFile(identifier, dst string) error {
	blocks := svr.muFileHs[identifier]

	if err := os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	off := int64(0)
	maxIdx := len(blocks)
	idx := 1
	for {
		b, ok := blocks[idx]
		if b != nil {
			println("write offset:", off)
			out.WriteAt(b.Bytes(), off)
			off += int64(b.Len())
		}
		if !ok || idx > maxIdx {
			break
		}
		idx += 1
	}
	delete(svr.muFileHs, identifier)
	return nil
}
