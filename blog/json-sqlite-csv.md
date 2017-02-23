JSON to SQLite to csv with Python
Basic data transformations using python
2016-11-20
# JSON to SQLite to csv with python.

So if you read my last post about webscraping and saving a whole bunch of JSON files, then there's a small chance that you might have just that - a whole bunch of similar JSON files. In this post I'll run you through my process to take JSON data, saved across multiple files and transform that data into SQLite tables. In my case it's a whole bunch of football stats saved in JSON format that I want to enter into a highly normalised SQL database. I'll then run through how to get the SQLite data into csv format. We'll use the command line SQLite tool for this. 

## JSON to objects

Absolutely the first thing you want to do is to try to describe the domain, to make a domain model. And then to think about how these models fit together. Think of it as Domain Driven Design (DDD) light. This isn't so much about decoupling business logic from service layers as would the case in proper DDD. At the end of the day this is scripting, we're not building software here.

The focus will be on describing the domain and using composition to handle the creation of objects. I want to make a highly normalised database. For the kind of thing where you are going to be writing very few queries, but where the queries are complex, it makes sense to normalise our data. That means that goals will be stored separately from matches and to get the goals from any given match, will require a table join. It also means to get all the goals for any given player, we won't need to join through a Match table or anything like that. So for example, where many football databases save information about Red Cards, or Yellow Cards right in the match table, I like to have those in their own tables with a Foreign Key to the match table and to the player table and to the team table. 

For each python class I recommend a minimum of three methods

    class myClass:
        def __init__():
            #Initialise the object, obvious choice.
        def __str__():
            #Handle printing behaviour. Really useful to make sure the objects have been created properly. 
        def Insert():
            #Handle inserting the object data into the SQLite database.

The Insert() method is really going to clean up the code in your script, especially when you have a list of objects

    for myObject in myObjects:
        myObject.Insert()

Simple. So basically, I recommend breaking the objects down into as many separate things as is practical. We don't want to store xml strings or anything in our databases because we're trying to cram too much information into our tables. In my database I settled for the following tables, and thus objects.

* Season - Contains the name of the competition and the year. 
* Referee
* Team
* Player
* Stadium - Probably should have gone in the team table, seeing as I'm only dealing with league data.
* Game
* Appearance - A relationship table for players and games, due to the many to many relationship.
* Goal
* Assist - Could be stored in the goal table.
* Yellow Card
* Red Card - Yellow and red card could be stored in the same table.

A final note - where you need an Autoincrementing primary key for your table, you don't want to store that in your python object, nor do you want to store Foreign Key references in the object. Instead pass those as parameters to the Insert method.

## Creating the sqlite database and tables.

At the top of your script you want to connect to your SQLite database, if the database doesn't exist, SQLite will make it for you. Man I love that feature. Then you want to create the tables, like so:

    import sqlite3 as lite

    con = lite.connect('test.db')
    cur = con.cursor()

    cur.execute("CREATE TABLE IF NOT EXISTS Season (id INTEGER PRIMARY KEY AUTOINCREMENT, Competition TEXT, Year INTEGER)")

    cur.execute("CREATE TABLE IF NOT EXISTS Referee (Name TEXT PRIMARY KEY)")

    cur.execute("CREATE TABLE IF NOT EXISTS Team (Name Text PRIMARY KEY)")

    cur.execute("CREATE TABLE IF NOT EXISTS Player (Name TEXT PRIMARY KEY, DOB INTEGER, Nation TEXT, Goalkeeper INTEGER)")

    cur.execute("CREATE TABLE IF NOT EXISTS Stadium (Name TEXT PRIMARY KEY, TeamName TEXT, FOREIGN KEY(TeamName) REFERENCES Team(Name))")
    # and so on and so on...

Notice that I don't use any kind of ORM. Personally, I like SQL, and I prefer to write it directly. I think it's good skill to have in data science, you won't always be able to use an ORM, but you can always query with SQL.

