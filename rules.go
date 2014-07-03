package gomobiledetect

const (
	IPHONE = iota
	BLACKBERRY
	HTC
	NEXUS
	DELL
	MOTOROLA
	SAMSUNG
	LG
	SONY
	ASUS
	MICROMAX
	PALM
	VERTU
	PANTECH
	FLY
	IMOBILE
	SIMVALLEY
	GENERICPHONE
)

const (
	IPAD = iota
	SAMSUNGTABLET
	KINDLE
	SURFACETABLET
	HPTABLET
	ASUSTABLET
	BLACKBERRYTABLET
	HTCTABLET
	MOTOROLATABLET
	NOOKTABLET
	ACERTABLET
	TOSHIBATABLET
	LGTABLET
	FUJITSUTABLET
	PRESTIGIOTABLET
	LENOVOTABLET
	YARVIKTABLET
	MEDIONTABLET
	ARNOVATABLET
	INTENSOTABLET
	IRUTABLET
	MEGAFONTABLET
	EBODATABLET
	ALLVIEWTABLET
	ARCHOSTABLET
	AINOLTABLET
	SONYTABLET
	CUBETABLET
	COBYTABLET
	MIDTABLET
	SMITTABLET
	ROCKCHIPTABLET
	FLYTABLET
	BQTABLET
	HUAWEITABLET
	NECTABLET
	PANTECHTABLET
	BRONCHOTABLET
	VERSUSTABLET
	ZYNCTABLET
	POSITIVOTABLET
	NABITABLET
	KOBOTABLET
	DANEWTABLET
	TEXETTABLET
	PLAYSTATIONTABLET
	TREKSTORTABLET
	PYLEAUDIOTABLET
	ADVANTABLET
	DANYTECHTABLET
	GALAPADTABLET
	MICROMAXTABLET
	KARBONNTABLET
	ALLFINETABLET
	PROSCANTABLET
	YONESTABLET
	CHANGJIATABLET
	GUTABLET
	POINTOFVIEWTABLET
	OVERMAXTABLET
	HCLTABLET
	DPSTABLET
	CRESTATABLET
	MEDIATEKTABLET
	CONCORDETABLET
	GOCLEVERTABLET
	MODECOMTABLET
	VONINOTABLET
	ECSTABLET
	STOREXTABLET
	VODAFONETABLET
	ESSENTIELBTABLET
	ROSSMOORTABLET
	IMOBILETABLET
	TOLINOTABLET
	AUDIOSONICTABLET
	AMPETABLET
	SKKTABLET
	TECNOTABLET
	JXDTABLET
	IJOYTABLET
	HUDL
	TELSTRATABLET
	VISTURETABLET
)

const (
	ANDROIDOS = iota
	BLACKBERRYOS
	PALMOS
	SYMBIANOS
	WINDOWSMOBILEOS
	WINDOWSPHONEOS
	IOS
	MEEGOOS
	MAEMOOS
	JAVAOS
	WEBOS
	BADAOS
	BREWOS
)

const (
	CHROME = iota
	DOLFIN
	OPERA
	SKYFIRE
	IE
	FIREFOX
	BOLT
	TEASHARK
	BLAZER
	SAFARI
	TIZEN
	UCBROWSER
	DIIGOBROWSER
	PUFFIN
	MERCURY
	GENERICBROWSER
)

