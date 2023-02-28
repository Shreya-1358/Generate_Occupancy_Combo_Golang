# Generating Room level occupancy combinations for a hotel for given occupancy using Golang.

## Steps to follow:-

### Prerequisites:-

#### Install go1.19:
To avoid any surprises in the project install go 1.19, please use the commands below to install go 1.19 in your machine:

```javascript
$ go install golang.org/dl/go1.19@latest
$ go1.19 download
```

#### Clone the project and install the dependencies:
Clone the project from Github using the following command:

```javascript
$ git clone git@github.com:Shreya-1358/Generate_Occupancy_Combo_Golang.git
```

Next run the following command to install dependencies.

```javascript
$ go get
```

#### Install Mockgen:
Mockgen is a mock generator tool which we are going to use to generate our mocks. To install run the following command:

```javascript
go install github.com/golang/mock/mockgen@v1.6.0
```

## API Reference

#### POST Create the occupancy combinations

```http
  POST /generate_occupancy
```

## Project Outline :-
* I've generated the occupancy combinations of adult and child for a given occupancy in golang. 
* I've taken BaseAdult, MaxAdult, BaseChild, MaxChild, MaxGuest and the ChildAgeRange as request and generating the occupancy combinations as response.
* I've also performed testing using mockgen on the GenerateOccupancy, GenerateCombination function and benchmarking on GenerateOccupancy function.
