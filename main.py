#from data.connection import Connection
from Data.connection import Connection
from Data.parser import Parser
#con = Connection()
#c = con.Select(tableName="books", nameRows=("title", "author"))
#print(c)
par = Parser()
p = par.ReadData('Books')