var (
	phoneDevices = [...]string{
		IPHONE:     `\biPhone.*Mobile|\biPod`, // |\biTunes
		BLACKBERRY: `BlackBerry|\bBB10\b|rim[0-9]+`,
		HTC:        `HTC|HTC.*(Sensation|Evo|Vision|Explorer|6800|8100|8900|A7272|S510e|C110e|Legend|Desire|T8282)|APX515CKT|Qtek9090|APA9292KT|HD_mini|Sensation.*Z710e|PG86100|Z715e|Desire.*(A8181|HD)|ADR6200|ADR6400L|ADR6425|001HT|Inspire 4G|Android.*\bEVO\b|T-Mobile G1|Z520m`,
		NEXUS:      `Nexus One|Nexus S|Galaxy.*Nexus|Android.*Nexus.*Mobile`,
		// @todo: Is `Dell Streak` a tablet or a phone? ;)
		DELL:     `Dell.*Streak|Dell.*Aero|Dell.*Venue|DELL.*Venue Pro|Dell Flash|Dell Smoke|Dell Mini 3iX|XCD28|XCD35|\b001DL\b|\b101DL\b|\bGS01\b`,
		MOTOROLA: `Motorola|DROIDX|DROID BIONIC|\bDroid\b.*Build|Android.*Xoom|HRI39|MOT-|A1260|A1680|A555|A853|A855|A953|A955|A956|Motorola.*ELECTRIFY|Motorola.*i1|i867|i940|MB200|MB300|MB501|MB502|MB508|MB511|MB520|MB525|MB526|MB611|MB612|MB632|MB810|MB855|MB860|MB861|MB865|MB870|ME501|ME502|ME511|ME525|ME600|ME632|ME722|ME811|ME860|ME863|ME865|MT620|MT710|MT716|MT720|MT810|MT870|MT917|Motorola.*TITANIUM|WX435|WX445|XT300|XT301|XT311|XT316|XT317|XT319|XT320|XT390|XT502|XT530|XT531|XT532|XT535|XT603|XT610|XT611|XT615|XT681|XT701|XT702|XT711|XT720|XT800|XT806|XT860|XT862|XT875|XT882|XT883|XT894|XT901|XT907|XT909|XT910|XT912|XT928|XT926|XT915|XT919|XT925`,
		SAMSUNG:  `Samsung|SGH-I337|BGT-S5230|GT-B2100|GT-B2700|GT-B2710|GT-B3210|GT-B3310|GT-B3410|GT-B3730|GT-B3740|GT-B5510|GT-B5512|GT-B5722|GT-B6520|GT-B7300|GT-B7320|GT-B7330|GT-B7350|GT-B7510|GT-B7722|GT-B7800|GT-C3010|GT-C3011|GT-C3060|GT-C3200|GT-C3212|GT-C3212I|GT-C3262|GT-C3222|GT-C3300|GT-C3300K|GT-C3303|GT-C3303K|GT-C3310|GT-C3322|GT-C3330|GT-C3350|GT-C3500|GT-C3510|GT-C3530|GT-C3630|GT-C3780|GT-C5010|GT-C5212|GT-C6620|GT-C6625|GT-C6712|GT-E1050|GT-E1070|GT-E1075|GT-E1080|GT-E1081|GT-E1085|GT-E1087|GT-E1100|GT-E1107|GT-E1110|GT-E1120|GT-E1125|GT-E1130|GT-E1160|GT-E1170|GT-E1175|GT-E1180|GT-E1182|GT-E1200|GT-E1210|GT-E1225|GT-E1230|GT-E1390|GT-E2100|GT-E2120|GT-E2121|GT-E2152|GT-E2220|GT-E2222|GT-E2230|GT-E2232|GT-E2250|GT-E2370|GT-E2550|GT-E2652|GT-E3210|GT-E3213|GT-I5500|GT-I5503|GT-I5700|GT-I5800|GT-I5801|GT-I6410|GT-I6420|GT-I7110|GT-I7410|GT-I7500|GT-I8000|GT-I8150|GT-I8160|GT-I8190|GT-I8320|GT-I8330|GT-I8350|GT-I8530|GT-I8700|GT-I8703|GT-I8910|GT-I9000|GT-I9001|GT-I9003|GT-I9010|GT-I9020|GT-I9023|GT-I9070|GT-I9082|GT-I9100|GT-I9103|GT-I9220|GT-I9250|GT-I9300|GT-I9305|GT-I9500|GT-I9505|GT-M3510|GT-M5650|GT-M7500|GT-M7600|GT-M7603|GT-M8800|GT-M8910|GT-N7000|GT-S3110|GT-S3310|GT-S3350|GT-S3353|GT-S3370|GT-S3650|GT-S3653|GT-S3770|GT-S3850|GT-S5210|GT-S5220|GT-S5229|GT-S5230|GT-S5233|GT-S5250|GT-S5253|GT-S5260|GT-S5263|GT-S5270|GT-S5300|GT-S5330|GT-S5350|GT-S5360|GT-S5363|GT-S5369|GT-S5380|GT-S5380D|GT-S5560|GT-S5570|GT-S5600|GT-S5603|GT-S5610|GT-S5620|GT-S5660|GT-S5670|GT-S5690|GT-S5750|GT-S5780|GT-S5830|GT-S5839|GT-S6102|GT-S6500|GT-S7070|GT-S7200|GT-S7220|GT-S7230|GT-S7233|GT-S7250|GT-S7500|GT-S7530|GT-S7550|GT-S7562|GT-S7710|GT-S8000|GT-S8003|GT-S8500|GT-S8530|GT-S8600|SCH-A310|SCH-A530|SCH-A570|SCH-A610|SCH-A630|SCH-A650|SCH-A790|SCH-A795|SCH-A850|SCH-A870|SCH-A890|SCH-A930|SCH-A950|SCH-A970|SCH-A990|SCH-I100|SCH-I110|SCH-I400|SCH-I405|SCH-I500|SCH-I510|SCH-I515|SCH-I600|SCH-I730|SCH-I760|SCH-I770|SCH-I830|SCH-I910|SCH-I920|SCH-I959|SCH-LC11|SCH-N150|SCH-N300|SCH-R100|SCH-R300|SCH-R351|SCH-R400|SCH-R410|SCH-T300|SCH-U310|SCH-U320|SCH-U350|SCH-U360|SCH-U365|SCH-U370|SCH-U380|SCH-U410|SCH-U430|SCH-U450|SCH-U460|SCH-U470|SCH-U490|SCH-U540|SCH-U550|SCH-U620|SCH-U640|SCH-U650|SCH-U660|SCH-U700|SCH-U740|SCH-U750|SCH-U810|SCH-U820|SCH-U900|SCH-U940|SCH-U960|SCS-26UC|SGH-A107|SGH-A117|SGH-A127|SGH-A137|SGH-A157|SGH-A167|SGH-A177|SGH-A187|SGH-A197|SGH-A227|SGH-A237|SGH-A257|SGH-A437|SGH-A517|SGH-A597|SGH-A637|SGH-A657|SGH-A667|SGH-A687|SGH-A697|SGH-A707|SGH-A717|SGH-A727|SGH-A737|SGH-A747|SGH-A767|SGH-A777|SGH-A797|SGH-A817|SGH-A827|SGH-A837|SGH-A847|SGH-A867|SGH-A877|SGH-A887|SGH-A897|SGH-A927|SGH-B100|SGH-B130|SGH-B200|SGH-B220|SGH-C100|SGH-C110|SGH-C120|SGH-C130|SGH-C140|SGH-C160|SGH-C170|SGH-C180|SGH-C200|SGH-C207|SGH-C210|SGH-C225|SGH-C230|SGH-C417|SGH-C450|SGH-D307|SGH-D347|SGH-D357|SGH-D407|SGH-D415|SGH-D780|SGH-D807|SGH-D980|SGH-E105|SGH-E200|SGH-E315|SGH-E316|SGH-E317|SGH-E335|SGH-E590|SGH-E635|SGH-E715|SGH-E890|SGH-F300|SGH-F480|SGH-I200|SGH-I300|SGH-I320|SGH-I550|SGH-I577|SGH-I600|SGH-I607|SGH-I617|SGH-I627|SGH-I637|SGH-I677|SGH-I700|SGH-I717|SGH-I727|SGH-i747M|SGH-I777|SGH-I780|SGH-I827|SGH-I847|SGH-I857|SGH-I896|SGH-I897|SGH-I900|SGH-I907|SGH-I917|SGH-I927|SGH-I937|SGH-I997|SGH-J150|SGH-J200|SGH-L170|SGH-L700|SGH-M110|SGH-M150|SGH-M200|SGH-N105|SGH-N500|SGH-N600|SGH-N620|SGH-N625|SGH-N700|SGH-N710|SGH-P107|SGH-P207|SGH-P300|SGH-P310|SGH-P520|SGH-P735|SGH-P777|SGH-Q105|SGH-R210|SGH-R220|SGH-R225|SGH-S105|SGH-S307|SGH-T109|SGH-T119|SGH-T139|SGH-T209|SGH-T219|SGH-T229|SGH-T239|SGH-T249|SGH-T259|SGH-T309|SGH-T319|SGH-T329|SGH-T339|SGH-T349|SGH-T359|SGH-T369|SGH-T379|SGH-T409|SGH-T429|SGH-T439|SGH-T459|SGH-T469|SGH-T479|SGH-T499|SGH-T509|SGH-T519|SGH-T539|SGH-T559|SGH-T589|SGH-T609|SGH-T619|SGH-T629|SGH-T639|SGH-T659|SGH-T669|SGH-T679|SGH-T709|SGH-T719|SGH-T729|SGH-T739|SGH-T746|SGH-T749|SGH-T759|SGH-T769|SGH-T809|SGH-T819|SGH-T839|SGH-T919|SGH-T929|SGH-T939|SGH-T959|SGH-T989|SGH-U100|SGH-U200|SGH-U800|SGH-V205|SGH-V206|SGH-X100|SGH-X105|SGH-X120|SGH-X140|SGH-X426|SGH-X427|SGH-X475|SGH-X495|SGH-X497|SGH-X507|SGH-X600|SGH-X610|SGH-X620|SGH-X630|SGH-X700|SGH-X820|SGH-X890|SGH-Z130|SGH-Z150|SGH-Z170|SGH-ZX10|SGH-ZX20|SHW-M110|SPH-A120|SPH-A400|SPH-A420|SPH-A460|SPH-A500|SPH-A560|SPH-A600|SPH-A620|SPH-A660|SPH-A700|SPH-A740|SPH-A760|SPH-A790|SPH-A800|SPH-A820|SPH-A840|SPH-A880|SPH-A900|SPH-A940|SPH-A960|SPH-D600|SPH-D700|SPH-D710|SPH-D720|SPH-I300|SPH-I325|SPH-I330|SPH-I350|SPH-I500|SPH-I600|SPH-I700|SPH-L700|SPH-M100|SPH-M220|SPH-M240|SPH-M300|SPH-M305|SPH-M320|SPH-M330|SPH-M350|SPH-M360|SPH-M370|SPH-M380|SPH-M510|SPH-M540|SPH-M550|SPH-M560|SPH-M570|SPH-M580|SPH-M610|SPH-M620|SPH-M630|SPH-M800|SPH-M810|SPH-M850|SPH-M900|SPH-M910|SPH-M920|SPH-M930|SPH-N100|SPH-N200|SPH-N240|SPH-N300|SPH-N400|SPH-Z400|SWC-E100|SCH-i909|GT-N7100|GT-N7105|SCH-I535|SM-N900A|SGH-I317|SGH-T999L|GT-S5360B|GT-I8262|GT-S6802|GT-S6312|GT-S6310|GT-S5312|GT-S5310|GT-I9105|GT-I8510|GT-S6790N|SM-G7105|SM-N9005|GT-S5301|GT-I9295|GT-I9195|SM-C101|GT-S7392|GT-S7560|GT-B7610|GT-I5510|GT-S7582|GT-S7530E`,
		LG:       `\bLG\b;|LG[- ]?(C800|C900|E400|E610|E900|E-900|F160|F180K|F180L|F180S|730|855|L160|LS840|LS970|LU6200|MS690|MS695|MS770|MS840|MS870|MS910|P500|P700|P705|VM696|AS680|AS695|AX840|C729|E970|GS505|272|C395|E739BK|E960|L55C|L75C|LS696|LS860|P769BK|P350|P500|P509|P870|UN272|US730|VS840|VS950|LN272|LN510|LS670|LS855|LW690|MN270|MN510|P509|P769|P930|UN200|UN270|UN510|UN610|US670|US740|US760|UX265|UX840|VN271|VN530|VS660|VS700|VS740|VS750|VS910|VS920|VS930|VX9200|VX11000|AX840A|LW770|P506|P925|P999|E612|D955|D802)`,
		SONY:     `SonyST|SonyLT|SonyEricsson|SonyEricssonLT15iv|LT18i|E10i|LT28h|LT26w|SonyEricssonMT27i`,
		ASUS:     `Asus.*Galaxy|PadFone.*Mobile`,
		// @ref: http://www.micromaxinfo.com/mobiles/smartphones
		// Added because the codes might conflict with Acer Tablets.
		MICROMAX: `Micromax.*\b(A210|A92|A88|A72|A111|A110Q|A115|A116|A110|A90S|A26|A51|A35|A54|A25|A27|A89|A68|A65|A57|A90)\b`,
		PALM:     `PalmSource|Palm`,                                                                                           // avantgo|blazer|elaine|hiptop|plucker|xiino ; @todo - complete the regex.
		VERTU:    `Vertu|Vertu.*Ltd|Vertu.*Ascent|Vertu.*Ayxta|Vertu.*Constellation(F|Quest)?|Vertu.*Monika|Vertu.*Signature`, // Just for fun ;)
		// @ref: http://www.pantech.co.kr/en/prod/prodList.do?gbrand=VEGA (PANTECH)
		// Most of the VEGA devices are legacy. PANTECH seem to be newer devices based on Android.
		PANTECH: `PANTECH|IM-A850S|IM-A840S|IM-A830L|IM-A830K|IM-A830S|IM-A820L|IM-A810K|IM-A810S|IM-A800S|IM-T100K|IM-A725L|IM-A780L|IM-A775C|IM-A770K|IM-A760S|IM-A750K|IM-A740S|IM-A730S|IM-A720L|IM-A710K|IM-A690L|IM-A690S|IM-A650S|IM-A630K|IM-A600S|VEGA PTL21|PT003|P8010|ADR910L|P6030|P6020|P9070|P4100|P9060|P5000|CDM8992|TXT8045|ADR8995|IS11PT|P2030|P6010|P8000|PT002|IS06|CDM8999|P9050|PT001|TXT8040|P2020|P9020|P2000|P7040|P7000|C790`,
		// @ref: http://www.fly-phone.com/devices/smartphones/ ; Included only smartphones.
		FLY:     `IQ230|IQ444|IQ450|IQ440|IQ442|IQ441|IQ245|IQ256|IQ236|IQ255|IQ235|IQ245|IQ275|IQ240|IQ285|IQ280|IQ270|IQ260|IQ250`,
		IMOBILE: `i-mobile (IQ|i-STYLE|idea|ZAA|Hitz)`,
		// Added simvalley mobile just for fun. They have some interesting devices.
		// @ref: http://www.simvalley.fr/telephonie---gps-_22_telephonie-mobile_telephones_.html
		SIMVALLEY: `\b(SP-80|XT-930|SX-340|XT-930|SX-310|SP-360|SP60|SPT-800|SP-120|SPT-800|SP-140|SPX-5|SPX-8|SP-100|SPX-8|SPX-12)\b`,
		// @Tapatalk is a mobile app; @ref: http://support.tapatalk.com/threads/smf-2-0-2-os-and-browser-detection-plugin-and-tapatalk.15565/#post-79039
		GENERICPHONE: `Tapatalk|PDA;|SAGEM|\bmmp\b|pocket|\bpsp\b|symbian|Smartphone|smartfon|treo|up.browser|up.link|vodafone|\bwap\b|nokia|Series40|Series60|S60|SonyEricsson|N900|MAUI.*WAP.*Browser`,
	}
	tabletDevices = [...]string{
		IPAD: `iPad|iPad.*Mobile`, // @todo: check for mobile friendly emails topic.
		//NexusTablet:   `^.*Android.*Nexus(((?:(?!Mobile))|(?:(\s(7|10).+))).)*$`,
		SAMSUNGTABLET: `SAMSUNG.*Tablet|Galaxy.*Tab|SC-01C|GT-P1000|GT-P1003|GT-P1010|GT-P3105|GT-P6210|GT-P6800|GT-P6810|GT-P7100|GT-P7300|GT-P7310|GT-P7500|GT-P7510|SCH-I800|SCH-I815|SCH-I905|SGH-I957|SGH-I987|SGH-T849|SGH-T859|SGH-T869|SPH-P100|GT-P3100|GT-P3108|GT-P3110|GT-P5100|GT-P5110|GT-P6200|GT-P7320|GT-P7511|GT-N8000|GT-P8510|SGH-I497|SPH-P500|SGH-T779|SCH-I705|SCH-I915|GT-N8013|GT-P3113|GT-P5113|GT-P8110|GT-N8010|GT-N8005|GT-N8020|GT-P1013|GT-P6201|GT-P7501|GT-N5100|GT-N5110|SHV-E140K|SHV-E140L|SHV-E140S|SHV-E150S|SHV-E230K|SHV-E230L|SHV-E230S|SHW-M180K|SHW-M180L|SHW-M180S|SHW-M180W|SHW-M300W|SHW-M305W|SHW-M380K|SHW-M380S|SHW-M380W|SHW-M430W|SHW-M480K|SHW-M480S|SHW-M480W|SHW-M485W|SHW-M486W|SHW-M500W|GT-I9228|SCH-P739|SCH-I925|GT-I9200|GT-I9205|GT-P5200|GT-P5210|SM-T311|SM-T310|SM-T210|SM-T210R|SM-T211|SM-P600|SM-P601|SM-P605|SM-P900|SM-T217|SM-T217A|SM-T217S|SM-P6000|SM-T3100|SGH-I467|XE500`,
		// @reference: http://www.labnol.org/software/kindle-user-agent-string/20378/
		KINDLE: `Kindle|Silk.*Accelerated|Android.*\b(KFOT|KFTT|KFJWI|KFJWA|KFOTE|KFSOWI|KFTHWI|KFTHWA|KFAPWI|KFAPWA|WFJWAE)\b`,
		// Only the Surface tablets with Windows RT are considered mobile.
		// @ref: http://msdn.microsoft.com/en-us/library/ie/hh920767(v=vs.85).aspx
		SURFACETABLET: `Windows NT [0-9.]+; ARM;`,
		// @ref: http://shopping1.hp.com/is-bin/INTERSHOP.enfinity/WFS/WW-USSMBPublicStore-Site/en_US/-/USD/ViewStandardCatalog-Browse?CatalogCategoryID=JfIQ7EN5lqMAAAEyDcJUDwMT
		HPTABLET: `HP Slate 7|HP ElitePad 900|hp-tablet|EliteBook.*Touch`,
		// @note: watch out for PadFone, see #132
		//AsusTablet:       `^.*PadFone((?!Mobile).)*$|Transformer|TF101|TF101G|TF300T|TF300TG|TF300TL|TF700T|TF700KL|TF701T|TF810C|ME171|ME301T|ME302C|ME371MG|ME370T|ME372MG|ME172V|ME173X|ME400C|Slider SL101|\bK00F\b|TX201LA`,
		ASUSTABLET:       `^.*PadFone.*$|Transformer|TF101|TF101G|TF300T|TF300TG|TF300TL|TF700T|TF700KL|TF701T|TF810C|ME171|ME301T|ME302C|ME371MG|ME370T|ME372MG|ME172V|ME173X|ME400C|Slider SL101|\bK00F\b|TX201LA`,
		BLACKBERRYTABLET: `PlayBook|RIM Tablet`,
		HTCTABLET:        `HTC Flyer|HTC Jetstream|HTC-P715a|HTC EVO View 4G|PG41200`,
		MOTOROLATABLET:   `xoom|sholest|MZ615|MZ605|MZ505|MZ601|MZ602|MZ603|MZ604|MZ606|MZ607|MZ608|MZ609|MZ615|MZ616|MZ617`,
		NOOKTABLET:       `Android.*Nook|NookColor|nook browser|BNRV200|BNRV200A|BNTV250|BNTV250A|BNTV400|BNTV600|LogicPD Zoom2`,
		// @ref: http://www.acer.ro/ac/ro/RO/content/drivers
		// @ref: http://www.packardbell.co.uk/pb/en/GB/content/download (Packard Bell is part of Acer)
		// @ref: http://us.acer.com/ac/en/US/content/group/tablets
		// @note: Can conflict with Micromax and Motorola phones codes.
		ACERTABLET: `Android.*; \b(A100|A101|A110|A200|A210|A211|A500|A501|A510|A511|A700|A701|W500|W500P|W501|W501P|W510|W511|W700|G100|G100W|B1-A71|B1-710|B1-711|A1-810)\b|W3-810|\bA3-A10\b`,
		// @ref: http://eu.computers.toshiba-europe.com/innovation/family/Tablets/1098744/banner_id/footerlink/
		// @ref: http://us.toshiba.com/tablets/tablet-finder
		// @ref: http://www.toshiba.co.jp/regza/tablet/
		TOSHIBATABLET: `Android.*(AT100|AT105|AT200|AT205|AT270|AT275|AT300|AT305|AT1S5|AT500|AT570|AT700|AT830)|TOSHIBA.*FOLIO`,
		// @ref: http://www.nttdocomo.co.jp/english/service/developer/smart_phone/technical_info/spec/index.html
		LGTABLET:      `\bL-06C|LG-V900|LG-V500|LG-V909\b`,
		FUJITSUTABLET: `Android.*\b(F-01D|F-05E|F-10D|M532|Q572)\b`,
		// Prestigio Tablets http://www.prestigio.com/support
		PRESTIGIOTABLET: `PMP3170B|PMP3270B|PMP3470B|PMP7170B|PMP3370B|PMP3570C|PMP5870C|PMP3670B|PMP5570C|PMP5770D|PMP3970B|PMP3870C|PMP5580C|PMP5880D|PMP5780D|PMP5588C|PMP7280C|PMP7280C3G|PMP7280|PMP7880D|PMP5597D|PMP5597|PMP7100D|PER3464|PER3274|PER3574|PER3884|PER5274|PER5474|PMP5097CPRO|PMP5097|PMP7380D|PMP5297C|PMP5297C_QUAD`,
		// @ref: http://support.lenovo.com/en_GB/downloads/default.page?#
		LENOVOTABLET: `IdeaTab|S2110|S6000|K3011|A3000|A1000|A2107|A2109|A1107|ThinkPad([ ]+)?Tablet`,
		// @ref: http://www.yarvik.com/en/matrix/tablets/
		YARVIKTABLET: `Android.*\b(TAB210|TAB211|TAB224|TAB250|TAB260|TAB264|TAB310|TAB360|TAB364|TAB410|TAB411|TAB420|TAB424|TAB450|TAB460|TAB461|TAB464|TAB465|TAB467|TAB468|TAB07-100|TAB07-101|TAB07-150|TAB07-151|TAB07-152|TAB07-200|TAB07-201-3G|TAB07-210|TAB07-211|TAB07-212|TAB07-214|TAB07-220|TAB07-400|TAB07-485|TAB08-150|TAB08-200|TAB08-201-3G|TAB08-201-30|TAB09-100|TAB09-211|TAB09-410|TAB10-150|TAB10-201|TAB10-211|TAB10-400|TAB10-410|TAB13-201|TAB274EUK|TAB275EUK|TAB374EUK|TAB462EUK|TAB474EUK|TAB9-200)\b`,
		MEDIONTABLET: `Android.*\bOYO\b|LIFE.*(P9212|P9514|P9516|S9512)|LIFETAB`,
		ARNOVATABLET: `AN10G2|AN7bG3|AN7fG3|AN8G3|AN8cG3|AN7G3|AN9G3|AN7dG3|AN7dG3ST|AN7dG3ChildPad|AN10bG3|AN10bG3DT`,
		// http://www.intenso.de/kategorie_en.php?kategorie=33
		// @todo: http://www.nbhkdz.com/read/b8e64202f92a2df129126bff.html - investigate
		INTENSOTABLET: `INM8002KP|INM1010FP|INM805ND|Intenso Tab`,
		// IRU.ru Tablets http://www.iru.ru/catalog/soho/planetable/
		IRUTABLET:     `M702pro`,
		MEGAFONTABLET: `MegaFon V9|\bZTE V9\b|Android.*\bMT7A\b`,
		// @ref: http://www.e-boda.ro/tablete-pc.html
		EBODATABLET: `E-Boda (Supreme|Impresspeed|Izzycomm|Essential)`,
		// @ref: http://www.allview.ro/produse/droseries/lista-tablete-pc/
		ALLVIEWTABLET: `Allview.*(Viva|Alldro|City|Speed|All TV|Frenzy|Quasar|Shine|TX1|AX1|AX2)`,
		// @reference: http://wiki.archosfans.com/index.php?title=Main_Page
		ARCHOSTABLET: `\b(101G9|80G9|A101IT)\b|Qilive 97R|ARCHOS 101G10`,
		// @ref: http://www.ainol.com/plugin.php?identifier=ainol&module=product
		AINOLTABLET: `NOVO7|NOVO8|NOVO10|Novo7Aurora|Novo7Basic|NOVO7PALADIN|novo9-Spark`,
		// @todo: inspect http://esupport.sony.com/US/p/select-system.pl?DIRECTOR=DRIVER
		// @ref: Readers http://www.atsuhiro-me.net/ebook/sony-reader/sony-reader-web-browser
		// @ref: http://www.sony.jp/support/tablet/
		SONYTABLET: `Sony.*Tablet|Xperia Tablet|Sony Tablet S|SO-03E|SGPT12|SGPT13|SGPT114|SGPT121|SGPT122|SGPT123|SGPT111|SGPT112|SGPT113|SGPT131|SGPT132|SGPT133|SGPT211|SGPT212|SGPT213|SGP311|SGP312|SGP321|EBRD1101|EBRD1102|EBRD1201`,
		// @ref: db + http://www.cube-tablet.com/buy-products.html
		CUBETABLET: `Android.*(K8GT|U9GT|U10GT|U16GT|U17GT|U18GT|U19GT|U20GT|U23GT|U30GT)|CUBE U8GT`,
		// @ref: http://www.cobyusa.com/?p=pcat&pcat_id=3001
		COBYTABLET: `MID1042|MID1045|MID1125|MID1126|MID7012|MID7014|MID7015|MID7034|MID7035|MID7036|MID7042|MID7048|MID7127|MID8042|MID8048|MID8127|MID9042|MID9740|MID9742|MID7022|MID7010`,
		// @ref: http://www.match.net.cn/products.asp
		MIDTABLET: `M9701|M9000|M9100|M806|M1052|M806|T703|MID701|MID713|MID710|MID727|MID760|MID830|MID728|MID933|MID125|MID810|MID732|MID120|MID930|MID800|MID731|MID900|MID100|MID820|MID735|MID980|MID130|MID833|MID737|MID960|MID135|MID860|MID736|MID140|MID930|MID835|MID733`,
		// @ref: http://pdadb.net/index.php?m=pdalist&list=SMiT (NoName Chinese Tablets)
		// @ref: http://www.imp3.net/14/show.php?itemid=20454
		SMITTABLET: `Android.*(\bMID\b|MID-560|MTV-T1200|MTV-PND531|MTV-P1101|MTV-PND530)`,
		// @ref: http://www.rock-chips.com/index.php?do=prod&pid=2
		ROCKCHIPTABLET: `Android.*(RK2818|RK2808A|RK2918|RK3066)|RK2738|RK2808A`,
		// @ref: http://www.fly-phone.com/devices/tablets/ ; http://www.fly-phone.com/service/
		FLYTABLET: `IQ310|Fly Vision`,
		// @ref: http://www.bqreaders.com/gb/tablets-prices-sale.html
		BQTABLET: `bq.*(Elcano|Curie|Edison|Maxwell|Kepler|Pascal|Tesla|Hypatia|Platon|Newton|Livingstone|Cervantes|Avant)|Maxwell.*Lite|Maxwell.*Plus`,
		// @ref: http://www.huaweidevice.com/worldwide/productFamily.do?method=index&directoryId=5011&treeId=3290
		// @ref: http://www.huaweidevice.com/worldwide/downloadCenter.do?method=index&directoryId=3372&treeId=0&tb=1&type=software (including legacy tablets)
		HUAWEITABLET: `MediaPad|MediaPad 7 Youth|IDEOS S7|S7-201c|S7-202u|S7-101|S7-103|S7-104|S7-105|S7-106|S7-201|S7-Slim`,
		// Nec or Medias Tab
		NECTABLET: `\bN-06D|\bN-08D`,
		// Pantech Tablets: http://www.pantechusa.com/phones/
		PANTECHTABLET: `Pantech.*P4100`,
		// Broncho Tablets: http://www.broncho.cn/ (hard to find)
		BRONCHOTABLET: `Broncho.*(N701|N708|N802|a710)`,
		// @ref: http://versusuk.com/support.html
		VERSUSTABLET: `TOUCHPAD.*[78910]|\bTOUCHTAB\b`,
		// @ref: http://www.zync.in/index.php/our-products/tablet-phablets
		ZYNCTABLET: `z1000|Z99 2G|z99|z930|z999|z990|z909|Z919|z900`,
		// @ref: http://www.positivoinformatica.com.br/www/pessoal/tablet-ypy/
		POSITIVOTABLET: `TB07STA|TB10STA|TB07FTA|TB10FTA`,
		// @ref: https://www.nabitablet.com/
		NABITABLET: `Android.*\bNabi`,
		KOBOTABLET: `Kobo Touch|\bK080\b|\bVox\b Build|\bArc\b Build`,
		// French Danew Tablets http://www.danew.com/produits-tablette.php
		DANEWTABLET: `DSlide.*\b(700|701R|702|703R|704|802|970|971|972|973|974|1010|1012)\b`,
		// Texet Tablets and Readers http://www.texet.ru/tablet/
		TEXETTABLET: `NaviPad|TB-772A|TM-7045|TM-7055|TM-9750|TM-7016|TM-7024|TM-7026|TM-7041|TM-7043|TM-7047|TM-8041|TM-9741|TM-9747|TM-9748|TM-9751|TM-7022|TM-7021|TM-7020|TM-7011|TM-7010|TM-7023|TM-7025|TM-7037W|TM-7038W|TM-7027W|TM-9720|TM-9725|TM-9737W|TM-1020|TM-9738W|TM-9740|TM-9743W|TB-807A|TB-771A|TB-727A|TB-725A|TB-719A|TB-823A|TB-805A|TB-723A|TB-715A|TB-707A|TB-705A|TB-709A|TB-711A|TB-890HD|TB-880HD|TB-790HD|TB-780HD|TB-770HD|TB-721HD|TB-710HD|TB-434HD|TB-860HD|TB-840HD|TB-760HD|TB-750HD|TB-740HD|TB-730HD|TB-722HD|TB-720HD|TB-700HD|TB-500HD|TB-470HD|TB-431HD|TB-430HD|TB-506|TB-504|TB-446|TB-436|TB-416|TB-146SE|TB-126SE`,
		// @note: Avoid detecting `PLAYSTATION 3` as mobile.
		PLAYSTATIONTABLET: `Playstation.*(Portable|Vita)`,
		// @ref: http://www.trekstor.de/surftabs.html
		TREKSTORTABLET: `ST10416-1|VT10416-1|ST70408-1|ST702xx-1|ST702xx-2|ST80208|ST97216|ST70104-2`,
		// @ref: http://www.pyleaudio.com/Products.aspx?%2fproducts%2fPersonal-Electronics%2fTablets
		PYLEAUDIOTABLET: `\b(PTBL10CEU|PTBL10C|PTBL72BC|PTBL72BCEU|PTBL7CEU|PTBL7C|PTBL92BC|PTBL92BCEU|PTBL9CEU|PTBL9CUK|PTBL9C)\b`,
		// @ref: http://www.advandigital.com/index.php?link=content-product&jns=JP001
		ADVANTABLET: `Android.* \b(E3A|T3X|T5C|T5B|T3E|T3C|T3B|T1J|T1F|T2A|T1H|T1i|E1C|T1-E|T5-A|T4|E1-B|T2Ci|T1-B|T1-D|O1-A|E1-A|T1-A|T3A|T4i)\b `,
		// @ref: http://www.danytech.com/category/tablet-pc
		DANYTECHTABLET: `Genius Tab G3|Genius Tab S2|Genius Tab Q3|Genius Tab G4|Genius Tab Q4|Genius Tab G-II|Genius TAB GII|Genius TAB GIII|Genius Tab S1`,
		// @ref: http://www.galapad.net/product.html
		GALAPADTABLET: `Android.*\bG1\b`,
		// @ref: http://www.micromaxinfo.com/tablet/funbook
		MICROMAXTABLET: `Funbook|Micromax.*\b(P250|P560|P360|P362|P600|P300|P350|P500|P275)\b`,
		// http://www.karbonnmobiles.com/products_tablet.php
		KARBONNTABLET: `Android.*\b(A39|A37|A34|ST8|ST10|ST7|Smart Tab3|Smart Tab2)\b`,
		// @ref: http://www.myallfine.com/Products.asp
		ALLFINETABLET: `Fine7 Genius|Fine7 Shine|Fine7 Air|Fine8 Style|Fine9 More|Fine10 Joy|Fine11 Wide`,
		// @ref: http://www.proscanvideo.com/products-search.asp?itemClass=TABLET&itemnmbr=
		PROSCANTABLET: `\b(PEM63|PLT1023G|PLT1041|PLT1044|PLT1044G|PLT1091|PLT4311|PLT4311PL|PLT4315|PLT7030|PLT7033|PLT7033D|PLT7035|PLT7035D|PLT7044K|PLT7045K|PLT7045KB|PLT7071KG|PLT7072|PLT7223G|PLT7225G|PLT7777G|PLT7810K|PLT7849G|PLT7851G|PLT7852G|PLT8015|PLT8031|PLT8034|PLT8036|PLT8080K|PLT8082|PLT8088|PLT8223G|PLT8234G|PLT8235G|PLT8816K|PLT9011|PLT9045K|PLT9233G|PLT9735|PLT9760G|PLT9770G)\b`,
		// @ref: http://www.yonesnav.com/products/products.php
		YONESTABLET: `BQ1078|BC1003|BC1077|RK9702|BC9730|BC9001|IT9001|BC7008|BC7010|BC708|BC728|BC7012|BC7030|BC7027|BC7026`,
		// @ref: http://www.cjshowroom.com/eproducts.aspx?classcode=004001001
		// China manufacturer makes tablets for different small brands (eg. http://www.zeepad.net/index.html)
		CHANGJIATABLET: `TPC7102|TPC7103|TPC7105|TPC7106|TPC7107|TPC7201|TPC7203|TPC7205|TPC7210|TPC7708|TPC7709|TPC7712|TPC7110|TPC8101|TPC8103|TPC8105|TPC8106|TPC8203|TPC8205|TPC8503|TPC9106|TPC9701|TPC97101|TPC97103|TPC97105|TPC97106|TPC97111|TPC97113|TPC97203|TPC97603|TPC97809|TPC97205|TPC10101|TPC10103|TPC10106|TPC10111|TPC10203|TPC10205|TPC10503`,
		// @ref: http://www.gloryunion.cn/products.asp
		// @ref: http://www.allwinnertech.com/en/apply/mobile.html
		// @ref: http://www.ptcl.com.pk/pd_content.php?pd_id=284 (EVOTAB)
		// @todo: Softwiner tablets?
		// aka. Cute or Cool tablets. Not sure yet, must research to avoid collisions.
		GUTABLET: `TX-A1301|TX-M9002|Q702|kf026`, // A12R|D75A|D77|D79|R83|A95|A106C|R15|A75|A76|D71|D72|R71|R73|R77|D82|R85|D92|A97|D92|R91|A10F|A77F|W71F|A78F|W78F|W81F|A97F|W91F|W97F|R16G|C72|C73E|K72|K73|R96G
		// @ref: http://www.pointofview-online.com/showroom.php?shop_mode=product_listing&category_id=118
		POINTOFVIEWTABLET: `TAB-P506|TAB-navi-7-3G-M|TAB-P517|TAB-P-527|TAB-P701|TAB-P703|TAB-P721|TAB-P731N|TAB-P741|TAB-P825|TAB-P905|TAB-P925|TAB-PR945|TAB-PL1015|TAB-P1025|TAB-PI1045|TAB-P1325|TAB-PROTAB[0-9]+|TAB-PROTAB25|TAB-PROTAB26|TAB-PROTAB27|TAB-PROTAB26XL|TAB-PROTAB2-IPS9|TAB-PROTAB30-IPS9|TAB-PROTAB25XXL|TAB-PROTAB26-IPS10|TAB-PROTAB30-IPS10`,
		// @ref: http://www.overmax.pl/pl/katalog-produktow,p8/tablety,c14/
		// @todo: add more tests.
		OVERMAXTABLET: `OV-(SteelCore|NewBase|Basecore|Baseone|Exellen|Quattor|EduTab|Solution|ACTION|BasicTab|TeddyTab|MagicTab|Stream|TB-08|TB-09)`,
		// @ref: http://hclmetablet.com/India/index.php
		HCLTABLET: `HCL.*Tablet|Connect-3G-2.0|Connect-2G-2.0|ME Tablet U1|ME Tablet U2|ME Tablet G1|ME Tablet X1|ME Tablet Y2|ME Tablet Sync`,
		// @ref: http://www.edigital.hu/es_e-book_olvaso/Tablet-c18385.html
		DPSTABLET: `DPS Dream 9|DPS Dual 7`,
		// @ref: http://www.visture.com/index.asp
		VISTURETABLET: `V97 HD|i75 3G|Visture V4( HD)?|Visture V5( HD)?|Visture V10`,
		// @ref: http://www.mijncresta.nl/tablet
		CRESTATABLET: `CTP(-)?810|CTP(-)?818|CTP(-)?828|CTP(-)?838|CTP(-)?888|CTP(-)?978|CTP(-)?980|CTP(-)?987|CTP(-)?988|CTP(-)?989`,
		// MediaTek - http://www.mediatek.com/_en/01_products/02_proSys.php?cata_sn=1&cata1_sn=1&cata2_sn=309
		MEDIATEKTABLET: `\bMT8125|MT8389|MT8135|MT8377\b`,
		// Concorde tab
		CONCORDETABLET: `Concorde([ ]+)?Tab|ConCorde ReadMan`,
		// GoClever Tablets - http://www.goclever.com/uk/products,c1/tablet,c5/
		GOCLEVERTABLET: `GOCLEVER TAB|A7GOCLEVER|M1042|M7841|M742|R1042BK|R1041|TAB A975|TAB A7842|TAB A741|TAB A741L|TAB M723G|TAB M721|TAB A1021|TAB I921|TAB R721|TAB I720|TAB T76|TAB R70|TAB R76.2|TAB R106|TAB R83.2|TAB M813G|TAB I721|GCTA722|TAB I70|TAB I71|TAB S73|TAB R73|TAB R74|TAB R93|TAB R75|TAB R76.1|TAB A73|TAB A93|TAB A93.2|TAB T72|TAB R83|TAB R974|TAB R973|TAB A101|TAB A103|TAB A104|TAB A104.2|R105BK|M713G|A972BK|TAB A971|TAB R974.2|TAB R104|TAB R83.3|TAB A1042`,
		// Modecom Tablets - http://www.modecom.eu/tablets/portal/
		MODECOMTABLET: `FreeTAB 9000|FreeTAB 7.4|FreeTAB 7004|FreeTAB 7800|FreeTAB 2096|FreeTAB 7.5|FreeTAB 1014|FreeTAB 1001 |FreeTAB 8001|FreeTAB 9706|FreeTAB 9702|FreeTAB 7003|FreeTAB 7002|FreeTAB 1002|FreeTAB 7801|FreeTAB 1331|FreeTAB 1004|FreeTAB 8002|FreeTAB 8014|FreeTAB 9704|FreeTAB 1003`,
		// Vonino Tablets - http://www.vonino.eu/tablets
		VONINOTABLET: `\b(Argus[ _]?S|Diamond[ _]?79HD|Emerald[ _]?78E|Luna[ _]?70C|Onyx[ _]?S|Onyx[ _]?Z|Orin[ _]?HD|Orin[ _]?S|Otis[ _]?S|SpeedStar[ _]?S|Magnet[ _]?M9|Primus[ _]?94[ _]?3G|Primus[ _]?94HD|Primus[ _]?QS|Android.*\bQ8\b|Sirius[ _]?EVO[ _]?QS|Sirius[ _]?QS|Spirit[ _]?S)\b`,
		// ECS Tablets - http://www.ecs.com.tw/ECSWebSite/Product/Product_List.aspx?CategoryID=14&MenuID=107&childid=M_107&LanID=0
		ECSTABLET: `V07OT2|TM105A|S10OT1|TR10CS1`,
		// Storex Tablets - http://storex.fr/espace_client/support.html
		// @note: no need to add all the tablet codes since they are guided by the first regex.
		STOREXTABLET: `eZee[_']?(Tab|Go)[0-9]+|TabLC7|Looney Tunes Tab`,
		// Generic Vodafone tablets.
		VODAFONETABLET: `SmartTab([ ]+)?[0-9]+|SmartTabII10`,
		// French tablets - Essentiel B http://www.boulanger.fr/tablette_tactile_e-book/tablette_tactile_essentiel_b/cl_68908.htm?multiChoiceToDelete=brand&mc_brand=essentielb
		// Aka: http://www.essentielb.fr/
		ESSENTIELBTABLET: `Smart[ ']?TAB[ ]+?[0-9]+|Family[ ']?TAB2`,
		// Ross & Moor - http://ross-moor.ru/
		ROSSMOORTABLET: `RM-790|RM-997|RMD-878G|RMD-974R|RMT-705A|RMT-701|RME-601|RMT-501|RMT-711`,
		// i-mobile http://product.i-mobilephone.com/Mobile_Device
		IMOBILETABLET: `i-mobile i-note`,
		// @ref: http://www.tolino.de/de/vergleichen/
		TOLINOTABLET:     `tolino tab [0-9.]+|tolino shine`,
		AUDIOSONICTABLET: `\\bC-22Q|T7-QC|T-17B|T-17P\\b`,
		AMPETABLET:       `Android.* A78 `,
		SKKTABLET:        `Android.* (SKYPAD|PHOENIX|CYCLOPS)`,
		TECNOTABLET:      `TECNO P9`,
		JXDTABLET:        `Android.*\b(F3000|A3300|JXD5000|JXD3000|JXD2000|JXD300B|JXD300|S5800|S7800|S602b|S5110b|S7300|S5300|S602|S603|S5100|S5110|S601|S7100a|P3000F|P3000s|P101|P200s|P1000m|P200m|P9100|P1000s|S6600b|S908|P1000|P300|S18|S6600|S9100)\b`,
		IJOYTABLET:       `Tablet (Spirit 7|Essentia|Galatea|Fusion|Onix 7|Landa|Titan|Scooby|Deox|Stella|Themis|Argon|Unique 7|Sygnus|Hexen|Finity 7|Cream|Cream X2|Jade|Neon 7|Neron 7|Kandy|Scape|Saphyr 7|Rebel|Biox|Rebel|Rebel 8GB|Myst|Draco 7|Myst|Tab7-004|Myst|Tadeo Jones|Tablet Boing|Arrow|Draco Dual Cam|Aurix|Mint|Amity|Revolution|Finity 9|Neon 9|T9w|Amity 4GB Dual Cam|Stone 4GB|Stone 8GB|Andromeda|Silken|X2|Andromeda II|Halley|Flame|Saphyr 9,7|Touch 8|Planet|Triton|Unique 10|Hexen 10|Memphis 4GB|Memphis 8GB|Onix 10)`,
		// @ref: http://www.tesco.com/direct/hudl/
		HUDL: `Hudl HT7S3`,
		// @ref: http://www.telstra.com.au/home-phone/thub-2/
		TELSTRATABLET: `T-Hub2`,
		//`GenericTablet`: `Android.*\b97D\b|Tablet(?!.*PC)|ViewPad7|BNTV250A|MID-WCDMA|LogicPD Zoom2|\bA7EB\b|CatNova8|A1_07|CT704|CT1002|\bM721\b|rk30sdk|\bEVOTAB\b|M758A|ET904|ALUMIUM10|Smartfren Tab`,

	}
	operatingSystems = [...]string{
		ANDROIDOS:    `Android`,
		BLACKBERRYOS: `blackberry|\bBB10\b|rim tablet os`,
		PALMOS:       `PalmOS|avantgo|blazer|elaine|hiptop|palm|plucker|xiino`,
		SYMBIANOS:    `Symbian|SymbOS|Series60|Series40|SYB-[0-9]+|\bS60\b`,
		// @reference: http://en.wikipedia.org/wiki/Windows_Mobile
		WINDOWSMOBILEOS: `Windows CE.*(PPC|Smartphone|Mobile|[0-9]{3}x[0-9]{3})|Window Mobile|Windows Phone [0-9.]+|WCE;`,
		// @reference: http://en.wikipedia.org/wiki/Windows_Phone
		// http://wifeng.cn/?r=blog&a=view&id=106
		// http://nicksnettravels.builttoroam.com/post/2011/01/10/Bogus-Windows-Phone-7-User-Agent-String.aspx
		WINDOWSPHONEOS: `Windows Phone 8.0|Windows Phone OS|XBLWP7|ZuneWP7`,
		IOS:            `\biPhone.*Mobile|\biPod|\biPad`,
		// http://en.wikipedia.org/wiki/MeeGo
		// @todo: research MeeGo in UAs
		MEEGOOS: `MeeGo`,
		// http://en.wikipedia.org/wiki/Maemo
		// @todo: research Maemo in UAs
		MAEMOOS: `Maemo`,
		JAVAOS:  `J2ME/|\bMIDP\b|\bCLDC\b`, // `|Java/` produces bug #135
		WEBOS:   `webOS|hpwOS`,
		BADAOS:  `\bBada\b`,
		BREWOS:  `BREW`,
	}
	browsers = [...]string{
		// @reference: https://developers.google.com/chrome/mobile/docs/user-agent
		CHROME:   `\bCrMo\b|CriOS|Android.*Chrome/[.0-9]* (Mobile)?`,
		DOLFIN:   `\bDolfin\b`,
		OPERA:    `Opera.*Mini|Opera.*Mobi|Android.*Opera|Mobile.*OPR/[0-9.]+|Coast/[0-9.]+`,
		SKYFIRE:  `Skyfire`,
		IE:       `IEMobile|MSIEMobile`, // |Trident/[.0-9]+
		FIREFOX:  `fennec|firefox.*maemo|(Mobile|Tablet).*Firefox|Firefox.*Mobile`,
		BOLT:     `bolt`,
		TEASHARK: `teashark`,
		BLAZER:   `Blazer`,
		// @reference: http://developer.apple.com/library/safari/#documentation/AppleApplications/Reference/SafariWebContent/OptimizingforSafarioniPhone/OptimizingforSafarioniPhone.html#//apple_ref/doc/uid/TP40006517-SW3
		SAFARI: `Version.*Mobile.*Safari|Safari.*Mobile`,
		// @ref: http://en.wikipedia.org/wiki/Midori_(web_browser)
		//`Midori`          : `midori`,
		TIZEN:     `Tizen`,
		UCBROWSER: `UC.*Browser|UCWEB`,
		// @ref: https://github.com/serbanghita/Mobile-Detect/issues/7
		DIIGOBROWSER: `DiigoBrowser`,
		// http://www.puffinbrowser.com/index.php
		PUFFIN: `Puffin`,
		// @ref: http://mercury-browser.com/index.html
		MERCURY: `\bMercury\b`,
		// @reference: http://en.wikipedia.org/wiki/Minimo
		// http://en.wikipedia.org/wiki/Vision_Mobile_Browser
		GENERICBROWSER: `NokiaBrowser|OviBrowser|OneBrowser|TwonkyBeamBrowser|SEMC.*Browser|FlyFlow|Minimo|NetFront|Novarra-Vision|MQQBrowser|MicroMessenger`,
	}

	phoneKeyToName = [...]string{
		IPHONE:       `iphone`,
		BLACKBERRY:   `blackberry`,
		HTC:          `htc`,
		NEXUS:        `nexus`,
		DELL:         `dell`,
		MOTOROLA:     `motorola`,
		SAMSUNG:      `samsung`,
		LG:           `lg`,
		SONY:         `sony`,
		ASUS:         `asus`,
		MICROMAX:     `micromax`,
		PALM:         `palm`,
		VERTU:        `vertu`,
		PANTECH:      `pantech`,
		FLY:          `fly`,
		IMOBILE:      `imobile`,
		SIMVALLEY:    `simvalley`,
		GENERICPHONE: `genericphone`,
	}

	tabletKeyToName = [...]string{
		IPAD:              `ipad`,
		SAMSUNGTABLET:     `samsungtablet`,
		KINDLE:            `kindle`,
		SURFACETABLET:     `surfacetablet`,
		HPTABLET:          `hptablet`,
		ASUSTABLET:        `asustablet`,
		BLACKBERRYTABLET:  `blackberrytablet`,
		HTCTABLET:         `htctablet`,
		MOTOROLATABLET:    `motorolatablet`,
		NOOKTABLET:        `nooktablet`,
		ACERTABLET:        `acertablet`,
		TOSHIBATABLET:     `toshibatablet`,
		LGTABLET:          `lgtablet`,
		FUJITSUTABLET:     `fujitsutablet`,
		PRESTIGIOTABLET:   `prestigiotablet`,
		LENOVOTABLET:      `lenovotablet`,
		YARVIKTABLET:      `yarviktablet`,
		MEDIONTABLET:      `mediontablet`,
		ARNOVATABLET:      `arnovatablet`,
		INTENSOTABLET:     `intensotablet`,
		IRUTABLET:         `irutablet`,
		MEGAFONTABLET:     `megafontablet`,
		EBODATABLET:       `ebodatablet`,
		ALLVIEWTABLET:     `allviewtablet`,
		ARCHOSTABLET:      `archostablet`,
		AINOLTABLET:       `ainoltablet`,
		SONYTABLET:        `sonytablet`,
		CUBETABLET:        `cubetablet`,
		COBYTABLET:        `cobytablet`,
		MIDTABLET:         `midtablet`,
		SMITTABLET:        `smittablet`,
		ROCKCHIPTABLET:    `rockchiptablet`,
		FLYTABLET:         `flytablet`,
		BQTABLET:          `bqtablet`,
		HUAWEITABLET:      `huaweitablet`,
		NECTABLET:         `nectablet`,
		PANTECHTABLET:     `pantechtablet`,
		BRONCHOTABLET:     `bronchotablet`,
		VERSUSTABLET:      `versustablet`,
		ZYNCTABLET:        `zynctablet`,
		POSITIVOTABLET:    `positivotablet`,
		NABITABLET:        `nabitablet`,
		KOBOTABLET:        `kobotablet`,
		DANEWTABLET:       `danewtablet`,
		TEXETTABLET:       `texettablet`,
		PLAYSTATIONTABLET: `playstationtablet`,
		TREKSTORTABLET:    `trekstortablet`,
		PYLEAUDIOTABLET:   `pyleaudiotablet`,
		ADVANTABLET:       `advantablet`,
		DANYTECHTABLET:    `danytechtablet`,
		GALAPADTABLET:     `galapadtablet`,
		MICROMAXTABLET:    `micromaxtablet`,
		KARBONNTABLET:     `karbonntablet`,
		ALLFINETABLET:     `allfinetablet`,
		YONESTABLET:       `yonestablet`,
		CHANGJIATABLET:    `changjiatablet`,
		GUTABLET:          `gutablet`,
		POINTOFVIEWTABLET: `pointofviewtablet`,
		OVERMAXTABLET:     `overmaxtablet`,
		HCLTABLET:         `hcltablet`,
		DPSTABLET:         `dpstablet`,
		MEDIATEKTABLET:    `mediatektablet`,
		CONCORDETABLET:    `concordetablet`,
		GOCLEVERTABLET:    `goclevertablet`,
		MODECOMTABLET:     `modecomtablet`,
		VONINOTABLET:      `voninotablet`,
		ECSTABLET:         `ecstablet`,
		STOREXTABLET:      `storextablet`,
		VODAFONETABLET:    `vodafonetablet`,
		ESSENTIELBTABLET:  `essentielbtablet`,
		ROSSMOORTABLET:    `rossmoortablet`,
		IMOBILETABLET:     `imobiletablet`,
		TOLINOTABLET:      `tolinotablet`,
		AUDIOSONICTABLET:  `audiosonictablet`,
		AMPETABLET:        `ampetablet`,
		SKKTABLET:         `skktablet`,
		TECNOTABLET:       `tecnotablet`,
		JXDTABLET:         `jxdtablet`,
		IJOYTABLET:        `ijoytablet`,
		HUDL:              `hudl`,
		TELSTRATABLET:     `telstratablet`,
		PROSCANTABLET:     `proscantablet`,
		VISTURETABLET:     `visturetablet`,
		CRESTATABLET:      `crestatablet`,
	}

	osKeyToName = [...]string{
		ANDROIDOS:       `androidos`,
		BLACKBERRYOS:    `blackberryos`,
		PALMOS:          `palmos`,
		SYMBIANOS:       `symbianos`,
		WINDOWSMOBILEOS: `windowsmobileos`,
		WINDOWSPHONEOS:  `windowsphoneos`,
		IOS:             `ios`,
		MEEGOOS:         `meegoos`,
		MAEMOOS:         `maemoos`,
		JAVAOS:          `javaos`,
		WEBOS:           `webos`,
		BADAOS:          `badaos`,
		BREWOS:          `brewos`,
	}

	browserKeyToName = [...]string{
		CHROME:         `chrome`,
		DOLFIN:         `dolfin`,
		OPERA:          `opera`,
		SKYFIRE:        `skyfire`,
		IE:             `ie`,
		FIREFOX:        `firefox`,
		BOLT:           `bolt`,
		TEASHARK:       `teashark`,
		BLAZER:         `blazer`,
		SAFARI:         `safari`,
		TIZEN:          `tizen`,
		UCBROWSER:      `ucbrowser`,
		DIIGOBROWSER:   `diigobrowser`,
		PUFFIN:         `puffin`,
		MERCURY:        `mercury`,
		GENERICBROWSER: `genericbrowser`,
	}
)

