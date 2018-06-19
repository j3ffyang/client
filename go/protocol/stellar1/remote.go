// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/remote.avdl

package stellar1

import (
	"errors"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PaymentDirectPost struct {
	FromDeviceID      keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	To                *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	DisplayAmount     string                `codec:"displayAmount" json:"displayAmount"`
	DisplayCurrency   string                `codec:"displayCurrency" json:"displayCurrency"`
	NoteB64           string                `codec:"noteB64" json:"noteB64"`
	SignedTransaction string                `codec:"signedTransaction" json:"signedTransaction"`
	QuickReturn       bool                  `codec:"quickReturn" json:"quickReturn"`
}

func (o PaymentDirectPost) DeepCopy() PaymentDirectPost {
	return PaymentDirectPost{
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		DisplayAmount:     o.DisplayAmount,
		DisplayCurrency:   o.DisplayCurrency,
		NoteB64:           o.NoteB64,
		SignedTransaction: o.SignedTransaction,
		QuickReturn:       o.QuickReturn,
	}
}

type PaymentRelayPost struct {
	FromDeviceID      keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	To                *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	ToAssertion       string                `codec:"toAssertion" json:"toAssertion"`
	RelayAccount      AccountID             `codec:"relayAccount" json:"relayAccount"`
	TeamID            keybase1.TeamID       `codec:"teamID" json:"teamID"`
	DisplayAmount     string                `codec:"displayAmount" json:"displayAmount"`
	DisplayCurrency   string                `codec:"displayCurrency" json:"displayCurrency"`
	BoxB64            string                `codec:"boxB64" json:"boxB64"`
	SignedTransaction string                `codec:"signedTransaction" json:"signedTransaction"`
	QuickReturn       bool                  `codec:"quickReturn" json:"quickReturn"`
}

func (o PaymentRelayPost) DeepCopy() PaymentRelayPost {
	return PaymentRelayPost{
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		ToAssertion:       o.ToAssertion,
		RelayAccount:      o.RelayAccount.DeepCopy(),
		TeamID:            o.TeamID.DeepCopy(),
		DisplayAmount:     o.DisplayAmount,
		DisplayCurrency:   o.DisplayCurrency,
		BoxB64:            o.BoxB64,
		SignedTransaction: o.SignedTransaction,
		QuickReturn:       o.QuickReturn,
	}
}

type RelayClaimPost struct {
	KeybaseID         KeybaseTransactionID `codec:"keybaseID" json:"keybaseID"`
	Dir               RelayDirection       `codec:"dir" json:"dir"`
	SignedTransaction string               `codec:"signedTransaction" json:"signedTransaction"`
	AutoClaimToken    *string              `codec:"autoClaimToken,omitempty" json:"autoClaimToken,omitempty"`
}

func (o RelayClaimPost) DeepCopy() RelayClaimPost {
	return RelayClaimPost{
		KeybaseID:         o.KeybaseID.DeepCopy(),
		Dir:               o.Dir.DeepCopy(),
		SignedTransaction: o.SignedTransaction,
		AutoClaimToken: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.AutoClaimToken),
	}
}

type PaymentSummaryType int

const (
	PaymentSummaryType_NONE    PaymentSummaryType = 0
	PaymentSummaryType_STELLAR PaymentSummaryType = 1
	PaymentSummaryType_DIRECT  PaymentSummaryType = 2
	PaymentSummaryType_RELAY   PaymentSummaryType = 3
)

func (o PaymentSummaryType) DeepCopy() PaymentSummaryType { return o }

var PaymentSummaryTypeMap = map[string]PaymentSummaryType{
	"NONE":    0,
	"STELLAR": 1,
	"DIRECT":  2,
	"RELAY":   3,
}

var PaymentSummaryTypeRevMap = map[PaymentSummaryType]string{
	0: "NONE",
	1: "STELLAR",
	2: "DIRECT",
	3: "RELAY",
}

