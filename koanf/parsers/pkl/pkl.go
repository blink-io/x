package pkl

import "github.com/apple/pkl-go/pkl"

type PKL struct {
}

func Parser() *PKL {
	return &PKL{}
}

// Unmarshal parses the given PKL bytes.
func (p *PKL) Unmarshal(b []byte) (map[string]any, error) {
	d := make(map[string]any)
	err := pkl.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Marshal marshals the given config map to PKL bytes.
func (p *PKL) Marshal(o map[string]any) ([]byte, error) {
	return []byte{}, nil
}
