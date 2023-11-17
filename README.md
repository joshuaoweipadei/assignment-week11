# assignment-week11
Student Name:
AARTI -500224115 
SUKHNEET KAUR -500224802 
YASH KELKAR -500223746 
OWEIPADEI JOSHUA BAYEFA -500221880 
DENNIS SAMUEL ANUMBA -500222264  
KARAN -500224005
 SUMIT -500224003 
HARDIK KUMAR -500223366 


Go Application steps and functionality
Step1:
•	We created a github repository() and sent an invite to all group members
Step2:
•	We added the project Git repository URl as a remote to our local repository
•	We created a readme.md file,main.go file, main_test.go file, and docker file and committed the files to the project github repository
•	We created a function that calls an external web API to retrieve the current weather in Toronto
•	This process required a weather API key to authenticate the request to the openweathermap API so we had to created an account in openweathermap.org and the api key was sent to the confirmation email.

Code Explanation:
•	we imported the necessary packages needed to work with JSON, making http request, and interacting with operating system 
•	We declared constants for openweather API Key and the baseURL for making API request. This was done for various reason such as  to minimize errors when updating or modifying the API key and Base URL at any point in time, reading and understanding the code easily, using the constants accross all functions in the program, and avoid making request to external service
•	created a weatherdata struct that represent the structure of the JSON response after making the OpenWeatherMap API request. Upon Request, the response will contain statistics such as  temperature, humidity, wind speed , pressure, and description of the current weather weather.
•	The above weatherObj is defined as an object to hold the description and temprrature of the Toronto weather.
•	The function Toronto weather is a function that calls the external api (open weather api) "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)”, to get the current weather of the Toronto location.we then checked if there is an error from the api else we continue. This function returns the weather description and the weather temperature which we de structure from the JSON objects gotten from the external api.
•	Toronto time handler function returns the temperature and description which is then called in a main fuction. 
•	The url to test t his code is localhost:5000/api/toronto-weather.

Test Function:
•	Test Function Structure:
Test functions in Go start with Test followed by the name of the function being tested.
Takes a single argument of type *testing.T.
•	Mocking the API Response:
Use a string (mockResponse) to simulate the JSON response from the OpenWeatherMap API.
•	Creating a Test Server:
httptest.NewServer creates a test HTTP server.
http.HandlerFunc allows you to define a function to handle incoming HTTP requests.
defer ts.Close() ensures that the test server is closed when the test function exits.
•	Overriding the Weather URL:
ts.URL provides the base URL of the test server.
Override the weatherURL variable in the main code with the test server URL.
•	Testing the Function:
Call the function to be tested (torontoWeather in this case).
Compare the result with the expected values using t.Errorf to report failures.
•	Test Execution:
Run the tests using the go test command in the terminal.

Docker Image
•	Link to the Docker image - docker pull yash247/nov17:v0.1