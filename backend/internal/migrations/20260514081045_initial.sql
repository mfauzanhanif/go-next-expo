-- Create "tenants" table
CREATE TABLE "tenants" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "schema_name" character varying NOT NULL,
  "domain" character varying NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "tenants_domain_key" to table: "tenants"
CREATE UNIQUE INDEX "tenants_domain_key" ON "tenants" ("domain");
-- Create index "tenants_schema_name_key" to table: "tenants"
CREATE UNIQUE INDEX "tenants_schema_name_key" ON "tenants" ("schema_name");
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL,
  "email" character varying NOT NULL,
  "phone_number" character varying NULL,
  "full_name" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_phone_number_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_number_key" ON "users" ("phone_number");
-- Create "tenant_users" table
CREATE TABLE "tenant_users" (
  "id" uuid NOT NULL,
  "role" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "tenant_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "tenant_users_tenants_tenant_users" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "tenant_users_users_tenant_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
