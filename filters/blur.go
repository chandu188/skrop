package filters

import (
	log "github.com/Sirupsen/logrus"
	"github.com/zalando/skipper/filters"
	"gopkg.in/h2non/bimg.v1"
)

const BlurName = "blur"

type blur struct {
	Sigma   float64
	MinAmpl float64
}

func NewBlur() filters.Spec {
	return &blur{}
}

func (r *blur) Name() string {
	return BlurName
}

func (r *blur) CreateOptions(image *bimg.Image) (*bimg.Options, error) {
	log.Debug("Create options for blurring ", r)

	blur := bimg.GaussianBlur{Sigma: r.Sigma, MinAmpl: r.MinAmpl}

	return &bimg.Options{GaussianBlur: blur}, nil
}

func (r *blur) CreateFilter(args []interface{}) (filters.Filter, error) {
	var err error

	if len(args) != 2 {
		return nil, filters.ErrInvalidFilterParameters
	}

	f := &blur{}

	f.Sigma, err = parseEskipFloatArg(args[0])
	if err != nil {
		return nil, err
	}

	f.MinAmpl, err = parseEskipFloatArg(args[1])
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *blur) Request(ctx filters.FilterContext) {}

func (r *blur) Response(ctx filters.FilterContext) {
	handleResponse(ctx, r)
}