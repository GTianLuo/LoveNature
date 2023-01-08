package util

import "C"
import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"lovenature/conf"
	"mime/multipart"
)

func auth() string {
	accessKey := conf.C.QiNiu.AccessKey
	secretKey := conf.C.QiNiu.SecretKey
	bucket := conf.C.QiNiu.Bucket
	//鉴权
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 7200,
	}

	return putPolicy.UploadToken(mac)
}

func UploadImg(file multipart.File, fileSize int64) (string, error) {

	accessKey := conf.C.QiNiu.AccessKey
	secretKey := conf.C.QiNiu.SecretKey
	bucket := conf.C.QiNiu.Bucket
	qiNiuServer := conf.C.QiNiu.QiNiuServer
	//鉴权
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 7200,
	}

	uploadToken := putPolicy.UploadToken(mac)
	// 上传Config对象
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadongZheJiang2,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, uploadToken, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	return qiNiuServer + ret.Key, nil
}

func DelImg(url string) error {
	accessKey := conf.C.QiNiu.AccessKey
	secretKey := conf.C.QiNiu.SecretKey
	bucket := conf.C.QiNiu.Bucket
	qiNiuServer := conf.C.QiNiu.QiNiuServer
	//鉴权
	mac := qbox.NewMac(accessKey, secretKey)
	//配置属性
	cfg := storage.Config{
		UseHTTPS: false,
		Zone:     &storage.ZoneHuadongZheJiang2,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	//从url中解析出key并删除
	return bucketManager.Delete(bucket, url[len(qiNiuServer):])
}
