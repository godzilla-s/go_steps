// 椭圆曲率加密算法
import (
    "bytes"
    "compress/gzip"
    "crypto/ecdsa"
    "crypto/elliptic"
    "encoding/hex"
    "errors"
    "fmt"
    "math/big"
    "strings"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/config"
)

var (
    runMode  string
    cfg      config.Configer //全局配置文件
    randKey  string
    randSign string
    prk      *ecdsa.PrivateKey
    puk      ecdsa.PublicKey
    curve    elliptic.Curve
)

func init() {
    var err error
    cfg, err = config.NewConfig("ini", "conf/app.conf")
    if err != nil {
        return
    }
    randSign = cfg.String("RandSign")
    if len(randSign) == 0 {
        return
    }
    randKey = cfg.String("randKey")
    if len(randKey) == 0 {
        return
    }
    beego.Trace("Rand Key =", randKey)
    beego.Trace("Rand Sign =", randSign)
    //根据rand长度，使用相应的加密椭圆参数
    length := len([]byte(randKey))
    if length &lt; 224/8 {
        beego.Error("The length of Rand Key is too small, Crypt init failed, Please reset it again !")
        return
    }
    if length &gt;= 521/8+8 {
        beego.Notice("Rand length =", length, "Using 521 level !")
        curve = elliptic.P521()
    } else if length &gt;= 384/8+8 {
        beego.Notice("Rand length =", length, "Using 384 level !")
        curve = elliptic.P384()
    } else if length &gt;= 256/8+8 {
        beego.Notice("Rand length =", length, "Using 256 level !")
        curve = elliptic.P256()
    } else if length &gt;= 224/8+8 {
        beego.Notice("Rand length =", length, "Using 244 level !")
        curve = elliptic.P224()
    }
    //创建密匙对
    prk, err = ecdsa.GenerateKey(curve, strings.NewReader(randKey))
    if err != nil {
        beego.Error("Crypt init fail,", err, " need = ", curve.Params().BitSize)
        return
    }
    puk = prk.PublicKey
}

//Encrypt 对Text进行加密，返回加密后的字节流
func Sign(text string) (string, error) {
    r, s, err := ecdsa.Sign(strings.NewReader(randSign), prk, []byte(text))
    if err != nil {
        return "", err
    }
    rt, err := r.MarshalText()
    if err != nil {
        return "", err
    }
    st, err := s.MarshalText()
    if err != nil {
        return "", err
    }
    var b bytes.Buffer
    w := gzip.NewWriter(&amp;b)
    defer w.Close()
    _, err = w.Write([]byte(string(rt) + "+" + string(st)))
    if err != nil {
        return "", err
    }
    w.Flush()
    return hex.EncodeToString(b.Bytes()), nil
}

//解密
func getSign(text, byterun []byte) (rint, sint big.Int, err error) {
    r, err := gzip.NewReader(bytes.NewBuffer(byterun))
    if err != nil {
        err = errors.New("decode error," + err.Error())
        return
    }
    defer r.Close()
    buf := make([]byte, 1024)
    count, err := r.Read(buf)
    if err != nil {
        fmt.Println("decode =", err)
        err = errors.New("decode read error," + err.Error())
        return
    }
    rs := strings.Split(string(buf[:count]), "+")
    if len(rs) != 2 {
        err = errors.New("decode fail")
        return
    }
    err = rint.UnmarshalText([]byte(rs[0]))
    if err != nil {
        err = errors.New("decrypt rint fail, " + err.Error())
        return
    }
    err = sint.UnmarshalText([]byte(rs[1]))
    if err != nil {
        err = errors.New("decrypt sint fail, " + err.Error())
        return
    }
    return
}

//Verify 对密文和明文进行匹配校验
func Verify(text, passwd string) (bool, error) {
    byterun, err := hex.DecodeString(passwd)
    if err != nil {
        return false, err
    }
    rint, sint, err := getSign([]byte(text), byterun)
    if err != nil {
        return false, err
    }
    result := ecdsa.Verify(&amp;puk, []byte(text), &amp;rint, &amp;sint)
    return result, nil
}
