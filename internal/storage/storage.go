package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type storage struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

var schema = `
	CREATE TABLE IF NOT EXISTS users (
	    							 id serial primary key,
	    							 email text not null,
	                                 password_hash text not null,
	                                 password_salt text not null,
	                                 is_recruiter boolean not null default false,
	                                 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS vacancies (
	    							 id serial primary key,
	    							 recruiter_id integer not null references users(id),
	    							 experience text,
	    							 city text,
	    							 employment_type text,
	    							 salary_from integer,
	    							 salary_to integer,
	    							 company_name text,
	    							 archived boolean,
	    							 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS applications (
	    							 id serial primary key,
	    							 vacancy_id bigint not null references vacancies(id),
	    							 user_id bigint not null references users(id),
	    							 stage text,
	    							 archived boolean,
	    							 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS resume (
	    							 id serial primary key,
	    							 user_id bigint not null references users(id),
	    							 job_name text not null,
									 gender text not null,
									 address text, 
									 birth_date timestamptz,
									 phone_number text,
									 salary_from integer,
	    							 salary_to integer,
									 education text, 
									 skills []text, 
									 nationality text,
									 disability boolean,
									 archived boolean,
	    							 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS experience (
	    							 id serial primary key,
	    							 resume_id bigint not null references resume(id),
	    							 company_name text,
	    							 time_from timestamptz,
	    							 time_to timestamptz,
									 position text,
									 work_exp_description text,
	    							 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS notifications (
	    							 id serial primary key,
	    							 user_id bigint not null references users(id),
	    							 type text,
	    							 message text,
	    							 archived boolean,
	    							 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS meeting (
	    							 id serial primary key,
	    							 application_id bigint not null references applications(id),
	    							 time timestamptz,
	    							 name text,
	    							 description text,
	    							 archived boolean,
	    							 created_at timestamptz not null default now()
	);

	CREATE TABLE IF NOT EXISTS time_slots (
	    							 id serial primary key,
	    							 recruiter_id bigint not null references users(id),
	    							 time_from timestamptz,
	    							 time_to timestamptz,
	    							 occupied boolean,
	    							 archived boolean,
	    							 created_at timestamptz not null default now()
	);
`

// New creates a new instance of database layer using Postgres
func New(path string, logger *zap.Logger) *storage {
	// Wait until database initialize in container
	time.Sleep(time.Second * 2)

	config, err := pgxpool.ParseConfig(path)
	if err != nil {
		logger.Fatal("Unable to parse config", zap.Error(err))
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}

	return &storage{
		pool:   conn,
		logger: logger,
	}
}

// CreateSchema executes needed schema
func (s storage) CreateSchema() {
	_, err := s.pool.Exec(context.Background(), schema)
	if err != nil {
		s.logger.Fatal("Error occurred while executing schema", zap.Error(err))
	}

	s.logger.Info("Schema successfully created/updated")
}
