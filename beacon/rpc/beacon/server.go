package beacon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/berachain/beacon-kit/runtime/service"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum Beacon Chain.
type Server struct {
	ContextGetter func(height int64, prove bool) (sdk.Context, error)
	Service       service.BeaconStorageBackend
}

func (s *Server) RegisterRoutes(router *mux.Router, contextGetter func(height int64, prove bool) (sdk.Context, error)) {
	router.HandleFunc("/eth/v1/beacon/states/{state_id}/randao", s.GetRandao).Methods(http.MethodGet)
}

// GetRandao fetches the RANDAO mix for the requested epoch from the state identified by state_id.
// If an epoch is not specified then the RANDAO mix for the state's current epoch will be returned.
// By adjusting the state_id parameter you can query for any historic value of the RANDAO mix.
// Ordinarily states from the same epoch will mutate the RANDAO mix for that epoch as blocks are applied.
func (s *Server) GetRandao(w http.ResponseWriter, r *http.Request) {
	stateId := mux.Vars(r)["state_id"]
	if stateId == "" {
		HandleError(w, "state_id is required in URL params", http.StatusBadRequest)
		return
	}

	stateIdAsInt, err := strconv.ParseUint(stateId, 10, 64)

	ctx, err := s.ContextGetter(int64(stateIdAsInt), false)
	if err != nil {
		HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	randao, err := s.Service.BeaconState(ctx).RandaoMixAtIndex(stateIdAsInt)
	if err != nil {
		HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &GetRandaoResponse{
		Data:                &Randao{Randao: hexutil.Encode(randao[:])},
		ExecutionOptimistic: true,
		Finalized:           true,
	}

	WriteJson(w, resp)
}

type GetRandaoResponse struct {
	ExecutionOptimistic bool    `json:"execution_optimistic"`
	Finalized           bool    `json:"finalized"`
	Data                *Randao `json:"data"`
}

type Randao struct {
	Randao string `json:"randao"`
}

const JsonMediaType = "application/json"

// WriteJson writes the response message in JSON format.
func WriteJson(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", JsonMediaType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		logrus.WithError(err).Error("Could not write response message")
	}
}

func HandleError(w http.ResponseWriter, message string, code int) {
	errJson := &DefaultJsonError{
		Message: message,
		Code:    code,
	}
	WriteError(w, errJson)
}

type HasStatusCode interface {
	StatusCode() int
}

// WriteError writes the error by manipulating headers and the body of the final response.
func WriteError(w http.ResponseWriter, errJson HasStatusCode) {
	j, err := json.Marshal(errJson)
	if err != nil {
		logrus.WithError(err).Error("Could not marshal error message")
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.Header().Set("Content-Type", JsonMediaType)
	w.WriteHeader(errJson.StatusCode())
	if _, err := io.Copy(w, io.NopCloser(bytes.NewReader(j))); err != nil {
		logrus.WithError(err).Error("Could not write error message")
	}
}

// DefaultJsonError is a JSON representation of a simple error value, containing only a message and an error code.
type DefaultJsonError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *DefaultJsonError) StatusCode() int {
	return e.Code
}

func (e *DefaultJsonError) Error() string {
	return fmt.Sprintf("HTTP request unsuccessful (%d: %s)", e.Code, e.Message)
}
