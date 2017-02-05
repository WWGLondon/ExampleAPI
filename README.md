# Session 2 Example API

This repository contains the example API from session 2 and should be the basis for the homework assignment.  

To start a webserver in GO we need to first register a route.  The route is accessed via http://localhost:9000/list in the browser which then executes the associated function `ListKittens`

```go
http.HandleFunc("/list", ListKittens)
```

This command registers a route on the `http.DefaultServeMux` which we use when we start the server like so.

```go
http.ListenAndServe(":9000", http.DefaultServeMux)
```

Now everytime you request /list the function ListKittens is executed.

```go
func ListKittens(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(pets)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
```

If we take a look at this function the first thing we are doing is converting the Go struct pets into a slice of bytes which can be written to the ResponseWriter `[]byte`.  If the conversion fails then we are catching the error and writing a status code 501 to the browser and exiting the function.

If the conversion succeedes and it should in our simple example we are writing the returned data to the ResponseWriter which in turn returns the data to your browser as JSON.

## Interfaces
When is a Dog a Kitten?  When they both are Pets, we use interfaces in Go to make objects interchangeable, the interface Pet in the animals package implements the following behaviour.

```go
type Pet interface {
	SetName(name string)
	GetName() string
	SetHobbies(hobbies []string)
	GetHobbies() []string
}
```

For a Kitten and a Dog to be defined as pets they need to both implement the methods defined in the interface.  If we take a look at when we are creating our slice of Pet in the main method you can see how this happens.

```go
pets = []animals.Pet{
		&animals.Kitten{
			Name: "Mr Tiggles",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
		&animals.Kitten{
			Name: "Top Cat",
			Hobbies: []string{
				"Taunting Officer Dibble",
			},
		},
		&animals.Dog{
			Name: "Fido",
			Hobbies: []string{
				"Barking",
			},
		},
	}
```

Look carefully at the `&` which precedes `animals.Dog` and `animals.Pet` and remember this is Go syntax for creating a reference.  If we did not add this when we are adding our pets to the slice then the compiler would complain.  The reason for this is that interfaces are always references, so it is not valid to attempt to assign a value type to a slice which is of interface even if the types implement the interface correctly.

## Building the API
```bash
go build main.go
./main //mac
main.exe // windows
```

## Running with go run
```bash
go run main.go
```

## Running tests
```bash
go test ./...
```
 
