Prerequisites:-
Install go1.19:
To avoid any surprises in the project install go 1.19, please use the commands below to install go 1.19 in your machine:
go install golang.org/dl/go1.19@latest
go1.19 download

Clone the project and install the dependencies:
Clone the project from Github and then run the following command to install dependencies.
go get

Install Mockgen:
mockgen is a mock generator tool which we are going to use to generate our mocks. To install run the following command:
go install github.com/golang/mock/mockgen@v1.6.0

Project Outline:-
I've generated the occupancy combinations of adult and child for a given occupancy in golang. 
I've taken BaseAdult, MaxAdult, BaseChild, MaxChild, MaxGuest and the ChildAgeRange as request and generating the occupancy combinations as response.
I've also performed testing using mockgen on the GenerateOccupancy, GenerateCombination function and benchmarking on GenerateOccupancy function.
