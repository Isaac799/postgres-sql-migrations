# You will need to create env.bash with 
# proper info and run it at least once each time you visit this terminal session.
# Make sure it's part of the .gitignore.

# Note: To load the environment variables into your current shell session,
# use the following command:
# . env.bash
#
# Remember to run this command every time you start a new terminal session
# to ensure the environment variables are set.

export dev_DBNAME="my_site_dev"
export dev_DBUSER="postgres"
export dev_DBPASS="postgres"
export dev_DBHOST="localhost"
export dev_DBPORT="5432"
export dev_DBSSL="disable"

export test_DBNAME="my_site_test"
export test_DBUSER="postgres"
export test_DBPASS="postgres"
export test_DBHOST="localhost"
export test_DBPORT="5432"
export test_DBSSL="disable"

export prod_DBNAME="my_site_prod"
export prod_DBUSER="postgres"
export prod_DBPASS="postgres"
export prod_DBHOST="localhost"
export prod_DBPORT="5432"
export prod_DBSSL="require"
