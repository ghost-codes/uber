CREATE TABLE "cabLocation" (
  "id" bigserial PRIMARY KEY,
  "driver" bigint NOT NULL,
  "cell_id" varchar NOT NULL,
  "position" point NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "cabLocation" ADD FOREIGN KEY ("driver") REFERENCES "driver" ("id");
