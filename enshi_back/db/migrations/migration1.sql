-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "categories" table
CREATE TABLE "public"."categories" ("category_id" integer NOT NULL, "category_name" character varying(50) NOT NULL, PRIMARY KEY ("category_id"), CONSTRAINT "categories_category_name_key" UNIQUE ("category_name"));
-- Create "users" table
CREATE TABLE "public"."users" ("user_id" bigint NOT NULL, "username" character varying(50) NOT NULL, "email" character varying(100) NOT NULL, "password" character varying(255) NOT NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "is_admin" boolean NOT NULL, PRIMARY KEY ("user_id"), CONSTRAINT "users_email_key" UNIQUE ("email"), CONSTRAINT "users_username_key" UNIQUE ("username"));
-- Create "blogs" table
CREATE TABLE "public"."blogs" ("blog_id" bigint NOT NULL, "user_id" bigint NOT NULL, "title" character varying(255) NULL, "description" text NULL, "category_id" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("blog_id"), CONSTRAINT "blogs_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("category_id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "blogs_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "posts" table
CREATE TABLE "public"."posts" ("post_id" bigint NOT NULL, "blog_id" bigint NULL, "user_id" bigint NULL, "title" character varying(255) NULL, "content" text NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("post_id"), CONSTRAINT "posts_blog_id_fkey" FOREIGN KEY ("blog_id") REFERENCES "public"."blogs" ("blog_id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "posts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "bookmarks" table
CREATE TABLE "public"."bookmarks" ("user_id" bigint NOT NULL, "post_id" bigint NOT NULL, "bookmarked_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("user_id", "post_id"), CONSTRAINT "bookmarks_post_id_fkey" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("post_id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "bookmarks_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "comments" table
CREATE TABLE "public"."comments" ("comment_id" bigint NOT NULL, "post_id" bigint NULL, "user_id" bigint NULL, "content" text NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("comment_id"), CONSTRAINT "comments_post_id_fkey" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("post_id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "comments_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "favorites" table
CREATE TABLE "public"."favorites" ("user_id" bigint NOT NULL, "blog_id" bigint NOT NULL, "favorited_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("user_id", "blog_id"), CONSTRAINT "favorites_blog_id_fkey" FOREIGN KEY ("blog_id") REFERENCES "public"."blogs" ("blog_id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "favorites_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "likes" table
CREATE TABLE "public"."likes" ("like_id" bigint NOT NULL, "user_id" bigint NULL, "comment_id" bigint NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("like_id"), CONSTRAINT "likes_comment_id_fkey" FOREIGN KEY ("comment_id") REFERENCES "public"."comments" ("comment_id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "likes_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "tags" table
CREATE TABLE "public"."tags" ("tag_id" integer NOT NULL, "tag_name" character varying(50) NOT NULL, PRIMARY KEY ("tag_id"), CONSTRAINT "tags_tag_name_key" UNIQUE ("tag_name"));
-- Create "post_tags" table
CREATE TABLE "public"."post_tags" ("post_id" bigint NOT NULL, "tag_id" integer NOT NULL, PRIMARY KEY ("post_id", "tag_id"), CONSTRAINT "post_tags_post_id_fkey" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("post_id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "post_tags_tag_id_fkey" FOREIGN KEY ("tag_id") REFERENCES "public"."tags" ("tag_id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "post_votes" table
CREATE TABLE "public"."post_votes" ("post_id" bigint NOT NULL, "user_id" bigint NOT NULL, "vote" boolean NOT NULL, PRIMARY KEY ("post_id", "user_id"), CONSTRAINT "post_votes_post_id_fkey" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("post_id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "post_votes_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "profiles" table
CREATE TABLE "public"."profiles" ("profile_id" bigint NOT NULL, "user_id" bigint NULL, "bio" text NULL, "avatar_url" character varying(255) NULL, "website_url" character varying(100) NULL, PRIMARY KEY ("profile_id"), CONSTRAINT "profiles_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
