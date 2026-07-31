package proteintranslation

import "errors"

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid base")
)

var aminos = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var res []string
	for i := 0; i+3 <= len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			return res, nil
		}
		if err != nil {
			return nil, err
		}
		res = append(res, protein)
	}
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := aminos[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}