func (e PaymentSummaryType) String() string {
	if v, ok := PaymentSummaryTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type PaymentSummary struct {
	Typ__     PaymentSummaryType     `codec:"typ" json:"typ"`
	Stellar__ *PaymentSummaryStellar `codec:"stellar,omitempty" json:"stellar,omitempty"`
	Direct__  *PaymentSummaryDirect  `codec:"direct,omitempty" json:"direct,omitempty"`
	Relay__   *PaymentSummaryRelay   `codec:"relay,omitempty" json:"relay,omitempty"`
}

func (o *PaymentSummary) Typ() (ret PaymentSummaryType, err error) {
	switch o.Typ__ {
	case PaymentSummaryType_STELLAR:
		if o.Stellar__ == nil {
			err = errors.New("unexpected nil value for Stellar__")
			return ret, err
		}
	case PaymentSummaryType_DIRECT:
		if o.Direct__ == nil {
			err = errors.New("unexpected nil value for Direct__")
			return ret, err
		}
	case PaymentSummaryType_RELAY:
		if o.Relay__ == nil {
			err = errors.New("unexpected nil value for Relay__")
			return ret, err
		}
	}
	return o.Typ__, nil
}

func (o PaymentSummary) Stellar() (res PaymentSummaryStellar) {
	if o.Typ__ != PaymentSummaryType_STELLAR {
		panic("wrong case accessed")
	}
	if o.Stellar__ == nil {
		return
	}
	return *o.Stellar__
}

func (o PaymentSummary) Direct() (res PaymentSummaryDirect) {
	if o.Typ__ != PaymentSummaryType_DIRECT {
		panic("wrong case accessed")
	}
	if o.Direct__ == nil {
		return
	}
	return *o.Direct__
}

func (o PaymentSummary) Relay() (res PaymentSummaryRelay) {
	if o.Typ__ != PaymentSummaryType_RELAY {
		panic("wrong case accessed")
	}
	if o.Relay__ == nil {
		return
	}
	return *o.Relay__
}

func NewPaymentSummaryWithStellar(v PaymentSummaryStellar) PaymentSummary {
	return PaymentSummary{
		Typ__:     PaymentSummaryType_STELLAR,
		Stellar__: &v,
	}
}

func NewPaymentSummaryWithDirect(v PaymentSummaryDirect) PaymentSummary {
	return PaymentSummary{
		Typ__:    PaymentSummaryType_DIRECT,
		Direct__: &v,
	}
}

func NewPaymentSummaryWithRelay(v PaymentSummaryRelay) PaymentSummary {
	return PaymentSummary{
		Typ__:   PaymentSummaryType_RELAY,
		Relay__: &v,
	}
}

func (o PaymentSummary) DeepCopy() PaymentSummary {
	return PaymentSummary{
		Typ__: o.Typ__.DeepCopy(),
		Stellar__: (func(x *PaymentSummaryStellar) *PaymentSummaryStellar {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Stellar__),
		Direct__: (func(x *PaymentSummaryDirect) *PaymentSummaryDirect {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Direct__),
		Relay__: (func(x *PaymentSummaryRelay) *PaymentSummaryRelay {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Relay__),
	}
}

type PaymentSummaryStellar struct {
	TxID        TransactionID `codec:"txID" json:"txID"`
	From        AccountID     `codec:"from" json:"from"`
	To          AccountID     `codec:"to" json:"to"`
	Amount      string        `codec:"amount" json:"amount"`
	Asset       Asset         `codec:"asset" json:"asset"`
	OperationID uint64        `codec:"operationID" json:"operationID"`
	Ctime       TimeMs        `codec:"ctime" json:"ctime"`
	CursorToken string        `codec:"cursorToken" json:"cursorToken"`
}

func (o PaymentSummaryStellar) DeepCopy() PaymentSummaryStellar {
	return PaymentSummaryStellar{
		TxID:        o.TxID.DeepCopy(),
		From:        o.From.DeepCopy(),
		To:          o.To.DeepCopy(),
		Amount:      o.Amount,
		Asset:       o.Asset.DeepCopy(),
		OperationID: o.OperationID,
		Ctime:       o.Ctime.DeepCopy(),
		CursorToken: o.CursorToken,
	}
}

type PaymentSummaryDirect struct {
	KbTxID          KeybaseTransactionID  `codec:"kbTxID" json:"kbTxID"`
	TxID            TransactionID         `codec:"txID" json:"txID"`
	TxStatus        TransactionStatus     `codec:"txStatus" json:"txStatus"`
	TxErrMsg        string                `codec:"txErrMsg" json:"txErrMsg"`
	FromStellar     AccountID             `codec:"fromStellar" json:"fromStellar"`
	From            keybase1.UserVersion  `codec:"from" json:"from"`
	FromDeviceID    keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	ToStellar       AccountID             `codec:"toStellar" json:"toStellar"`
	To              *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	Amount          string                `codec:"amount" json:"amount"`
	Asset           Asset                 `codec:"asset" json:"asset"`
	DisplayAmount   *string               `codec:"displayAmount,omitempty" json:"displayAmount,omitempty"`
	DisplayCurrency *string               `codec:"displayCurrency,omitempty" json:"displayCurrency,omitempty"`
	NoteB64         string                `codec:"noteB64" json:"noteB64"`
	Ctime           TimeMs                `codec:"ctime" json:"ctime"`
	Rtime           TimeMs                `codec:"rtime" json:"rtime"`
	CursorToken     string                `codec:"cursorToken" json:"cursorToken"`
}

func (o PaymentSummaryDirect) DeepCopy() PaymentSummaryDirect {
	return PaymentSummaryDirect{
		KbTxID:       o.KbTxID.DeepCopy(),
		TxID:         o.TxID.DeepCopy(),
		TxStatus:     o.TxStatus.DeepCopy(),
		TxErrMsg:     o.TxErrMsg,
		FromStellar:  o.FromStellar.DeepCopy(),
		From:         o.From.DeepCopy(),
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		ToStellar:    o.ToStellar.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		Amount: o.Amount,
		Asset:  o.Asset.DeepCopy(),
		DisplayAmount: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayAmount),
		DisplayCurrency: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayCurrency),
		NoteB64:     o.NoteB64,
		Ctime:       o.Ctime.DeepCopy(),
		Rtime:       o.Rtime.DeepCopy(),
		CursorToken: o.CursorToken,
	}
}

type PaymentSummaryRelay struct {
	KbTxID          KeybaseTransactionID  `codec:"kbTxID" json:"kbTxID"`
	TxID            TransactionID         `codec:"txID" json:"txID"`
	TxStatus        TransactionStatus     `codec:"txStatus" json:"txStatus"`
	TxErrMsg        string                `codec:"txErrMsg" json:"txErrMsg"`
	FromStellar     AccountID             `codec:"fromStellar" json:"fromStellar"`
	From            keybase1.UserVersion  `codec:"from" json:"from"`
	FromDeviceID    keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	To              *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	ToAssertion     string                `codec:"toAssertion" json:"toAssertion"`
	RelayAccount    AccountID             `codec:"relayAccount" json:"relayAccount"`
	Amount          string                `codec:"amount" json:"amount"`
	DisplayAmount   *string               `codec:"displayAmount,omitempty" json:"displayAmount,omitempty"`
	DisplayCurrency *string               `codec:"displayCurrency,omitempty" json:"displayCurrency,omitempty"`
	Ctime           TimeMs                `codec:"ctime" json:"ctime"`
	Rtime           TimeMs                `codec:"rtime" json:"rtime"`
	BoxB64          string                `codec:"boxB64" json:"boxB64"`
	TeamID          keybase1.TeamID       `codec:"teamID" json:"teamID"`
	Claim           *ClaimSummary         `codec:"claim,omitempty" json:"claim,omitempty"`
	CursorToken     string                `codec:"cursorToken" json:"cursorToken"`
}

func (o PaymentSummaryRelay) DeepCopy() PaymentSummaryRelay {
	return PaymentSummaryRelay{
		KbTxID:       o.KbTxID.DeepCopy(),
		TxID:         o.TxID.DeepCopy(),
		TxStatus:     o.TxStatus.DeepCopy(),
		TxErrMsg:     o.TxErrMsg,
		FromStellar:  o.FromStellar.DeepCopy(),
		From:         o.From.DeepCopy(),
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		ToAssertion:  o.ToAssertion,
		RelayAccount: o.RelayAccount.DeepCopy(),
		Amount:       o.Amount,
		DisplayAmount: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayAmount),
		DisplayCurrency: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayCurrency),
		Ctime:  o.Ctime.DeepCopy(),
		Rtime:  o.Rtime.DeepCopy(),
		BoxB64: o.BoxB64,
		TeamID: o.TeamID.DeepCopy(),
		Claim: (func(x *ClaimSummary) *ClaimSummary {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Claim),
		CursorToken: o.CursorToken,
	}
}

type ClaimSummary struct {
	TxID      TransactionID        `codec:"txID" json:"txID"`
	TxStatus  TransactionStatus    `codec:"txStatus" json:"txStatus"`
	TxErrMsg  string               `codec:"txErrMsg" json:"txErrMsg"`
	Dir       RelayDirection       `codec:"dir" json:"dir"`
	ToStellar AccountID            `codec:"toStellar" json:"toStellar"`
	To        keybase1.UserVersion `codec:"to" json:"to"`
}

func (o ClaimSummary) DeepCopy() ClaimSummary {
	return ClaimSummary{
		TxID:      o.TxID.DeepCopy(),
		TxStatus:  o.TxStatus.DeepCopy(),
		TxErrMsg:  o.TxErrMsg,
		Dir:       o.Dir.DeepCopy(),
		ToStellar: o.ToStellar.DeepCopy(),
		To:        o.To.DeepCopy(),
	}
}

type PaymentDetails struct {
	Summary  PaymentSummary `codec:"summary" json:"summary"`
	Memo     string         `codec:"memo" json:"memo"`
	MemoType string         `codec:"memoType" json:"memoType"`
}

func (o PaymentDetails) DeepCopy() PaymentDetails {
	return PaymentDetails{
		Summary:  o.Summary.DeepCopy(),
		Memo:     o.Memo,
		MemoType: o.MemoType,
	}
}

type AccountDetails struct {
	AccountID     AccountID `codec:"accountID" json:"accountID"`
	Seqno         string    `codec:"seqno" json:"seqno"`
	Balances      []Balance `codec:"balances" json:"balances"`
	SubentryCount int       `codec:"subentryCount" json:"subentryCount"`
	Available     string    `codec:"available" json:"available"`
}

func (o AccountDetails) DeepCopy() AccountDetails {
	return AccountDetails{
		AccountID: o.AccountID.DeepCopy(),
		Seqno:     o.Seqno,
		Balances: (func(x []Balance) []Balance {
			if x == nil {
				return nil
			}
			var ret []Balance
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Balances),
		SubentryCount: o.SubentryCount,
		Available:     o.Available,
	}
}

type PaymentsPage struct {
	Payments []PaymentSummary `codec:"payments" json:"payments"`
	Cursor   *PageCursor      `codec:"cursor,omitempty" json:"cursor,omitempty"`
}

func (o PaymentsPage) DeepCopy() PaymentsPage {
	return PaymentsPage{
		Payments: (func(x []PaymentSummary) []PaymentSummary {
			if x == nil {
				return nil
			}
			var ret []PaymentSummary
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Payments),
		Cursor: (func(x *PageCursor) *PageCursor {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Cursor),
	}
}

type AutoClaim struct {
	KbTxID KeybaseTransactionID `codec:"kbTxID" json:"kbTxID"`
}

func (o AutoClaim) DeepCopy() AutoClaim {
	return AutoClaim{
		KbTxID: o.KbTxID.DeepCopy(),
	}
}

type PaymentStatusMsg struct {
	KeybaseID   KeybaseTransactionID `codec:"keybaseID" json:"keybaseID"`
	StellarID   TransactionID        `codec:"stellarID" json:"stellarID"`
	Pending     bool                 `codec:"pending" json:"pending"`
	Relay       bool                 `codec:"relay" json:"relay"`
	AutoClaimed bool                 `codec:"autoClaimed" json:"autoClaimed"`
	ErrMsg      *string              `codec:"errMsg,omitempty" json:"errMsg,omitempty"`
}

func (o PaymentStatusMsg) DeepCopy() PaymentStatusMsg {
	return PaymentStatusMsg{
		KeybaseID:   o.KeybaseID.DeepCopy(),
		StellarID:   o.StellarID.DeepCopy(),
		Pending:     o.Pending,
		Relay:       o.Relay,
		AutoClaimed: o.AutoClaimed,
		ErrMsg: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ErrMsg),
	}
}

type BalancesArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type DetailsArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type RecentPaymentsArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
	Cursor    *PageCursor          `codec:"cursor,omitempty" json:"cursor,omitempty"`
	Limit     int                  `codec:"limit" json:"limit"`
}

type PaymentDetailsArg struct {
	Caller keybase1.UserVersion `codec:"caller" json:"caller"`
	TxID   string               `codec:"txID" json:"txID"`
}

type AccountSeqnoArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type SubmitPaymentArg struct {
	Caller  keybase1.UserVersion `codec:"caller" json:"caller"`
	Payment PaymentDirectPost    `codec:"payment" json:"payment"`
}

type SubmitRelayPaymentArg struct {
	Caller  keybase1.UserVersion `codec:"caller" json:"caller"`
	Payment PaymentRelayPost     `codec:"payment" json:"payment"`
}

type SubmitRelayClaimArg struct {
	Caller keybase1.UserVersion `codec:"caller" json:"caller"`
	Claim  RelayClaimPost       `codec:"claim" json:"claim"`
}

type AcquireAutoClaimLockArg struct {
	Caller keybase1.UserVersion `codec:"caller" json:"caller"`
}

type ReleaseAutoClaimLockArg struct {
	Caller keybase1.UserVersion `codec:"caller" json:"caller"`
	Token  string               `codec:"token" json:"token"`
}

type NextAutoClaimArg struct {
	Caller keybase1.UserVersion `codec:"caller" json:"caller"`
}

type IsMasterKeyActiveArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type PingArg struct {
}

type RemoteInterface interface {
	Balances(context.Context, BalancesArg) ([]Balance, error)
	Details(context.Context, DetailsArg) (AccountDetails, error)
	RecentPayments(context.Context, RecentPaymentsArg) (PaymentsPage, error)
	PaymentDetails(context.Context, PaymentDetailsArg) (PaymentDetails, error)
	AccountSeqno(context.Context, AccountSeqnoArg) (string, error)
	SubmitPayment(context.Context, SubmitPaymentArg) (PaymentResult, error)
	SubmitRelayPayment(context.Context, SubmitRelayPaymentArg) (PaymentResult, error)
	SubmitRelayClaim(context.Context, SubmitRelayClaimArg) (RelayClaimResult, error)
	AcquireAutoClaimLock(context.Context, keybase1.UserVersion) (string, error)
	ReleaseAutoClaimLock(context.Context, ReleaseAutoClaimLockArg) error
	NextAutoClaim(context.Context, keybase1.UserVersion) (*AutoClaim, error)
	IsMasterKeyActive(context.Context, IsMasterKeyActiveArg) (bool, error)
	Ping(context.Context) (string, error)
}

func RemoteProtocol(i RemoteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "stellar.1.remote",
		Methods: map[string]rpc.ServeHandlerDescription{
			"balances": {
				MakeArg: func() interface{} {
					ret := make([]BalancesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]BalancesArg)
					if !ok {
						err = rpc.NewTypeError((*[]BalancesArg)(nil), args)
						return
					}
					ret, err = i.Balances(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"details": {
				MakeArg: func() interface{} {
					ret := make([]DetailsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DetailsArg)
					if !ok {
						err = rpc.NewTypeError((*[]DetailsArg)(nil), args)
						return
					}
					ret, err = i.Details(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"recentPayments": {
				MakeArg: func() interface{} {
					ret := make([]RecentPaymentsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RecentPaymentsArg)
					if !ok {
						err = rpc.NewTypeError((*[]RecentPaymentsArg)(nil), args)
						return
					}
					ret, err = i.RecentPayments(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"paymentDetails": {
				MakeArg: func() interface{} {
					ret := make([]PaymentDetailsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PaymentDetailsArg)
					if !ok {
						err = rpc.NewTypeError((*[]PaymentDetailsArg)(nil), args)
						return
					}
					ret, err = i.PaymentDetails(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"accountSeqno": {
				MakeArg: func() interface{} {
					ret := make([]AccountSeqnoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AccountSeqnoArg)
					if !ok {
						err = rpc.NewTypeError((*[]AccountSeqnoArg)(nil), args)
						return
					}
					ret, err = i.AccountSeqno(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"submitPayment": {
				MakeArg: func() interface{} {
					ret := make([]SubmitPaymentArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SubmitPaymentArg)
					if !ok {
						err = rpc.NewTypeError((*[]SubmitPaymentArg)(nil), args)
						return
					}
					ret, err = i.SubmitPayment(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"submitRelayPayment": {
				MakeArg: func() interface{} {
					ret := make([]SubmitRelayPaymentArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SubmitRelayPaymentArg)
					if !ok {
						err = rpc.NewTypeError((*[]SubmitRelayPaymentArg)(nil), args)
						return
					}
					ret, err = i.SubmitRelayPayment(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"submitRelayClaim": {
				MakeArg: func() interface{} {
					ret := make([]SubmitRelayClaimArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SubmitRelayClaimArg)
					if !ok {
						err = rpc.NewTypeError((*[]SubmitRelayClaimArg)(nil), args)
						return
					}
					ret, err = i.SubmitRelayClaim(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"acquireAutoClaimLock": {
				MakeArg: func() interface{} {
					ret := make([]AcquireAutoClaimLockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AcquireAutoClaimLockArg)
					if !ok {
						err = rpc.NewTypeError((*[]AcquireAutoClaimLockArg)(nil), args)
						return
					}
					ret, err = i.AcquireAutoClaimLock(ctx, (*typedArgs)[0].Caller)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"releaseAutoClaimLock": {
				MakeArg: func() interface{} {
					ret := make([]ReleaseAutoClaimLockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ReleaseAutoClaimLockArg)
					if !ok {
						err = rpc.NewTypeError((*[]ReleaseAutoClaimLockArg)(nil), args)
						return
					}
					err = i.ReleaseAutoClaimLock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"nextAutoClaim": {
				MakeArg: func() interface{} {
					ret := make([]NextAutoClaimArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NextAutoClaimArg)
					if !ok {
						err = rpc.NewTypeError((*[]NextAutoClaimArg)(nil), args)
						return
					}
					ret, err = i.NextAutoClaim(ctx, (*typedArgs)[0].Caller)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"isMasterKeyActive": {
				MakeArg: func() interface{} {
					ret := make([]IsMasterKeyActiveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]IsMasterKeyActiveArg)
					if !ok {
						err = rpc.NewTypeError((*[]IsMasterKeyActiveArg)(nil), args)
						return
					}
					ret, err = i.IsMasterKeyActive(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"ping": {
				MakeArg: func() interface{} {
					ret := make([]PingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.Ping(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RemoteClient struct {
	Cli rpc.GenericClient
}

func (c RemoteClient) Balances(ctx context.Context, __arg BalancesArg) (res []Balance, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.balances", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) Details(ctx context.Context, __arg DetailsArg) (res AccountDetails, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.details", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) RecentPayments(ctx context.Context, __arg RecentPaymentsArg) (res PaymentsPage, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.recentPayments", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) PaymentDetails(ctx context.Context, __arg PaymentDetailsArg) (res PaymentDetails, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.paymentDetails", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) AccountSeqno(ctx context.Context, __arg AccountSeqnoArg) (res string, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.accountSeqno", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) SubmitPayment(ctx context.Context, __arg SubmitPaymentArg) (res PaymentResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.submitPayment", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) SubmitRelayPayment(ctx context.Context, __arg SubmitRelayPaymentArg) (res PaymentResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.submitRelayPayment", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) SubmitRelayClaim(ctx context.Context, __arg SubmitRelayClaimArg) (res RelayClaimResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.submitRelayClaim", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) AcquireAutoClaimLock(ctx context.Context, caller keybase1.UserVersion) (res string, err error) {
	__arg := AcquireAutoClaimLockArg{Caller: caller}
	err = c.Cli.Call(ctx, "stellar.1.remote.acquireAutoClaimLock", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) ReleaseAutoClaimLock(ctx context.Context, __arg ReleaseAutoClaimLockArg) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.releaseAutoClaimLock", []interface{}{__arg}, nil)
	return
}

func (c RemoteClient) NextAutoClaim(ctx context.Context, caller keybase1.UserVersion) (res *AutoClaim, err error) {
	__arg := NextAutoClaimArg{Caller: caller}
	err = c.Cli.Call(ctx, "stellar.1.remote.nextAutoClaim", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) IsMasterKeyActive(ctx context.Context, __arg IsMasterKeyActiveArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.isMasterKeyActive", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) Ping(ctx context.Context) (res string, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.ping", []interface{}{PingArg{}}, &res)
	return
}
