/*
 Navicat Premium Data Transfer

 Source Server         : postgres_pro
 Source Server Type    : PostgreSQL
 Source Server Version : 110005
 Source Host           : 127.0.0.1:5432
 Source Catalog        : feitago
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110005
 File Encoding         : 65001

 Date: 30/03/2020 14:31:46
*/


-- ----------------------------
-- Table structure for large_admin_casbin
-- ----------------------------
DROP TABLE IF EXISTS "public"."large_admin_casbin";
CREATE TABLE "public"."large_admin_casbin" (
  "p_type" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "v0" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "v1" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "v2" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "v3" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "v4" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "v5" varchar(256) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;

-- ----------------------------
-- Indexes structure for table large_admin_casbin
-- ----------------------------
CREATE INDEX "idx_large_admin_casbin_p_type" ON "public"."large_admin_casbin" USING btree (
  "p_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_large_admin_casbin_v0" ON "public"."large_admin_casbin" USING btree (
  "v0" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_large_admin_casbin_v1" ON "public"."large_admin_casbin" USING btree (
  "v1" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_large_admin_casbin_v2" ON "public"."large_admin_casbin" USING btree (
  "v2" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_large_admin_casbin_v3" ON "public"."large_admin_casbin" USING btree (
  "v3" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_large_admin_casbin_v4" ON "public"."large_admin_casbin" USING btree (
  "v4" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_large_admin_casbin_v5" ON "public"."large_admin_casbin" USING btree (
  "v5" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
