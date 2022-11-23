package babyjub

import (
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"math/rand"
	"testing"

	"staterecovery/crypto/constants"
	"staterecovery/crypto/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAdd1(t *testing.T) {
	a := &Point{X: big.NewInt(0), Y: big.NewInt(1)}
	b := &Point{X: big.NewInt(0), Y: big.NewInt(1)}

	c := NewPoint().Projective().Add(a.Projective(), b.Projective())
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t, "0", c.X.String())
	assert.Equal(t, "1", c.Y.String())
}

func ToAnyHexadecimal(numberStr string, currentHexadecimal, toHexadecimal int) (string, error) {
	bigCurrent, ok := new(big.Int).SetString(numberStr, currentHexadecimal)
	if !ok {
		log.Println("in ToAnyHexadecimal SetString:", ok)
		return "", errors.New("dee")
	}
	toNumberStr := bigCurrent.Text(toHexadecimal)
	return toNumberStr, nil
}

func TestCompress(t *testing.T) {
	x, _ := decimal.NewFromString("3490403304705816298967278020004010442111351094036683346883781626388274458942")
	y, _ := decimal.NewFromString("15097368018802186105341826992046618355217377505365242188231270000890858939327")
	a := &Point{X: x.BigInt(), Y: y.BigInt()}
	compressed := a.Compress()
	compressedHex := common.Bytes2Hex(compressed[:])
	log.Println("compressedHex:", compressedHex)
	compressedInt, _ := ToAnyHexadecimal(compressedHex, 16, 10)
	log.Println("compressedInt:", compressedInt)

	compressedByte := common.FromHex("2160ceb725341c107a1cb4214ee5201a094dcd8820d6345425569a27f76ba3bf")
	test := [32]byte{}
	copy(test[:], compressedByte[:])
	reverse := [32]byte{}
	for i := 0; i < 32; i ++ {
		reverse[31 - i] = test[i]
	}
	testResult, err := a.Decompress(reverse)
	if err != nil {
		log.Println("err:", err.Error())
		return
	}
	log.Println("x:", testResult.X.String())
	log.Println("y:", testResult.Y.String())
}

func TestAdd2(t *testing.T) {
	aX := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	aY := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	a := &Point{X: aX, Y: aY}

	bX := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	bY := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	b := &Point{X: bX, Y: bY}

	c := NewPoint().Projective().Add(a.Projective(), b.Projective()).Affine()
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t,
		"6890855772600357754907169075114257697580319025794532037257385534741338397365",
		c.X.String())
	assert.Equal(t,
		"4338620300185947561074059802482547481416142213883829469920100239455078257889",
		c.Y.String())

	d := NewPointProjective().Add(c.Projective(), c.Projective()).Affine()
	assert.Equal(t,
		"2f6458832049e917c95867185a96621336df33e13c98e81d1ef4928cdbb77772",
		hex.EncodeToString(d.X.Bytes()))

	// Projective
	aP := a.Projective()
	bP := b.Projective()
	cP := NewPointProjective().Add(aP, bP)
	c2 := cP.Affine()
	assert.Equal(t, c, c2)
}

func TestAdd3(t *testing.T) {
	aX := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	aY := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	a := &Point{X: aX, Y: aY}

	bX := utils.NewIntFromString(
		"16540640123574156134436876038791482806971768689494387082833631921987005038935")
	bY := utils.NewIntFromString(
		"20819045374670962167435360035096875258406992893633759881276124905556507972311")
	b := &Point{X: bX, Y: bY}

	c := NewPoint().Projective().Add(a.Projective(), b.Projective()).Affine()
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t,
		"7916061937171219682591368294088513039687205273691143098332585753343424131937",
		c.X.String())
	assert.Equal(t,
		"14035240266687799601661095864649209771790948434046947201833777492504781204499",
		c.Y.String())
}

