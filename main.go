package main

import (
	"context"
	"pgx-protobuf/db/queries"
	"pgx-protobuf/proto"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/pgx-proto")
	if err != nil {
		logrus.Fatalf("can't connect to db: %v", err)
	}

	q := queries.New(db)

	err = q.CreateTest(context.Background(), queries.CreateTestParams{
		Data: &proto.Test{
			Timestamp: timestamppb.Now(),
			Name:      "test",
		},
	})

	if err != nil {
		logrus.Fatalf("can't create test: %v", err)
	}

	tests, err := q.GetTests(context.Background())
	if err != nil {
		logrus.Fatalf("can't get tests: %v", err)
	}

	logrus.Infof("tests: %v", tests)

	defer db.Close()
}
