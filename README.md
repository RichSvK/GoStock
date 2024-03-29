## Indonesia Stock Holding Composition with Go

### Stock Market: Indonesia
### Data Source: KSEI
### Link: https://www.ksei.co.id/archive_download/holding_composition?setLocale=en-US

### System Requirements
- Software used in developing this program:
  - Go 1.22
  - MySQL 8.3.0
  - Text Editor: Visual Studio Code

### Program Description
This program allows users to export a CSV file containing the holder composition of scriptless shares for a chosen stock to a folder named "Output". Additionally, it provides functionality to insert data, sourced from KSEI in its original txt file format, into a MySQL database. The program stores the source txt files in a dedicated folder named "data".

### Program Preparation
1. Install all the required software such as Go and MySQL.
2. Get MySQL Driver using command go get github.com/go-sql-driver/mysql/.

### Database Setup
1. Create database using the commands in `StokDatabase.sql`.

### Program Flow
1. Make a connection to the database
2. Program will ask database name, username, and password with some validation
    * You can check the user and host of the database using "SELECT User, Host FROM mysql.user;" command in MySQL.
    * If the inputted database name, username, or password are incorrect, the connection is failed and the program will prompt the user to input them again.
      <br>
      ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/09bb6e6a-0931-4d06-b42f-1bc6142f6444)
      <br>
    * If the inputted database name, username, and password are correct, the connection will be established.
      <br>
      ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/9f6e418e-6eea-41be-a2ee-32c527a827ec)
      <br>  
3.  After successfully login, the program will display the main menu.
    <br>
    ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/ae6c237c-b05d-4c6c-af4b-dfeb387147aa)
    <br><br>

    - If user inputted invalid value, the program will prompt the user to input them again
      <br>
      ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/c77f6df9-fd60-491d-92b1-a50c46379378)
      <br>

    - If the input is 1, the user enters the insert menu and displays all files in the "data" folder:
      * If the user inputs an invalid or out of range value, the program will prompt user to input again.
      * If the input is valid but the data is already inserted in the database, the program will display an exception of duplicate entry and return to the main menu.
      For example, I already insert "Desember.txt" to my database. If I want to insert it "Desember.txt" again to my database, it will be rejected because of duplicate entry of primary key
        <br>
        ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/b03d5a67-178e-46be-8cc7-b9b45ad33337)
        <br>
      * If the input is valid and the data is yet to be inserted in the database, the program will insert the data from the txt file original format from KSEI to the database.
        <br>
        ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/6cc6dc18-4c04-4707-a8cf-18fb192826dc)
        <br>
      * Program will go back to main menu
   
    - If the input is "2", the user enters the export menu:
      * If the user inputs an invalid or unavailable stock , the program will prompt them to input again.
      * If the user inputs an unavailable stock , the program will return to main menu.
        <br>
        ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/a5485235-cc79-41af-8c13-ccafb7f3f8f9)
        <br>
      * If the input is a valid stock name available in the database, the program will export a "stockName.csv" file to the "Output" folder.
        <br>
        ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/7519b510-fb6a-49c5-9b1d-bf9ecd7eb85b)
        <br>
        ![image](https://github.com/RichSvK/Stock_Holder_Composition_Go/assets/87809864/12c6b1a0-53e6-4419-84bc-5353dc74b0d0)
        <br>
      * Program will go back to main menu
      <br>

    - If the input is "3", the program will be exited.
