// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package dispatch

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/core/option"
	"github.com/mongodb/mongo-go-driver/core/readconcern"
	"github.com/mongodb/mongo-go-driver/core/writeconcern"
)

// ErrUnacknowledgedWrite is returned from functions that have an unacknowledged
// write concern.
var ErrUnacknowledgedWrite = errors.New("unacknowledged write")

func writeConcernOption(wc *writeconcern.WriteConcern) (option.OptWriteConcern, error) {
	elem, err := wc.MarshalBSONElement()
	if err != nil {
		return option.OptWriteConcern{}, err
	}
	return option.OptWriteConcern{WriteConcern: elem, Acknowledged: wc.Acknowledged()}, nil
}

func readConcernOption(rc *readconcern.ReadConcern) (option.OptReadConcern, error) {
	elem, err := rc.MarshalBSONElement()
	if err != nil {
		return option.OptReadConcern{}, err
	}
	return option.OptReadConcern{ReadConcern: elem}, nil
}
