package translater

import (
	"io"

	"github.com/htamakos/contran/types"
)

type Translater interface {
	Input(values []byte) (*types.CommonSpecs, error)
	Output(specs *types.CommonSpecs) ([]byte, error)
}

type TranslateService struct {
	input  Translater
	output Translater
	out    io.Writer
}

func NewTranslateService(input, output Translater, out io.Writer) *TranslateService {
	return &TranslateService{
		input:  input,
		output: output,
		out:    out,
	}
}

func (s *TranslateService) Translate(values []byte) error {
	specs, err := s.input.Input(values)
	if err != nil {
		return err
	}

	results, err := s.output.Output(specs)
	if err != nil {
		return err
	}

	_, err = s.out.Write(results)
	if err != nil {
		return err
	}

	return nil
}
