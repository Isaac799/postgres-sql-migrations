# PostgreSQL Migrations

A utility for managing PostgreSQL database changes via SQL file migrations. Inspired by Ecto, this tool offers a simple command-line interface for creating and applying migrations. 

## Getting Set Up

### 1. Go to the Migrations Folder  
Make sure you’re in the right spot:
```bash
cd /cmd/migration
```

### 2. Install the PostgreSQL Library  
Run this to grab the necessary PostgreSQL library:
```bash
go get
```

### 3. Environment Variables

Create an `env.bash` file with your database credentials. It is part of the `.gitignore` by default. Check out `example.env.bash` for the expected format.

#### Loading Environment Variables

To load your environment variables into the current shell session, run:

```bash
. env.bash
```

## Key Concepts

1. **Migrations Folder**  
   Check out the `migrations` (same level as executable) directory where your SQL files live. Each file represents a migration, or incremental change to the database. Such as creating a table or altering a column.

2. **Migrations Table**  
   A `migrations` table is stored on the connected database, and keeps a record of which migrations have been applied. This helps you track changes across different environments.

3. **Undoing Changes**  
   Avoid deleting migrations, as doing so ruins your ability to sync up environments. This helps you *fix mistakes instead of replacing them*, especially when you can’t just wipe everything (e.g. prod).

4. **Migration Content**  
   This application will handle migrations and create the migration files for you. However, it’s your responsibility to fill them out with the appropriate SQL code.

### Naming Your Migration Files  
When naming your migration files, use a verb-noun format to clarify their purpose. For example, names like `create_user_table` or `add_email_to_user` make it easy to understand each migration's intent at a glance.

**Tip**: All migrations run inside a transaction. If one fails, none of the changes will be applied, preventing a half-finished update. That said, creating and dropping the databases are not protected by this.

## Usage

To get started with the tool, just run:

```bash
go run . [option]
```

For now, chaining flags is not supported. This is part of why I created `-full-reset`

| Option                | Shorthand | Description                                                                                                     |
| --------------------- | --------- | --------------------------------------------------------------------------------------------------------------- |
| `-help`               | `-h`      | Show help information.                                                                                          |
| `-generate-migration` | `-gm`     | Generate a new migration file with the given name. [Naming Your Migration Files](#naming-your-migration-files). |
| `-env`                | `-e`      | Pick the environment (`dev` or `prod`, should align with `env.bash`).                                           |
| `-list-migrations`    | `-lm`     | List migrations currently in the database.                                                                      |
| `-summary`            | `-s`      | Provides a summary of applied and local migrations.                                                             |
| `-create`             | `-c`      | Creates the database (initial setup).                                                                           |
| `-drop`               | `-d`      | Delete the database (wipes all data). Confirmation prompt given.                                                |
| `-migrate`            | `-m`      | Apply migrations.                                                                                               |
| `-migrate-dry`        | `-m-dry`  | Dry run migrations without applying changes. Useful for previewing migrations.                                  |
| `-full-reset`         | `-fr`     | Drops, creates, migrates, and gives summary.                                                                    |

##  Migration Commands: FAQs and Solutions

### Generate a Migration
**Q: How do I create a new `user` table?**  
A: Use the command below to generate a migration file for defining the new table.  
```bash
go run . -gm create_user_table
```
*This will create a file named `<timestamp>_create_user_table.sql`. You will need to fill out this file with the necessary SQL create table statement.*

---

### List Current Migrations in Development
**Q: How can I check the migrations applied in development?**  
A: Run the following command to see all migrations that have been applied in the development environment.  
```bash
go run . -e dev -lm
```
*This helps you identify when a migration was made to a database.*

---

### Create the Database for Testing
**Q: What command do I use to set up a fresh database for testing?**  
A: Execute this command to create a new testing database.  
```bash
go run . -e test -create
```
*Now you have a clean database for your testing needs.*

---

### Run Migrations in Development
**Q: How do I apply pending migrations in development?**  
A: Use the command below after modifying your migration files.  
```bash
go run . -e dev -migrate
```
*This will apply any outstanding migrations in your development environment.*

---

### Rollback a Migration
**Q: How can I rollback the last migration that created a `user` table?**  
A: Use the command below to generate a migration file for dropping the table.  
```bash
go run . -gm drop_user_table
```
*This creates a file named `drop_user_table.sql`. You will need to fill out this file with the necessary SQL drop table statement.*

---

### Drop the Database in Development
**Q: What if I need to drop my development database?**  
A: To drop all data and reset the database, run the following command.  
```bash
go run . -e dev -drop
```
*This will remove everything in your development database.*

---


### Full Reset the Development Database
**Q: How do I perform a full reset of my development database?**  
A: Use this command to drop the existing database, create a new one, and apply all migrations.  
```bash
go run . -e dev -full-reset
```
*This is ideal for establishing a clean slate for quicker development cycles.*

---

### Summarize Applied and Local Migrations
**Q: How can I summarize the applied migrations on the server versus the local migrations on my computer?**  
A: Run the following command to get a quick overview of your migration status.  
```bash
go run . -e dev -summary
```
*This will show you how many migrations have been applied on the server (those currently in use) and how many are waiting to be run locally on your computer. This helps you identify if your local environment is up to date with the server. Note: it does not compare the content of those migrations*

---

### Dry Run Migrations
**Q: Is there a way to preview migration changes without applying them?**  
A: Yes, you can preview changes with this command.  
```bash
go run . -e dev -migrate-dry
```
*This allows you to see the changes that would be made without actually altering the database.*
