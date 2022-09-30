### CV Kabayan Golang Test
Another Titanic is about to sink. There are [n] numbers of people aboard the ship. Each person has a
property of name and age. The captain of the ship gives an order to evacuate and list the names of the
most people with the same age sorted alphabetically.


Input (.json) :
````
[
{"name": "John","age": 15},
{"name": "Peter","age": 12},
{"name": "Roger","age": 12},
{"name": "Anne","age": 44},
{"name": "Marry","age": 15},
{"name": "Nancy","age": 15}
]
````

Output :
````
 ['John', 'Mary', 'Nancy']
````
Explanation : Because the most occurring age in the data is 15 (occured 3 times), the output should be an array of strings with the three people's name, in this case it should be ['John', 'Mary', 'Nancy'].

Exception :
- In the case there are multiple "age" with the same maximum-occurence count, i need to split the data and print them into different arrays (i.e when 'Anne' age is 12, the result is: ['John', 'Mary', 'Nancy'], ['Anne','Peter','Roger']


Here's the summarize of the problem :
- we're trying to find the most frequent "age" from struct that comes from the decoded .json file (containing string "name" and integer "age").

- After that we need to print the "name" based on the maximum occurence frequency of "age".

- The printed "name" based on the maximum-occurence of "age" needs to be sorted alpabethically


### My Approach 
Break down the problem into sub-problem :
1. given backend-titanic-test.json as input  
  1.1  we need to read and parse the .json file into golang     
  1.2  store the parsed file into golang struct
2. to find the most frequent "age" we need to do several approach  
  2.1 group all name by it's age  
  2.2 compare the group size based on it's common age  
  2.3 if there is multiple group with the same size, we need to split the group individually
3. we need to sort the grouped name with its most frequent age alpabethically  
  3.1 for instance we can use the sort.Slice built in package
 
 ### To Run This Program
 ````
 git clone https://github.com/ariejanuarb/kabayan-backend-test
 ````
 Open the directory in the terminal, and then run the main.go
 ````
 go run main.go
 ````
 To check if the condition if there's any multiple age with same maximum frequency, you can try to load both json files by changing the code on line 19 :

````
content, err := os.ReadFile("backend-titanic-test.json")
content, err := os.ReadFile("backend-titanic-test-1.json")
````
### Result (Scenario 1 and 2)
#### scenario 1 : only 1 age number with maximum frequency
input : "backend-titanic-test.json"

output :
````
[ Alisa  Cash  Clay  Cline  Daphne  Hester  Irwin  Madden  Osborne  Sonja ]
````
#### scenario 2 : multiple age number with same maximum frequency
json input : "backend-titanic-test-1.json"

output :
````
[ Alisa  Cash  Clay  Cline  Daphne  Hester  Irwin  Madden  Osborne  Sonja ][ Dudley  Green  Holloway  Lottie  Munoz  Nieves  Richards  Tamara  Tracey  Yang ]
````