func TestAdd4(t *testing.T) {
	aX := utils.NewIntFromString(
		"0")
	aY := utils.NewIntFromString(
		"1")
	a := &Point{X: aX, Y: aY}

	bX := utils.NewIntFromString(
		"16540640123574156134436876038791482806971768689494387082833631921987005038935")
	bY := utils.NewIntFromString(
		"20819045374670962167435360035096875258406992893633759881276124905556507972311")
	b := &Point{X: bX, Y: bY}

	c := NewPoint().Projective().Add(a.Projective(), b.Projective()).Affine()
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t,
		"16540640123574156134436876038791482806971768689494387082833631921987005038935",
		c.X.String())
	assert.Equal(t,
		"20819045374670962167435360035096875258406992893633759881276124905556507972311",
		c.Y.String())
}

func TestInCurve1(t *testing.T) {
	p := &Point{X: big.NewInt(0), Y: big.NewInt(1)}
	assert.Equal(t, true, p.InCurve())
}

func TestInCurve2(t *testing.T) {
	p := &Point{X: big.NewInt(1), Y: big.NewInt(0)}
	assert.Equal(t, false, p.InCurve())
}

func TestMul0(t *testing.T) {
	x := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	s := utils.NewIntFromString("3")

	r2 := NewPoint().Projective().Add(p.Projective(), p.Projective()).Affine()
	r2 = NewPoint().Projective().Add(r2.Projective(), p.Projective()).Affine()
	r := NewPoint().Mul(s, p)
	assert.Equal(t, r2.X.String(), r.X.String())
	assert.Equal(t, r2.Y.String(), r.Y.String())

	assert.Equal(t,
		"19372461775513343691590086534037741906533799473648040012278229434133483800898",
		r.X.String())
	assert.Equal(t,
		"9458658722007214007257525444427903161243386465067105737478306991484593958249",
		r.Y.String())
}

func TestMul1(t *testing.T) {
	x := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	s := utils.NewIntFromString(
		"14035240266687799601661095864649209771790948434046947201833777492504781204499")
	r := NewPoint().Mul(s, p)
	assert.Equal(t,
		"17070357974431721403481313912716834497662307308519659060910483826664480189605",
		r.X.String())
	assert.Equal(t,
		"4014745322800118607127020275658861516666525056516280575712425373174125159339",
		r.Y.String())
}

func TestMul2(t *testing.T) {
	x := utils.NewIntFromString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := utils.NewIntFromString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}
	s := utils.NewIntFromString(
		"20819045374670962167435360035096875258406992893633759881276124905556507972311")
	r := NewPoint().Mul(s, p)
	assert.Equal(t,
		"13563888653650925984868671744672725781658357821216877865297235725727006259983",
		r.X.String())
	assert.Equal(t,
		"8442587202676550862664528699803615547505326611544120184665036919364004251662",
		r.Y.String())
}

func TestInCurve3(t *testing.T) {
	x := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InCurve())
}

func TestInCurve4(t *testing.T) {
	x := utils.NewIntFromString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := utils.NewIntFromString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InCurve())
}

func TestInSubGroup1(t *testing.T) {
	x := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InSubGroup())
}

func TestInSubGroup2(t *testing.T) {
	x := utils.NewIntFromString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := utils.NewIntFromString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InSubGroup())
}

func TestPointFromSignAndy(t *testing.T) {
	x := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}

	sign := PointCoordSign(p.X)
	p2, err := PointFromSignAndY(sign, p.Y)
	assert.Equal(t, nil, err)
	assert.Equal(t, p.X.String(), p2.X.String())
	assert.Equal(t, p.Y.String(), p2.Y.String())
}

func TestPackAndUnpackSignY(t *testing.T) {
	x := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	pComp := p.Compress()

	s, y := UnpackSignY(pComp)

	pComp2 := PackSignY(s, y)
	assert.Equal(t, pComp, pComp2)

	emptyPointComp := [32]byte{}
	s, y = UnpackSignY(emptyPointComp)

	pComp2 = PackSignY(s, y)
	assert.Equal(t, emptyPointComp, pComp2)
}

