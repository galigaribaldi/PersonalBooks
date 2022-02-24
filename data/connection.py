import psycopg2
import pandas as pd
class Connection():
    
    url_Connection = "postgresql://nnckmgzasylosx:99b170dfe60bb4771b0ac9edf19939df386bc88cf1222aa7fe69ce6446a4a254@ec2-44-196-68-164.compute-1.amazonaws.com:5432/ddaui1ic2hdida"
    
    def __init__(self):
        self.conn = psycopg2.connect(self.url_Connection)
        
    def Insert(self,tableName, nameRows, values):
        stringValues = "(%s"
        for i in values:
            stringValues = stringValues +",%s"
        query = "INSERT INTO "+tableName+"("+nameRows+")" + "VALUES"+stringValues+")"
        cur = self.conn.cursor()
        cur.execute(query, values)
        self.con.commit()
        self.con.close()
        
    def Select(self, tableName, nameRows=()):
        query = "SELECT * FROM "+ tableName
        if len(nameRows) == 0:
            res = pd.read_sql(query, con=self.conn)
        else:
            query = "SELECT "+ ",".join(nameRows) + " FROM " + tableName 
            res = pd.read_sql(query, con=self.conn)
        return res