package entities

import "github.com/tanyudii/balance-api/internal/pkg/errutil"

type AccountDaftarRequest struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

func (r *AccountDaftarRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.Nama == "" {
		fields["nama"] = "nama tidak boleh kosong"
	}
	if r.Nik == "" {
		fields["nik"] = "nik tidak boleh kosong"
	}
	if r.NoHp == "" {
		fields["no_hp"] = "no hp tidak boleh kosong"
	}
	return errutil.BadRequestOrNil(fields)
}

type AccountMutationRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

func (r *AccountMutationRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.NoRekening == "" {
		fields["no_rekening"] = "no rekening tidak boleh kosong"
	}
	if r.Nominal <= 0 {
		fields["nominal"] = "nominal harus lebih dari 0"
	}
	return errutil.BadRequestOrNil(fields)
}