// rules of detection for each kind of browser
type rules struct {
	phoneDevices         [len(phoneDevices)]string
	tabletDevices        [len(tabletDevices)]string
	operatingSystems     [len(operatingSystems)]string
	browsers             [len(browsers)]string
	mobileDetectionRules map[string]string
}

// NewRules creates a object with all rules necessary to figure out a browser from a User Agent string
func NewRules() *rules {
	rules := &rules{phoneDevices: phoneDevices, tabletDevices: tabletDevices, operatingSystems: operatingSystems, browsers: browsers}
	rules.setMobileDetectionRules()
	return rules
}

func (r *rules) MobileDetectionRules() map[string]string {
	return r.mobileDetectionRules
}

//Method sets the mobile detection rules. This method is used for the magic methods $detect->is*().
func (rules *rules) setMobileDetectionRules() {
	rulesCount := len(phoneDevices) +
		len(tabletDevices) +
		len(operatingSystems) +
		len(browsers)

	mobileDetectionRules := make(map[string]string, rulesCount)
	for ruleKey, ruleValue := range rules.phoneDevices {
		ruleName := phoneKeyToName[ruleKey]
		mobileDetectionRules[ruleName] = ruleValue
	}

	for ruleKey, ruleValue := range rules.tabletDevices {
		ruleName := tabletKeyToName[ruleKey]
		mobileDetectionRules[ruleName] = ruleValue
	}

	for ruleKey, ruleValue := range rules.operatingSystems {
		ruleName := osKeyToName[ruleKey]
		mobileDetectionRules[ruleName] = ruleValue
	}

	//TODO: why is it suppose to be here?
	for ruleKey, ruleValue := range rules.browsers {
		ruleName := browserKeyToName[ruleKey]
		mobileDetectionRules[ruleName] = ruleValue
	}

	rules.mobileDetectionRules = mobileDetectionRules
}
