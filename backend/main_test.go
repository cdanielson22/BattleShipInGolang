package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getCurrentBoard(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getCurrentBoard(tt.args.context)
		})
	}
}

func Test_addPlayerShip(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addPlayerShip(tt.args.context)
		})
	}
}
