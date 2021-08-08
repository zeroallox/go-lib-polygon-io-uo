package polyconst

// Exchanges

type USEquityExchangeId int

const (
	USEQEX_INVALID        USEquityExchangeId = -1
	USEQEX_NONE           USEquityExchangeId = 0
	USEQEX_NYSE_AMEX      USEquityExchangeId = 1
	USEQEX_NASDAQ_OMX_BX  USEquityExchangeId = 2
	USEQEX_NYSE_NATIONAL  USEquityExchangeId = 3
	USEQEX_FINRA          USEquityExchangeId = 4
	USEQEX_CQS            USEquityExchangeId = 5
	USEQEX_ISE            USEquityExchangeId = 6
	USEQEX_EDGA           USEquityExchangeId = 7
	USEQEX_EDGX           USEquityExchangeId = 8
	USEQEX_CHICAGO        USEquityExchangeId = 9
	USEQEX_NYSE           USEquityExchangeId = 10
	USEQEX_NYSE_ARCA      USEquityExchangeId = 11
	USEQEX_NASDAQ         USEquityExchangeId = 12
	USEQEX_CTS            USEquityExchangeId = 13
	USEQEX_LTSE           USEquityExchangeId = 14
	USEQEX_IEX            USEquityExchangeId = 15
	USEQEX_CBSX           USEquityExchangeId = 16
	USEQEX_NASDAQ_OMX_PSX USEquityExchangeId = 17
	USEQEX_BATS_BYX       USEquityExchangeId = 18
	USEQEX_BATS_BZX       USEquityExchangeId = 19
	USEQEX_MIAX_PEARL     USEquityExchangeId = 20
	USEQEX_MEMX           USEquityExchangeId = 21
	USEQEX_NASDAQ_TAPE_B  USEquityExchangeId = 32
)

// TRFs and Other Stuff
type USEquityTRFId int

const (
	USEQTRF_FINRA_NYSE_TRF            USEquityTRFId = 201
	USEQTRF_FINRA_NASDAQ_TRF_CARTERET USEquityTRFId = 202
	USEQTRF_FINRA_NASDAQ_TRF_CHICAGO  USEquityTRFId = 203

	USEQTRF_SIP = 500
)

type USEquityTapeId int

const (
	USEQTP_INVALID USEquityTapeId = -1
	USEQTP_NONE    USEquityTapeId = 0
	USEQTP_CTA_1   USEquityTapeId = 1
	USEQTP_UTP_3   USEquityTapeId = 3
)

type USEquityTradeCondition int

