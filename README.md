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

## Query parameters
While there is not an example of query parameters in this project we did cover them in the second session and they are part of the homework.  If we remind our selfs of the structure of a URL.

```
http://www.google.com?q=search&something=else
```

This url can be broken down into the following components:
* protocol = http://
* server = www.google.com
* querystring = ?q=search&something=else

The querystring is a grouping of key value pairs separated by `=` which follow the `?` multiple key value pairs are separated by `&`.  Deconstructing our example query string we would get the following pairs:
* field: q value: search
* field: sometihng value: else

If we need to access these values we can use the method on the `http.Request` object, given a handler which has the signature:

```
func ListKittens(rw http.ResponseWriter, r *http.Request)
```

We are passed a reference to the http.Request object which is named `r`, to access the query string we can therfore use the following syntax.

```
v := r.URL.Query().Get("q")
```

With our example query string the value of v would be `search`.

(https://golang.org/pkg/net/http/#Request)[https://golang.org/pkg/net/http/#Request]

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
 
