// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package command

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/description"
	"github.com/mongodb/mongo-go-driver/core/option"
	"github.com/mongodb/mongo-go-driver/core/wiremessage"
)

// DropIndexes represents the dropIndexes command.
//
// The dropIndexes command drops indexes for a namespace.
type DropIndexes struct {
	NS     Namespace
	Index  string
	Opts   []option.DropIndexesOptioner
	result bson.Reader
	err    error
}

// Encode will encode this command into a wire message for the given server description.
func (di *DropIndexes) Encode(desc description.SelectedServer) (wiremessage.WireMessage, error) {
	cmd := bson.NewDocument(
		bson.EC.String("dropIndexes", di.NS.Collection),
		bson.EC.String("index", di.Index),
	)

	for _, opt := range di.Opts {
		if opt == nil {
			continue
		}
		err := opt.Option(cmd)
		if err != nil {
			return nil, err
		}
	}

	return (&Command{DB: di.NS.DB, Command: cmd, isWrite: true}).Encode(desc)
}

// Decode will decode the wire message using the provided server description. Errors during decoding
// are deferred until either the Result or Err methods are called.
func (di *DropIndexes) Decode(desc description.SelectedServer, wm wiremessage.WireMessage) *DropIndexes {
	di.result, di.err = (&Command{}).Decode(desc, wm).Result()
	return di
}

// Result returns the result of a decoded wire message and server description.
func (di *DropIndexes) Result() (bson.Reader, error) {
	if di.err != nil {
		return nil, di.err
	}
	return di.result, nil
}

// Err returns the error set on this command.
func (di *DropIndexes) Err() error { return di.err }

// RoundTrip handles the execution of this command using the provided wiremessage.ReadWriter.
func (di *DropIndexes) RoundTrip(ctx context.Context, desc description.SelectedServer, rw wiremessage.ReadWriter) (bson.Reader, error) {
	wm, err := di.Encode(desc)
	if err != nil {
		return nil, err
	}

	err = rw.WriteWireMessage(ctx, wm)
	if err != nil {
		return nil, err
	}
	wm, err = rw.ReadWireMessage(ctx)
	if err != nil {
		return nil, err
	}
	return di.Decode(desc, wm).Result()
}
