package mixin

import (
	"context"
	"encoding/json"
	"fmt"
)

// ModifyPIN modify pin
func (user User) ModifyPIN(ctx context.Context, oldPIN, pin string) (*User, error) {
	if pin == oldPIN {
		return nil, nil
	}

	pinEncrypted, err := user.signPIN(oldPIN)
	if err != nil {
		return nil, requestError(err)
	}

	paras := map[string]interface{}{
		"old_pin": pinEncrypted,
	}

	data, err := user.RequestWithPIN(ctx, "POST", "/pin/update", paras, pin)
	if err != nil {
		return nil, requestError(err)
	}

	var resp struct {
		User  *User  `json:"data"`
		Error *Error `json:"error"`
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, requestError(err)
	} else if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.User, nil
}

// VerifyPIN verify user pin
func (user User) VerifyPIN(ctx context.Context, pin string) (*User, error) {
	data, err := user.RequestWithPIN(ctx, "POST", "/pin/verify", nil, pin)

	var resp struct {
		User  *User  `json:"data"`
		Error *Error `json:"error"`
	}

	fmt.Printf("xxxxx\n")
	if err = json.Unmarshal(data, &resp); err != nil {
		fmt.Printf("xx%v", resp.User)
		return nil, requestError(err)
	} else if resp.Error != nil {
		fmt.Printf("xx%v", resp.Error)
		return nil, resp.Error
	}
	return resp.User, nil
}