Then create your classes. For example:

	class Competition:
		def __init__(self, league, year):
			self.league = league
			self.year = year

		def __str__(self):
			return "League: {}, Year:{}".format(self.league, self.year)

		def insert(self, con, cur):
			cur.execute("INSERT OR IGNORE INTO Season (Competition, Year) VALUES (?, ?)", (self.league, self.year))
			con.commit()


	class Referee:
		def __init__(self, name):
			self.name = name

		def __str__(self):
			return "Referee: {}".format(self.name)

		def insert(self, con, cur):
			cur.execute("INSERT OR IGNORE INTO Referee (Name) VALUES (?)", (self.name,))
			con.commit()


	class Team:
		def __init__(self, name):
			self.name = name

		def __str__(self):
			return "Team: {}".format(self.name)

		def insert(self, con, cur):
			cur.execute("INSERT OR IGNORE INTO Team (Name) VALUES (?)", (self.name,))
			con.commit()


	class Player:
		def __init__(self, name, DOB, nationality, goalkeeper):
			self.name = name
			self.DOB = DOB
			self.nationality = nationality
			#Goalkeeper is a boolean value
			self.goalkeeper = goalkeeper

		def __str__(self):
			return "PLAYER - Name: {}, DOB: {}, Nationality: {}, Is Goalkeeper: {}".format(self.name, self.DOB, self.nationality, self.goalkeeper)

		def insert(self, con, cur):
			cur.execute("INSERT OR IGNORE INTO Player (Name, DOB, Nation, Goalkeeper) VALUES (?, ?, ?, ?)", (self.name, self.DOB, self.nationality, self.goalkeeper))
			con.commit()

The init and str methods are pretty standard fare. You'll notice that on the insert methods, we pass the cursor and the connection to the method. It's just a good way of being concise. It also make sure we don't initialise a database connection every time we insert data. We do however commit after every insertion. I like to use judicious printing and regular committing, so I can pick up where I left off if something goes wrong.

Other things to note, I use player names, team names and other such names as primary keys. You see, under the hood SQLite3 doesn't actually make primary keys, it just makes unique indices. I know there were no two players called exactly the same thing so that also simplifies things. It also means we don't have to insert and then retrieve the id when doing insertions. And we can do a simple "INSERT OR IGNORE" to prevent duplicates.

## Looping through

So now that you've created your tables, made your classes and given each class an insert method, you just need to loop through the JSON files and make all the insertions. So it starts like this

	for index in range(380): #For 380 games in a premier league season.
		file = 'game{}.json'.format(index)
		with open(file, 'r') as data_file:
			data = json.load(data_file)

Then the insertion might look something like this. Remember, JSON objects are mapped to dicts so it works just like accessing nested elements in a dict.

	players = []
	for teamList in data["teamLists"]:
		for player in teamList["lineup"]:
			if player["matchPosition"] == "G":
				players.append(Player(player["name"]["display"], player["birth"]["date"]["millis"], player["birth"]["country"]["country"], 1))
			else:
				players.append(Player(player["name"]["display"], player["birth"]["date"]["millis"], player["birth"]["country"]["country"], 0))

	for player in players:
		player.insert(con, cur)

Apologies if the code wrapping makes that tricky to read. So we make a whole bunch of player objects and put them in a list called players, then we loop through the players list and insert the data into the database. Simples.

## There you go

So this wasn't exactly a walkthrough, but I hope you understand a possible process to scrape JSON data from a website, map it to a domain using basic Domain Driven Design, and creating a normalised SQLite database from the data. The only thing left to do is to query the data and save the results as a csv file that you can then use in a D3.js animation. 

## Querying the sqlite database
So I recommend creating an SQL file. Call it something clever like my_query.sql. Then in the terminal, you want to test your query on stdout first. so you would type

    sqlite3 mydatabase.db
    >>> .mode column
    >>> .headers ON
    >>> .read my_query.sql

If you don't like the results, tweak your query and run it again.

    >>> .read my_query.sql

Now your happy with it. For example this to get a list of all players and how many goals they scored. 

	SELECT player.name AS players, 
	       COUNT(goal.id) AS goals 
	  FROM player 
	       LEFT JOIN goal 
	       ON player.name = goal.player
	 GROUP BY players
	 ORDER BY goals DESC;

Then in the sqlite command you would do this.

    >>> .mode csv
    >>> .output my_results.csv
    >>> .read my_query.sql
    >>> .exit
    open my_results.csv

And there you have it, the easiest way to get csv data from SQlite. No python imports, no worrying about UTF-8, no csv writers etc. Just simple. Now, [creating a graph from the csv data.](/D3-CDF-with-hover)