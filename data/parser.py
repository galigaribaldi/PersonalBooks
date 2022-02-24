import pandas as pd

class Parser():
    def ReadData(self, sheetName):
        data_all = pd.ExcelFile("Data/files/Initial_data.xlsx")
        data = data_all.parse(sheetName)
        print(data)