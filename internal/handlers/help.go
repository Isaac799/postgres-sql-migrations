package handlers

import (
	"fmt"
	"log"
	"strings"
)

func ShowHelp() {
	var helpMessage strings.Builder

	helpMessage.WriteString("\n")
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "Option", "Short", "Description"))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-------------------", "--------", "-------------------------------------------------"))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-help", "-h", "Show help information."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-generate-migration", "-gm", "Generate a new migration file with the given name."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-env", "-e", "Pick the environment (should align with `env.bash`)."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-list-migrations", "-lm", "List migrations currently in the database."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-summary", "-s", "Provides a summary of applied and local migrations."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-create", "-c", "Creates the database (initial setup)."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-drop", "-d", "Delete the database (wipes all data)."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-migrate", "-m", "Apply migrations."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-full-reset", "-fr", "Drops, creates, migrates, and gives summary."))
	helpMessage.WriteString(fmt.Sprintf("%-20s %-8s %-50s\n", "-migrate-dry", "-m-dry", "Preview migration changes."))
	helpMessage.WriteString("\nNote: Flags cannot be chained together\n")
	helpMessage.WriteString("Note: Ensure the 'migrations' folder is at the same level as the executable.\n")

	log.Println(helpMessage.String())
}
