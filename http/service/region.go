package service

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"focus/proto"
	"focus/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	protobuf "google.golang.org/protobuf/proto"
)

// 区域数据
type RegionData struct {
	RegionQuery *proto.QueryCurrRegionHttpRsp // Proto
	Base64      string                        // Base64
}

var regions = make(map[string]*RegionData) // 区域数据列表
var regionListResponse string              // 区域列表返回内容

func init() {
	// 接口地址 格式: http(s)://IP:Port
	dispatchDomain := fmt.Sprintf("https://%s:%d", "localhost", 443)
	servers := make([]*proto.RegionSimpleInfo, 0) // 服务器列表

	// 区域临时配置
	regionName := "os_usa"
	regionTitle := "focus"
	regionIp := "localhost"
	regionPort := 22102
	// 区域标识符
	identifier := &proto.RegionSimpleInfo{
		Name:        regionName,
		Title:       regionTitle,
		Type:        "DEV_PUBLIC",
		DispatchUrl: dispatchDomain + "/query_cur_region/" + regionName,
	}
	servers = append(servers, identifier)
	// 区域数据
	regionInfo := &proto.RegionInfo{
		GateserverIp:   regionIp,
		GateserverPort: uint32(regionPort),
		SecretKey:      utils.DISPATCH_SEED,
	}
	updatedQuery := &proto.QueryCurrRegionHttpRsp{
		RegionInfo: regionInfo,
	}
	queryMarshaled, err := protobuf.Marshal(updatedQuery) // 序列化
	if err != nil {
		log.Println("[Region] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}
	// 记录区域数据
	regions[regionName] = &RegionData{
		RegionQuery: updatedQuery,
		Base64:      base64.StdEncoding.EncodeToString(queryMarshaled),
	}

	// 客户端配置文件
	customConfig := []byte("{\"sdkenv\":\"2\",\"checkdevice\":\"false\",\"loadPatch\":\"false\",\"showexception\":\"false\",\"regionConfig\":\"pm|fk|add\",\"downloadMode\":\"0\"}")
	customConfig = utils.XOR(customConfig, utils.DISPATCH_KEY) // XOR加密

	// 总区域列表
	updatedRegionList := &proto.QueryRegionListHttpRsp{
		RegionList:                  servers,
		ClientSecretKey:             utils.DISPATCH_SEED,
		ClientCustomConfigEncrypted: customConfig,
		EnableLoginPc:               true,
	}
	regionMarshaled, err := protobuf.Marshal(updatedRegionList) // 序列化
	if err != nil {
		log.Println("[Region] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}
	// 设置区域列表返回信息
	regionListResponse = base64.StdEncoding.EncodeToString(regionMarshaled)
}

// 查询区域列表服务
func QueryRegionList(c *gin.Context) {
	c.String(http.StatusOK, regionListResponse)
	log.Printf("[Region] 客户端 %s 访问: query_region_list\n", c.RemoteIP())
}

// 查询区域证书服务
func QueryCurRegion(c *gin.Context) {
	regionName := c.Param("region") // 区域名称
	version := c.Query("version")   // 版本号
	region := regions[regionName]   // 区域数据

	// 获取区域数据
	regionData := "CAESGE5vdCBGb3VuZCB2ZXJzaW9uIGNvbmZpZw=="
	if region != nil {
		regionData = region.Base64
	}

	// 判断版本号是否正确
	if strings.Contains(version, "3.1.") {
		if c.Query("dispatchSeed") == "" {
			// 兼容UA补丁?
			c.JSON(http.StatusOK, gin.H{
				"content": "",
				"sign":    "TW9yZSBsb3ZlIGZvciBVQSBQYXRjaCBwbGF5ZXJz",
			})
			return
		}
		keyId := c.Query("key_id") // 公钥编号
		var publicKey *rsa.PublicKey
		// 确定使用的公钥
		if keyId == "3" {
			publicKey = utils.CUR_OS_ENCRYPT_KEY // 国际服
		} else {
			publicKey = utils.CUR_CN_ENCRYPT_KEY // 国服
		}
		regionInfo, err := base64.StdEncoding.DecodeString(regionData)
		if err != nil {
			log.Println("[Region] 解密失败! regionData Base64解密错误, error: ", err.Error())
			return
		}

		// 证书加密
		encryptedBytes, err := utils.RsaEncryptBlock(regionInfo, publicKey)
		if err != nil {
			log.Println("[Region] 加密失败! regionInfo Rsa加密错误, error: ", err.Error())
			return
		}
		// 证书签名
		h := crypto.SHA256.New()
		h.Write([]byte(regionInfo))
		digest := h.Sum(nil)
		signature, err := rsa.SignPKCS1v15(nil, utils.CUR_SIGNING_KEY, crypto.SHA256, digest)
		if err != nil {
			log.Println("[Region] 加密失败! regionInfo PKCS签名错误, error: ", err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"content": base64.StdEncoding.EncodeToString(encryptedBytes),
			"sign":    base64.StdEncoding.EncodeToString(signature),
		})
	} else {
		c.String(http.StatusOK, regionData)
	}
	log.Printf("[Region] 客户端 %s 访问: query_cur_region/%s\n", c.RemoteIP(), regionName)
}

// 获取当前区域的查询
func GetCurrentRegion() *proto.QueryCurrRegionHttpRsp {
	region, ok := regions["os_usa"]
	if ok {
		return region.RegionQuery
	}
	return nil
}
