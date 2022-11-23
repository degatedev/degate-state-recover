package poseidon

import (
	"log"
	"math"
	"math/big"
	"staterecovery/common/util"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	// "github.com/stretchr/testify/assert"
)


func TestPoseidonHash(t *testing.T) {
	// b0 := big.NewInt(0)
	// b1 := big.NewInt(1)
	// b2 := big.NewInt(2)

	// h, err := Hash([]*big.Int{b1})
	// assert.Nil(t, err)
	// assert.Equal(t,
	// 	"11316722965829087614032985243432266723826890185209218714357779037968059437034",
	// 	h.String())

	// h, err = Hash([]*big.Int{b1, b2})
	// assert.Nil(t, err)
	// assert.Equal(t,
	// 	"18034868597434240293665220970421168445584131937984445797953356852217236273181",
	// 	h.String())

	// h, err = Hash([]*big.Int{b0, b1, b2, b0, b1, b2, b0, b1, b2})
	// assert.Nil(t, err)
	// assert.Equal(t,
	// 	"901171399164243716242986386548401611619556238528027833916499064730590161382",
	// 	h.String())

	// h, err = Hash([]*big.Int{b0, b1, b2, b0, b1, b2})
	// assert.Nil(t, err)
	// assert.Equal(t,
	// 	"13102070988478037395154308865607405548746274688434317574093002894058697028363",
	// 	h.String())

	// std::string owner = "301049204097205545172088324898475594154498399555";
    // std::string pubX = "12838628619158778216939803703555672017729495653381372941964259226476861150226";
    // std::string pubY = "18804586892836094061286413079049561190038087497348695885035576206651318166610";
    // std::string appPubX = "0";
    // std::string appPubY = "0";
    // std::string nonce = "11";
    // std::string disableAppKeySpotTrade = "0";
    // std::string disableAppKeyWithdraw = "0";
    // std::string disableAppKeyTransferToOther = "0";
    // std::string balanceRoot = "14787015035826428939028881958841326039511687960717659315562040032908955599541";
    // std::string storageRoot = "19888613002804776873994574691676194087433727064046896651235023891396063669836";
	b0, _ := decimal.NewFromString("301049204097205545172088324898475594154498399555")
	b1, _ := decimal.NewFromString("12838628619158778216939803703555672017729495653381372941964259226476861150226")
	b2, _ := decimal.NewFromString("18804586892836094061286413079049561190038087497348695885035576206651318166610")
	b3, _ := decimal.NewFromString("0")
	b4, _ := decimal.NewFromString("0")
	b5, _ := decimal.NewFromString("11")
	b6, _ := decimal.NewFromString("0")
	b7, _ := decimal.NewFromString("0")
	b8, _ := decimal.NewFromString("0")
	b9, _ := decimal.NewFromString("14787015035826428939028881958841326039511687960717659315562040032908955599541")
	b10, _ := decimal.NewFromString("19888613002804776873994574691676194087433727064046896651235023891396063669836")
	second := time.Now().UnixNano()
	h, err := Hash([]*big.Int{
		b0.BigInt(), 
		b1.BigInt(), 
		b2.BigInt(), 
		b3.BigInt(), 
		b4.BigInt(), 
		b5.BigInt(), 
		b6.BigInt(), 
		b7.BigInt(), 
		b8.BigInt(), 
		b9.BigInt(), 
		b10.BigInt(),
	})
	if err != nil {
		log.Println("err:", err.Error())
		return
	}
	// 0.489000 ms
	log.Println("time:", strconv.FormatInt(time.Now().UnixNano() - second, 10))
	log.Println("h:", h.String())


	a0, _ := decimal.NewFromString("20000")
	h, err = HashAccurate([]*big.Int{
		a0.BigInt(), 
		// a0.BigInt(), 
		// a0.BigInt(), 
		// a0.BigInt(), 
	}, 5)
	if err != nil {
		log.Println("err:", err.Error())
		return
	}
	log.Println("h:", h.String())
}

