package configs

var DefaultConfig = Config{
	Port:             8532,
	MetricsPort:      8533,
	Sentry:           "",
	ProviderPriority: []string{"binance", "huobi", "coingecko", "kucoin", "bitfinex", "kraken", "okx", "osmosis", "bitstamp", "bybit", "exchangerate", "frankfurter", "fer"},
	Providers: map[string]ProviderConfig{
		"binance": {
			Symbols: []string{
				"BTCUSDT",
				"ETHUSDT",
				"BNBUSDT",
				"NEOUSDT",
				"LTCUSDT",
				"QTUMUSDT",
				"ADAUSDT",
				"XRPUSDT",
				"EOSUSDT",
				"TUSDUSDT",
				"IOTAUSDT",
				"XLMUSDT",
				"ONTUSDT",
				"TRXUSDT",
				"ETCUSDT",
				"ICXUSDT",
				"VETUSDT",
				"BNBUSDC",
				"BTCUSDC",
				"ETHUSDC",
				"USDCUSDT",
				"LINKUSDT",
				"WAVESUSDT",
				"HOTUSDT",
				"ZILUSDT",
				"ZRXUSDT",
				"FETUSDT",
				"BATUSDT",
				"XMRUSDT",
				"ZECUSDT",
				"IOSTUSDT",
				"CELRUSDT",
				"DASHUSDT",
				"OMGUSDT",
				"THETAUSDT",
				"ENJUSDT",
				"MATICUSDT",
				"ATOMUSDT",
				"ONEUSDT",
				"FTMUSDT",
				"ALGOUSDT",
				"DOGEUSDT",
				"DUSKUSDT",
				"ANKRUSDT",
				"COCOSUSDT",
				"MTLUSDT",
				"TOMOUSDT",
				"DENTUSDT",
				"CHZUSDT",
				"BANDUSDT",
				"BUSDUSDT",
				"XTZUSDT",
				"RENUSDT",
				"RVNUSDT",
				"HBARUSDT",
				"NKNUSDT",
				"STXUSDT",
				"KAVAUSDT",
				"ARPAUSDT",
				"IOTXUSDT",
				"RLCUSDT",
				"BCHUSDT",
				"OGNUSDT",
				"COTIUSDT",
				"SOLUSDT",
				"CTSIUSDT",
				"CHRUSDT",
				"STMXUSDT",
				"KNCUSDT",
				"LRCUSDT",
				"COMPUSDT",
				"ZENUSDT",
				"SNXUSDT",
				"DGBUSDT",
				"SXPUSDT",
				"MKRUSDT",
				"STORJUSDT",
				"MANAUSDT",
				"YFIUSDT",
				"BALUSDT",
				"BLZUSDT",
				"ANTUSDT",
				"CRVUSDT",
				"SANDUSDT",
				"OCEANUSDT",
				"DOTUSDT",
				"RSRUSDT",
				"PAXGUSDT",
				"TRBUSDT",
				"SUSHIUSDT",
				"KSMUSDT",
				"EGLDUSDT",
				"RUNEUSDT",
				"BELUSDT",
				"UNIUSDT",
				"AVAXUSDT",
				"FLMUSDT",
				"ALPHAUSDT",
				"AAVEUSDT",
				"NEARUSDT",
				"FILUSDT",
				"INJUSDT",
				"AUDIOUSDT",
				"CTKUSDT",
				"AXSUSDT",
				"UNFIUSDT",
				"ROSEUSDT",
				"XEMUSDT",
				"SKLUSDT",
				"GRTUSDT",
				"1INCHUSDT",
				"REEFUSDT",
				"CELOUSDT",
				"TRUUSDT",
				"CKBUSDT",
				"TWTUSDT",
				"LITUSDT",
				"SFPUSDT",
				"DODOUSDT",
				"CAKEUSDT",
				"ALICEUSDT",
				"LINAUSDT",
				"PERPUSDT",
				"CFXUSDT",
				"TLMUSDT",
				"BAKEUSDT",
				"SHIBUSDT",
				"ICPUSDT",
				"ARUSDT",
				"MASKUSDT",
				"LPTUSDT",
				"ATAUSDT",
				"GTCUSDT",
				"KLAYUSDT",
				"C98USDT",
				"QNTUSDT",
				"FLOWUSDT",
				"MINAUSDT",
				"XECUSDT",
				"DYDXUSDT",
				"USDPUSDT",
				"GALAUSDT",
				"DARUSDT",
				"BNXUSDT",
				"ENSUSDT",
				"JASMYUSDT",
				"RNDRUSDT",
				"FXSUSDT",
				"HIGHUSDT",
				"CVXUSDT",
				"PEOPLEUSDT",
				"SPELLUSDT",
				"ACHUSDT",
				"IMXUSDT",
				"API3USDT",
				"WOOUSDT",
				"TUSDT",
				"ASTRUSDT",
				"GMTUSDT",
				"APEUSDT",
				"NEXOUSDT",
				"GALUSDT",
				"LDOUSDT",
				"OPUSDT",
				"LEVERUSDT",
				"STGUSDT",
				"LUNCUSDT",
				"GMXUSDT",
				"APTUSDT",
				"OSMOUSDT",
				"PHBUSDT",
				"HOOKUSDT",
				"MAGICUSDT",
				"RPLUSDT",
				"AGIXUSDT",
				"SSVUSDT",
				"LQTYUSDT",
				"AMBUSDT",
			},
		},
		"bitstamp": {
			Interval: 30,
			Timeout:  5,
			Symbols: []string{
				"1incheur",
				"1inchusd",
				"aaveeur",
				"aaveusd",
				"adaeur",
				"adausd",
				"algoeur",
				"algousd",
				"apeeur",
				"apeusd",
				"avaxeur",
				"avaxusd",
				"axseur",
				"axsusd",
				"bateur",
				"batusd",
				"bcheur",
				"bchusd",
				"btceur",
				"btcusd",
				"chzeur",
				"chzusd",
				"crveur",
				"crvusd",
				"cvxeur",
				"cvxusd",
				"daiusd",
				"dogeeur",
				"dogeusd",
				"doteur",
				"dotusd",
				"dydxeur",
				"dydxusd",
				"enjeur",
				"enjusd",
				"etheur",
				"ethusd",
				"flreur",
				"flrusd",
				"ftmeur",
				"ftmusd",
				"grteur",
				"grtusd",
				"gusdusd",
				"hbareur",
				"hbarusd",
				"imxeur",
				"imxusd",
				"injeur",
				"injusd",
				"linkeur",
				"linkusd",
				"lrceur",
				"lrcusd",
				"ltceur",
				"ltcusd",
				"manaeur",
				"manausd",
				"maticeur",
				"maticusd",
				"mkreur",
				"mkrusd",
				"neareur",
				"nearusd",
				"rndreur",
				"rndrusd",
				"sandeur",
				"sandusd",
				"shibeur",
				"shibusd",
				"snxeur",
				"snxusd",
				"soleur",
				"solusd",
				"unieur",
				"uniusd",
				"usdceur",
				"usdcusd",
				"usdteur",
				"usdtusd",
				"xlmeur",
				"xlmusd",
				"xrpeur",
				"xrpusd",
			},
		},
		"huobi": {
			Symbols: []string{
				"1inchusdt",
				"aaveusdt",
				"achusdt",
				"adausdt",
				"agixusdt",
				"algousdt",
				"ankrusdt",
				"apeusdt",
				"aptusdc",
				"aptusdt",
				"atomusdt",
				"avaxusdt",
				"axsusdc",
				"axsusdt",
				"bchusdc",
				"bchusdt",
				"bitusdt",
				"blurusdt",
				"bnbusdt",
				"bsvusdt",
				"btcusdc",
				"btcusdt",
				"bttusdc",
				"bttusdt",
				"c98usdt",
				"cakeusdt",
				"chzusdt",
				"coreusdt",
				"crvusdt",
				"csprusdt",
				"cvxusdt",
				"daiusdt",
				"dashusdt",
				"dogeusdt",
				"dotusdt",
				"dydxusdt",
				"egldusdt",
				"enjusdt",
				"eosusdt",
				"etcusdc",
				"etcusdt",
				"ethusdc",
				"ethusdt",
				"fetusdt",
				"filusdt",
				"flokiusdt",
				"flowusdc",
				"flowusdt",
				"flrusdt",
				"ftmusdt",
				"galusdt",
				"gmtusdt",
				"gmxusdt",
				"grtusdt",
				"gtusdt",
				"hbarusdc",
				"hbarusdt",
				"htusdt",
				"icpusdt",
				"imxusdt",
				"jstusdt",
				"kavausdt",
				"klayusdt",
				"ksmusdt",
				"ldousdt",
				"linkusdc",
				"linkusdt",
				"lrcusdc",
				"lrcusdt",
				"ltcusdt",
				"lunausdt",
				"luncusdt",
				"magicusdt",
				"manausdt",
				"maskusdt",
				"maticusdt",
				"minausdt",
				"mkrusdt",
				"nearusdt",
				"neousdt",
				"nexousdt",
				"oneusdt",
				"opusdc",
				"opusdt",
				"peopleusdt",
				"renusdt",
				"rndrusdt",
				"rplusdt",
				"sandusdc",
				"sandusdt",
				"shibusdt",
				"snxusdt",
				"solusdt",
				"ssvusdt",
				"sushiusdt",
				"thetausdc",
				"thetausdt",
				"tonusdt",
				"trxusdc",
				"trxusdt",
				"tusdusdt",
				"uniusdt",
				"usdcusdt",
				"usddusdc",
				"usddusdt",
				"usdpusdt",
				"vetusdc",
				"vetusdt",
				"wavesusdt",
				"winusdt",
				"woousdt",
				"xdcusdt",
				"xecusdt",
				"xlmusdt",
				"xmrusdt",
				"xrpusdc",
				"xrpusdt",
				"xtzusdc",
				"xtzusdt",
				"yfiusdt",
				"zecusdt",
				"zilusdt",
			},
		},
		"kucoin": {
			Symbols: []string{
				"1INCH-USDT",
				"AAVE-USDT",
				"ACH-USDT",
				"ADA-USDC",
				"ADA-USDT",
				"AGIX-USDT",
				"AGLD-USDT",
				"ALGO-USDC",
				"ALGO-USDT",
				"ALICE-USDT",
				"ANKR-USDT",
				"APE-USDC",
				"APE-USDT",
				"API3-USDT",
				"APT-USDT",
				"AR-USDT",
				"ARB-USDT",
				"ARPA-USDT",
				"ASTR-USDT",
				"ATOM-USDC",
				"ATOM-USDT",
				"AUDIO-USDT",
				"AVAX-USDC",
				"AVAX-USDT",
				"AXS-USDT",
				"BAL-USDT",
				"BAND-USDT",
				"BAT-USDT",
				"BCH-USDC",
				"BCH-USDT",
				"BCHSV-USDC",
				"BCHSV-USDT",
				"BLUR-USDT",
				"BNB-USDC",
				"BNB-USDT",
				"BTC-DAI",
				"BTC-USDC",
				"BTC-USDT",
				"BTT-USDT",
				"BUSD-USDC",
				"BUSD-USDT",
				"C98-USDT",
				"CAKE-USDT",
				"CELO-USDT",
				"CELR-USDT",
				"CFX-USDT",
				"CHR-USDT",
				"CHZ-USDT",
				"CKB-USDT",
				"COCOS-USDT",
				"COMP-USDT",
				"CRO-USDT",
				"CRV-USDT",
				"CSPR-USDT",
				"CTSI-USDT",
				"CVX-USDT",
				"DASH-USDT",
				"DGB-USDT",
				"DOGE-USDC",
				"DOGE-USDT",
				"DOT-USDC",
				"DOT-USDT",
				"DUSK-USDT",
				"DYDX-USDT",
				"EGLD-USDT",
				"ENJ-USDT",
				"ENS-USDT",
				"EOS-USDC",
				"EOS-USDT",
				"ETC-USDC",
				"ETC-USDT",
				"ETH-DAI",
				"ETH-USDC",
				"ETH-USDT",
				"FET-USDT",
				"FIL-USDT",
				"FLOKI-USDC",
				"FLOKI-USDT",
				"FLOW-USDT",
				"FLR-USDC",
				"FLR-USDT",
				"FTM-USDC",
				"FTM-USDT",
				"FXS-USDT",
				"GAL-USDT",
				"GALAX-USDT",
				"GMT-USDC",
				"GMT-USDT",
				"GMX-USDT",
				"GRT-USDT",
				"HBAR-USDT",
				"HIGH-USDT",
				"HT-USDT",
				"ICP-USDT",
				"IMX-USDT",
				"INJ-USDT",
				"IOST-USDT",
				"IOTX-USDT",
				"JASMY-USDC",
				"JASMY-USDT",
				"JST-USDT",
				"KAVA-USDT",
				"KCS-USDC",
				"KCS-USDT",
				"KDA-USDC",
				"KDA-USDT",
				"KLAY-USDT",
				"KNC-USDT",
				"KSM-USDT",
				"LDO-USDC",
				"LDO-USDT",
				"LINA-USDT",
				"LINK-USDC",
				"LINK-USDT",
				"LIT-USDT",
				"LOOKS-USDT",
				"LQTY-USDT",
				"LRC-USDT",
				"LTC-USDC",
				"LTC-USDT",
				"LUNA-USDC",
				"LUNA-USDT",
				"LUNC-USDC",
				"LUNC-USDT",
				"MAGIC-USDT",
				"MANA-USDT",
				"MASK-USDT",
				"MATIC-USDC",
				"MATIC-USDT",
				"MINA-USDT",
				"MKR-DAI",
				"MKR-USDT",
				"NEAR-USDC",
				"NEAR-USDT",
				"NEO-USDT",
				"NFT-USDT",
				"OCEAN-USDT",
				"OGN-USDT",
				"OMG-USDT",
				"ONE-USDT",
				"ONT-USDT",
				"OP-USDC",
				"OP-USDT",
				"OSMO-USDT",
				"PAXG-USDT",
				"PEOPLE-USDT",
				"PERP-USDT",
				"QNT-USDT",
				"REEF-USDT",
				"REN-USDT",
				"RNDR-USDT",
				"ROSE-USDT",
				"RPL-USDT",
				"RSR-USDT",
				"RUNE-USDC",
				"RUNE-USDT",
				"RVN-USDT",
				"SAND-USDT",
				"SHIB-USDC",
				"SHIB-USDT",
				"SLP-USDT",
				"SNX-USDT",
				"SOL-USDC",
				"SOL-USDT",
				"SSV-USDT",
				"STORJ-USDT",
				"STX-USDT",
				"SUSHI-USDT",
				"SXP-USDT",
				"T-USDT",
				"THETA-USDT",
				"TLM-USDT",
				"TON-USDT",
				"TRB-USDT",
				"TRU-USDT",
				"TRX-USDC",
				"TRX-USDT",
				"TWT-USDT",
				"UNFI-USDT",
				"UNI-USDT",
				"USDC-USDT",
				"USDD-USDC",
				"USDD-USDT",
				"USDP-USDT",
				"USDT-DAI",
				"USDT-USDC",
				"VET-USDT",
				"VRA-USDC",
				"VRA-USDT",
				"WAVES-USDT",
				"WAX-USDT",
				"WOO-USDT",
				"XCN-USDC",
				"XCN-USDT",
				"XDC-USDT",
				"XEC-USDT",
				"XEM-USDT",
				"XLM-USDT",
				"XMR-USDT",
				"XRP-USDC",
				"XRP-USDT",
				"XTZ-USDT",
				"YFI-USDT",
				"YGG-USDT",
				"ZEC-USDT",
				"ZIL-USDC",
				"ZIL-USDT",
			},
		},
		"kraken": {
			Symbols: []string{
				"1INCH/USD",
				"AAVE/USD",
				"ADA/USD",
				"ADA/USDT",
				"ALGO/USD",
				"ALGO/USDT",
				"APE/USD",
				"APE/USDT",
				"APT/USD",
				"ARB/USD",
				"ATOM/USD",
				"ATOM/USDT",
				"AVAX/USD",
				"AVAX/USDT",
				"AXS/USD",
				"BCH/USD",
				"BCH/USDT",
				"BIT/USD",
				"BTT/USD",
				"CHZ/USD",
				"CRV/USD",
				"CVX/USD",
				"DAI/USD",
				"DAI/USDT",
				"DASH/USD",
				"DOT/USD",
				"DOT/USDT",
				"DYDX/USD",
				"EGLD/USD",
				"ENJ/USD",
				"EOS/USD",
				"EOS/USDT",
				"ETC/USD",
				"ETH/DAI",
				"ETH/USD",
				"ETH/USDC",
				"ETH/USDT",
				"FIL/USD",
				"FLOW/USD",
				"FLR/USD",
				"FTM/USD",
				"FXS/USD",
				"GMX/USD",
				"GRT/USD",
				"ICP/USD",
				"IMX/USD",
				"KAVA/USD",
				"LDO/USD",
				"LINK/USD",
				"LINK/USDT",
				"LRC/USD",
				"LTC/USD",
				"LTC/USDT",
				"MANA/USD",
				"MANA/USDT",
				"MASK/USD",
				"MATIC/USD",
				"MATIC/USDT",
				"MINA/USD",
				"MKR/USD",
				"NEAR/USD",
				"PAXG/USD",
				"QNT/USD",
				"RNDR/USD",
				"RPL/USD",
				"RUNE/USD",
				"SAND/USD",
				"SHIB/USD",
				"SHIB/USDT",
				"SNX/USD",
				"SOL/USD",
				"SOL/USDT",
				"STX/USD",
				"TRX/USD",
				"UNI/USD",
				"USDC/USD",
				"USDC/USDT",
				"USDT/USD",
				"WBTC/USD",
				"XBT/DAI",
				"XBT/USD",
				"XBT/USDC",
				"XBT/USDT",
				"XDG/USD",
				"XDG/USDT",
				"XLM/USD",
				"XMR/USD",
				"XMR/USDT",
				"XRP/USD",
				"XRP/USDT",
				"XTZ/USD",
				"XTZ/USDT",
				"ZEC/USD",
			},
		},
		"bitfinex": {
			Symbols: []string{
				"tADAUSD",
				"tAPTUSD",
				"tATOUSD",
				"tAVAX:USD",
				"tBTCUSD",
				"tDAIUSD",
				"tDOGE:USD",
				"tDOTUSD",
				"tETCUSD",
				"tETHUSD",
				"tFILUSD",
				"tLEOUSD",
				"tLINK:USD",
				"tLTCUSD",
				"tMATIC:USD",
				"tSHIB:USD",
				"tSOLUSD",
				"tTRXUSD",
				"tUDCUSD",
				"tUNIUSD",
				"tUSTUSD",
				"tWBTUSD",
				"tXLMUSD",
				"tXMRUSD",
				"tXRPUSD",
				"tXTZUSD",
				"tZECUSD",
				"tSUSHI:USD",
				"tSAND:USD",
				"tAXSUSD",
			},
		},
		"okx": {
			Symbols: []string{
				"1INCH-USDC",
				"1INCH-USDT",
				"AAVE-USDC",
				"AAVE-USDT",
				"ADA-USDC",
				"ADA-USDT",
				"AGLD-USDC",
				"AGLD-USDT",
				"ALGO-USDC",
				"ALGO-USDT",
				"ALPHA-USDT",
				"ANT-USDC",
				"ANT-USDT",
				"APE-USDC",
				"APE-USDT",
				"API3-USDC",
				"API3-USDT",
				"APT-USDC",
				"APT-USDT",
				"AR-USDC",
				"AR-USDT",
				"ARB-USDT",
				"ATOM-USDC",
				"ATOM-USDT",
				"AVAX-USDC",
				"AVAX-USDT",
				"AXS-USDC",
				"AXS-USDT",
				"BADGER-USDT",
				"BAL-USDT",
				"BAND-USDT",
				"BAT-USDC",
				"BAT-USDT",
				"BCH-USDC",
				"BCH-USDT",
				"BICO-USDC",
				"BICO-USDT",
				"BLUR-USDC",
				"BLUR-USDT",
				"BNB-USDC",
				"BNB-USDT",
				"BNT-USDT",
				"BSV-USDC",
				"BSV-USDT",
				"BTC-DAI",
				"BTC-USDC",
				"BTC-USDT",
				"BTT-USDT",
				"CEL-USDC",
				"CEL-USDT",
				"CELO-USDC",
				"CELO-USDT",
				"CFX-USDT",
				"CHZ-USDC",
				"CHZ-USDT",
				"COMP-USDC",
				"COMP-USDT",
				"CORE-USDC",
				"CORE-USDT",
				"CRO-USDC",
				"CRO-USDT",
				"CRV-USDC",
				"CRV-USDT",
				"CSPR-USDC",
				"CSPR-USDT",
				"CVC-USDT",
				"CVX-USDT",
				"DAI-USDT",
				"DASH-USDC",
				"DASH-USDT",
				"DOGE-USDC",
				"DOGE-USDT",
				"DORA-USDT",
				"DOT-USDC",
				"DOT-USDT",
				"DYDX-USDC",
				"DYDX-USDT",
				"EGLD-USDC",
				"EGLD-USDT",
				"ENJ-USDT",
				"ENS-USDC",
				"ENS-USDT",
				"EOS-USDC",
				"EOS-USDT",
				"ETC-USDC",
				"ETC-USDT",
				"ETH-DAI",
				"ETH-USDC",
				"ETH-USDT",
				"ETHW-USDC",
				"ETHW-USDT",
				"FIL-USDC",
				"FIL-USDT",
				"FITFI-USDC",
				"FITFI-USDT",
				"FLM-USDC",
				"FLM-USDT",
				"FLOKI-USDC",
				"FLOKI-USDT",
				"FLOW-USDC",
				"FLOW-USDT",
				"FLR-USDC",
				"FLR-USDT",
				"FTM-USDC",
				"FTM-USDT",
				"GALA-USDC",
				"GALA-USDT",
				"GFT-USDC",
				"GFT-USDT",
				"GMT-USDC",
				"GMT-USDT",
				"GMX-USDC",
				"GMX-USDT",
				"GODS-USDC",
				"GODS-USDT",
				"GRT-USDC",
				"GRT-USDT",
				"HBAR-USDC",
				"HBAR-USDT",
				"ICP-USDC",
				"ICP-USDT",
				"IMX-USDC",
				"IMX-USDT",
				"IOST-USDC",
				"IOST-USDT",
				"IOTA-USDT",
				"JST-USDC",
				"JST-USDT",
				"KISHU-USDC",
				"KISHU-USDT",
				"KLAY-USDC",
				"KLAY-USDT",
				"KNC-USDC",
				"KNC-USDT",
				"KSM-USDC",
				"KSM-USDT",
				"LDO-USDT",
				"LEO-USDT",
				"LINK-USDC",
				"LINK-USDT",
				"LOOKS-USDC",
				"LOOKS-USDT",
				"LPT-USDT",
				"LRC-USDC",
				"LRC-USDT",
				"LTC-USDC",
				"LTC-USDT",
				"LUNA-USDC",
				"LUNA-USDT",
				"LUNC-USDC",
				"LUNC-USDT",
				"MAGIC-USDT",
				"MANA-USDC",
				"MANA-USDT",
				"MASK-USDC",
				"MASK-USDT",
				"MATIC-USDC",
				"MATIC-USDT",
				"MINA-USDC",
				"MINA-USDT",
				"MKR-USDC",
				"MKR-USDT",
				"NEAR-USDC",
				"NEAR-USDT",
				"NEO-USDT",
				"NFT-USDC",
				"NFT-USDT",
				"OKB-USDC",
				"OKB-USDT",
				"OMG-USDC",
				"OMG-USDT",
				"ONT-USDT",
				"OP-USDC",
				"OP-USDT",
				"PEOPLE-USDC",
				"PEOPLE-USDT",
				"PERP-USDC",
				"PERP-USDT",
				"QTUM-USDT",
				"REN-USDT",
				"RPL-USDC",
				"RPL-USDT",
				"RSR-USDC",
				"RSR-USDT",
				"RVN-USDT",
				"SAND-USDC",
				"SAND-USDT",
				"SHIB-USDC",
				"SHIB-USDT",
				"SLP-USDC",
				"SLP-USDT",
				"SNX-USDC",
				"SNX-USDT",
				"SOL-USDC",
				"SOL-USDT",
				"SSV-USDC",
				"SSV-USDT",
				"STARL-USDC",
				"STARL-USDT",
				"STORJ-USDC",
				"STORJ-USDT",
				"STX-USDC",
				"STX-USDT",
				"SUSHI-USDC",
				"SUSHI-USDT",
				"SWEAT-USDT",
				"THETA-USDC",
				"THETA-USDT",
				"TON-USDC",
				"TON-USDT",
				"TRB-USDC",
				"TRB-USDT",
				"TRX-USDC",
				"TRX-USDT",
				"UMA-USDT",
				"UNI-USDC",
				"UNI-USDT",
				"USDC-USDT",
				"USDT-USDC",
				"USTC-USDC",
				"USTC-USDT",
				"WAVES-USDC",
				"WAVES-USDT",
				"WBTC-USDT",
				"WOO-USDT",
				"XCH-USDC",
				"XCH-USDT",
				"XEC-USDT",
				"XLM-USDC",
				"XLM-USDT",
				"XMR-USDC",
				"XMR-USDT",
				"XRP-USDC",
				"XRP-USDT",
				"XTZ-USDC",
				"XTZ-USDT",
				"YFI-USDC",
				"YFI-USDT",
				"YFII-USDC",
				"YFII-USDT",
				"YGG-USDC",
				"YGG-USDT",
				"ZEC-USDC",
				"ZEC-USDT",
				"ZEN-USDT",
				"ZIL-USDC",
				"ZIL-USDT",
				"ZRX-USDT",
			},
		},
		"coingecko": {
			Interval: 6,
			Timeout:  5,
			Symbols: []string{
				"bitcoin",
				"ethereum",
				"binancecoin",
				"tether",
				"usd-coin",
				"binance-usd",
				"dai",
				"ripple",
				"dogecoin",
				"cardano",
				"matic-network",
				"polkadot",
				"litecoin",
				"staked-ether",
				"okb",
				"shiba-inu",
				"solana",
				"tron",
				"uniswap",
				"avalanche-2",
				"chainlink",
				"ethereum-classic",
				"the-open-network",
				"monero",
				"stellar",
				"algorand",
				"quant-network",
				"filecoin",
				"near",
				"vechain",
				"flow",
				"apecoin",
				"internet-computer",
				"elrond-erd-2",
				"chiliz",
				"eos",
				"chain-2",
				"tezos",
				"lido-dao",
				"the-sandbox",
				"theta-token",
				"aave",
				"axie-infinity",
				"decentraland",
				"iota",
				"maker",
				"pancakeswap-token",
				"aptos",
				"ecash",
				"zcash",
				"klay-token",
				"gatechain-token",
				"neo",
				"arweave",
				"dash",
				"fantom",
				"the-graph",
				"havven",
				"mina-protocol",
				"curve-dao-token",
				"nexo",
				"radix",
				"gmx",
				"basic-attention-token",
				"zilliqa",
				"ethereum-name-service",
				"1inch",
				"helium",
				"xdce-crowd-sale",
				"frax-share",
				"blockstack",
				"convex-finance",
				"enjincoin",
				"immutable-x",
				"loopring",
				"decred",
				"defichain",
				"theta-fuel",
				"amp-token",
				"compound-governance-token",
				"nxm",
				"dydx",
				"cosmos",
				"terra-luna-2",
				"crypto-com-chain",
				"osmosis",
				"thorchain",
				"evmos",
				"kava",
				"oec-token",
				"ankr",
				"kadena",
				"terra-luna",
				"terrausd",
				"injective-protocol",
				"secret",
				"juno-network",
				"stargaze",
				"akash-network",
			},
		},
		"osmosis": {
			Interval: 6,
			Symbols: []string{
				"ATOM/USDC",
				"AKT/USDC",
				"JUNO/USDC",
				"SCRT/USDC",
				"STARS/USDC",
				"OSMO/USDC",
				"INJ/USDC",
				"LUNA/USDC",
				"KAVA/USDC",
				"LINK/USDC",
				"DOT/USDC",
				"LUNC/USDC",
			},
		},
		"coinbase": {
			Symbols: []string{
				"1INCH-USD",
				"AAVE-USD",
				"ADA-USD",
				"ALGO-USD",
				"APE-USD",
				"APT-USD",
				"ARB-USD",
				"ATOM-USD",
				"AVAX-USD",
				"AXS-USD",
				"BAT-USD",
				"BCH-USD",
				"BIT-USD",
				"BTC-USD",
				"CHZ-USD",
				"CRO-USD",
				"CRV-USD",
				"DAI-USD",
				"DASH-USD",
				"DOGE-USD",
				"DOT-USD",
				"EGLD-USD",
				"ENJ-USD",
				"EOS-USD",
				"ETC-USD",
				"ETH-USD",
				"FIL-USD",
				"FLOW-USD",
				"GRT-USD",
				"GUSD-USD",
				"HBAR-USD",
				"ICP-USD",
				"IMX-USD",
				"INJ-USD",
				"KAVA-USD",
				"LDO-USD",
				"LINK-USD",
				"LRC-USD",
				"LTC-USD",
				"MANA-USD",
				"MASK-USD",
				"MATIC-USD",
				"MINA-USD",
				"MKR-USD",
				"NEAR-USD",
				"OP-USD",
				"QNT-USD",
				"RNDR-USD",
				"RPL-USD",
				"SAND-USD",
				"SHIB-USD",
				"SNX-USD",
				"SOL-USD",
				"STX-USD",
				"UNI-USD",
				"USDT-USD",
				"WBTC-USD",
				"XLM-USD",
				"XTZ-USD",
				"ZEC-USD",
			},
		},
		"bybit": {
			Symbols: []string{
				"1INCHUSDT",
				"AAVEUSDT",
				"ACHUSDT",
				"ADAUSDT",
				"AGIXUSDT",
				"AGLDUSDT",
				"ALGOUSDT",
				"ANKRUSDT",
				"APEUSDT",
				"APTUSDT",
				"ARBUSDT",
				"ARUSDT",
				"ATOMUSDT",
				"AVAXUSDT",
				"AXSUSDT",
				"BATUSDT",
				"BCHUSDT",
				"BELUSDT",
				"BICOUSDT",
				"BITUSDT",
				"BLURUSDT",
				"BNBUSDT",
				"BOBAUSDT",
				"BTCUSDT",
				"BTTUSDT",
				"BUSDUSDT",
				"C98USDT",
				"CELOUSDT",
				"CHZUSDT",
				"COMPUSDT",
				"COREUSDT",
				"CRVUSDT",
				"CTCUSDT",
				"DASHUSDT",
				"DGBUSDT",
				"DOGEUSDT",
				"DOTUSDT",
				"DYDXUSDT",
				"EGLDUSDT",
				"ENJUSDT",
				"ENSUSDT",
				"EOSUSDT",
				"ETCUSDT",
				"ETHUSDT",
				"ETHWUSDT",
				"FILUSDT",
				"FITFIUSDT",
				"FLOWUSDT",
				"FLRUSDT",
				"FTMUSDT",
				"FXSUSDT",
				"GALAUSDT",
				"GALUSDT",
				"GLMRUSDT",
				"GMTUSDT",
				"GMUSDT",
				"GMXUSDT",
				"GPTUSDT",
				"GRTUSDT",
				"HBARUSDT",
				"HFTUSDT",
				"HNTUSDT",
				"HOOKUSDT",
				"HOTUSDT",
				"ICPUSDT",
				"ICXUSDT",
				"IDUSDT",
				"IMXUSDT",
				"INJUSDT",
				"JASMYUSDT",
				"JSTUSDT",
				"KDAUSDT",
				"KLAYUSDT",
				"KSMUSDT",
				"LDOUSDT",
				"LINKUSDT",
				"LISUSDT",
				"LOOKSUSDT",
				"LRCUSDT",
				"LTCUSDT",
				"MAGICUSDT",
				"MANAUSDT",
				"MASKUSDT",
				"MATICUSDT",
				"MINAUSDT",
				"MKRUSDT",
				"NEARUSDT",
				"OMGUSDT",
				"ONEUSDT",
				"OPUSDT",
				"PAXGUSDT",
				"PEOPLEUSDT",
				"QTUMUSDT",
				"RENUSDT",
				"RNDRUSDT",
				"ROSEUSDT",
				"RPLUSDT",
				"RSS3USDT",
				"RUNEUSDT",
				"RVNUSDT",
				"SANDUSDT",
				"SCRTUSDT",
				"SCUSDT",
				"SDUSDT",
				"SLPUSDT",
				"SNXUSDT",
				"SOLUSDT",
				"SPELLUSDT",
				"SSVUSDT",
				"STGUSDT",
				"STXUSDT",
				"SUNUSDT",
				"SUSHIUSDT",
				"SWEATUSDT",
				"THETAUSDT",
				"TRXUSDT",
				"TUSDT",
				"TWTUSDT",
				"UNIUSDT",
				"USDCUSDT",
				"USDDUSDT",
				"WAVESUSDT",
				"WOOUSDT",
				"XEMUSDT",
				"XLMUSDT",
				"XRPUSDT",
				"XTZUSDT",
				"YFIUSDT",
				"ZECUSDT",
				"ZENUSDT",
				"ZILUSDT",
				"ZRXUSDT",
			},
		},
		"exchangerate": {
			Interval: 30,
			Timeout:  5,
			Symbols:  FiatCoins,
		},
		"frankfurter": {
			Interval: 30,
			Timeout:  5,
			Symbols:  FiatCoins,
		},
		"fer": {
			Interval: 30,
			Timeout:  5,
			Symbols:  FiatCoins,
		},
	},
}
