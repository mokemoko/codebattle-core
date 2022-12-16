package models

import (
	"context"
	"database/sql"
	"reflect"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func SetupDatabase(isDebug bool) error {
	// TODO: be variable
	db, err := sql.Open("sqlite3", "../sql/db.sqlite3")
	if err != nil {
		return err
	}
	if isDebug {
		boil.DebugMode = true
	}
	boil.SetDB(db)

	// MEMO: sqliteでtimestampカラムが自動更新されないため、hookで対応
	addHook()

	return nil
}

func addHook() {
	AddEntryHook(boil.BeforeUpdateHook, func(ctx context.Context, executor boil.ContextExecutor, entity *Entry) error {
		return updateHook(entity)
	})
	AddEntryHook(boil.BeforeInsertHook, func(ctx context.Context, executor boil.ContextExecutor, entity *Entry) error {
		return insertHook(entity)
	})
	AddContestHook(boil.BeforeUpdateHook, func(ctx context.Context, executor boil.ContextExecutor, entity *Contest) error {
		return updateHook(entity)
	})
	AddContestHook(boil.BeforeInsertHook, func(ctx context.Context, executor boil.ContextExecutor, entity *Contest) error {
		return insertHook(entity)
	})
	AddMatchHook(boil.BeforeUpdateHook, func(ctx context.Context, executor boil.ContextExecutor, entity *Match) error {
		return updateHook(entity)
	})
	AddMatchHook(boil.BeforeInsertHook, func(ctx context.Context, executor boil.ContextExecutor, entity *Match) error {
		return insertHook(entity)
	})
	AddUserHook(boil.BeforeUpdateHook, func(ctx context.Context, executor boil.ContextExecutor, entity *User) error {
		return updateHook(entity)
	})
	AddUserHook(boil.BeforeInsertHook, func(ctx context.Context, executor boil.ContextExecutor, entity *User) error {
		return insertHook(entity)
	})
}

func updateHook(entity interface{}) error {
	return hookTimestamp(entity, []string{"UpdatedAt"})
}

func insertHook(entity interface{}) error {
	return hookTimestamp(entity, []string{"CreatedAt", "UpdatedAt"})
}

func hookTimestamp(entity interface{}, fieldNames []string) error {
	ts := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	rv := reflect.ValueOf(entity).Elem()
	for _, fieldName := range fieldNames {
		field := rv.FieldByName(fieldName)
		field.SetString(ts)
	}
	return nil
}