const (
	USEQTC_Invalid            USEquityTradeCondition = -1
	USEQTC_Regular            USEquityTradeCondition = 0
	USEQTC_Acquisition        USEquityTradeCondition = 1
	USEQTC_AveragePrice       USEquityTradeCondition = 2
	USEQTC_AutomaticExecution USEquityTradeCondition = 3
	USEQTC_Bunched            USEquityTradeCondition = 4
	USEQTC_BunchSold          USEquityTradeCondition = 5
	USEQTC_CAPElection        USEquityTradeCondition = 6
	USEQTC_CashTrade          USEquityTradeCondition = 7
	USEQTC_Closing            USEquityTradeCondition = 8
	USEQTC_Cross              USEquityTradeCondition = 9
	USEQTC_DerivativelyPriced USEquityTradeCondition = 10
	USEQTC_Distribution       USEquityTradeCondition = 11
	USEQTC_FormTExtendedHours USEquityTradeCondition = 12
	USEQTC_FormTOutOfSequence USEquityTradeCondition = 13
	USEQTC_InterMarketSweep   USEquityTradeCondition = 14
	USEQTC_OfficialClose      USEquityTradeCondition = 15
	USEQTC_OfficialOpen       USEquityTradeCondition = 16
	USEQTC_Opening            USEquityTradeCondition = 17
	USEQTC_Reopening          USEquityTradeCondition = 18
	//	19 <<< DUPLICATE
	USEQTC_NextDay             USEquityTradeCondition = 20
	USEQTC_PriceVariation      USEquityTradeCondition = 21
	USEQTC_PriorReferencePrice USEquityTradeCondition = 22
	USEQTC_Rule155AMEX         USEquityTradeCondition = 23
	USEQTC_Rule127NYSE         USEquityTradeCondition = 24
	//	25 <<< DUPLICATE
	USEQTC_Opened              USEquityTradeCondition = 26
	USEQTC_RegularStoppedStock USEquityTradeCondition = 27
	//	28 <<< DUPLICATE
	USEQTC_Seller                            USEquityTradeCondition = 29
	USEQTC_SoldLast                          USEquityTradeCondition = 30
	USEQTC_SoldLastStoppedStock              USEquityTradeCondition = 31
	USEQTC_SoldOutOfSequence                 USEquityTradeCondition = 32
	USEQTC_SoldOutOfSequenceStoppedStock     USEquityTradeCondition = 33
	USEQTC_Split                             USEquityTradeCondition = 34
	USEQTC_StockOption                       USEquityTradeCondition = 35
	USEQTC_YellowFlag                        USEquityTradeCondition = 36
	USEQTC_OddLot                            USEquityTradeCondition = 37
	USEQTC_CorrectedConsolidatedClosingPrice USEquityTradeCondition = 38
	USEQTC_Unknown                           USEquityTradeCondition = 39
	USEQTC_Held                              USEquityTradeCondition = 40
	USEQTC_TradeThruExempt                   USEquityTradeCondition = 41
	USEQTC_NonEligible                       USEquityTradeCondition = 42
	USEQTC_NonEligibleExtended               USEquityTradeCondition = 43
	USEQTC_Cancelled                         USEquityTradeCondition = 44
	USEQTC_Recovery                          USEquityTradeCondition = 45
	USEQTC_Correction                        USEquityTradeCondition = 46
	USEQTC_AsOf                              USEquityTradeCondition = 47
	USEQTC_AsOfCorrection                    USEquityTradeCondition = 48
	USEQTC_AsOfCancel                        USEquityTradeCondition = 49
	USEQTC_OOB                               USEquityTradeCondition = 50
	USEQTC_Summary                           USEquityTradeCondition = 51
	USEQTC_Contingent                        USEquityTradeCondition = 52
	USEQTC_ContingentQualified               USEquityTradeCondition = 53
	USEQTC_Errored                           USEquityTradeCondition = 54
	USEQTC_OpeningReopeningTradeDetail       USEquityTradeCondition = 55
	USEQTC_IntradayTradeDetail               USEquityTradeCondition = 56
	//
	USEQTC_SSRActivated   USEquityTradeCondition = 57
	USEQTC_SSRContinued   USEquityTradeCondition = 58
	USEQTC_SSRDeactivated USEquityTradeCondition = 59
	USEQTC_SSRInEffect    USEquityTradeCondition = 60
	//
	USEQTC_FinancialStatusNormal                             USEquityTradeCondition = 61
	USEQTC_FinancialStatusBankrupt                           USEquityTradeCondition = 62
	USEQTC_FinancialStatusDeficient                          USEquityTradeCondition = 63
	USEQTC_FinancialStatusDelinquent                         USEquityTradeCondition = 64
	USEQTC_FinancialStatusBankruptAndDeficient               USEquityTradeCondition = 65
	USEQTC_FinancialStatusBankruptAndDelinquent              USEquityTradeCondition = 66
	USEQTC_FinancialStatusDeficientAndDelinquent             USEquityTradeCondition = 67
	USEQTC_FinancialStatusDeficientDelinquentAndBankrupt     USEquityTradeCondition = 68
	USEQTC_FinancialStatusLiquidation                        USEquityTradeCondition = 69
	USEQTC_FinancialStatusCreationsSuspended                 USEquityTradeCondition = 70
	USEQTC_FinancialStatusRedemptionsSuspended               USEquityTradeCondition = 71
	USEQTC_FinancialStatusCreationsAndOrRedemptionsSuspended USEquityTradeCondition = 72
)
