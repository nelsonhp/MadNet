package objs

import (
	mdefs "github.com/MadBase/MadNet/consensus/objs/capn"
	"github.com/MadBase/MadNet/consensus/objs/validator"
	"github.com/MadBase/MadNet/errorz"
	gUtils "github.com/MadBase/MadNet/utils"
	capnp "zombiezen.com/go/capnproto2"
)

// Validator ...
type Validator struct {
	VAddr      []byte
	GroupShare []byte
}

// UnmarshalBinary takes a byte slice and returns the corresponding
// Validator object
func (b *Validator) UnmarshalBinary(data []byte) error {
	bh, err := validator.Unmarshal(data)
	if err != nil {
		return err
	}
	defer bh.Struct.Segment().Message().Reset(nil)
	return b.UnmarshalCapn(bh)
}

// UnmarshalCapn unmarshals the capnproto definition of the object
func (b *Validator) UnmarshalCapn(bh mdefs.Validator) error {
	err := validator.Validate(bh)
	if err != nil {
		return err
	}
	b.VAddr = gUtils.CopySlice(bh.VAddr())
	b.GroupShare = gUtils.CopySlice(bh.GroupShare())
	return nil
}

// MarshalBinary takes the Validator object and returns the canonical
// byte slice
func (b *Validator) MarshalBinary() ([]byte, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("not initialized")
	}
	bh, err := b.MarshalCapn(nil)
	if err != nil {
		return nil, err
	}
	defer bh.Struct.Segment().Message().Reset(nil)
	return validator.Marshal(bh)
}

// MarshalCapn marshals the object into its capnproto definition
func (b *Validator) MarshalCapn(seg *capnp.Segment) (mdefs.Validator, error) {
	if b == nil {
		return mdefs.Validator{}, errorz.ErrInvalid{}.New("not initialized")
	}
	var bh mdefs.Validator
	if seg == nil {
		_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		if err != nil {
			return bh, err
		}
		tmp, err := mdefs.NewRootValidator(seg)
		if err != nil {
			return bh, err
		}
		bh = tmp
	} else {
		tmp, err := mdefs.NewValidator(seg)
		if err != nil {
			return bh, err
		}
		bh = tmp
	}
	if err := bh.SetVAddr(b.VAddr); err != nil {
		return bh, err
	}
	if err := bh.SetGroupShare(b.GroupShare); err != nil {
		return bh, err
	}
	return bh, nil
}
