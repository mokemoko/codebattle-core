generate:
	sqlboiler -o ../models sqlite3
reset:
	sqlite3 db.sqlite3 < schema.sql
	sqlite3 db.sqlite3 < seeds.sql
