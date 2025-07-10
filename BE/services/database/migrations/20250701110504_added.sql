-- Create "devices" table
CREATE TABLE "public"."devices" (
  "id" serial NOT NULL,
  "device_id" character varying(100) NOT NULL,
  "location" character varying(255) NULL,
  "installed_at" timestamptz NULL DEFAULT now(),
  "active" boolean NULL DEFAULT true,
  PRIMARY KEY ("id"),
  CONSTRAINT "devices_device_id_key" UNIQUE ("device_id")
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" serial NOT NULL,
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NULL,
  "phone" character varying(50) NULL,
  "created_at" timestamptz NULL DEFAULT now(),
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_email_key" UNIQUE ("email")
);
-- Create "cards" table
CREATE TABLE "public"."cards" (
  "id" serial NOT NULL,
  "card_id" character varying(100) NOT NULL,
  "user_id" integer NOT NULL,
  "active" boolean NULL DEFAULT true,
  "assigned_at" timestamptz NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "cards_card_id_key" UNIQUE ("card_id"),
  CONSTRAINT "cards_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_cards_card_id" to table: "cards"
CREATE INDEX "idx_cards_card_id" ON "public"."cards" ("card_id");
-- Create "balances" table
CREATE TABLE "public"."balances" (
  "id" serial NOT NULL,
  "user_id" integer NULL,
  "card_id" integer NOT NULL,
  "ride_cost" numeric(10,2) NOT NULL,
  "balance" numeric(10,2) NOT NULL DEFAULT 0,
  "updated_at" timestamptz NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "balances_card_id_key" UNIQUE ("card_id"),
  CONSTRAINT "unique_user_card" UNIQUE ("user_id", "card_id"),
  CONSTRAINT "balances_card_id_fkey" FOREIGN KEY ("card_id") REFERENCES "public"."cards" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "balances_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "charges" table
CREATE TABLE "public"."charges" (
  "id" serial NOT NULL,
  "user_id" integer NULL,
  "amount" numeric(10,2) NOT NULL,
  "type" character varying(20) NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "charges_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "charges_type_check" CHECK ((type)::text = ANY ((ARRAY['topup'::character varying, 'ride'::character varying])::text[]))
);
-- Create index "idx_charges_user_id" to table: "charges"
CREATE INDEX "idx_charges_user_id" ON "public"."charges" ("user_id");
-- Create "trips" table
CREATE TABLE "public"."trips" (
  "id" serial NOT NULL,
  "user_id" integer NULL,
  "card_id" integer NULL,
  "device_id" character varying(100) NOT NULL,
  "floor" integer NULL,
  "timestamp" timestamptz NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "trips_card_id_fkey" FOREIGN KEY ("card_id") REFERENCES "public"."cards" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "trips_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "idx_trips_user_id" to table: "trips"
CREATE INDEX "idx_trips_user_id" ON "public"."trips" ("user_id");

-- AUTH USERS TABLE FOR AUTHENTICATION
CREATE TABLE IF NOT EXISTS auth_users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'customer')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