func TestMerkleTree(t *testing.T) {
	// std::string one = "6224207326931433006311810030489209471396486193319112129783570885466938759508";
    // std::string two = "14210355018590017149382026284573545675731681758516104177472367813701353963238";
    // std::string three = "14210355018590017149382026284573545675731681758516104177472367813701353963238";
    // std::string four = "14210355018590017149382026284573545675731681758516104177472367813701353963238";
	a0, _ := decimal.NewFromString("6224207326931433006311810030489209471396486193319112129783570885466938759508")
	a1, _ := decimal.NewFromString("14210355018590017149382026284573545675731681758516104177472367813701353963238")
	a2, _ := decimal.NewFromString("14210355018590017149382026284573545675731681758516104177472367813701353963238")
	a3, _ := decimal.NewFromString("14210355018590017149382026284573545675731681758516104177472367813701353963238")
	h, err := HashAccurate([]*big.Int{
		a0.BigInt(), 
		a1.BigInt(), 
		a2.BigInt(), 
		a3.BigInt(), 
	}, 5)
	if err != nil {
		log.Println("err:", err.Error())
		return
	}
	log.Println("h:", h.String())
}

func TestPublicKeyOne(t *testing.T) {
	pubX, pubY := UnpackPublickKey("2160ceb725341c107a1cb4214ee5201a094dcd8820d6345425569a27f76ba3bf")
	log.Println("pubX:", pubX.String())
	log.Println("pubY:", pubY.String())

	test1, _ := decimal.NewFromString("21888242871839275222246405745257275088548364400416034343698204186575808495617")
	test2, _ := decimal.NewFromString("18460562194448123717098299273647690723188589856320307089893489725445607041947")
	test3 := test1.Sub(test2)
	log.Println("test3:", test3.String())
	
	test4, _ := decimal.NewFromString("7993856734532335104975857550077232487906787448363864460152476640414119364540")
	p, _ := decimal.NewFromString("21888242871839275222246405745257275088548364400416034343698204186575808495617")
	new(big.Int).ModSqrt(test4.BigInt(), p.BigInt())
	log.Println("test4:", test4.String())
	// 3490403304705816298967278020004010442111351094036683346883781626388274458942

	one := util.NewDecimalFromString("1")
	
	test5 := new(big.Int).Sub(one.BigInt(), test2.BigInt())
	test6 := new(big.Int).Mod(test5, p.BigInt())
	log.Println("test6:", test6.String())

}

func TestPublicKey(t *testing.T) {
	// compressedKeyIntStr := "9515267948657196726831090518272995529784296264381513362535790499183065198592"
	// compressedKey := util.IntStrToBytes(compressedKeyIntStr)
	// if len(compressedKey) != 32 {
	// 	panic("in unpackPoint compressedKey invalid length")
	// }
	// compressedKeyBits := util.Bytes2Bits(compressedKey)
	// sign := false
	// if compressedKeyBits[31 * 8] == 1 {
	// 	sign = true
	// 	compressedKeyBits[31 * 8] = 0
	// }
	// // if compressedKeyBits[0] == 1 {
	// // 	sign = true
	// // 	compressedKeyBits[0] = 0
	// // }
	// compressedKeyInt := util.BitsToIntStr(compressedKeyBits)
	// publicKeyY = util.NewDecimalFromString(compressedKeyInt)
	publicKeyY, _ := decimal.NewFromString("9515267948657196726831090518272995529784296264381513362535790499183065198592")
	d := util.NewDecimalFromString("168696")
	a := util.NewDecimalFromString("168700")
	one := decimal.NewFromInt(1)
	// publicKeyX2 := (one.Sub(publicKeyY)).Div(a.Sub(d.Mul(publicKeyY.Mul(publicKeyY))))
	publicKeyY2 := publicKeyY.Mul(publicKeyY)
	test1 := publicKeyY2.Sub(one)
	test2 := d.Mul(publicKeyY2)
	log.Println("test1:", test1.String())
	log.Println("test2:", test2.String())
	publicKeyX2 := (publicKeyY2.Sub(one)).Div(d.Mul(publicKeyY2).Sub(a))
	log.Println("publicKeyX2:", publicKeyX2.String())
	// if sign {
	// 	publicKeyX2 = publicKeyX2.Neg()
	// }
	publicKeyXFloat64 := math.Sqrt(publicKeyX2.InexactFloat64())
	publicKeyX := util.NewDecimalFromString(strconv.FormatFloat(publicKeyXFloat64, 'f', -1, 64))
	log.Println("publicKeyX:", publicKeyX.String())
}