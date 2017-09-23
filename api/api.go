package api

import "github.com/st3v/coup/fleet"

type Request struct {
	Scooters []int `json:"scooters"`
	C        int   `json:"C"`
	P        int   `json:"P"`
}

type Response struct {
	NumFE int `json:"fleet_engineers"`
}

func (req *Request) Handle() (Response, error) {
	n, err := fleet.MinNumFE(req.Scooters, req.C, req.P)
	if err != nil {
		return Response{}, err
	}

	return Response{n}, nil
}
