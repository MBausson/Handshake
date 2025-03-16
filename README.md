# Handshake

## Presentation

This project is an attempt to implement the [Six degrees of separation](https://en.wikipedia.org/wiki/Six_degrees_of_separation) theory.  
The *Six degrees of separation* idea asserts that anybody in the world can be reached within 6 or fewer connections.  

The project is based on Steam friends, but can easily be adapted to another social platform.

## Usage

In order to use this program, you need :
- A Steam API key, [retrievable here](https://steamcommunity.com/dev/apikey)
- A starting point which is represented by a user's Steam ID
- A user to reach, based on his Steam ID

```shell
go build
```

```shell
./handshake --key=<Api Key> --from=<Steam ID> --to=<Steam ID> --depth=<Maximum depth, default is 6>
```

The program will then output diverse data while doing its work.  
In the end, if a user is reached, the chain of handshakes will be displayed.    


## Improvements

For the moment, the implementation is quite slow and memory consuming, especially for deep relations (> 3, 4).  
I'm planning on improving the algorithm & data structure used to address this issue.  

