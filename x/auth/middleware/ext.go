package middleware

import (
	"context"

	abci "github.com/tendermint/tendermint/abci/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type HasExtensionOptionsTx interface {
	GetExtensionOptions() []*codectypes.Any
	GetNonCriticalExtensionOptions() []*codectypes.Any
}

type rejectExtensionOptionsMiddleware struct {
	next tx.Handler
}

// NewRejectExtensionOptionsMiddleware creates a new rejectExtensionOptionsMiddleware.
// rejectExtensionOptionsMiddleware is a middleware that rejects all extension
// options which can optionally be included in protobuf transactions. Users that
// need extension options should create a custom middleware chain that handles
// needed extension options properly and rejects unknown ones.
func RejectExtensionOptionsMiddleware(txh tx.Handler) tx.Handler {
	return rejectExtensionOptionsMiddleware{
		next: txh,
	}
}

var _ tx.Handler = rejectExtensionOptionsMiddleware{}

func checkEctOpts(tx sdk.Tx) error {
	if hasExtOptsTx, ok := tx.(HasExtensionOptionsTx); ok {
		if len(hasExtOptsTx.GetExtensionOptions()) != 0 {
			return sdkerrors.ErrUnknownExtensionOptions
		}
	}

	return nil
}

// CheckTx implements tx.Handler.CheckTx.
func (txh rejectExtensionOptionsMiddleware) CheckTx(ctx context.Context, tx sdk.Tx, req abci.RequestCheckTx) (abci.ResponseCheckTx, error) {
	if err := checkEctOpts(tx); err != nil {
		return abci.ResponseCheckTx{}, err
	}

	return txh.next.CheckTx(ctx, tx, req)
}

// DeliverTx implements tx.Handler.DeliverTx.
func (txh rejectExtensionOptionsMiddleware) DeliverTx(ctx context.Context, tx sdk.Tx, req abci.RequestDeliverTx) (abci.ResponseDeliverTx, error) {
	if err := checkEctOpts(tx); err != nil {
		return abci.ResponseDeliverTx{}, err
	}

	return txh.next.DeliverTx(ctx, tx, req)
}

// SimulateTx implements tx.Handler.SimulateTx method.
func (txh rejectExtensionOptionsMiddleware) SimulateTx(ctx context.Context, sdkTx sdk.Tx, req tx.RequestSimulateTx) (tx.ResponseSimulateTx, error) {
	if err := checkEctOpts(sdkTx); err != nil {
		return tx.ResponseSimulateTx{}, err
	}

	return txh.next.SimulateTx(ctx, sdkTx, req)
}