func TestCompressDecompress1(t *testing.T) {
	x := utils.NewIntFromString(
		"8616796765096162578520928440127092306210952615345657010489183608073496444026")
	y := utils.NewIntFromString(
		"12402052136221678618341947144652713479139644703569747758310794733558115261840")
	p := &Point{X: x, Y: y}

	buf := p.Compress()
	assert.Equal(t,
		"53b81ed5bffe9545b54016234682e7b2f699bd42a5e9eae27ff4051bc698ce85",
		hex.EncodeToString(buf[:]))

	p2, err := NewPoint().Decompress(buf)
	assert.Equal(t, nil, err)
	assert.Equal(t, p.X.String(), p2.X.String())
	assert.Equal(t, p.Y.String(), p2.Y.String())
}

func TestCompressDecompress2(t *testing.T) {
	x := utils.NewIntFromString(
		"3490403304705816298967278020004010442111351094036683346883781626388274458942")
	y := utils.NewIntFromString(
		"15097368018802186105341826992046618355217377505365242188231270000890858939327")
	p := &Point{X: x, Y: y}

	buf := p.Compress()
	assert.Equal(t,
		"bfa36bf7279a56255434d62088cd4d091a20e54e21b41c7a101c3425b7ce6021",
		hex.EncodeToString(buf[:]))

	p2, err := NewPoint().Decompress(buf)
	assert.Equal(t, nil, err)
	assert.Equal(t, p.X.String(), p2.X.String())
	assert.Equal(t, p.Y.String(), p2.Y.String())
}

func TestCompressDecompressRnd(t *testing.T) {
	for i := 0; i < 64; i++ {
		p1 := NewPoint().Mul(big.NewInt(int64(i)), B8)
		buf := p1.Compress()
		p2, err := NewPoint().Decompress(buf)
		assert.Equal(t, nil, err)
		assert.Equal(t, p1.X.Bytes(), p2.X.Bytes())
		assert.Equal(t, p1.Y.Bytes(), p2.Y.Bytes())
	}
}

func BenchmarkBabyjub(b *testing.B) {
	const n = 256

	rnd := rand.New(rand.NewSource(42)) //nolint:gosec

	var badpoints [n]*Point
	for i := 0; i < n; i++ {
		x := new(big.Int).Rand(rnd, constants.Q)
		y := new(big.Int).Rand(rnd, constants.Q)
		badpoints[i] = &Point{X: x, Y: y}
	}

	var points [n]*Point
	var pointsProj [n]*PointProjective
	baseX := utils.NewIntFromString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	baseY := utils.NewIntFromString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	base := &Point{X: baseX, Y: baseY}
	for i := 0; i < n; i++ {
		s := new(big.Int).Rand(rnd, constants.Q)
		points[i] = NewPoint().Mul(s, base)
		pointsProj[i] = NewPoint().Mul(s, base).Projective()
	}

	var scalars [n]*big.Int
	for i := 0; i < n; i++ {
		scalars[i] = new(big.Int).Rand(rnd, constants.Q)
	}

	b.Run("AddConst", func(b *testing.B) {
		p0 := &Point{X: big.NewInt(0), Y: big.NewInt(1)}
		p1 := &Point{X: big.NewInt(0), Y: big.NewInt(1)}
		p0Proj := p0.Projective()
		p1Proj := p1.Projective()

		p2 := NewPoint().Projective()
		for i := 0; i < b.N; i++ {
			p2.Add(p0Proj, p1Proj)
		}
	})

	b.Run("AddRnd", func(b *testing.B) {
		res := NewPoint().Projective()
		for i := 0; i < b.N; i++ {
			res.Add(pointsProj[i%(n/2)], pointsProj[i%(n/2)+1])
		}
	})

	b.Run("MulRnd", func(b *testing.B) {
		res := NewPoint()
		for i := 0; i < b.N; i++ {
			res.Mul(scalars[i%n], points[i%n])
		}
	})

	b.Run("Compress", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[i%n].Compress()
		}
	})

	b.Run("InCurve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			badpoints[i%n].InCurve()
		}
	})

	b.Run("InSubGroup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[i%n].InCurve()
		}
	})
}
