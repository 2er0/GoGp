package main

import (
	"fmt"
	"github.com/2er0/GoGp/agent"
	"github.com/2er0/GoGp/engine"
	"github.com/2er0/GoGp/fitness"
	"github.com/2er0/GoGp/generator"
	"github.com/2er0/GoGp/mutator"
	"github.com/2er0/GoGp/selection"
	"github.com/2er0/GoGp/solution"
	"github.com/2er0/GoGp/solutioncheck"
	"github.com/2er0/GoGp/utils"
	"time"
	"github.com/2er0/GoGp/crossover"
)

var e = *engine.New()

func main() {
	start := time.Now()
	esDriver()
	gaDriver()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	e.GetAll()
}

func gaDriver() {
	boothFunc := "(x1 + 2 * x2 - 7) ** 2 + (2 * x1 + x2 - 5) ** 2"

	a1 := agent.NewGA(10, 20, 200, 2, false,
		fitness.NewFunc(boothFunc), generator.NewRealGen(), crossover.NewSimple(),
		mutator.NewReal(), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a1)
	e.Run()

	a2 := agent.NewGA(10, 20, 200, 2, true,
		fitness.NewFunc(boothFunc), generator.NewRealGen(), crossover.NewSimple(),
		mutator.NewReal(), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a2)
	e.Run()

	a1_2 := agent.NewGA(20, 50, 200, 2, false,
		fitness.NewFunc(boothFunc), generator.NewRealGen(), crossover.NewSimple(),
		mutator.NewReal(), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a1_2)
	e.Run()

	a2_2 := agent.NewGA(20, 50, 200, 2, true,
		fitness.NewFunc(boothFunc), generator.NewRealGen(), crossover.NewSimple(),
		mutator.NewReal(), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a2_2)
	e.Run()

	/*fmt.Println(boothFunc, "==> best @ [1 3] = 0")
	fmt.Println(a1.Result())
	fmt.Println(a2.Result())*/

	path := map[int]utils.Point{
		0:  {288, 149},
		1:  {288, 129},
		2:  {270, 133},
		3:  {256, 141},
		4:  {256, 157},
		5:  {246, 157},
		6:  {236, 169},
		7:  {228, 169},
		8:  {228, 161},
		9:  {220, 169},
		10: {212, 169},
		11: {204, 169},
		12: {196, 169},
		13: {188, 169},
		14: {196, 161},
		15: {188, 145},
		16: {172, 145},
		17: {164, 145},
		18: {156, 145},
		19: {148, 145},
	}

	a3 := agent.NewGA(10, 20, 200, 20, false,
		fitness.NewTsp(path), generator.NewTsp(), crossover.NewSimple(),
		mutator.NewTsp(20), selection.NewWithOut(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a3)
	e.Run()

	a4 := agent.NewGA(10, 20, 200, 20, false,
		fitness.NewTsp(path), generator.NewTsp(), crossover.NewPattern(),
		mutator.NewTsp(20), selection.NewWithOut(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a4)
	e.Run()

	a5 := agent.NewGA(10, 20, 200, 20, true,
		fitness.NewTsp(path), generator.NewTsp(), crossover.NewSimple(),
		mutator.NewTsp(20), selection.NewWithOut(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a5)

	a6 := agent.NewGA(10, 20, 200, 20, true,
		fitness.NewTsp(path), generator.NewTsp(), crossover.NewPattern(),
		mutator.NewTsp(20), selection.NewWithOut(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a6)

	e.RunAll()

	C130 := map[int]utils.Point{
		1:   {334.5909245845, 161.7809319139},
		2:   {397.6446634067, 262.8165330708},
		3:   {503.8741827107, 172.8741151168},
		4:   {444.0479403502, 384.6491809647},
		5:   {311.6137146746, 2.0091699828},
		6:   {662.8551011379, 549.2301263653},
		7:   {40.0979030612, 187.2375430791},
		8:   {526.8941409181, 215.7079092185},
		9:   {209.1887938487, 691.0262291948},
		10:  {683.2674131973, 414.2096286906},
		11:  {280.7494438748, 5.9206392047},
		12:  {252.7493090080, 535.7430385019},
		13:  {698.7850451923, 348.4413729766},
		14:  {678.7574678104, 410.7256424438},
		15:  {220.0041131179, 409.1225812873},
		16:  {355.1528556851, 76.3912076444},
		17:  {296.9724227786, 313.1312792361},
		18:  {504.5154071733, 240.8866564499},
		19:  {224.1079496785, 358.4872228907},
		20:  {470.6801296968, 309.6259188406},
		21:  {554.2530513223, 279.4242466521},
		22:  {567.6332684419, 352.7162027273},
		23:  {599.0532671093, 361.0948690386},
		24:  {240.5232959211, 430.6036007844},
		25:  {32.0825972787, 345.8551009775},
		26:  {91.0538736891, 148.7213270256},
		27:  {248.2179894723, 343.9528017384},
		28:  {488.8909044347, 3.6122311393},
		29:  {206.0467939820, 437.7639406167},
		30:  {575.8409415632, 141.9670960195},
		31:  {282.6089948164, 329.4183805862},
		32:  {27.6581484868, 424.7684581747},
		33:  {568.5737309870, 287.0975660546},
		34:  {269.4638933331, 295.9464636385},
		35:  {417.8004856811, 341.2596589955},
		36:  {32.1680938737, 448.8998721172},
		37:  {561.4775136009, 357.3543930067},
		38:  {342.9482167470, 492.3321423839},
		39:  {399.6752075383, 156.8435035519},
		40:  {571.7371050025, 375.7575350833},
		41:  {370.7559842751, 151.9060751898},
		42:  {509.7093253204, 435.7975189314},
		43:  {177.0206999750, 295.6044772584},
		44:  {526.1674198605, 409.4859418161},
		45:  {316.5725171854, 65.6400108214},
		46:  {469.2908100279, 281.9891445025},
		47:  {572.7630641427, 373.3208821255},
		48:  {29.5176994283, 330.0382309000},
		49:  {454.0082936692, 537.2178547659},
		50:  {416.1546762271, 227.6133100741},
		51:  {535.2514330806, 471.0648643744},
		52:  {265.4455533675, 684.9987192464},
		53:  {478.0542110167, 509.6452028741},
		54:  {370.4781203413, 332.5390063041},
		55:  {598.3479202004, 446.8693279856},
		56:  {201.1521139175, 649.0260268945},
		57:  {193.6925360026, 680.2322840744},
		58:  {448.5792598859, 532.7934059740},
		59:  {603.2853485624, 134.4006473609},
		60:  {543.0102490781, 481.5168231148},
		61:  {214.5750793346, 43.6460117543},
		62:  {426.3501451825, 61.7285415996},
		63:  {89.0447037063, 277.1158385868},
		64:  {84.4920100219, 31.8474816424},
		65:  {220.0468614154, 623.0778103080},
		66:  {688.4613313444, 0.4702312726},
		67:  {687.2857531630, 373.5346236130},
		68:  {75.4934933967, 312.9175377486},
		69:  {63.4170993511, 23.7039309674},
		70:  {97.9363495877, 211.0910930878},
		71:  {399.5255884970, 170.8221968365},
		72:  {456.3167017346, 597.1937161677},
		73:  {319.8855102422, 626.8396604886},
		74:  {295.9250894897, 664.6291554845},
		75:  {288.4868857235, 667.7284070537},
		76:  {268.3951858954, 52.9010181645},
		77:  {140.4709056068, 513.5566720960},
		78:  {689.8079027159, 167.5947003748},
		79:  {280.5784506848, 458.7533546925},
		80:  {453.3884433554, 282.9082328989},
		81:  {213.5704943432, 525.8681817779},
		82:  {133.6953004520, 677.1757808026},
		83:  {521.1658690522, 132.8617086506},
		84:  {30.2657946347, 450.0754502986},
		85:  {657.0199585283, 39.7772908299},
		86:  {6.9252241961, 23.8749241575},
		87:  {252.4286967767, 535.1659364856},
		88:  {42.8551682504, 63.8232081774},
		89:  {145.8999393902, 399.5255884970},
		90:  {638.4885715591, 62.6262558472},
		91:  {489.2756391122, 665.3131282446},
		92:  {361.2231139311, 564.2347787901},
		93:  {519.9475425732, 347.9711417040},
		94:  {129.3349741063, 435.6692740389},
		95:  {259.7172815016, 454.6495181318},
		96:  {676.3421890013, 371.0979706551},
		97:  {84.5133841706, 183.3260738572},
		98:  {77.7164048671, 354.3833863300},
		99:  {335.9802442534, 660.6321896676},
		100: {264.3554717810, 377.5743377274},
		101: {51.6826916855, 676.0429509187},
		102: {692.1376849300, 543.8010925819},
		103: {169.2191356800, 547.8194325476},
		104: {194.0131482339, 263.4791316822},
		105: {415.1928395332, 78.9133571973},
		106: {415.0432204919, 479.0801701569},
		107: {169.8389859939, 245.6103433244},
		108: {525.0987124228, 213.5063718969},
		109: {238.6851191283, 33.4932910965},
		110: {116.2112467718, 363.5742702940},
		111: {16.9283258126, 656.5711014044},
		112: {434.3440768162, 92.6996831431},
		113: {40.5253860363, 424.6829615797},
		114: {530.4849979086, 183.8390534273},
		115: {484.3595848990, 49.2460387276},
		116: {263.6501248722, 426.5852608187},
		117: {450.2891917862, 126.3853415784},
		118: {441.7822805823, 299.7724362653},
		119: {24.2169105375, 500.3474481664},
		120: {503.7886861157, 514.6895019799},
		121: {635.5389390312, 200.9811207275},
		122: {614.5922732529, 418.8691931188},
		123: {21.7161351334, 660.9741760476},
		124: {143.8266469611, 92.6996831431},
		125: {637.7191022040, 54.2048412384},
		126: {566.5645610042, 199.9551615873},
		127: {196.6849168280, 221.8209157619},
		128: {384.9270448985, 87.4630166986},
		129: {178.1107815614, 104.6905805938},
		130: {403.2874386776, 205.8971749407},
	}

	a7 := agent.NewGA(20, 50, 200, 130, false,
		fitness.NewTsp(C130), generator.NewTsp(), crossover.NewSimple(),
		mutator.NewTsp(130), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a7)
	e.Run()

	a8 := agent.NewGA(20, 50, 200, 130, false,
		fitness.NewTsp(C130), generator.NewTsp(), crossover.NewPattern(),
		mutator.NewTsp(130), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a8)
	e.Run()

	a9 := agent.NewGA(20, 50, 200, 130, true,
		fitness.NewTsp(C130), generator.NewTsp(), crossover.NewSimple(),
		mutator.NewTsp(130), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a9)

	a10 := agent.NewGA(20, 50, 200, 130, true,
		fitness.NewTsp(C130), generator.NewTsp(), crossover.NewPattern(),
		mutator.NewTsp(130), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a10)

	e.RunAll()

	a11 := agent.NewGA(100, 500, 5000, 130, true,
		fitness.NewTsp(C130), generator.NewTsp(), crossover.NewPattern(),
		mutator.NewTsp(130), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(a11)
	e.Run()

	/*fmt.Println(a3.Result())
	fmt.Println(a4.Result())
	fmt.Println(a5.Result())
	fmt.Println(a6.Result())

	fmt.Println(a7.Result())
	fmt.Println(a8.Result())
	fmt.Println(a9.Result())
	fmt.Println(a10.Result())
	fmt.Println(a11.Result())*/

	// Known best C130 solution
	sol := generator.NewTsp().New(130)
	val := []int{1, 41,  39,  117,  112,  115,  28,  62,  105,  128,  16,  45,  5,  11,  76,  109,  61,  129,  124,  64,
		69,  86,  88,  26,  7,  97,  70,  107,  127,  104,  43,  34,  17,  31,  27,  19,  100,  15,  29,  24,  116,  95,  79,
		87,  12,  81,  103,  77,  94,  89,  110,  98,  68,  63,  48,  25,  113,  32,  36,  84,  119,  111,  123,  101,  82,
		57,  9,  56,  65,  52,  75,  74,  99,  73,  92,  38,  106,  53,  120,  58,  49,  72,  91,  6,  102,  10,  14,  67,
		13,  96,  122,  55,  60,  51,  42,  44,  93,  37,  22,  47,  40,  23,  33,  21,  126,  121,  78,  66,  85, 125, 90,
		59, 30, 83, 3, 114, 108, 8, 18, 46, 80, 118, 20, 4, 35, 54, 2, 50, 130, 71}
	for i := range val {
		val[i] = val[i] - 1
	}
	sol.SetValues(val)
	fit := fitness.NewTsp(C130)
	fmt.Println(fit.Calc(sol))
}

func esDriver() {
	baseFunc := "x1 ** 2 + 1"

	a3 := agent.NewES(10, 20, 100, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithOut(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())

	e.Add(a3)
	e.Run()

	a4 := agent.NewES(10, 20, 100, 1,
		fitness.NewFunc(baseFunc), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWithOut(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a4)

	a5 := agent.NewES(10, 20, 100, 1,
		fitness.NewFunc(baseFunc), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWith(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a5)

	a6 := agent.NewES(10, 20, 100, 1,
		fitness.NewFunc(baseFunc), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a6)

	e.RunAll()

	a_3_1 := agent.NewES(1, 1, 100, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a_3_1)

	a_3_4 := agent.NewES(1, 10, 100, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a_3_4)

	a_3_5 := agent.NewES(10, 10, 100, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a_3_5)

	a_3 := agent.NewES(10, 20, 100, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a_3)

	a_3_2 := agent.NewES(50, 70, 200, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a_3_2)

	a_3_3 := agent.NewES(100, 300, 200, 1,
		fitness.NewTestFunc(1), generator.NewFloatGen(), // x1 * x1 + 1
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a_3_3)
	e.RunAll()

	boothFunc := "(x1 + 2 * x2 - 7) ** 2 + (2 * x1 + x2 - 5) ** 2"

	a7 := agent.NewES(10, 20, 200, 2,
		fitness.NewFunc(boothFunc), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWithOut(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a7)
	e.Run()

	a8 := agent.NewES(10, 20, 200, 2,
		fitness.NewFunc(boothFunc), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWith(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a8)
	e.Run()

	a9 := agent.NewES(10, 20, 200, 2,
		fitness.NewFunc(boothFunc), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a9)
	e.Run()

	testFunc2 := "4 * x1 - x2 * x3"

	a10 := agent.NewES(10, 20, 200, 3,
		fitness.NewFunc(testFunc2), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWith(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a10)

	a11 := agent.NewES(10, 20, 200, 3,
		fitness.NewFunc(testFunc2), generator.NewFloatGen(),
		mutator.NewFloat(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a11)

	e.RunAll()

	a12 := agent.NewES(10, 20, 200, 2,
		fitness.NewFunc(boothFunc), generator.NewRealGen(),
		mutator.NewReal(), selection.NewWithOut(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a12)

	a13 := agent.NewES(10, 20, 200, 2,
		fitness.NewFunc(boothFunc), generator.NewRealGen(),
		mutator.NewReal(), selection.NewWith(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a13)

	a14 := agent.NewES(10, 20, 200, 2,
		fitness.NewFunc(boothFunc), generator.NewRealGen(),
		mutator.NewReal(), selection.NewWithoutElit(), mutator.NewSigma15(),
		solutioncheck.NewNone(), solution.NewFloatCompMin())
	e.Add(a14)

	e.RunAll()

	path := map[int]utils.Point{
		0:  {288, 149},
		1:  {288, 129},
		2:  {270, 133},
		3:  {256, 141},
		4:  {256, 157},
		5:  {246, 157},
		6:  {236, 169},
		7:  {228, 169},
		8:  {228, 161},
		9:  {220, 169},
		10: {212, 169},
		11: {204, 169},
		12: {196, 169},
		13: {188, 169},
		14: {196, 161},
		15: {188, 145},
		16: {172, 145},
		17: {164, 145},
		18: {156, 145},
		19: {148, 145},
	}

	tsp1 := agent.NewES(10, 20, 200, 20,
		fitness.NewTsp(path), generator.NewTsp(),
		mutator.NewTsp(20), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(tsp1)

	tsp2 := agent.NewES(1, 1, 200, 20,
		fitness.NewTsp(path), generator.NewTsp(),
		mutator.NewTsp(20), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(tsp2)

	tsp3 := agent.NewES(50, 100, 200, 20,
		fitness.NewTsp(path), generator.NewTsp(),
		mutator.NewTsp(20), selection.NewWithoutElit(), mutator.NewSigmaOri(),
		solutioncheck.NewTsp(), solution.NewFloatCompMin())
	e.Add(tsp3)

	/*fmt.Println()
	fmt.Println(baseFunc, "==> best @ 0 = 1")
	fmt.Println(a3.Result())
	fmt.Println(a4.Result())
	fmt.Println(a5.Result())
	fmt.Println(a6.Result())

	fmt.Println(boothFunc, "==> best @ [1 3] = 0")
	fmt.Println(a7.Result())
	fmt.Println(a8.Result())
	fmt.Println(a9.Result())
	fmt.Println(a12.Result())
	fmt.Println(a13.Result())
	fmt.Println(a14.Result())

	fmt.Println(testFunc2, "==> Angabe test func")
	fmt.Println(a10.Result())
	fmt.Println(a11.Result())*/
}
