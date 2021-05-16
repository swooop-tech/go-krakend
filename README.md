# go-krakend

This repository is a sample project to simply test krakend api gateway with a very simple api backend simulation.

To run, make sure **krakend cli** is installed in your machine and run:

`krakend run -c krakend-config.json` -d

You can then call the endpoint configured in the json through postman and krakend will automatically call and delegate the backend response to the client.
