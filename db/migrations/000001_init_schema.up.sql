CREATE TABLE "userMetaData" (
  "id" varchar PRIMARY KEY,
  "phone_number" varchar UNIQUE NOT NULL,
  "date_of_birth" timestamptz NOT NULL,
  "created_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rideHistory" (
  "id" bigserial PRIMARY KEY,
  "source" point NOT NULL,
  "destination" point NOT NULL,
  "user" varchar NOT NULL,
  "payment_id" bigint,
  "driver" bigint NOT NULL,
  "board_time" timestamptz NOT NULL,
  "arrival_time" timestamptz,
  "status" varchar NOT NULL
);

CREATE TABLE "paymentHistory" (
  "id" bigserial PRIMARY KEY,
  "amount_cents" int NOT NULL,
  "payment_method" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pendingRide" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "driver_id" bigint NOT NULL,
  "source" point NOT NULL,
  "destination" json NOT NULL,
  "created_at" timestamptz NOT NULL,
  "estimated_price_cents" int NOT NULL
);

CREATE TABLE "driver" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "hashed_password" varchar NOT NULL,
  "contact" varchar NOT NULL,
  "car_number" varchar NOT NULL,
  "car_brand" varchar NOT NULL,
  "car_color" varchar NOT NULL,
  "profile_picture" varchar
);

ALTER TABLE "rideHistory" ADD FOREIGN KEY ("user") REFERENCES "userMetaData" ("id");

ALTER TABLE "rideHistory" ADD FOREIGN KEY ("payment_id") REFERENCES "paymentHistory" ("id");

ALTER TABLE "rideHistory" ADD FOREIGN KEY ("driver") REFERENCES "driver" ("id");

ALTER TABLE "pendingRide" ADD FOREIGN KEY ("user_id") REFERENCES "userMetaData" ("id");

ALTER TABLE "pendingRide" ADD FOREIGN KEY ("driver_id") REFERENCES "driver" ("id");

